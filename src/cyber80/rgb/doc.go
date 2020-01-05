/*
Package cyber80_rgb provides tools working with 3-byte — i.e., 24-bit — RGB colors.

Creation

One way of creating a cyber80_rgb.Type is:

	red   :=  41 // == 0x29
	green := 173 // == 0xad
	blue  := 255 // == 0xff
	
	rgb := cyber80_rgb.RGB(red, green, blue)

Another way of creatig a cyber80_rgb.Type is:

	var colors []byte = []bytes{
		0x29,0xad,0xff, // blue
		0x83,0x76,0x9c, // indigo
		0xff,0x77,0xab, // pink
		0xff,0xcc,0xaa, // peach
	}
	
	p := buffer[3]
	
	rgb := cyber80_rgb.Bytes(p)

Formatted Printing

cyber80_rgb.Type supports formatted printing via Go's built-in "fmt" package.

%x will output a lower-case hexidecimal value. For example:

	89abcd

%X will output an upper-case hexidecimal value.

	89ABCD

%c will output an RGB value

	rgb(137,171,205)

Web Hexadecimal Color Code

If one wanted to create a hexadecimal color code suitable for using for the Web (in HTML, CSS, or JavaScript) — i.e.,
the type that begins with a “#” symbol — then something similar to the following could be done:

	fmt.Sprintf("#%x") // Note te “#” symbol before the “%x” verb.

*/
package cyber80_rgb
