package main

import (
	"fmt"
	"os/exec"
)


func main() {
    pm, err := checkSysPackageManager()
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    fmt.Println("Detected package manager:", pm)

	// TODO: Handle anything other then apt
}

func checkSysPackageManager() (string, error) {
	// Check what package manager the system uses
	// Check for the presence of various package managers
    packageManagers := []string{"apt", "dnf", "yum", "zypper", "pacman"}

    for _, pm := range packageManagers {
        _, err := exec.LookPath(pm)
        if err == nil {
            return pm, nil
        }
    }

    // If no package manager is found, return an error
    return "", fmt.Errorf("no package manager found")
}

