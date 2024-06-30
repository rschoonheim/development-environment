package installation

import (
	"log/slog"
	"os/exec"
)

// InstallDocker - installs docker on the host
func InstallDocker() {

	if isDockerInstalled() {
		slog.Info("Docker is already installed.")
		return
	}

	slog.Info("Starting docker installation...")

	addGPGKey()
	installDocker()
}

// addGPGKey - adds Docker's official GPG key
func addGPGKey() {
	script := "# Add Docker's official GPG key:\nsudo apt-get update\nsudo apt-get install ca-certificates curl\nsudo install -m 0755 -d /etc/apt/keyrings\nsudo curl -fsSL https://download.docker.com/linux/ubuntu/gpg -o /etc/apt/keyrings/docker.asc\nsudo chmod a+r /etc/apt/keyrings/docker.asc\n\n# Add the repository to Apt sources:\necho \\\n  \"deb [arch=$(dpkg --print-architecture) signed-by=/etc/apt/keyrings/docker.asc] https://download.docker.com/linux/ubuntu \\\n  $(. /etc/os-release && echo \"$VERSION_CODENAME\") stable\" | \\\n  sudo tee /etc/apt/sources.list.d/docker.list > /dev/null\nsudo apt-get update"
	cmd := exec.Command("bash", "-c", script)
	err := cmd.Run()
	if err != nil {
		panic(err)
	}
}

// installDocker - installs Docker
func installDocker() {
	script := "sudo apt-get install -y docker-ce docker-ce-cli containerd.io docker-buildx-plugin docker-compose-plugin"
	cmd := exec.Command("bash", "-c", script)
	err := cmd.Run()
	if err != nil {
		panic(err)
	}
}

// isDockerInstalled - checks if docker is installed on the host
func isDockerInstalled() bool {
	cmd := exec.Command("docker", "--version")
	err := cmd.Run()
	if err != nil {
		return false
	}
	return true
}
