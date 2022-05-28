package core

import "os/exec"

func Command(cmd string) (string, error) {
	c := exec.Command("bash", "-c", cmd)
	output, err := c.CombinedOutput()
	return string(output), err
}
