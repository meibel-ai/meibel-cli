package version

// Set via ldflags at build time by GoReleaser.
var (
	Version   = "dev"
	Commit    = "none"
	BuildDate = "unknown"
)
