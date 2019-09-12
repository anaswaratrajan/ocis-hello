package command

import (
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// Health is the entrypoint for the health command.
func Health() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "health",
		Short: "Check health status",
		Long:  "",
		Run: func(cmd *cobra.Command, args []string) {
			log.Info().
				Str("addr", viper.GetString("metrics.addr")).
				Msg("Executed health command")
		},
	}

	cmd.Flags().String("metrics-addr", "", "Address to metrics endpoint")
	viper.BindPFlag("metrics.addr", cmd.Flags().Lookup("metrics-addr"))
	viper.BindEnv("metrics.addr", "HELLO_METRICS_ADDR")

	return cmd
}
