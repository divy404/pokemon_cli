package cmd

import (
    "fmt"
    "github.com/spf13/cobra"
    "os"
)

var rootCmd = &cobra.Command{
    Use:   "pokecli",
    Short: "A CLI tool to explore Pokemon from PokeAPI",
    Long:  "pokecli lets you fetch and explore Pokemon data from PokeAPI using the terminal.",

}

func Execute() {
    if err := rootCmd.Execute(); err != nil {
        fmt.Fprintf(os.Stderr, "Oops. An error while executing:'%s'\n", err)
        os.Exit(1)
    }
}
func init() {
    
}