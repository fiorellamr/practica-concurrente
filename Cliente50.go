package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

type TCPClient50 struct {
	SERVERIP             string
	SERVERPORT           int
	mMessageListener     func(string)
	mRun                 bool
	out                  *bufio.Writer
	incoming             *bufio.Reader
}

func NewTCPClient50(ip string, messageReceivedCallback func(string)) *TCPClient50 {
	return &TCPClient50{
		SERVERIP:         ip,
		SERVERPORT:       4444,
		mMessageListener: messageReceivedCallback,
		mRun:             false,
		out:              nil,
		incoming:         nil,
	}
}

func (c *TCPClient50) SendMessage(message string) {
	if c.out != nil {
		c.out.WriteString(message + "\n")
		c.out.Flush()
	}
}

func (c *TCPClient50) StopClient() {
	c.mRun = false
}

func (c *TCPClient50) Run() {
	c.mRun = true
	fmt.Println("TCP Client" + "C: Conectando...")
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", c.SERVERIP, c.SERVERPORT))
	if err != nil {
		fmt.Println("TCP" + "C: Error", err)
		return
	}
	defer conn.Close()

	c.out = bufio.NewWriter(conn)
	fmt.Println("TCP Client" + "C: Sent.")
	fmt.Println("TCP Client" + "C: Done.")

	c.incoming = bufio.NewReader(conn)
	for c.mRun {
		serverMessage, err := c.incoming.ReadString('\n')
		if err != nil {
			fmt.Println("TCP" + "C: Error", err)
			return
		}
		serverMessage = strings.TrimRight(serverMessage, "\n")
		if serverMessage != "" && c.mMessageListener != nil {
			c.mMessageListener(serverMessage)
		}
	}
}

type Cliente50 struct {
	mTcpClient *TCPClient50
}

func NewCliente50() *Cliente50 {
	return &Cliente50{
		mTcpClient: nil,
	}
}

func (c *Cliente50) Main() {
	objcli := NewCliente50()
	objcli.Iniciar()
}

func (c *Cliente50) Iniciar() {
	onMessageReceived := func(message string) {
		c.ClienteRecibe(message)
	}

	tcpClientThread := func() {
		c.mTcpClient = NewTCPClient50("192.168.18.5", onMessageReceived)
		c.mTcpClient.Run()
	}

	go tcpClientThread()

	salir := "n"
	fmt.Println("Cliente bandera 01")
	for salir != "s" {
		sc := bufio.NewScanner(os.Stdin)
		sc.Scan()
		salir = sc.Text()
		c.ClienteEnvia(salir)
	}
	fmt.Println("Cliente bandera 02")
}

func (c *Cliente50) ClienteRecibe(llego string) {
	fmt.Println("CLINTE50 El mensaje::" + llego)
}

func (c *Cliente50) ClienteEnvia(envia string) {
	if c.mTcpClient != nil {
		c.mTcpClient.SendMessage(envia)
	}
}

func main() {
	objcli := NewCliente50()
	objcli.Iniciar()
}
