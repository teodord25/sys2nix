package scanner

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

type Distro string

const (
	Arch    Distro = "archlinux"
	Debian  Distro = "debian"
	Ubuntu  Distro = "ubuntu"
	Manjaro Distro = "manjaro"
	NixOS   Distro = "nixos"
)

func DetectDistro() (Distro, error) {
	data, err := os.ReadFile("/etc/os-release")
	if err != nil {
		return "", fmt.Errorf("reading os-release: %w", err)
	}

	for line := range strings.SplitSeq(string(data), "\n") {
		key, val, ok := strings.Cut(line, "=")
		if ok && key == "ID" {
			return Distro(strings.Trim(val, `"`)), nil
		}
	}

	return "", fmt.Errorf("no ID field found in os-release")
}

type DistroManager struct {
	distro Distro
}

type PackageManager string

const (
	Pacman PackageManager = "pacman"
	Apt    PackageManager = "apt"
	Nix    PackageManager = "nix"
)

func checkPresent(pm PackageManager, d Distro) (PackageManager, error) {
	if _, err := exec.LookPath(string(pm)); err == nil {
		return pm, nil
	} else {
		// TODO: maybe just log these as warnings
		return "", fmt.Errorf(
			"Failed to find (%s) for (%s) (%w)", string(pm), d, err,
		)
	}
}

func DetectPrimaryManager(distro Distro) (PackageManager, error) {
	switch distro {
	case Arch:
		return checkPresent(Pacman, distro)
	case NixOS:
		return checkPresent(Nix, distro)
	}

	return "", errors.New("Failed to detect primary package manager.")
	// TODO: is this even an issue? just go through all the existing ones
}

var secondaryManagers = []struct {
	name string
	bin  string
}{
	{"flatpak", "flatpak"},
	{"snap", "snap"},
	{"pip", "pip3"},
	{"cargo", "cargo"},
	{"npm", "npm"},
	{"brew", "brew"},
}

func DetectSecondaryManagers() []string {
	var found []string
	for _, m := range secondaryManagers {
		if _, err := exec.LookPath(m.bin); err == nil {
			found = append(found, m.name)
		}
	}
	return found
}
