### Using show_int and show_float, we determine that the integer 3510593 has hexadecimal representation 0x00359141, while the floating-point number 3510593.0 has hexadecimal representation 0x4A564504.

**A. Write the binary representations of these two hexadecimal values.**

0x00359141: 00000000001101011001000101000001

0x4A564504: 01001010010101100100010100000100

**B. Shift these two strings relative to one another to maximize the number of
matching bits. How many bits match? **

<pre>
00000000001101011001000101000001
  01001010010101100100010100000100
</pre>

总共有 21 位匹配

**C. What parts of the strings do not match? **

0x4A564504 二进制串的前 9 位和后 2 位

