package main

import (
	"fmt"
	"math/rand"
	"path/filepath"
	"runtime"
	"sync/atomic"
	"time"
	
)

var (
	IsCloseDebugInfo bool
	IsCloseErrorInfo bool
	
	// use to generate UUID
	// need to init hardwareAddr and clockSeq
	timeBase = time.Date(1582, time.October, 15, 0, 0, 0, 0, time.UTC).Unix()
	hardwareAddr []byte
	clockSeq uint32

)

func init() {
	IsCloseDebugInfo = false
	IsCloseErrorInfo = false
	
	// Init variable for generate UUID
	hardwareAddr = []byte{48,49,50,51,52,53}		//  生產環境需要改成虛擬機的MAC地址
	clockSeq = uint32(rand.Int31())
	
	// TODO
}

func DebugInfo(debugInfo string) {
	if IsCloseDebugInfo {
		return
	}

	_, file, line, _ := runtime.Caller(1)
	file = filepath.Base(file)
	fmt.Printf("%s:%d:%s\n\r", file, line, debugInfo)
}

func ErrorInfo(errInfo string) {
	if IsCloseErrorInfo {
		return
	}

	_, file, line, _ := runtime.Caller(1)
	file = filepath.Base(file)
	fmt.Printf("%s:%d:%s\n\r", file, line, errInfo)
}

type UUID [16]byte

func (u UUID) String() string {
	var offsets = [...]int{0, 2, 4, 6, 9, 11, 14, 16, 19, 21, 24, 26, 28, 30, 32, 34}
	const hexString = "0123456789abcdef"
	r := make([]byte, 36)
	for i, b := range u {
		r[offsets[i]] = hexString[b>>4]
		r[offsets[i]+1] = hexString[b&0xF]
	}
	r[8] = '-'
	r[13] = '-'
	r[18] = '-'
	r[23] = '-'
	return string(r)
}


func GenerateUUID() UUID {
	return FromTime(time.Now())
}

func FromTime(aTime time.Time) UUID {
	var u UUID

	utcTime := aTime.In(time.UTC)
	t := uint64(utcTime.Unix()-timeBase)*10000000 + uint64(utcTime.Nanosecond()/100)
	u[0], u[1], u[2], u[3] = byte(t>>24), byte(t>>16), byte(t>>8), byte(t)
	u[4], u[5] = byte(t>>40), byte(t>>32)
	u[6], u[7] = byte(t>>56)&0x0F, byte(t>>48)

	clock := atomic.AddUint32(&clockSeq, 1)
	u[8] = byte(clock >> 8)
	u[9] = byte(clock)

	copy(u[10:], hardwareAddr)

	u[6] |= 0x10 // set version to 1 (time based uuid)
	u[8] &= 0x3F // clear variant
	u[8] |= 0x80 // set to IETF variant

	return u
}


type IMSI [16]byte
func (imsi *IMSI)String() string{
	return ""
}
func GenerateIMSI(){
	//model.PlmnId{}
}
