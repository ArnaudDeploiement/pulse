package cmds

import (
	"fmt"
	"os"

	"pulse/fn"

	"github.com/spf13/cobra"
)



func CreateCmd() *cobra.Command{

var groupName string
var relayAddr string
var outdir string

	
	cmd := &cobra.Command{
		Use:   "create",
		Short: "Create a protocol",
		Run: func(cmd *cobra.Command, args []string) {
			if groupName == "" {
				fmt.Println("You have to give a name with -n")
				os.Exit(1)
			}
			if relayAddr == "" {
				fmt.Println("You have to give a relay with -r")
				os.Exit(1)
			}
		
		  fmt.Println(fn.FnCreate(groupName, relayAddr, outdir))	
		},
	}

	cmd.Flags().StringVarP(&groupName, "name", "n", "", "Group's name")
	cmd.Flags().StringVarP(&relayAddr, "relay", "r", "", "Relay's address")
	cmd.Flags().StringVarP(&outdir, "outdir", "o", "", "Outdir path")
	cmd.MarkFlagRequired("name")
	cmd.MarkFlagRequired("relay")



	return cmd
}

