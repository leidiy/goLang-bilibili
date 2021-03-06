package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"

)

func demo1() {
	file, err := os.Open("./main.go")
	if err != nil {
		fmt.Printf("open file failed, err: %v\n", err)
	}
	defer file.Close()
}

// 循环读取
func demo2() {
	// 只读方式打开当前目录下的main.go文件
	file, err := os.Open("./main.go")
	if err != nil {
		fmt.Println("open file failed!, err:", err)
		return
	}
	defer file.Close()
	// 使用Read方法读取数据
	tmp := make([]byte, 128)
	var content []byte
	for {
		n, err := file.Read(tmp)
		if err ==  io.EOF {
			fmt.Println("文件读完了")
			break
		}
		if err != nil {
			fmt.Println("read file failed!, err:", err)
		}
		fmt.Printf("读取了%d字节数据\n", n)
		content = append(content, tmp[:n]...)
	}
	fmt.Println(string(content))
}

// bufio是在file的基础上封装了一层API，支持更多的功能。
func bufioDemo() {
	file, err := os.Open("./main.go")
	if err != nil {
		fmt.Println("open file failed, err:", err)
		return
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			fmt.Print(line)
			return
		}
		if err != nil {
			fmt.Println("read file failed, err:", err)
			return
		}
		fmt.Print(line)
	}
}

// ioutil.ReadFile读取整个文件
func readWholeFile() {
	content, err := ioutil.ReadFile("./main.go")
	if err != nil {
		fmt.Println("read file failed, err:", err)
		return
	}
	fmt.Println(string(content))
}

func writeDemo1() {
	file, err := os.OpenFile("xx1.txt", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("open file failed, err:", err)
		return
	}
	defer file.Close()
	str := "hello 沙河"
	file.Write([]byte(str))
	file.WriteString(str)
}

func writeDemo2() {
	file, err := os.OpenFile("xx2.txt", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("open file failed, err:", err)
		return
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	for i:=0; i<10; i++ {
		writer.WriteString("hello 沙河")  //将数据先写入缓存
	}
	writer.Flush()  //将缓存中的内容写入文件
}

func writeDemo3() {
	str := "ahello 沙河a"
	err := ioutil.WriteFile("./xx3.txt", []byte(str), 0666)
	if err != nil {
		fmt.Println("open file failed, err:", err)
		return
	}
}

func main() {
	demo1()
	println()
	demo2()
	println()
	bufioDemo()
	println()
	readWholeFile()
	println()
	writeDemo1()
	println()
	writeDemo2()
	println()
	writeDemo3()
}