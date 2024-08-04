package maps

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"log"
	"strings"

	"github.com/baldurstod/go-source1-tools/repository"
	"github.com/luckcolors/lzma"
)

const BSP_MAGIC = 0x50534256 //VBSP
const BSP_HEADER_LUMPS_COUNT = 64
const BYTES_PER_LUMP_HEADER = 16
const BYTES_PER_PLANE = 20

var mapCache = func() map[string]map[string]*BspMap { return make(map[string]map[string]*BspMap) }()

func findMap(repository string, filename string) *BspMap {
	r, ok := mapCache[repository]
	if !ok {
		return nil
	}
	m, ok := r[filename]
	if !ok {
		return nil
	}
	return m
}

func addMap(repository string, filename string, m *BspMap) {
	r, ok := mapCache[repository]
	if !ok {
		r = make(map[string]*BspMap)
		mapCache[repository] = r
	}

	r[filename] = m
}

func GetMap(repository string, filename string) (*BspMap, error) {
	filename = strings.TrimSuffix(filename, ".bsp")
	filename += ".bsp"

	if m := findMap(repository, filename); m != nil {
		return m, nil
	}

	m, err := loadBsp(repository, filename)
	if err != nil {
		return nil, err
	}
	/*
		m := maps.NewBspMap()
		m.SetFile(file)
	*/
	addMap(repository, filename, m)

	return m, nil
}

func loadBsp(repo string, filename string) (*BspMap, error) {
	reader := repository.GetRepository(repo)
	if reader == nil {
		return nil, errors.New("unknown repository")
	}

	b, err := reader.ReadFile(filename)
	if err != nil {
		return nil, errors.New("unable to read file " + filename)
	}

	loader := newBspLoader(repo, b)

	if err := loader.parseBsp(); err != nil {
		return nil, err
	}

	return loader.m, nil
}

type BspLoader struct {
	reader io.ReadSeeker
	b      []byte
	repo   string
	m      *BspMap
}

func newBspLoader(repo string, b []byte) *BspLoader {
	return &BspLoader{
		reader: bytes.NewReader(b),
		b:      b,
		repo:   repo,
		m:      newBspMap(),
	}
}

func (loader *BspLoader) parseBsp() error {
	if err := loader.parseHeader(); err != nil {
		return err
	}
	if err := loader.parseLumps(); err != nil {
		return err
	}
	return nil
}

func (loader *BspLoader) parseHeader() error {
	var magic uint32
	err := binary.Read(loader.reader, binary.LittleEndian, &magic)
	if err != nil {
		return fmt.Errorf("failed to read magic in parseHeader: <%w>", err)
	}

	if magic != BSP_MAGIC {
		return fmt.Errorf("wrong magic: <%d>", err)
	}

	if err = binary.Read(loader.reader, binary.LittleEndian, &loader.m.fileVersion); err != nil {
		return fmt.Errorf("failed to read magic in parseHeader: <%w>", err)
	}

	if err := loader.parseLumpDirectory(); err != nil {
		return err
	}

	if err = binary.Read(loader.reader, binary.LittleEndian, &loader.m.mapRevision); err != nil {
		return fmt.Errorf("failed to read magic in parseHeader: <%w>", err)
	}

	/*

		#parseHeader(reader, bsp) {
			reader.seek(4); //skip first 4 char TODO: check == 'VBSP' ?

			bsp.bspFileVersion = reader.getInt32();
			this._parseLumpDirectory(reader, bsp);
			bsp.mapRevision = reader.getInt32();
		}*/

	return nil
}

func (loader *BspLoader) parseLumpDirectory() error {
	var err error
	//startOffset, _ := loader.reader.Seek(0, io.SeekCurrent)
	for i := 0; i < BSP_HEADER_LUMPS_COUNT; i++ {
		/*if _, err = loader.reader.Seek(startOffset+int64(i*BYTES_PER_LUMP_HEADER), io.SeekStart); err != nil {
			return err
		}*/

		lump := newBspLump(i)

		if err = binary.Read(loader.reader, binary.LittleEndian, &lump.offset); err != nil {
			return fmt.Errorf("failed to read offset in parseLumpDirectory: <%w>", err)
		}
		if err = binary.Read(loader.reader, binary.LittleEndian, &lump.len); err != nil {
			return fmt.Errorf("failed to read len in parseLumpDirectory: <%w>", err)
		}

		loader.reader.Seek(8, io.SeekCurrent)

		loader.m.addLump(lump)
	}

	return nil
}

func (loader *BspLoader) parseLumps() error {
	for _, lump := range loader.m.lumps {
		if err := loader.parseLump(lump); err != nil {
			return err
		}
	}
	return nil
}

func (loader *BspLoader) parseLump(lump *bspLump) error {
	loader.reader.Seek(int64(lump.offset), io.SeekStart)

	var reader io.ReadSeeker
	var err error
	if reader, err = loader.initLZMALump(); err != nil {
		return err
	}

	switch lump.lumpType {
	case LUMP_ENTITIES:
		loader.parseLumpEntities(reader, lump)
	case LUMP_PLANES:
		loader.parseLumpPlanes(reader, lump)
	}

	return nil
}

func (loader *BspLoader) parseLumpEntities(reader io.ReadSeeker, lump *bspLump) error {
	buf := make([]byte, lump.len)
	if _, err := reader.Read(buf); err != nil {
		return err
	}
	lumpData := lump.data.(bspLumpEntities)
	lumpData.text = string(buf[:])

	return nil
}

func (loader *BspLoader) parseLumpPlanes(reader io.ReadSeeker, lump *bspLump) error {
	/*
		#parseLumpPlanes(reader, lump) {
			reader.seek(lump.lumpOffset);
			const BYTES_PER_PLANE = 20;
			const planesCount = lump.getLumpLen() / BYTES_PER_PLANE;
			const lumpData = [];
			for (let planeIndex = 0; planeIndex < planesCount; planeIndex++) {
				let plane = new SourceBSPLumpPlane();
				plane.normal = reader.getVector3();
				plane.dist = reader.getFloat32();
				plane.type = reader.getInt32();
				lumpData.push(plane);
			}
			lump.setLumpData(lumpData);
		}
	*/

	/*
		buf := make([]byte, lump.len)
		if _, err := reader.Read(buf); err != nil {
			return err
		}
		lumpData := lump.data.(bspLumpEntities)
		lumpData.text = string(buf[:])
	*/

	return nil
}

func (loader *BspLoader) initLZMALump() (io.ReadSeeker, error) {
	var lzmaMagic uint32
	var err error
	if err = binary.Read(loader.reader, binary.LittleEndian, &lzmaMagic); err != nil {
		return nil, fmt.Errorf("failed to read lzmaMagic in initLZMALump: <%w>", err)
	}

	if lzmaMagic == 0x414d5a4c { //LZMA
		var uncompressedSize uint32
		var compressedSize uint32
		var properties [5]byte

		if err = binary.Read(loader.reader, binary.LittleEndian, &uncompressedSize); err != nil {
			return nil, fmt.Errorf("failed to read uncompressedSize in initLZMALump: <%w>", err)
		}
		if err = binary.Read(loader.reader, binary.LittleEndian, &compressedSize); err != nil {
			return nil, fmt.Errorf("failed to read uncompressedSize in initLZMALump: <%w>", err)
		}
		/*
			if err = binary.Read(loader.reader, binary.LittleEndian, &properties); err != nil {
				return nil, fmt.Errorf("failed to read uncompressedSize in initLZMALump: <%w>", err)
			}*/

		log.Println(uncompressedSize, compressedSize, properties)
		var offset int64
		if offset, err = loader.reader.Seek(0, io.SeekCurrent); err != nil {
			return nil, err
		}

		/*var r io.ReadCloser
		if r, err = lzma.NewReader(bytes.NewReader(loader.b[offset:])); err != nil {
			return nil, err
		}*/
		r := lzma.NewReader(bytes.NewReader(loader.b[offset:]))
		var p [1000]byte
		if _, err := r.Read(p[:]); err != nil {
			return nil, err
		}
		log.Println(p)

		//bytes.NewReader(b)

		/*
			const uncompressedSize = reader.getUint32();
			const compressedSize = reader.getUint32();
			const properties = reader.getBytes(5);
			const compressedDatas = reader.getBytes(compressedSize);// 4 + 4 + 4 + 5

			reader = new BinaryReader(DecompressLZMA(properties, compressedDatas, uncompressedSize));

			lump.lumpOffset = 0;
			lump.lumpLen = uncompressedSize;
		*/

	}

	return loader.reader, nil
}

//func initLZMALump()
/*
function InitLZMALump(reader, lump) {
	if (reader.getString(4, lump.lumpOffset) === 'LZMA') {
		const uncompressedSize = reader.getUint32();
		const compressedSize = reader.getUint32();
		const properties = reader.getBytes(5);
		const compressedDatas = reader.getBytes(compressedSize);// 4 + 4 + 4 + 5

		reader = new BinaryReader(DecompressLZMA(properties, compressedDatas, uncompressedSize));

		lump.lumpOffset = 0;
		lump.lumpLen = uncompressedSize;
	}
	return reader;
}
*/
