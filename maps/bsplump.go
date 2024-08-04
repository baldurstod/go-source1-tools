package maps

type BspLumpData interface {
}

type bspLump struct {
	lumpType int
	offset   uint32
	len      uint32
	data     BspLumpData
}

func newBspLump(t int) *bspLump {
	var data BspLumpData
	switch t {
	case LUMP_ENTITIES:
		data = bspLumpEntities{}
	case LUMP_PLANES:
		data = bspLumpPlanes{}
	default:
		//panic("unknonw lump type " + strconv.Itoa(t))
	}

	return &bspLump{
		lumpType: t,
		data:     data,
	}
}

type bspLumpEntities struct {
	text string
	//BspLumpBase
}

type bspLumpPlanes struct {
	//BspLumpBase
}
