### In the output from Program 1.4, what happened to the processes with process IDs 852 and 853.

```c
// Programm 1.4 Print the process ID
#include "ourhdr.h"

int
main(void)
{
    printf("hello world from process ID %d\n", getpid());
    exit(0);
}
```

如果运行这个程序，会产生如下输出：

```bash
$ a.out
hello world from process ID 851
$ a.out
hello world from process ID 854
```

PID 为 852 和 853 的进程已经结束了