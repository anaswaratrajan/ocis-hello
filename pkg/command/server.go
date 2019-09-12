package command

import (
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// Server is the entrypoint for the server command.
func Server() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "server",
		Short: "Start integrated server",
		Long:  "",
		Run: func(cmd *cobra.Command, args []string) {
			log.Info().
				Str("addr", viper.GetString("server.addr")).
				Msg("Executed server command")
		},
	}

	cmd.Flags().String("server-addr", "", "Address to bind the server")
	viper.BindPFlag("server.addr", cmd.Flags().Lookup("server-addr"))
	viper.BindEnv("server.addr", "HELLO_SERVER_ADDR")

	return cmd
}
