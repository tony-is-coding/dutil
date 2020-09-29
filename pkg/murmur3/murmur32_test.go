package murmur3

import (
	"fmt"
	"strconv"
	"testing"
)

func TestMinMax(t *testing.T) {
	var min uint32
	var max uint32
	min, max = uint32((1<<32) -1) , 0

	for i := 0; i < 10000000; i++ {
		h32 := New32WithSeed(0x01)
		h32.Write([]byte(strconv.Itoa(i)))
		res := h32.Sum32()
		if res < min {
			min = res
		}
		if res > max {
			max = res
		}
	}
	fmt.Printf("min: %d \n max: %d ", min, max)

}
