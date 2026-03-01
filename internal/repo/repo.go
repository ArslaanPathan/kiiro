package repo

type PackageListing struct {
	Repository string `json:"repo"`
	KiirofileOverride *string `json:"kiirofile_override"`
	KiirotomlOverride *string `json:"kiirotoml_override"`
}
