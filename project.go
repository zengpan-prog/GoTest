package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//创建一个新文件，写入内容 5句 hello,Go语言
	filePath := "E:/Download/GoTest/gotest.txt"
	file, err := os.OpenFile(filePath, os.O_WRONLY | os.O_CREATE, 0666)
	if err != nil{
		fmt.Println("open file err",err)
	}
	//及时关闭file句柄
	defer file.Close()

	//写入文件时，使用带缓存的 *Writer
	write := bufio.NewWriter(file)
	for i := 0; i < 10; i++ {
		_, _ = write.WriteString("hello,Go语言 \n")
	}
	//Flush将缓存的文件真正写入到文件中
	_ = write.Flush()
}
