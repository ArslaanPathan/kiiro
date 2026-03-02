// Copyright (c) 2026 Arslaan Pathan
// This software is licensed under the ARPL. See LICENSE for details.

package sources

import (
	"fmt"
	"encoding/json"
	"net/http"
)

type PackageListing struct {
	Repository string `json:"repo"`
	KiirofileOverride *string `json:"kiirofile_override"`
	KiirotomlOverride *string `json:"kiirotoml_override"`
}

func FetchPackageListing(url string) (*PackageListing, error) {
	fmt.Println("Fetching package listing...")

	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	body, err := io.ReadAll(res.Body)
	res.Body.Close()

	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d", res.StatusCode)
		return nil, err
	}
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	var packageListing PackageListing
	err := json.Unmarshal(body, &packageListing)
	if err != nil {
		return nil, err
	}
	fmt.Println("Fetched package listing!")
	return &packageListing, nil
}
