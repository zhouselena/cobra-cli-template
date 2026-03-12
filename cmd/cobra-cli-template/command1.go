package main

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	cobraclitemplate "github.com/zhouselena/cobra-cli-template"
)

func newCommand1() *cobra.Command {

    /* INSTANTIATE COMMAND */

    cmd := &cobra.Command{
        Use: "command1",
        Short: "command1 is a test command that does various prints",
        // Use if there are required arguments
        // Args: func(cmd *cobra.Command, args []string) error {
        //     if len(args) < 1 {
        //         return fmt.Errorf("this command requires {arg}")
        //     }
        //     return nil
        // },
    }

    /* FLAGS */

    var flag1, flag2 string
    cmd.Flags().StringVar(&flag1, "flag1", "", `specify {description}, ex. "example"`) // required flag
    cmd.Flags().StringVar(&flag2, "flag2", "", `specify {description}, ex. "example"`) // optional flag

    // mark required flags
    for _, flag := range []string{"flag1"} {
        cmd.MarkFlagRequired(flag)
    }

	/* ATTACH SUBCOMMANDS */
	cmd.AddCommand(newSubcommand1())

    /* RUN COMMAND1 */

    cmd.Run = func(cmd *cobra.Command, args []string) {
        // Use if required env variables
        // uri := os.Getenv("REQUIRED_ENDPOINT")
        // if uri == "" {
        //     log.Fatal("REQUIRED_ENDPOINT is not set")
        // }

        err := runCommand1(cmd, args, flag1, flag2)
        if err != nil {
            log.Fatalf("failed to run command1: %v", err)
        }
    }

    return cmd

}

func runCommand1(cmd *cobra.Command, args []string, flag1 string, flag2 string) error {

    str, err := cobraclitemplate.Command1(flag1, flag2)
    if err != nil {
        return err
    }

    fmt.Println(str)
    return nil
}