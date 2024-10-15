package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

// 定义一个结构体，用于存储键值对
type Key_Value struct {
	data map[string]string
}

// 定义set方法，用于设置键值对
func (s *Key_Value) Set(key, value string) {
	s.data[key] = value
}

// 定义get方法，用于获取键值对
func (s *Key_Value) Get(key string) (string, bool) {
	value, ok := s.data[key]
	return value, ok
}

// 定义delete方法，用于删除键值对
func (s *Key_Value) Delete(key string) {
	delete(s.data, key)
}

// 定义handleConnection方法，用于处理连接
func handleConnection(conn net.Conn, store *Key_Value) {
	defer conn.Close()

	reader := bufio.NewReader(conn)
	// 创建一个新的读取器，将连接作为参数传入

	for {
		// 读取一行数据
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading from connection:", err)
			break
			// 如果报错则退出循环，关闭连接
		}

		// 去除字符串首尾的空格
		line = strings.TrimSpace(line)
		args := strings.Split(line, " ")

		// 根据命令执行相应的操作
		switch args[0] {

		case "PING":
			conn.Write([]byte("PONG\n"))
			// 当发送PING指令时，返回PONG

		case "ECHO":
			if len(args) > 1 {
				conn.Write([]byte(args[1] + "\n"))
				// 将第一个参数写入连接
			} else {
				conn.Write([]byte("\n"))
				// 如果没有参数，则写入一个空行
			}

		case "SET":
			if len(args) == 3 {
				store.Set(args[1], args[2])
				// 在存储中设置键值对
				conn.Write([]byte("DONE\n"))
				// 返回DONE表示设置成功
			} else {
				conn.Write([]byte("Wrong argument number for 'set'\n"))
				// 如果参数数量不为3，则返回错误信息
			}

		case "GET":
			if len(args) == 2 {
				value, ok := store.Get(args[1])
				// 从存储中获取键值对
				if ok {
					conn.Write([]byte(value + "\n"))
					// 当读取到键值对时，将值写入连接
				} else {
					conn.Write([]byte("NULL\n"))
					// 当没有读取到键值对时，返回NULL
				}
			} else {
				conn.Write([]byte("Wrong argument number for 'get'\n"))
				// 如果参数数量不为2，则返回错误信息
			}

		case "DEL":
			if len(args) == 2 {
				store.Delete(args[1])
				// 将键值对从存储中删除
				conn.Write([]byte("OK\n"))
				// Write OK to the connection
			} else {
				conn.Write([]byte("Wrong argument number for 'del'\n"))
				// 如果参数数量不为2，则返回错误信息
			}

		case "QUIT":
			conn.Write([]byte("I'll miss you!\n"))
			// 当发送QUIT指令时，返回I'll miss you!，表示成功推出
			return
			// 退出循环，关闭连接
		default:
			conn.Write([]byte("Unknown command: " + args[0] + "\n"))
			// 如果命令不在上述列表中，则返回错误信息
		}
	}
}

func main() {
	store := &Key_Value{data: make(map[string]string)}
	// 创建一个新的存储

	listener, err := net.Listen("tcp", ":6379")
	// 监听6379端口
	if err != nil {
		fmt.Println("Error listening, port:", err)
		return
		// 如果报错则退出程序
	}
	defer listener.Close()
	// 程序退出时关闭监听

	fmt.Println("Listening to port 6379")

	for {
		conn, err := listener.Accept()
		// 接受新的连接
		if err != nil {
			fmt.Println("Error connection:", err)
			continue
			// 如果报错则继续监听
		}

		fmt.Println("New connection from port", conn.RemoteAddr())

		go handleConnection(conn, store)
		// 创建一个新的协程，处理连接1
	}
}
