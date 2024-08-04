package source1_test

import (
	"bytes"
	"io"
	"log"
	"os"
	"testing"

	"github.com/baldurstod/go-source1-tools/maps"
	"github.com/luckcolors/lzma"
)

func TestMap(t *testing.T) {

	m, err := maps.GetMap("tf2", "itemtest.bsp") //cp_hadal
	if err != nil {
		t.Error(err)
		return
	}

	log.Println(m)
}

func TestLzma(t *testing.T) {
	/*

		m, err := maps.GetMap("tf2", "cp_hadal.bsp")
		if err != nil {
			t.Error(err)
			return
		}

		log.Println(m)
	*/
	var b bytes.Buffer
	w := lzma.NewWriter(&b)
	w.Write([]byte("hello, world\n"))
	w.Close()
	log.Println(b)

	// read that data back:
	//
	r := lzma.NewReader(&b)
	io.Copy(os.Stdout, r)
	r.Close()
}
