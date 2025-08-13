package cmd

import (
	"fmt"
	"pokemon_cli/api"
	"strings"
	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use: "get [nameOrID]",
	Short: "Get details of a Pokemon",
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		p, err := api.FetchPokemon(args[0])
		if err!= nil {
			fmt.Println("Error:", err)
			return
		}
		fmt.Printf("\n=== %s (ID: %d) ===\n", strings.Title(p.Name), p.ID)
		fmt.Printf("Height: %d | Weight: %d\n", p.Height, p.Weight)

		// Types
		fmt.Print("Types: ")
		for i, t := range p.Types {
			if i > 0 {
				fmt.Print(", ")
			}
			fmt.Print(strings.Title(t.Type.Name))
		}
		// abilities 
		fmt.Print("\nAbilities: ")
		for i, a := range p.Abilities {
			if i > 0 {
				fmt.Printf(", ")
			}
			fmt.Print(strings.Title(a.Ability.Name))
		}
		fmt.Println("\n=====================")
	},
}
func init () {
	rootCmd.AddCommand(getCmd)
}