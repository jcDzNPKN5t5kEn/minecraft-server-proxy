package Config

type Config struct {
	Remote        string
	Local         string
	Proxy         string
	WebUIPort     string
	OverwriteHost string
	OverwritePort int
}

var (
	CurConfig Config
)
