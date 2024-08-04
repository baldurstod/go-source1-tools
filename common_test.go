package source1_test

import (
	"log"

	"github.com/baldurstod/go-source1-tools/files"
	"github.com/baldurstod/go-source1-tools/repository"
)

const varFolder = "./var/"

var _ = func() bool {
	repository.AddRepository(
		repository.NewRepositoryFS("tf2",
			files.NewVpkFS(
				"S:\\Program Files\\Steam\\steamapps\\common\\Team Fortress 2\\tf\\tf2_misc_dir.vpk",
				"S:\\Program Files\\Steam\\steamapps\\common\\Team Fortress 2\\tf\\tf2_textures_dir.vpk",
			),
			files.NewFileFS("S:\\Program Files\\Steam\\steamapps\\common\\Team Fortress 2\\tf\\maps"),
		))
	return true
}()

var _ = func() bool { log.SetFlags(log.LstdFlags | log.Lshortfile); return true }()
