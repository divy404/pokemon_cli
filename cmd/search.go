package cmd

import (
	"fmt"
	"pokemon_cli/api"
	"strings"
	"github.com/spf13/cobra"
)



var typeName string

var searchCmd = &cobra.Command{
    Use:   "search",
    Short: "Search Pokémon by type (use --type)",
    Run: func(cmd *cobra.Command, args []string) {
        if typeName == "" {
            fmt.Println("Please provide a type with --type")
            return
        }
        data, err := api.FetchPokemonByType(typeName)
        if err != nil {
            fmt.Println("Error:", err)
            return
        }
        fmt.Printf("\nPokémon of type '%s':\n", strings.Title(typeName))
        for i, p := range data.Pokemon {
            fmt.Printf("%d. %s\n", i+1, strings.Title(p.Pokemon.Name))
        }
        fmt.Println()
    },
}

func init() {
    searchCmd.Flags().StringVarP(&typeName, "type", "t", "", "Type of Pokémon (e.g., fire, water)")
    rootCmd.AddCommand(searchCmd)
}