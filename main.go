package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"
)

func main() {
	start := time.Now()
	if c, err := exec.Command(os.Args[1], os.Args[2:]...).CombinedOutput(); err != nil {
		log.Fatal(err)
	} else {
		end := time.Now()
		elapsed := end.Sub(start)
		fmt.Printf("Elapsed: %s", elapsed)
		fmt.Printf("%s\n", c)
	}

}
