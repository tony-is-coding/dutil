package file

import (
	"filesystem/pkg/uidgen"
	"fmt"
	"log"
	_ "net/http/pprof"
	"os"
	"path/filepath"
)

var MaxReadSize = 8 * 1 << 20

// 读取单个大文件
// 切分为多份小文件在 [多个 goroutine中 单独进行IO读写]
func ReadFile(filePath string) (err error) {

	f, err := os.Open(filePath)
	defer f.Close()
	if err != nil {
		return err
	}
	split := make([]byte, MaxReadSize)
	fileCount := 0
	tmePath := fmt.Sprintf("%s/tmp/", filepath.Dir(filePath))
	if !FileExist(tmePath) {
		err = os.Mkdir(tmePath, os.ModePerm)
		if err != nil {
			return err
		}
	}
	err = RemoveUnderAll(tmePath)
	if err != nil {
		return err
	}

	// 长度为 10 的通道用于异常通信
	writeErr := make(chan error, 10)
	go func() {
		for {
			select {
			case err := <-writeErr:
				// TODO: 记录异常, 或者直接中断整次serve
				log.Println(err)
				// break
			}
		}
	}()

	for {
		switch n, err := f.Read(split[:]); {
		case n > 0:
			go func() {
				fileName := fmt.Sprintf("%s/tmp/%s.%d",
					filepath.Dir(filePath), filepath.Base(filePath), uidgen.SnowFlake())
				err := writeFile(split, fileName)
				if err != nil {
					writeErr <- err
				}
			}()
			fileCount++
			break
		case n == 0:
			return nil
		case n < 0:
			return err
		}
	}
}

// 将字节数据写入指定的路径中
func writeFile(p []byte, fileName string) error {

	file, err := os.Create(fileName)
	defer file.Close()
	if err != nil {
		return err
	}
	_, err = file.Write(p)
	if err != nil {
		return err
	}
	return nil
}
