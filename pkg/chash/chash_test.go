package chash

import (
	"fmt"
	"strconv"
	"testing"
)

func TestHash32WithSeed(t *testing.T) {
	//randInt := func() uint8 {
	//	return uint8(rand.Intn(255))
	//}

	addressTable := []string{
		"127.0.0.1:8000",
		"127.0.0.2:8000",
		"127.0.0.3:8000",
		"127.0.0.4:8000",
		"127.0.0.5:8000",
	}

	for i := 0; i < 100; i++ {
		res := hash32WithSeed([]byte(strconv.Itoa(i)), 0xa1)
		fmt.Println(addressTable[res%uint32(len(addressTable))])
	}
}
