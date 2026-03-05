// Copyright (c) 2026 Arslaan Pathan
// This software is licensed under the ARPL. See LICENSE for details.

package main

import (
	"fmt"
	"github.com/ArslaanPathan/kiiro/internal/sources"
	"log"
)

func main() {
	fmt.Println("kiiro development test")
	fmt.Println("fetching test package listing")
	packageListing, err := sources.FetchPackageListing("http://localhost:8081/test.json")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(packageListing.Repository)
	if packageListing.KiirofileOverride != nil {
		fmt.Println(*packageListing.KiirofileOverride)
	}
	if packageListing.KiirotomlOverride != nil {
		fmt.Println(*packageListing.KiirotomlOverride)
	}
}
