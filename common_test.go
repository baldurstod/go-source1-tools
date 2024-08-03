package source1_test

import (
	"log"

	"github.com/baldurstod/go-source1-tools/repository"
	"github.com/baldurstod/go-source1-tools/vpk"
)

const varFolder = "./var/"

var _ = func() bool {
	repository.AddRepository("tf2",
		vpk.NewVpkFS(
			"S:\\Program Files\\Steam\\steamapps\\common\\Team Fortress 2\\tf\\tf2_misc_dir.vpk",
			"S:\\Program Files\\Steam\\steamapps\\common\\Team Fortress 2\\tf\\tf2_textures_dir.vpk",
		))
	return true
}()

var _ = func() bool { log.SetFlags(log.LstdFlags | log.Lshortfile); return true }()
