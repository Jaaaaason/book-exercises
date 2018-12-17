Fill in the table below showing the effects of the different shift operations on single- byte quantities. The best way to think about shift operations is to work with binary representations. Convert the initial values to binary, perform the shifts, and then convert back to hexadecimal. Each of the answers should be 8 binary digits or 2 hexadecimal digits.

<pre>
                                       (Logical)      (Arithmetic)
     x                 x << 3           x >> 2           x >> 2
-------------     -------------    -------------    -------------
Hex    Binary     Binary    Hex    Binary    Hex    Binary    Hex
-----------------------------------------------------------------
0xC3   11000011   00011000  0x18   00110000  0x30   11110000  0xF0   
0x75   01110101   10101000  0xA8   00011101  0x1D   00011101  0x1D
0x87   10000111   00111000  0x38   00100001  0x21   11100001  0xE1
0x66   01100110   00110000  0x30   00011001  0x19   00011001  0x19
</pre>

