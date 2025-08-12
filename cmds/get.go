package cmds

import (
	"fmt"
	"pulse/fn"

	"github.com/spf13/cobra"
)

func GetCMD() *cobra.Command{
	var repoName string;
	var protocol string;

		cmd := &cobra.Command{
		Use:   "get",
		Short: "Get data",
		Run: func(cmd *cobra.Command, args []string) {
			if repoName == "" || protocol == "" {
				fmt.Printf("You must specify a repository path --d\nYou must specify the path to 'protocol.json' --p")
			} 
			fn.FnGet(protocol,repoName)	
		},
	}



	cmd.Flags().StringVarP(&repoName, "repository", "re", "", "Repository path")
	cmd.Flags().StringVarP(&protocol, "protocol", "p", "", "Protocol path")
	cmd.MarkFlagRequired("repository")
	cmd.MarkFlagRequired("protocol")

	return cmd

}