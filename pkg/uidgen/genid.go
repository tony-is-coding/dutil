package uidgen

import (
	"os"
	"sync"
	"time"
)

// 从 2008-08-08 08:00:00 开始
const StartEpoch = 1218182400000
const MachineNum = 16

var PID int

var Sequence  = 0

var L sync.Mutex

func init() {
	PID = os.Getpid()
}

func SnowFlake() uint64 {
	L.Lock()
	if Sequence == 1000 {
		Sequence = 0
	}
	num := Sequence
	Sequence++
	L.Unlock()
	return (uint64(time.Now().UnixNano()/1000000)-StartEpoch)<<22 | MachineNum<<10 | uint64(num)
}
