Write C expressions, in terms of variable x, for the following values. Your code should work for any word size w ≥ 8. For reference, we show the result of evalu- ating the expressions for x = 0x87654321, with w = 32. 

**A. The least significant byte of x, with all other bits set to 0. [0x00000021].**

```c
x = x & 255;
```

**B. All but the least significant byte of x complemented, with the least significant byte left unchanged. [0x789ABC21].**

```c
x = ~x ^ 255;
```

**C. The least significant byte set to all 1s, and all other bytes of x left unchanged. [0x876543FF]. **

```c
x = x | 255;
```

