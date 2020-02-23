package main

import (
	"log"
	"net"
	"time"
)

func main() {
	log.Println("begin dial...")
    addr, err := net.ResolveTCPAddr("tcp", ":8888")
    if err != nil {
        log.Println("dial error:", err)
        return
    }
	conn, err := net.DialTCP("tcp", nil, addr)
	if err != nil {
		log.Println("dial error:", err)
		return
	}

    conn.Close()

    go write(conn)
    go read(conn)

	time.Sleep(time.Second * 1000)
}

func write(conn *net.TCPConn) {
    for {
        // 每个2秒写一个
        time.Sleep(time.Second * 2)
        buf := []byte("client svr")
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
