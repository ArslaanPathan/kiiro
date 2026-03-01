// Copyright (c) 2026 Arslaan Pathan
// This software is licensed under the ARPL. See LICENSE for details.

package git

import "os/exec"

func Clone(url string, dest string) error {
	return exec.Command("git", "clone", url, dest).Run()
}
