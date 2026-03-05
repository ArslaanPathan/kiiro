// Copyright (c) 2026 Arslaan Pathan
// This software is licensed under the ARPL. See LICENSE for details.

package manifest

import (
	"fmt"
	"encoding/json"
	"net/http"
	"log"
	"io"
	"errors"
)

type KiiroToml struct {
	Name string `toml:"pkgname"`
	Version string `toml:"version"`
	Architectures string `toml:"archs"`
	Authors []string `toml:"authors"`
	Maintainers []string `toml:"maintainers"`
	Depends []string `toml:"deps"`
}

func FetchKiiroToml(url string) (*KiiroToml, error) {
	fmt.Println("Fetching kiiro.toml...")
	return nil, errors.New("not implemented")
}
