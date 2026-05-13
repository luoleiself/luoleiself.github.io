package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"
)

func main() {
	process, err := os.StartProcess("/usr/bin/ls", []string{"ls", "-l"}, &os.ProcAttr{
		Env:   os.Environ(),
		Files: []*os.File{os.Stdin, os.Stdout, os.Stderr},
	})
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
		process.Kill()
	}
	fmt.Printf("The process id is %v\n", process)

	cmd := exec.Command("ls", "-la")
	err = cmd.Run()
	if err != nil {
		log.Fatal(err)
		// os.Exit(1)
		cmd.Cancel()
	}
	fmt.Println(cmd.Environ())
	fmt.Printf("The command is %v\n", cmd)

	ctx, cancel := context.WithCancel(context.TODO())
	cmd1 := exec.CommandContext(ctx, "sleep", "5")
	time.Sleep(time.Second * 2)
	cancel()
	err = cmd1.Start()
	if err != nil {
		log.Fatal(err)
		cmd1.Cancel()
	}
	fmt.Printf("The commandContext is  %v\n", cmd1)
}
