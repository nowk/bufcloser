package bufcloser

import "io"
import "os"
import "testing"
import "github.com/nowk/assert"

func TestReader(t *testing.T) {
	f, _ := os.Open("./sample/data.csv")
	r := New(f)
	defer r.Close()

	i := 0
	for {
		_, err := r.ReadBytes('\n')
		i++
		if err == io.EOF {
			assert.Equal(t, 3, i)
			break
		}

		if err != nil {
			t.Fatal(err)
		}
	}
}

func TestCloser(t *testing.T) {
	f, _ := os.Open("./sample/data.csv")
	r := New(f)
	r.Close()

	_, err := r.ReadBytes('\n')
	assert.Equal(t, "read ./sample/data.csv: bad file descriptor", err.Error())
}
