We indicated in Figure 4.1 that SVR4 doesn't currently provide the S_ISLNK marco. But SVR4 does support symbolic links and defines S_IFLNK in <sys/stat.h>. (Perhaps someone forgot to define S_ISLNK?) Devise a way around this omission that can be placed in ourhdr.h, so any programd that needs the S_ISLNK marco can use it.

查看 <sys/stat.h> 文件得 S_ISLNK 的定义为：

`#define	S_ISLNK(m)	(((m) & S_IFMT) == S_IFLNK)	/* symbolic link */`

copy 进 ourhdr.h 即可