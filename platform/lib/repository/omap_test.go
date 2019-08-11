package repository

import (
	"fmt"
	"testing"
)

func TestNewStringKeyed(t *testing.T) {
	myOmap:= NewStringKeyed()
	
	// Test nil Map
	if myOmap.Delete("jerry"){
		t.Error("Test delete key from nil map failure")
	}
	if _,found:=myOmap.Find("Jerry");found{						// 这里应该查找失败
		t.Error("Test find in nil map failure")
	}
	
	
	if inserted:=myOmap.Insert("Neil","Neil is my son");!inserted{
		t.Error("Inserted Neil failure")
	}
	if inserted:=myOmap.Insert("Jerry","Jerry is neil's father.");!inserted{
		t.Error("Insert Jerry is failure")
	}
	
	if inserted:=myOmap.Insert("Dabao","Dabao is Jerry's file.");!inserted{
		t.Error("Insert Dabao is failure")
	}
	
	myOmap.Do(func(key,value interface{}){fmt.Printf("{%s: %s}\n",key,value)})
	
	if !myOmap.Delete("Jerry"){
		t.Error("Delete Jerry failure")
	}else{
		fmt.Println("Jerry is deleted.")
	}
	
	myOmap.Do(func(key,value interface{}){fmt.Printf("{%s: %s}\n",key,value)})
	
	// test fund
	if value,found:=myOmap.Find("Neil");!found{							// 这里应该查找成功
		t.Error("Test find in nil map failure")
	}else {
		fmt.Println(value)
	}
	
}