package source1_test

import (
	"log"
	"testing"

	"github.com/baldurstod/go-source1-tools/repository"
)

func TestRepositories(t *testing.T) {
	repo := repository.GetRepository("tf2")
	if repo == nil {
		t.Error("repo not found")
	}

	filenames := [...]string{"models/player/demo.mdl", "materials/models/player/demo/demoman_red.vtf"}

	for _, filename := range filenames {
		buf, err := repo.ReadFile(filename)
		if err != nil {
			t.Error(err)
			return
		}
		log.Println(filename, buf[0:200])
	}
}
