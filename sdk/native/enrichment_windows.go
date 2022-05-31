package native

import (
	"io/ioutil"
	"os"
	"strings"
)

const magicFileName = "dt_metadata_e617c525669e072eebe3d0f08212e8f2.properties"

// On Windows, os.Open (used by ioutil.ReadFile) uses CreateFileW, which is hooked by OneAgent.
func openEnrichmentMetadataFile() (*os.File, error) {
	data, err := ioutil.ReadFile(magicFileName)
	if err != nil {
		return nil, err
	}

	path := strings.TrimSuffix(string(data), "\n")
	return os.Open(path)
}
