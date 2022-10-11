package main

import write "demos/file_write"

func main() {
	// 测试文件写入
	if err := write.Begin(); err != nil {
		panic(err)
	}
}
