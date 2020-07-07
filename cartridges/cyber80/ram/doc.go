/*
Package cyber80_ram provides the RAM for a cyber80 fantasy computer.

In particular the cyber80's RAM is the variable:

	cyber80_ram.Value

(cyber80_ram.Value is of type cyber80_ram.Type)

Getting And Setting

Normal Go slice operations can be performed cyber80_ram.Value.

For example to set value at memory address ‘1015’ to the value ‘0x0c’ one could do something similar to:

	cyber80_ram.Value[1015] = 0x0c

And, for example, to get the value at memort ‘1015’ one could do something similar to:

	var value byte = cyber80_ram.Value[1015]

Size

The total size of cyber80_ram.Value (and all cyber80_ram.Type), in bytes, is given by:

	cyber80_ram.Size

Raster Display Memory

The raster display memory starts at memory address:

	cyber80_ram.BPRaster

The last byte of raster display memory is at memory address:

	cyber80_ram.BPRaster + cyber80_raster.Size - 1

To write at the 5th memory location of the raster display memory one could do something similar to:

	index := cyber80_ram.BPRaster + 5 - 1
	
	cyber80_ram.Value[index] = 0x0b

Although, since most people want to think is terms of (x,y) coordinates when dealing with pixels the raster display memory,
one could use the cyber80_raster.Offset() helper. For example:

	x := 8
	y : 16
	
	index := cyber80_ram.BPRaster + cyber80_raster.Offset(x, y)
	
	cyber80_ram.Value[index] = 0x0b
*/
package cyber80_ram
