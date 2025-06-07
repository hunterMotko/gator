package cmd

import "github.com/huntermotko/gator/internal/config"

func handlerInit() error {
	if err := config.InitConfig(); err != nil {
		return err
	}
	return nil
}
