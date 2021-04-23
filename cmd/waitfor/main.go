package main

import (
    "flag"
    "fmt"
    "os"
    "strings"
    "time"
)

func main() {
    var envVars string
    var sleepVar time.Duration

    flag.StringVar(&envVars, "env", "", "comma separated list of ENV variables to check")
    flag.DurationVar(&sleepVar, "sleep", 5 * time.Second, "duration value")
    flag.Parse()

    if envVars == "" {
        fmt.Println("! env flag is required")
        os.Exit(1)
    }
    envs := strings.Split(envVars, ",")
    for _, env := range envs {
        value := os.Getenv(env)
        if value == "" {
            // one of the env is blank/not-set - immediately exit
            fmt.Println("= ["+env + "] is not set yet")
            // sleep to wait for the env var to be ready
            fmt.Printf("= sleep for %s before exit \n", sleepVar)
            time.Sleep(sleepVar)
            os.Exit(1)
            return
        }
    }

    fmt.Println("= all "+envVars + " is detected")
    os.Exit(0)
}
