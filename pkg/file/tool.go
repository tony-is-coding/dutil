package file

import (
	"os"
	"path/filepath"
)

// 判断给定路径的文件是否存在
func FileExist(absPath string) bool {
	_, err := os.Stat(absPath)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}


// 移除指定路径目录下的 所有文件/文件夹 递归删除
func RemoveUnderAll(path string) error {
	files, err := filepath.Glob(filepath.Join(path, "*"))
	if err != nil {
		return err
	}
	for _, f := range files {
		return os.RemoveAll(f)
	}
	return nil
}