package links

import (
	"fmt"
	"os"
	"testing"
)

func TestSimple(t *testing.T) {
	t.Run("simple", testfile("testfiles/simple.html"))
	t.Run("hard", testfile("testfiles/hard.html"))
}

func testfile(fname string) func(t *testing.T) {
	return func(t *testing.T) {
		fd, oerr := os.Open(fname)
		if oerr != nil {
			t.Fatal("where is my file?")
		}
		defer fd.Close()
		//xtr := NewSubstringXtractor()
		xtr := NewParserXtractor()
		lelinks, err := xtr.Xtract(fd)
		if err != nil {
			t.Fatal("failed to extract:%s", err)
		}
		for _, l := range lelinks {
			fmt.Printf("%s contains:%+v\n", fname, l)
		}
	}
}
