package cyber80_ram

var (
	Value Type
)

func init() {
	// PALETTE
	{
		bp := BPPalette

		//                                         R    G    B
		Value[bp+ 0], Value[bp+ 1], Value[bp+ 2] = 0x00,0x00,0x00 //  0 — black
		Value[bp+ 3], Value[bp+ 4], Value[bp+ 5] = 0x1d,0x2b,0x53 //  1 — dark-blue
		Value[bp+ 6], Value[bp+ 7], Value[bp+ 8] = 0x7e,0x25,0x53 //  2 — dark-purple
		Value[bp+ 9], Value[bp+10], Value[bp+11] = 0x00,0x87,0x51 //  3 — dark-green

		Value[bp+12], Value[bp+13], Value[bp+14] = 0xab,0x52,0x36 //  4 — brown
		Value[bp+15], Value[bp+16], Value[bp+17] = 0x5f,0x57,0x4f //  5 — dark-gray
		Value[bp+18], Value[bp+19], Value[bp+20] = 0xc2,0xc3,0xc7 //  6 — light-gray
		Value[bp+21], Value[bp+22], Value[bp+23] = 0xff,0xf1,0xe8 //  7 — white

		Value[bp+24], Value[bp+25], Value[bp+26] = 0xff,0x00,0x4d //  8 — red
		Value[bp+27], Value[bp+28], Value[bp+29] = 0xff,0xa3,0x00 //  9 — orange
		Value[bp+30], Value[bp+31], Value[bp+32] = 0xff,0xec,0x27 // 10 — yellow
		Value[bp+33], Value[bp+34], Value[bp+35] = 0x00,0xe4,0x36 // 11 — green

		Value[bp+36], Value[bp+37], Value[bp+38] = 0x29,0xad,0xff // 12 — blue
		Value[bp+39], Value[bp+40], Value[bp+41] = 0x83,0x76,0x9c // 13 — indigo
		Value[bp+42], Value[bp+43], Value[bp+44] = 0xff,0x77,0xab // 14 — pink
		Value[bp+45], Value[bp+46], Value[bp+47] = 0xff,0xcc,0xaa // 15 — peach
	}
}
