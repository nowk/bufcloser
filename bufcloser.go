package bufcloser

import "bufio"
import "io"

// closer wraps a bufio.Reader and a io.Closer
type closer struct {
	*bufio.Reader
	io.Closer
}

// New returns new closer
func New(c io.ReadCloser) *closer {
	r := bufio.NewReader(c)

	return &closer{
		Reader: r,
		Closer: c,
	}
}
