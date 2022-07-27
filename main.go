/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/marianina8/eu2022/cmd"
)

func main() {
	cmd.Execute()

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGINT)
	go func() {
		<-c
		// handle exit gracefully
		fmt.Println("handle exit gracefully")
		os.Exit(0)
	}()
}
