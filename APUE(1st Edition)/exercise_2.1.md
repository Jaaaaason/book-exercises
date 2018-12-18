we mentioned in Section 2.7 that some of the primitive system data types are defined in more than oe header. For example, size_t is defined in six different headers. Since program could #include all six of these different headers and since ANSI C does not allow multiple typedefs for the same name, how must be headers be written?

在头文件中必须用预处理指令检查别名是否已经定义，但由于预处理只能检查宏是否定义，所以需要设置一个宏作为标记，标记别名是否已经定义。比如 size_t 在头文件中应该这样定义：

```c
#ifndef __SIZE_T
#define __SIZE_T
typedef [type] size_t;
#endif
```

