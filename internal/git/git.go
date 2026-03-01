package git

import "os/exec"

func Clone(url string, dest string) error {
	return exec.Command("git", "clone", url, dest).Run()
}
