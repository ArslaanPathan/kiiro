// Copyright (c) 2026 Arslaan Pathan
// This software is licensed under the ARPL. See LICENSE for details.

package manifest

type KiiroToml struct {
	Name string `toml:"pkgname"`
	Version string `toml:"version"`
	Architectures string `toml:"archs"`
	Authors []string `toml:"authors"`
	Maintainers []string `toml:"maintainers"`
	Depends []string `toml:"deps"`
}

