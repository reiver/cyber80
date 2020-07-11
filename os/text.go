package main

const (
	textWidth  = 32
	textHeight = 32
)

// Text represents a ‘textWidth’ × ‘textHeight’ = 32×32 matrix of characters.
//
// Conceptually, it is a left-handed coordinate system with (x,y)=(0,0) in the top-left corner,
// and (x,y)=(31,31) in the bottom-right corner.
//
// Text[0] corresponds to (x,y)=(0,0).
// Text[1] corresponds to (x,y)=(1,0).
// Text[2] corresponds to (x,y)=(2,0).
// ...
// Text[29] corresponds to (x,y)=(29,0).
// Text[30] corresponds to (x,y)=(30,0).
// Text[31] corresponds to (x,y)=(31,0).
// Text[32] corresponds to (x,y)=(0,1).
// Text[33] corresponds to (x,y)=(0,2).
// Text[34] corresponds to (x,y)=(0,3).
// ...
// Text[1021] corresponds to (x,y)=(29,31).
// Text[1022] corresponds to (x,y)=(30,31).
// Text[1023] corresponds to (x,y)=(31,31).
//
// Or, in other words:
//
// (x,y) corresponds to Text[y*32 + x]
type Text [textWidth*textHeight]rune

var text Text

func init() {
	text[0]  = 'W'
	text[1]  = 'e'
	text[2]  = 'l'
	text[3]  = 'c'
	text[4]  = 'o'
	text[5]  = 'm'
	text[6]  = 'e'
	text[7]  = ' '
	text[8]  = 't'
	text[9]  = 'o'
	text[10] = ' '
	text[11] = 't'
	text[12] = 'h'
	text[13] = 'e'
	text[14] = ' '
	text[15] = 'c'
	text[16] = 'y'
	text[17] = 'b'
	text[18] = 'e'
	text[19] = 'r'
	text[20] = '8'
	text[21] = '0'
	text[22] = '!'

	text[2*textWidth+ 0] = 'm'
	text[2*textWidth+ 1] = 'a'
	text[2*textWidth+ 2] = 'g'
	text[2*textWidth+ 3] = 'i'
	text[2*textWidth+ 4] = 'c'
	text[2*textWidth+ 5] = '2'
	text[2*textWidth+ 6] = '5'
	text[2*textWidth+ 7] = '6'
	text[2*textWidth+ 8] = ' '
	text[2*textWidth+ 9] = 'v'
	text[2*textWidth+10] = '0'
	text[2*textWidth+11] = '.'
	text[2*textWidth+12] = '0'
	text[2*textWidth+13] = '.'
	text[2*textWidth+14] = '1'
	text[2*textWidth+15] = '.'
	text[2*textWidth+16] = '0'
	text[2*textWidth+17] = '0'
	text[2*textWidth+18] = '0'
	text[2*textWidth+19] = '0'

	text[4*textWidth+ 0] = 'P'
	text[4*textWidth+ 1] = 'o'
	text[4*textWidth+ 2] = 'w'
	text[4*textWidth+ 3] = 'e'
	text[4*textWidth+ 4] = 'r'
	text[4*textWidth+ 5] = 'e'
	text[4*textWidth+ 6] = 'd'
	text[4*textWidth+ 7] = ' '
	text[4*textWidth+ 8] = 'b'
	text[4*textWidth+ 9] = 'y'
	text[4*textWidth+10] = ' '
	text[4*textWidth+11] = 'd'
	text[4*textWidth+12] = 'o'
	text[4*textWidth+13] = 't'
	text[4*textWidth+14] = '-'
	text[4*textWidth+15] = 'm'
	text[4*textWidth+16] = 'a'
	text[4*textWidth+17] = 't'
	text[4*textWidth+18] = 'r'
	text[4*textWidth+19] = 'i'
	text[4*textWidth+20] = 'x'
	text[4*textWidth+21] = ' '
	text[4*textWidth+22] = 'g'
	text[4*textWidth+23] = 'r'
	text[4*textWidth+24] = 'a'
	text[4*textWidth+25] = 'p'
	text[4*textWidth+26] = 'h'
	text[4*textWidth+27] = 'i'
	text[4*textWidth+28] = 'c'
	text[4*textWidth+29] = 's'
	text[4*textWidth+30] = '.'

	text[6*textWidth+ 0] = 'H'
	text[6*textWidth+ 1] = 'e'
	text[6*textWidth+ 2] = 'l'
	text[6*textWidth+ 3] = 'l'
	text[6*textWidth+ 4] = 'o'
	text[6*textWidth+ 5] = ' '
	text[6*textWidth+ 6] = 'w'
	text[6*textWidth+ 7] = 'o'
	text[6*textWidth+ 8] = 'r'
	text[6*textWidth+ 9] = 'l'
	text[6*textWidth+10] = 'd'
	text[6*textWidth+11] = '!'

	text[8*textWidth+0] = 'R'
	text[8*textWidth+1] = 'E'
	text[8*textWidth+2] = 'A'
	text[8*textWidth+3] = 'D'
	text[8*textWidth+4] = 'Y'
}
