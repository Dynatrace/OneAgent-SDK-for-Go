package native

import (
	"bufio"
	"strings"
)

// GetEnrichmentMetadata returns metadata that can be used to manually enrich
// log messages when unsupported logging frameworks are used. This function is
// like OneAgentSDK.GetEnrichmentMetadata, except that it does not require the
// Go code module be injected (but instead requires use of cgo on Linux only).
//
// See also https://www.dynatrace.com/support/help/shortlink/enrich-metrics
//
// Since SDK version 1.1.0 - OneAgent version 1.245
func GetEnrichmentMetadata() (map[string]string, error) {
	magicFile, err := openEnrichmentMetadataFile()
	if err != nil {
		return nil, err
	}
	defer func() { _ = magicFile.Close() }()

	result := make(map[string]string)
	lines := bufio.NewScanner(magicFile)
	for lines.Scan() {
		kv := strings.SplitN(lines.Text(), "=", 2)
		if len(kv) == 2 && kv[0] != "" {
			result[kv[0]] = kv[1]
		}
	}
	if lines.Err() != nil {
		return nil, lines.Err()
	}
	return result, nil
}
