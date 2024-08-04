package maps

type BspMap struct {
	repository  string
	filename    string
	fileVersion int32
	mapRevision int32
	lumps       []*bspLump
}

func newBspMap() *BspMap {
	return &BspMap{
		//lumpHeaders: make([]*BspLumpHeader, 0, 64),
		lumps: make([]*bspLump, 0, 64),
	}
}

/*
func (m *BspMap) addLumpHeader(h *BspLumpHeader) {
	m.lumpHeaders = append(m.lumpHeaders, h)
}
*/

func (m *BspMap) addLump(l *bspLump) {
	m.lumps = append(m.lumps, l)
}
