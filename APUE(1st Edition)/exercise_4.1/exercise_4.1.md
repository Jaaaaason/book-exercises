Modify Program 4.1 to use stat instead of lstat. What changes if one of the command-line arguments is symbolic link?

先看看书中对 stat 和 lstat 的描述：

> stat: the **stat** function returns a structure of information about the named file.

> lstat: the **lstat** function is similar to **stat**, but when the named file is symbolic link, lstat returns information about the symbolic link, not the file referenced by the symbolic link.

当应用于软链接时，**stat** 返回的是 **原文件** 的信息，若软链接指向的为 ”regular file“，则输出为 ”regular file“，若指向的为 “directory”，则输出为 “directory”，以此类推，永远不会输出 ”symbolic link“。

