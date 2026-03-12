package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

func newSubcommand1() *cobra.Command {

    var count int

    cmd := &cobra.Command{
        Use: "subcommand1",
        Short: "A subcommand under command1",
        RunE: func(cmd *cobra.Command, args []string)error {
            return runSubcommand1(cmd, args, count)
        },
    }

	// flags
	cmd.Flags().IntVarP(&count, "count", "c", 1, "Number of iterations")

	return cmd
}

func runSubcommand1(cmd *cobra.Command, args []string, count int) error {

	fmt.Printf("Running subcommand1 with count=%d\n", count)

	return nil
}