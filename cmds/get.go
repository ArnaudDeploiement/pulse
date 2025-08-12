package cmds

import (
	"fmt"
	"os"
	"pulse/fn"

	"github.com/spf13/cobra"
)

func GetCMD() *cobra.Command{
	var storeDir string;
	var protocol string;

		cmd := &cobra.Command{
		Use:   "get",
		Short: "Get data",
		Run: func(cmd *cobra.Command, args []string) {
			if storeDir == "" || protocol == "" {
				fmt.Printf("You must specify a repository path --d\nYou must specify the path to 'protocol.json' --p")
				os.Exit(1)
			} 
			fn.FnGet(protocol,storeDir)	
		},
	}



	cmd.Flags().StringVarP(&storeDir, "repository", "r", "", "Repository path")
	cmd.Flags().StringVarP(&protocol, "protocol", "p", "", "Protocol path")
	cmd.MarkFlagRequired("repository")
	cmd.MarkFlagRequired("protocol")

	return cmd

}