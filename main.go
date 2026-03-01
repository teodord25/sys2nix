package main

import (
	"os"

	"github.com/charmbracelet/log"
	"github.com/teodord25/sys2nix/internal/scanner"
	"github.com/teodord25/sys2nix/messenger"
)

func setupLogger() (*log.Logger, *os.File) {
	f, err := os.OpenFile("sys2nix.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		panic(err)
	}

	logger := log.New(f)
	logger.SetLevel(log.DebugLevel)
	logger.SetReportCaller(true)

	return logger, f
}

func main() {
	logger, f := setupLogger()
	defer f.Close()

	msg := messenger.NewMessenger(logger)

	distro, err := scanner.DetectDistro()
	if err != nil {
		msg.Warn("Failed to detect distro: %w", err)
	}
	msg.Success("Detected distro: %s", distro)

	msg.Info("Package managers found: %s", scanner.DetectSecondaryManagers())

	m, err := scanner.DetectPrimaryManager(distro)
	if err != nil {
		msg.Warn("Failed to detect primary package manager: %s", err)
	}
	msg.Info("Primary Package managers found: %s", m)
}
