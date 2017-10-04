package main

import (
  "bytes"
  "net"
  "fmt"
  "bufio"
  "os"
)

type Server struct {
  peers []string
  name string
}

func main() {
  var buffer bytes.Buffer
  PORT := os.Args[1]
  NAME := os.Args[2]

  s := &Server{[]string{PORT, }, NAME}
  fmt.Println("Launching server on port " + PORT)

  // listen on all interfaces
  buffer.WriteString(":")
  buffer.WriteString(PORT)
  listener, _ := net.Listen("tcp", buffer.String())
  defer listener.Close()

  // accept connection on port
  conn, _ := listener.Accept()

  // run loop forever (or until ctrl-c)
  for {
    // will listen for message to process ending in newline (\n)
    message, _ := bufio.NewReader(conn).ReadString('\n')
    // output message received
    fmt.Print("Message Received:", string(message))
    // sample process for string received
    newmessage := s.name
    // send new string back to client
    conn.Write([]byte(newmessage + "\n"))
  }
}
