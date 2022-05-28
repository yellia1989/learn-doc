//server.go

package main

import (
	"log"
	"net"
    "time"
)

const gReadSleep = time.Second * 10
const gWriteSleep = time.Second * 10

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

        c.Close()
        log.Println("close read & write")

		go read(c)
		go write(c)
	}
}

func write(conn *net.TCPConn) {

    time.Sleep(gWriteSleep)

    for {
        // 每个2秒写一个
        time.Sleep(time.Second * 2)
        buf := []byte("client hello")

        log.Println("write begin\n"+run_shell(`./stat.sh`))
    	n, err := conn.Write(buf)
	    if err != nil {
		    log.Println("write error:", err)
            log.Println("write end\n"+run_shell(`./stat.sh`))
            break
	    } else {
		    log.Printf("write %d bytes, content is %s\n", n, string(buf[:n]))
            log.Println("write end\n"+run_shell(`./stat.sh`))
        }
    }
}

func read(conn *net.TCPConn) {

    time.Sleep(gReadSleep)

    for {
        // 每个2秒读一个
        time.Sleep(time.Second * 2)
        buf := make([]byte, 100)

        log.Println("read begin\n"+run_shell(`./stat.sh`))
	    n, err := conn.Read(buf)
	    if err != nil {
	    	log.Println("read error:", err)
            log.Println("read end\n"+run_shell(`./stat.sh`))
            break
	    } else {
	    	log.Printf("read %d bytes, content is %s\n", n, string(buf[:n]))
            log.Println("read end\n"+run_shell(`./stat.sh`))
	    }
    }
}
