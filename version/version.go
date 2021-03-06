package version

const Maj = "0"
const Min = "1"
const Fix = "4"

var (
	// The full version string
	Version = "0.1.4"

	// GitCommit is set with --ldflags "-X main.gitCommit=$(git rev-parse HEAD)"
	GitCommit string
)

func init() {
	if GitCommit != "" {
		Version += "-" + GitCommit[:8]
	}
}
