package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	addr, err := net.ResolveTCPAddr("tcp", "localhost:9000")
	if err != nil {
		panic(err)
	}
	conn, err := net.DialTCP("tcp", nil, addr)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	fmt.Print("Type your nickname: ")
	rd := bufio.NewReader(os.Stdin)
	username, err := rd.ReadString('\n')
	username = strings.TrimSpace(username)
	if err != nil {
		panic(err)
	}
	fmt.Fprintln(conn, username)

	go func() {
		sc := bufio.NewReader(conn)
		for {
			fmt.Print()
			msg, err := sc.ReadString('\n')
			if err != nil {
				continue
			}
			msg = strings.TrimSpace(msg)
			fmt.Println()
			fmt.Println(msg)
			fmt.Printf("%s: ", username)
		}
	}()

	for {
		fmt.Printf("%s: ", username)
		msg, err := rd.ReadString('\n')
		if err != nil {
			continue
		}
		msg = strings.TrimSpace(msg)
		fmt.Fprintln(conn, msg)
	}
}
