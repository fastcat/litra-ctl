package cmd

import "github.com/spf13/cobra"

func newRoot() *cobra.Command {
	root := &cobra.Command{
		Use:           "litra-ctl",
		SilenceUsage:  true,
		SilenceErrors: true,
		Short:         "Control Logictech Litra Devices",
	}

	root.AddCommand(
		newOn(),
		newOff(),
		newBrightness(),
		newTemperature(),
		newScript(),
	)

	return root
}

func Execute() error {
	return newRoot().Execute()
}
