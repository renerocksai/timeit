package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"
)

func main() {
	cmd := exec.Command(os.Args[1], os.Args[2:]...)
	start := time.Now()
	if output, err := cmd.CombinedOutput(); err != nil {
		log.Fatal(err)
	} else {
		end := time.Now()
		elapsed := end.Sub(start)
		fmt.Printf("Elapsed: %s", elapsed)
		fmt.Printf("%s\n", output)
	}
}
