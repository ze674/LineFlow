package scanner

type Scanner interface {
	Scan() (string, error)
}
