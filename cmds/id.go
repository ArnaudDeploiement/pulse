package cmds

import (
	"fmt"

	"github.com/spf13/cobra"
)

func IdCmd() *cobra.Command {
	var mode string // variable pour stocker la valeur du flag

	cmd := &cobra.Command{
		Use:   "id",
		Short: "Get or Add PeerId",
		Run: func(cmd *cobra.Command, args []string) {
		
			switch mode {
			case "id":
				fmt.Println("get id")
			case "add":
			fmt.Printf("add id %s", args[0:])
			default:
				fmt.Println("erreur")
			}
		},
	}

	// Déclaration du flag UNE seule fois
	cmd.Flags().StringVarP(&mode, "mode", "m", "id", "Action: id or add")

	return cmd
}
