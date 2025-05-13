package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"sync"
)

type Msg struct {
	msg      string
	sender   string
	senderId int
}

var (
	users      = make(map[int]net.Conn)
	nextUserId = 1
	mutex      sync.Mutex
	broadcast  = make(chan Msg)
)

type Command struct {
	desc string
	fn   func(net.Conn)
}

var cmds = make(map[string]Command)

func initCommands() {
	cmds["!help"] = Command{
		desc: "Shows all commands list",
		fn: func(conn net.Conn) {
			text := ""
			for cmdName, cmd := range cmds {
				text += fmt.Sprintf("%s - %s\n", cmdName, cmd.desc)
			}
			conn.Write([]byte(text))
		},
	}
	cmds["!online"] = Command{
		desc: "Shows cuurent chat online",
		fn: func(conn net.Conn) {
			text := fmt.Sprintf("Current chat online: %d\n", len(users))
			conn.Write([]byte(text))
		},
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	mutex.Lock()
	userId := nextUserId
	nextUserId++
	users[userId] = conn
	mutex.Unlock()

	rd := bufio.NewReader(conn)
	username, err := rd.ReadString('\n')
	if err != nil || len([]rune(username)) > 20 {
		conn.Close()
	}
	username = strings.TrimSpace(username)

	broadcast <- Msg{
		msg:      fmt.Sprintf("%s joined the chat", username),
		sender:   "SERVER",
		senderId: 0,
	}

	conn.Write([]byte("Welcome to the chat! Type !help to see all commands list. Don't be rude and have fun conversations!"))

	for {
		msg, err := rd.ReadString('\n')
		if err != nil {
			break
		}
		msg = strings.TrimSpace(msg)
		if msg == "" {
			continue
		}
		if cmd, ok := cmds[msg]; ok {
			cmd.fn(conn)
			continue
		}

		broadcast <- Msg{
			msg:      msg,
			sender:   username,
			senderId: userId,
		}
	}

	broadcast <- Msg{
		msg:      fmt.Sprintf("%s left the chat", username),
		sender:   "SERVER",
		senderId: 0,
	}

	mutex.Lock()
	delete(users, userId)
	mutex.Unlock()
}

func broadcaster() {
	for {
		msg := <-broadcast
		for i, conn := range users {
			if msg.senderId != i {
				text := fmt.Sprintf("%s: %s\n", msg.sender, msg.msg)
				conn.Write([]byte(text))
				fmt.Print(text)
			}
		}
	}
}

func main() {
	addr, err := net.ResolveTCPAddr("tcp", "localhost:9000")
	if err != nil {
		panic(err)
	}
	ln, err := net.ListenTCP("tcp", addr)
	if err != nil {
		panic(err)
	}
	defer ln.Close()
	initCommands()

	fmt.Println("Chat running on localhost:9000")

	go broadcaster()

	for {
		conn, err := ln.AcceptTCP()
		if err != nil {
			continue
		}

		go handleConnection(conn)
	}
}
