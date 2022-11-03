package cmd

import (
	"fmt"

	"github.com/derickr/go-litra-driver"
	"github.com/spf13/cobra"
)

func withDevice(cmd *cobra.Command, dev **litra.LitraDevice) {
	cmd.PreRunE = func(cmd *cobra.Command, args []string) (err error) {
		// litra.New panics a lot
		defer func() {
			if recovered := recover(); recovered != nil {
				if recoveredErr, ok := recovered.(error); ok {
					err = fmt.Errorf("unable to open litra device: %w", recoveredErr)
				} else {
					err = fmt.Errorf("unable to open litra device: %v", recovered)
				}
			}
		}()
		*dev, err = litra.New()
		return err
	}
	cmd.PostRunE = func(cmd *cobra.Command, args []string) error {
		if *dev != nil {
			(*dev).Close()
		}
		return nil
	}
}
