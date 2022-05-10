package main

import (
    "runtime"
    "fmt"
    "regexp"
)

var a = PrintCallerName(0, "main.a")
var b = PrintCallerName(0, "main.b")

func init() {
    a = PrintCallerName(0, "main.init.a")
}

func init() {
    b = PrintCallerName(0, "main.init.b")
    func() {
        b = PrintCallerName(0, "main.init.b[1]")
    }()
}

func main() {
    /*
    fmt.Println("Caller::\n")

    for skip := 0; ; skip++ {
        _, file, line, ok := runtime.Caller(skip)
        if !ok {
            break
        }
        fmt.Printf("skip:%d %s:%d\n", skip, file, line)
    }

    fmt.Println("\nCallers::")

    pc := make([]uintptr, 10)
    for skip := 0; ; skip++ {
        n := runtime.Callers(skip, pc)
        if n == 0 {
            break
        }
        if frames := runtime.CallersFrames(pc[:n]); frames != nil {
            for {
                frame, more := frames.Next()
                fmt.Printf("skip:%d %s %d %s\n", skip, frame.File, frame.Line, frame.Function)
                if !more {
                    break
                }
            }
        }
    }
    */
    
    a = PrintCallerName(0, "main.main.a")
    b = PrintCallerName(0, "main.main.b")
    func() {
        b = PrintCallerName(0, "main.main.b[1]")
        func() {
            b = PrintCallerName(0, "main.main.b[1][1]")
        }()
        func() {
            b = PrintCallerName(0, "main.main.b[1][2]")
        }()
        b = PrintCallerName(0, "main.main.b[2]")
    }()
}

func CallerNameOld(skip int) (name, file string, line int, ok bool) {
    var pc uintptr
    if pc, file, line, ok = runtime.Caller(skip + 1); !ok {
        return
    }
    name = runtime.FuncForPC(pc).Name()

    return
}

func CallerName(skip int, noclosure bool) (name, file string, line int, ok bool) {
    var (
        reClosure = regexp.MustCompile(`\.func\d+$`) // main.func1
        reClosure2 = regexp.MustCompile(`\.\d+$`) // main.func1.1
        reClosure3 = regexp.MustCompile(`main\.init.\d+`) // main.init.0
    )
    for {
        var pc uintptr
        if pc, file, line, ok = runtime.Caller(skip + 1); !ok {
            return
        }
        name = runtime.FuncForPC(pc).Name()
        if !noclosure {
            return
        }
        if reClosure3.MatchString(name) {
            name = "main.init"
            fmt.Printf("skip:%d %s\n",skip, name)
            return
        }
        if reClosure.MatchString(name) || reClosure2.MatchString(name) {
            fmt.Printf("skip:%d %s\n",skip, name)
            skip++
            continue
        } else {
            fmt.Printf("skip:%d %s\n",skip, name)
        }
        return
    }
    return
}

func PrintCallerName(skip int, comment string) bool {
    /*for skip := 0; ; skip++ {
        _, file, line, ok := runtime.Caller(skip+1)
        if !ok {
            return false
        }
        fmt.Printf("skip:%d %s:%d\n", skip+1, file, line)
    }
    return true
    */

    name, file, line, ok := CallerName(skip + 1, false)
    if !ok {
        return false
    }
    fmt.Printf("skip = %v, comment = %s\n", skip, comment)
    fmt.Printf("  file = %v, line = %d\n", file, line)
    fmt.Printf("  name = %v\n", name)
    fmt.Printf("--------------------------------------\n")
    return true
}
