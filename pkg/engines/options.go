package engines

type Options struct {
	Domain   string
	Filetype string
	Key      string
	Command  string
	Block    []string
	Page     int
	Wildcard bool
}
