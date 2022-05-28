package main

import (
    "flag"
    "time"
    "fmt"
    "errors"
    "strings"
)

// 赋值给flag变量
var ipFlag = flag.String("ip", "127.0.0.1", "ip")

// 绑定变量到flag
var timesFlag int
func init() {
    flag.IntVar(&timesFlag, "times", 0, "times")
}

// 自定义类型
type interval []time.Duration
// String is the method to format the flag's value, part of the flag.Value interface.
// The String method's output will be used in diagnostics.
func (i *interval) String() string {
    return fmt.Sprint(*i)
}

// Set is the method to set the flag value, part of the flag.Value interface.
// Set's argument is a string to be parsed to set the flag.
// It's a comma-separated list, so we split it.
func (i *interval) Set(value string) error {
    // If we wanted to allow the flag to be set multiple times,
    // accumulating values, we would delete this if statement.
    // That would permit usages such as
    //  -deltaT 10s -deltaT 15s
    // and other combinations.
    if len(*i) > 0 {
        return errors.New("interval flag already set")
    }
    for _, dt := range strings.Split(value, ",") {
        duration, err := time.ParseDuration(dt)
        if err != nil {
            return err
        }
        *i = append(*i, duration)
    }
    return nil
}
var intervalFlag interval
func init() {
    flag.Var(&intervalFlag, "deltaT", "comma-separated list of intervals to use between events")
}

func main() {
    flag.Parse()
    flag.PrintDefaults()

    fmt.Printf("ip=%s,times=%d,interval=%s\n", *ipFlag, timesFlag, intervalFlag)
}
