// Copyright (c) 2026 Arslaan Pathan
// This software is licensed under the ARPL. See LICENSE for details.

package repo

type PackageListing struct {
	Repository string `json:"repo"`
	KiirofileOverride *string `json:"kiirofile_override"`
	KiirotomlOverride *string `json:"kiirotoml_override"`
}
