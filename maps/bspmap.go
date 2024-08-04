package maps

type BspMap struct {
	repository  string
	filename    string
	fileVersion int32
	mapRevision int32
	lumps       []*BspLump
}

func newBspMap() *BspMap {
	return &BspMap{
		//lumpHeaders: make([]*BspLumpHeader, 0, 64),
		lumps: make([]*BspLump, 0, 64),
	}
}

/*
func (m *BspMap) addLumpHeader(h *BspLumpHeader) {
	m.lumpHeaders = append(m.lumpHeaders, h)
}
*/

func (m *BspMap) addLump(l *BspLump) {
	m.lumps = append(m.lumps, l)
}
