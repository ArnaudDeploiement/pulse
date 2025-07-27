package cmds

import (
	"fmt"

	"github.com/spf13/cobra"
)

func ListCmd() *cobra.Command{

var listgroup []string;
		
	cmd := &cobra.Command{
		Use:   "Liste",
		Short: "Liste de vos groupes",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("voici la liste des groupes %s", listgroup)
	},
}

return cmd

}