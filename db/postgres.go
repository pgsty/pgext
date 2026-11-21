package db

// DefaultPGVersion is the recommended major for generated configurations.
// It is deliberately independent of the newest supported catalog target.
const DefaultPGVersion = 18

const (
	MinPGVersion = 15
	MaxPGVersion = 19
)

// SupportedPGVersions returns catalog targets, newest first.
func SupportedPGVersions() []int {
	versions := make([]int, 0, MaxPGVersion-MinPGVersion+1)
	for pg := MaxPGVersion; pg >= MinPGVersion; pg-- {
		versions = append(versions, pg)
	}
	return versions
}

func SupportsPGVersion(pg int) bool {
	return pg >= MinPGVersion && pg <= MaxPGVersion
}
