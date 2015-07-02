package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	job := Job{"GarbageCollection", log.New(os.Stdout, "LOG: ", log.Ldate)}
	fmt.Printf("Command: %s\nLog: %w+\n", job.Command, job.Logger)
}

type Job struct {
	Command string
	*log.Logger
}
