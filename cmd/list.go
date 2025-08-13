package cmd

import (
	"fmt"
	"pokemon_cli/api"
	"strings"

	"github.com/spf13/cobra"
)

var limit int 
var offset int 

var listCmd = &cobra.Command {
	Use: "list",
	Short: "List pokemon with pagination",
	Run: func(cmd *cobra.Command, args []string) {
		list, err := api.FetchPoekmonList(limit,offset)
		if err != nil {
			fmt.Println("Error:",err)
			return
		}
		fmt.Printf("\nShowing %d Pokemon (offset %d):\n",len(list.Results),offset)
		for i, p := range list.Results {
			fmt.Printf("%d, %s\n", offset+i+1, strings.Title(p.Name))
		}
		fmt.Println()
	},
}
func init() {
	listCmd.Flags().IntVarP(&limit, "limit", "l", 10, "Number of Pokémon to list")
    listCmd.Flags().IntVarP(&offset, "offset", "o", 0, "Offset for listing Pokémon")
    rootCmd.AddCommand(listCmd)

}