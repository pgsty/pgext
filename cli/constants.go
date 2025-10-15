/*
Copyright 2018-2025 Ruohang Feng <rh@vonng.com>
*/
package cli

// Constants for batch operations
const (
	// BatchSize defines the number of records to process in a single batch
	BatchSize = 1000

	// MaxFileSize defines the maximum allowed file size for downloads (500MB)
	MaxFileSize = 500 * 1024 * 1024

	// HTTPTimeout defines the default HTTP client timeout
	HTTPTimeout = 60

	// DefaultWorkers defines the default number of concurrent workers
	DefaultWorkers = 4
)