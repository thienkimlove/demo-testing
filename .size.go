package size

func Size(a int) string {
	GoCover_0.Count[0] = 1
	switch {
	case a < 0:
		GoCover_0.Count[2] = 1
		return "negative"
	case a == 0:
		GoCover_0.Count[3] = 1
		return "zero"
	case a < 10:
		GoCover_0.Count[4] = 1
		return "small"
	}
	GoCover_0.Count[1] = 1
	return "enormous"
}

var GoCover_0 = struct {
	Count   [5]uint32
	Pos     [3 * 5]uint32
	NumStmt [5]uint16
}{
	Pos: [3 * 5]uint32{
		3, 4, 0x9001a, // [0]
		12, 12, 0x130002, // [1]
		5, 6, 0x14000d, // [2]
		7, 8, 0x10000e, // [3]
		9, 10, 0x11000e, // [4]
	},
	NumStmt: [5]uint16{
		1, // 0
		1, // 1
		1, // 2
		1, // 3
		1, // 4
	},
}
