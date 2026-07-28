package export_chzcake_metadata_metadata

import "wit_component/chzcake_metadata_metadata"

// Get implements chzcake:metadata/metadata.get@1.0.0.
func Get() chzcake_metadata_metadata.Info {
	return chzcake_metadata_metadata.Info{
		Name:        "test-plugin",
		Version:     "1.0.0",
		Description: "Test plugin for the chzcake metadata contract",
	}
}
