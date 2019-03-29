package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"text/scanner"
	
)

func main() {
	// Scanner a string
	src := "I am students, and I love technology."
	localScanner:=scanner.Scanner{}
	
	localScanner.Init(bytes.NewReader([]byte(src)))
	
	tok:=localScanner.Scan()
	for tok!=scanner.EOF{
		fmt.Println(localScanner.TokenText())
		tok=localScanner.Scan()
	}
	
	// Read a file, and print according to line.
	myFile,err:=os.Open("C:\\jerryduren@github.com\\mygoproject\\src\\github.com\\jerryduren\\5gc\\demo\\demoscanner.go")
	if err!=nil{
		log.Fatal(err)
		return
	}
	buf:=bufio.NewReader(myFile)
	line,err:=buf.ReadString('\n')		// 這個間隔符也是讀到line裏面了的！
	for err==nil || err==io.EOF{
		fmt.Print(line)
		if err==io.EOF{
			break
		}
		line,err=buf.ReadString('\n')
	}
	myFile.Close()
	
	// Scanner a file
	myFile,err=os.Open("C:\\jerryduren@github.com\\mygoproject\\src\\github.com\\jerryduren\\5gc\\demo\\demoscanner.go")
	if err!=nil{
		log.Fatal(err)
		return
	}
	wordMap:=make(map[string]int)
	fileScanner:=scanner.Scanner{}
	fileScanner.Init(myFile).Mode = scanner.GoTokens	// 設置如何區分一個token
	tok=fileScanner.Scan()
	for tok!=scanner.EOF{
		wordMap[fileScanner.TokenText()]++
		tok=fileScanner.Scan()
	}
	myFile.Close()
	for k,v:=range wordMap{
		fmt.Println(k,":",v)
	}
	
	// demo bufio
	fmt.Println("Please input a line string, system will echo display it.")
	for {
		inputIO := bufio.NewReader(os.Stdin)
		outputIO := bufio.NewWriter(os.Stdout)
		str, err := inputIO.ReadString('\n')
		if err!=nil{
			break
		}
		if strings.HasPrefix(str,"exit"){
			break
		}
		outputIO.WriteString(str)
		outputIO.Flush()
	}
	
	// 接下來用bytes.Buffer對象來實現類似的功能
	// bufio.Reader/bufio.Writer與bytes.Buffer有啥區別？
	// 一個是bytes流的緩存隊列，一個是對I/O對象加緩衝
	src ="what is this? This is a desk. what is that? that is a pen."
	bf:=bytes.NewBufferString(src)
	line,err=bf.ReadString('?')
	fmt.Println(line)
	bf.WriteTo(os.Stdout)
	fmt.Println(src)
}
