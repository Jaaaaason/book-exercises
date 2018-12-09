### In Section 1.7 the argument to `perror` is defined with the ANSI C attribute const, while the integer to `strerror` isn't defined with this attribute. Why?

下面是两个函数的声明:

```c
char *strerror(int errnum);
void perror(const char *msg);
```

很明显 `perror` 接收的是指针类型的参数，对指针 `msg` 指向的值的操作会影响到外部变量，也就是该函数产生了副作用，所以用 `const` 避免修改 `msg` 指向的值。

而 `strerror` 接收是 int 型的，属于值传递，`errnum` 的地址与外部变量不同，也就不用担心会对外部变量产生影响。