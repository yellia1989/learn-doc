package main

// 测试conn的线程安全
// 结论:conn上的操作不是线程安全的

import (
    "fmt"
    "time"
    "sync"
    "context"
    "math/rand"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
)

func main() {
    db, err := sql.Open("mysql", "dev:777777@tcp(192.168.0.16:3306)/test")
    if err != nil {
        fmt.Printf("open db err: %v", err)
        return
    }

    readFn := func(conn *sql.Conn, thread string) {
        ctx := context.Background()
        rows, err := conn.QueryContext(ctx, "select id from student")
        if err != nil {
            fmt.Printf("thread %s err: %v\n", thread, err)
            return
        }
        defer rows.Close()

        for rows.Next() {
            var id int
            if err := rows.Scan(&id); err != nil {
                fmt.Printf("thread %s, scan err: %v\n", thread, err)
                return
            }
            fmt.Printf("thread %s, id: %d\n", thread, id)
        }
        conn.ExecContext(ctx, "delete from student")
    }

    writeFn := func(conn *sql.Conn, thread string) {
        id := rand.Intn(2000000)
        _, err := conn.ExecContext(context.Background(), "insert into student(id) values(?)", id)
        if err != nil {
            fmt.Printf("thread %s write err: %v\n", thread, err)
            return
        }
        fmt.Printf("thread %s write success: %d\n", thread, id)
    }

    var mutex sync.Mutex
    conn, err := db.Conn(context.Background())
    if err != nil {
        fmt.Printf("create conn err: %s\n", err)
        return
    }

    // 5个线程同时读
    for i := 0; i < 5; i++ {
        go func(id int) {
            for {
                mutex.Lock()
                readFn(conn, fmt.Sprintf("readt%d", id))
                mutex.Unlock()

                time.Sleep(time.Second * 1)
            }
        }(i+1)
    }

    // 2个线程写
    for i := 0; i < 2; i++ {
        go func(id int) {
            for {
                mutex.Lock()
                writeFn(conn, fmt.Sprintf("writet%d", id))
                mutex.Unlock()

                time.Sleep(time.Second * 1)
            }
        }(i+1)
    }

    for {
        t := time.NewTicker(time.Second * 1)
        select {
        case now := <- t.C :
            fmt.Printf("%v:%v\n", now, db.Stats())
        }
    }
}
