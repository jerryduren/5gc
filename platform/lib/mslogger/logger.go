package mslogger

import "fmt"

type MsLogger interface {
	WriteLog(format string,a ...interface{})
}

var microServiceLogger MsLogger

func SetMicroServiceLogger(msLogger MsLogger){
	microServiceLogger = msLogger
}
func init(){

}

func d()  {
	fmt.Printf()
}

func Printf(format string, a ...interface{})(int, error){

}