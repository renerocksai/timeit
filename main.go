package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"
)

func main() {
	// this works, but is not 100% compatible with all commands, so better wrap them in a cmd process.
	//    cmd := exec.Command(os.Args[1], os.Args[2:]...)
	os.Args[0] = "/c" // override timeit.exe with /c for cmd, keeps exec.Command(...) simple
	cmd := exec.Command("cmd", os.Args...)
	start := time.Now()
	if output, err := cmd.CombinedOutput(); err != nil {
		log.Fatal(err)
	} else {
		end := time.Now()
		elapsed := end.Sub(start)
		fmt.Printf("Elapsed: %s\n", elapsed)
		fmt.Printf("%s\n", output)
	}
}
