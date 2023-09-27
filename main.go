package main

import (
    "fmt"
    "github.com/spf13/cobra"
    "os"
)

func main() {
    var rootCmd = &cobra.Command{Use: "hello"}

    var cmdHello = &cobra.Command{
        Use:   "hi",
        Short: "Prints 'Hello, World!'",
        Run: func(cmd *cobra.Command, args []string) {
            fmt.Println("Hello, World!")
        },
    }

    rootCmd.AddCommand(cmdHello)

    if err := rootCmd.Execute(); err != nil {
        fmt.Fprintln(os.Stderr, err)
        os.Exit(1)
    }
}
