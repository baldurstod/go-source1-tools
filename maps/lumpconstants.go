package maps

const LUMP_ENTITIES = 0
const LUMP_PLANES = 1
const LUMP_TEXDATA = 2
const LUMP_VERTEXES = 3
const LUMP_VISIBILITY = 4
const LUMP_NODES = 5
const LUMP_TEXINFO = 6
const LUMP_FACES = 7
const LUMP_LIGHTING = 8
const LUMP_OCCLUSION = 9
const LUMP_LEAFS = 10
const LUMP_FACEIDS = 11
const LUMP_EDGES = 12
const LUMP_SURFEDGES = 13
const LUMP_MODELS = 14
const LUMP_WORLDLIGHTS = 15
const LUMP_LEAFFACES = 16
const LUMP_LEAFBRUSHES = 17
const LUMP_BRUSHES = 18
const LUMP_BRUSHSIDES = 19
const LUMP_AREAS = 20
const LUMP_AREAPORTALS = 21
const LUMP_UNUSED0 = 22
const LUMP_UNUSED1 = 23
const LUMP_UNUSED2 = 24
const LUMP_UNUSED3 = 25
const LUMP_DISPINFO = 26
const LUMP_ORIGINALFACES = 27
const LUMP_PHYSDISP = 28
const LUMP_PHYSCOLLIDE = 29
const LUMP_VERTNORMALS = 30
const LUMP_VERTNORMALINDICES = 31
const LUMP_DISP_LIGHTMAP_ALPHAS = 32
const LUMP_DISP_VERTS = 33                     // CDispVerts
const LUMP_DISP_LIGHTMAP_SAMPLE_POSITIONS = 34 // For each displacement
//		 For each lightmap sample
//				 byte for index
//				 if 255; then index = next byte + 255
//				 3 bytes for barycentric coordinates
// The game lump is a method of adding game-specific lumps
// FIXME: Eventually; all lumps could use the game lump system
const LUMP_GAME_LUMP = 35
const LUMP_LEAFWATERDATA = 36
const LUMP_PRIMITIVES = 37
const LUMP_PRIMVERTS = 38
const LUMP_PRIMINDICES = 39

// A pak file can be embedded in a .bsp now; and the file system will search the pak
//	file first for any referenced names; before deferring to the game directory
//	file system/pak files and finally the base directory file system/pak files.
const LUMP_PAKFILE = 40
const LUMP_CLIPPORTALVERTS = 41

// A map can have a number of cubemap entities in it which cause cubemap renders
// to be taken after running vrad.
const LUMP_CUBEMAPS = 42
const LUMP_TEXDATA_STRING_DATA = 43
const LUMP_TEXDATA_STRING_TABLE = 44
const LUMP_OVERLAYS = 45
const LUMP_LEAFMINDISTTOWATER = 46
const LUMP_FACE_MACRO_TEXTURE_INFO = 47
const LUMP_DISP_TRIS = 48
const LUMP_PHYSCOLLIDESURFACE = 49 // deprecated.	We no longer use win32-specific havok compression on terrain
const LUMP_WATEROVERLAYS = 50
const LUMP_LEAF_AMBIENT_INDEX_HDR = 51 // index of LUMP_LEAF_AMBIENT_LIGHTING_HDR
const LUMP_LEAF_AMBIENT_INDEX = 52     // index of LUMP_LEAF_AMBIENT_LIGHTING

// optional lumps for HDR
const LUMP_LIGHTING_HDR = 53
const LUMP_WORLDLIGHTS_HDR = 54
const LUMP_LEAF_AMBIENT_LIGHTING_HDR = 55 // NOTE: this data overrides part of the data stored in LUMP_LEAFS.
const LUMP_LEAF_AMBIENT_LIGHTING = 56     // NOTE: this data overrides part of the data stored in LUMP_LEAFS.

const LUMP_XZIPPAKFILE = 57   // deprecated. xbox 1: xzip version of pak file
const LUMP_FACES_HDR = 58     // HDR maps may have different face data.
const LUMP_MAP_FLAGS = 59     // extended level-wide flags. not present in all levels
const LUMP_OVERLAY_FADES = 60 // Fade distances for overlays
