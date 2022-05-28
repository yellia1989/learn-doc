package main

import (
	"log"
	"net"
	"time"
)

const gReadSleep = time.Second * 2
const gWriteSleep = time.Second * 2

func main() {

    log.Println("\n"+run_shell(`./stat.sh`))

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

    log.Println("\n"+run_shell(`./stat.sh`))

    //conn.Close()
    //log.Println("close read & write")

    log.Println("\n"+run_shell(`./stat.sh`))

    go write(conn)
    go read(conn)

	time.Sleep(time.Second * 1000)
}

func write(conn *net.TCPConn) {
    time.Sleep(gWriteSleep)
    for {
        // 每个2秒写一个
        time.Sleep(time.Second * 2)
        buf := []byte("hello svr")

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
