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

The size of cyber80_ram.Value (and all cyber80_ram.Type), in bytes, is given by:

	cyber80_ram.Size

*/
package cyber80_ram
