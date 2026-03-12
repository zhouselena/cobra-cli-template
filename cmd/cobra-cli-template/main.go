package main

import (
	"log"

	"github.com/spf13/cobra"
)

func main() {
    cmd := &cobra.Command{
        Use: "cobra-cli-template",
        Short: "cobra-cli-template is a template for writing Go Cobra CLI tools",
        Version: "0.0.0-alpha",
    }

    cmd.AddCommand(newCommand1())

    if err := cmd.Execute(); err != nil {
        log.Fatalf("error: %v", err)
    }
}