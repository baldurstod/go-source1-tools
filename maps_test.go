package source1_test

import (
	"bytes"
	"io"
	"log"
	"os"
	"testing"

	"github.com/baldurstod/go-source1-tools/maps"
	"github.com/ulikunitz/xz/lzma"
)

func TestMap(t *testing.T) {

	m, err := maps.GetMap("tf2", "cp_hadal.bsp") //cp_hadal
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
	w, _ := lzma.NewWriter(&b)
	w.Write([]byte("hello, world\n"))
	w.Close()
	log.Println(b)

	// read that data back:
	//
	r, _ := lzma.NewReader(&b)
	io.Copy(os.Stdout, r)
	//r.Close()
}
