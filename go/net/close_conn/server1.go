//server.go

package main

import (
	"log"
	"net"
    "time"
)

func main() {
    addr, err := net.ResolveTCPAddr("tcp", ":8888")
    if err != nil {
        log.Println("listen error:", err)
        return
    }
	l, err := net.ListenTCP("tcp", addr)
	if err != nil {
		log.Println("listen error:", err)
		return
	}

	for {
		c, err := l.AcceptTCP()
		if err != nil {
			log.Println("accept error:", err)
			break
		}
		// start a new goroutine to handle
		// the new connection.
		log.Println("accept a new connection")
		go read(c)
		go write(c)
	}
}

func write(conn *net.TCPConn) {

    time.Sleep(time.Second * 10)

    for {
        // 每个2秒写一个
        time.Sleep(time.Second * 2)
        buf := []byte("client hello")
    	n, err := conn.Write(buf)
	    if err != nil {
		    log.Println("write error:", err)
            break
	    } else {
		    log.Printf("write %d bytes, content is %s\n", n, string(buf[:n]))
        }
    }
}

func read(conn *net.TCPConn) {

    time.Sleep(time.Second * 20)

    for {
        // 每个2秒读一个
        time.Sleep(time.Second * 2)
        buf := make([]byte, 100)
	    n, err := conn.Read(buf)
	    if err != nil {
	    	log.Println("read error:", err)
            break
	    } else {
	    	log.Printf("read %d bytes, content is %s\n", n, string(buf[:n]))
	    }
    }
}
