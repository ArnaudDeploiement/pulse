package main

import (
	"pulse/cmds"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

func cli() {

	const banner = `

__/\\\\\\\\\\\\\____/\\\________/\\\__/\\\_________________/\\\\\\\\\\\____/\\\\\\\\\\\\\\\_        
 _\/\\\/////////\\\_\/\\\_______\/\\\_\/\\\_______________/\\\/////////\\\_\/\\\///////////__       
  _\/\\\_______\/\\\_\/\\\_______\/\\\_\/\\\______________\//\\\______\///__\/\\\_____________      
   _\/\\\\\\\\\\\\\/__\/\\\_______\/\\\_\/\\\_______________\////\\\_________\/\\\\\\\\\\\_____     
    _\/\\\/////////____\/\\\_______\/\\\_\/\\\__________________\////\\\______\/\\\///////______    
     _\/\\\_____________\/\\\_______\/\\\_\/\\\_____________________\////\\\___\/\\\_____________   
      _\/\\\_____________\//\\\______/\\\__\/\\\______________/\\\______\//\\\__\/\\\_____________  
       _\/\\\______________\///\\\\\\\\\/___\/\\\\\\\\\\\\\\\_\///\\\\\\\\\\\/___\/\\\\\\\\\\\\\\\_ 
        _\///_________________\/////////_____\///////////////____\///////////_____\///////////////__
		
`
	root := &cobra.Command{
		Use:   "Pulse",
		Short: "Pulse - Version 0.1",
		Long:  color.HiMagentaString(banner) + "\n\n" + "Partage de données privées en P2P, sans serveur et en toute sécurité." + "\n\n" + color.HiCyanString("Pulse remplace le stockage et le partage traditionnels par un échange privé, sécurisé et décentralisé."),
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}


	createCmd:= cmds.CreateCmd();
	sendCmd:=cmds.SendCmd();
	receiveCmd:= cmds.ReceiveCMD();
	listGCmd:=cmds.ListCmd();
	stopCmd:=cmds.StopCmd()


	root.CompletionOptions.DisableDefaultCmd = true

	root.AddCommand(
		createCmd,
		sendCmd,
		receiveCmd,
		listGCmd,
		stopCmd,
	)

	root.SetHelpTemplate(`

{{with .Long}}{{.}}{{end}}

Utiliser Pulse :

  Pulse + commandes

Les commandes:
{{range .Commands}}{{if .IsAvailableCommand}}
  {{rpad .Name .NamePadding }} {{.Short}}
{{end}}{{end}}
`)

	root.Execute()

}