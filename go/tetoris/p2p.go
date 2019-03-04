package main

import (
  "fmt"
  "os"
  "text/scanner"
  "github.com/sacOO7/socketcluster-client-go/scclient"
)

const (
  HOST = "signaling"
  PORT = 3000
  ROOM = "tetoris"
)

type MyEventData struct {
  Type string
}

func onConnect(client scclient.Client) {
  fmt.Println("Connected to server")
}

func onDisconnect(client scclient.Client, err error) {
  fmt.Printf("Error: %s\n", err.Error())
}

func onConnectError(client scclient.Client, err error) {
  fmt.Printf("Error: %s\n", err.Error())
}

func onSetAuthentication(client scclient.Client, token string) {
  fmt.Println("Auth token received :", token)
}

func onAuthentication(client scclient.Client, isAuthenticated bool) {
  fmt.Println("Client authenticated :", isAuthenticated)
  go startCode(client)
}

//func createClient(socket *gosocketio.Client, wg *sync.WaitGroup) {
//  for true {
//  socket.On("connect", func(evt *gosocketio.Channel) {
//    socket.Emit("enter", ROOM)
//    fmt.Println("socket.io connected. enter room")
//  })
//  socket.On("message", func(message *gosocketio.Channel) {
//    fmt.Println(message)
//  })
//  socket.On("user disconnected", func(evt string) {
//    fmt.Println("====user disconnected==== evt:%s", evt);
//  })
//  socket.Emit("message", MyEventData{"call me"})
//  }
//
//  wg.Done()
//}

func main() {
  var reader scanner.Scanner
  socket := scclient.New("ws://signaling:3000/socket.io/?EIO=3&transport=websocket");
  socket.SetBasicListener(onConnect, onConnectError, onDisconnect)
  socket.SetAuthenticationListener(onSetAuthentication, onAuthentication)
  socket.On("connect", func(eventName string, data interface{}) {
    fmt.Println(eventName)
  })
  go socket.Connect()

  fmt.Println("Enter any key to terminate the program")
  reader.Init(os.Stdin)
  reader.Next()
}

func startCode(socket scclient.Client) {
  // start writing your code from here
  // All emit, receive and publish events
  socket.Emit("enter", "kirby")
}
