### Consider the following three calls to show_bytes:

Int val = 0x87654321;

byte_pointer valp = (byte_pointer) &val;

A. show_bytes(valp, 1);

B. show_bytes(valp, 2);

C. show_bytes(valp, 3);



```c
typedef unsigned char *byte_pointer;

void show_bytes(byte_pointer start, int len) {
    int i;
    for(i = 0; i < len; i++)
        printf(" %.2x", start[i]);
    printf("\n");
}
```

Int val = 0x87654321; // 此时 val 在内存中的存储如下，假设开始地址为 0x100

Little endian: 

![](https://i.imgur.com/2JRhweK.png)

Big endian:

![](https://i.imgur.com/lyPNEAd.png)





A. Little endian: 21 &nbsp;&nbsp;&nbsp;&nbsp; Big endian: 87

B. Little endian: 21 43 &nbsp;&nbsp;&nbsp;&nbsp; Big endian: 87 65

C. Little endian: 21 43 65 &nbsp;&nbsp;&nbsp;&nbsp; Big endian: 87 65 43

