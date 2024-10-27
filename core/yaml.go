package core

import (
	"fmt"

	"github.com/gkwa/fewduck/core/config"
	"github.com/go-logr/logr"
)

func ReadYAML(logger logr.Logger, path string) error {
	reader := config.NewFileReader(logger)

	cfg, err := reader.Read(path)
	if err != nil {
		return fmt.Errorf("failed to read config: %w", err)
	}

	fmt.Printf("yaml value: %v\n", cfg.X)
	return nil
}
