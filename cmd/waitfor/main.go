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
    var durVar time.Duration
    flag.StringVar(&envVars, "env", "", "comma separated list of ENV variables to check")
    flag.DurationVar(&durVar, "sleep", 5 * time.Second, "duration value")
    flag.Parse()

    if envVars == "" {
        fmt.Println("env flag is required")
        os.Exit(1)
    }
    envs := strings.Split(envVars, ",")
    for _, env := range envs {
        value := os.Getenv(env)
        if value == "" {
            // one of the env is blank/not-set - immediately exit
            fmt.Println(env + " is not set yet")
            // sleep to wait for the env var to be ready
            fmt.Printf("Sleep for %s before exit \n", durVar)
            time.Sleep(durVar)
            os.Exit(1)
            return
        }
    }

    os.Exit(0)
}
