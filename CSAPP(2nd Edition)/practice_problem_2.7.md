### What would be printed as a result of the following call to show_bytes?

```c
const char *s = "abcdef";
show_bytes((byte_pointer) s, strlen(s));
```

show_bytes 定义如下：

```c
typedef unsigned char *byte_pointer; 4
void show_bytes(byte_pointer start, int len) {
	int i;
	for (i = 0; i < len; i++)
		printf(" %.2x", start[i]);
	printf("\n");
}
```

输出如下：

`  61 62 63 64 65 66`