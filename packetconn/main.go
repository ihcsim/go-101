package main

import (
	"log"
	"net"
	"os"
	"strconv"
)

var logger *log.Logger = log.New(os.Stdout, "golang-example:", log.LstdFlags)

func main() {
	host := "127.0.0.1:7000"

	stop := make(chan struct{})
	defer close(stop)

	go listenAt(host)

	numMsgs := 1000
	bufferLen := 2 // vary this value to observe the behaviour of buffered and unbuffered channel
	capturedMsgs := make(chan []byte, bufferLen)
	defer close(capturedMsgs)

	go sendMsgTo(host, numMsgs, capturedMsgs)
	go checkMsgs(numMsgs, capturedMsgs, stop)

	<-stop
}

func listenAt(host string) {
	pconn, err := net.ListenPacket("udp", host)
	if err != nil {
		logger.Panicf("[Server] Fail to connect to %s. Error: %s\n", host, err.Error())
	} else {
		logger.Printf("[Server] Listening for UDP traffic at %s\n", host)
	}
	defer pconn.Close()

	for {
		var b []byte
		n, addr, err := pconn.ReadFrom(b)
		if err != nil {
			logger.Panicf("[Server] Fail to read data from %s. Error: %s\n", addr.String(), err.Error())
		}
		if n > 0 {
			logger.Printf("[Server] Read %d bytes from %s\n", n, addr)
		}
	}
}

func sendMsgTo(host string, numMsgs int, capturedMsgs chan []byte) {
	conn, err := net.Dial("udp", host)
	if err != nil {
		logger.Panicf("[Client] Fail to connect to %s, Error: %s\n", host, err.Error())
	}
	defer conn.Close()

	for i := 0; i < numMsgs; i++ {
		b := []byte("Hello World " + strconv.Itoa(i))
		capturedMsgs <- b
		n, err := conn.Write(b)
		if err != nil {
			logger.Panicf("[Client] Fail to write data from %s. Error: %s\n", host, err.Error())
		}

		if n > 0 {
			logger.Printf("[Client] Wrote %d bytes to %s\n", n, host)
		}
	}
}

func checkMsgs(numMsgs int, capturedMsgs chan []byte, stop chan struct{}) {
	count := 0
	for data := range capturedMsgs {
		count++
		logger.Printf("Captured messages: %s", string(data))
		if count == numMsgs {
			stop <- struct{}{}
			return
		}
	}
}
