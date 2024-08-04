package maps

type BspLumpHeader struct {
	lumpType int
	offset   uint32
	len      uint32
}

type BspLumpData interface {
}

type BspLump struct {
	offset uint32
	len    uint32
	data   BspLumpData
}

func newBspLump(t int) *BspLump {
	switch t {
	case LUMP_ENTITIES:
		return &BspLump{
			data: BspLumpEntities{},
		}
	case LUMP_PLANES:
		return &BspLump{
			data: BspLumpPlanes{},
		}
	default:
		//panic("unknonw lump type " + strconv.Itoa(t))
		return &BspLump{}

	}
	return nil
}

type BspLumpEntities struct {
	//BspLumpBase
}

type BspLumpPlanes struct {
	//BspLumpBase
}
