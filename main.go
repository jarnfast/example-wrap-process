package main

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

func main() {

	if len(os.Args) == 1 {
		fmt.Fprintf(os.Stderr, "no command is given, example-wrap-process can't determine the entrypoint")
		os.Exit(1)
	}

	wrapCmd := os.Args[1:]

	binary, err := exec.LookPath(wrapCmd[0])
	if err != nil {
		fmt.Fprintf(os.Stderr, "binary to wrap (%s) not found: %v", wrapCmd[0], err)
		os.Exit(1)
	}

	fmt.Printf("spawning: %v\n", wrapCmd)

	err = syscall.Exec(binary, wrapCmd, os.Environ())
	if err != nil {
		fmt.Fprintf(os.Stderr, "exec failed: %v	", err)
		os.Exit(1)
	}
}
