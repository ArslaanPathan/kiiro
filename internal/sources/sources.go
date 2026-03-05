// Copyright (c) 2026 Arslaan Pathan
// This software is licensed under the ARPL. See LICENSE for details.

package sources

import (
	"fmt"
	"encoding/json"
	"net/http"
	"io"
)

type PackageListing struct {
	Repository string `json:"repo"`
	KiirofileOverride *string `json:"kiirofile_override"`
	KiirotomlOverride *string `json:"kiirotoml_override"`
}

func FetchPackageListing(url string) (*PackageListing, error) {
	fmt.Println("Fetching package listing...")

	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(res.Body)
	res.Body.Close()

	if res.StatusCode > 299 {
		return nil, err
	}
	if err != nil {
		return nil, err
	}

	packageListing, err := ParsePackageListing(body)
	if err != nil {
		return nil, err
	}
	fmt.Println("Fetched package listing!")
	return packageListing, nil
}

func ParsePackageListing(pkgListingJSON []byte) (*PackageListing, error) {
	var packageListing PackageListing
	err := json.Unmarshal(pkgListingJSON, &packageListing)
	if err != nil {
		return nil, err
	}
	return &packageListing, nil
}
