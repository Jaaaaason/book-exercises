Examine your system's headers and list the actual data types used to implement the primitive system data types.

我现在所使用的系统是 Mac OS Mojave，对应的 primitive system data types 如下：

```c
/* caddr_t
 * /usr/include/sys/_types/_caddr_t.h
 */
30	typedef	char *	caddr_t;

/* clock_t
 * /usr/include/sys/_types/_clock_t.h
 */
31	typedef __darwin_clock_t	clock_t;
/* /usr/include/i386/_types.h */
117	typedef unsigned long	__darwin_clock_t;

/* comp_t
 * /usr/include/sys/acct.h
 */
82	typedef u_int16_t	comp_t;
/* /usr/include/sys/_types/_u_int16_t.h */
30	typedef	unsigned short	u_int16_t;

/* dev_t
 * /usr/include/sys/_types/_dev_t.h
 */
31	typedef __darwin_dev_t	dev_t;
/* /usr/include/sys/_types.h */
57	typedef __int32_t	__darwin_dev_t;
/* /usr/include/i386/_types.h */
44	typedef int	__int32_t;

/* fd_set
 * /usr/include/sys/_types/_fd_def.h
 */
49	typedef	struct fd_set {
50		__int32_t	fds_bits[__DARWIN_howmany(__DARWIN_FD_SETSIZE, __DARWIN_NFDBITS)];
51	} fd_set;

/* fpos_t
 * /usr/include/_stdio.h
 */
81	typedef __darwin_off_t	fpos_t;
/* /usr/include/sys/_types.h */
71	typedef __int64_t	__darwin_off_t;
/* /usr/include/i386/_types.h */
46	typedef long long	__int64_t;

/* gid_t
 * /usr/include/sys/_types/_gid_t.h
 */
31	typedef __darwin_gid_t	gid_t;
/* /usr/include/sys/_types.h */
60	typedef __uint32_t	__darwin_gid_t;
/* /usr/include/i386/_types.h */
45	typedef unsigned int	__uint32_t;

/* ino_t
 * /usr/include/sys/_types/_ino_t.h
 */
31	typedef	__darwin_ino_t		ino_t;
/* /usr/include/sys/_types.h */
62	typedef __uint64_t	__darwin_ino64_t;
63	#if __DARWIN_64_BIT_INO_T
64	typedef __darwin_ino64_t __darwin_ino_t;
65	#else
66	typedef __uint32_t	__darwin_ino_t;
67	#endif
/* /usr/include/i386/_types.h */
45	typedef unsigned int		__uint32_t;
47	typedef unsigned long long	__uint64_t;

/* mode_t
 * /usr/include/sys/_types/_mode_t.h
 */
31	typedef	__darwin_mode_t	mode_t;
/* /usr/include/sys/_types.h */
70	typedef __uint16_t	__darwin_mode_t;
/* /usr/include/i386/_types.h */
43	typedef	unsigned short	__uint16_t;

/* nlink_t
 * /usr/include/sys/_types/_nlink_t.h
 */
31	typedef __uint16_t        nlink_t;
/* /usr/include/i386/_types.h */
43	typedef	unsigned short	__uint16_t;

/* off_t
 * /usr/include/sys/_types/_off_t.h
 */
31	typedef __darwin_off_t	off_t;
/* /usr/include/sys/_types.h */
71	typedef __int64_t	__darwin_off_t;
/* /usr/include/i386/_types.h */
46	typedef long long	__int64_t;

/* pid_t
 * /usr/include/sys/_types/_pid_t.h
 */
31	typedef __darwin_pid_t	pid_t;
/* /usr/include/sys/_types.h */
72	typedef __int32_t	__darwin_pid_t;
/* /usr/include/i386/_types.h */
44	typedef	int	__int32_t;

/* ptrdiff_t
 * /usr/include/sys/_types/_ptrdiff_t.h
 */
32 typedef __darwin_ptrdiff_t ptrdiff_t;
/* /usr/include/i386/_types.h */
83	#if defined(__PTRDIFF_TYPE__)
84	typedef __PTRDIFF_TYPE__	__darwin_ptrdiff_t;
85	#elif defined(__LP64__)
86	typedef long				__darwin_ptrdiff_t;
87	#else
88	typedef int					__darwin_ptrdiff_t;
89	#endif

/* rlim_t
 * /usr/include/sys/resource.h
 */
89	typedef __uint64_t	rlim_t;
/* /usr/include/i386/_types.h */
47	typedef unsigned long long	__uint64_t;

/* sig_atomic_t
 * /usr/include/i386/signal.h
 */
39	typedef int sig_atomic_t;

/* sigset_t
 * /usr/include/sys/_types/_sigset_t.h
 */
31	typedef __darwin_sigset_t	sigset_t;
/* /usr/include/sys/_types.h */
73	typedef __uint32_t	__darwin_sigset_t;
/* /usr/include/i386/_types.h */
45	typedef unsigned int		__uint32_t;

/* size_t
 * /usr/include/sys/_types/_size_t.h
 */
31	typedef __darwin_size_t	size_t;
/* /usr/include/i386/_types.h */
91	#if defined(__SIZE_TYPE__)
92	typedef __SIZE_TYPE__	__darwin_size_t;
93	#else
94	typedef unsigned long	__darwin_size_t;
95	#endif

/* ssize_t
 * /usr/include/sys/_types/_ssize_t.h
 */
31	typedef __darwin_ssize_t	ssize_t;
/* /usr/include/i386/_types.h */
119	typedef long	__darwin_ssize_t;

/* time_t
 * /usr/include/sys/_types/_time_t.h
 */
31	typedef __darwin_time_t	time_t;
/* /usr/include/i386/_types.h */
120	typedef long	__darwin_time_t;

/* uid_t
 * /usr/include/sys/_types/_uid_t.h
 */
31	typedef __darwin_uid_t	uid_t;
/* /usr/include/sys/_types.h */
75	typedef __uint32_t	__darwin_uid_t;
/* /usr/include/i386/_types.h */
45	typedef unsigned int		__uint32_t;

/* wchar_t
 * /usr/include/sys/_types/_wchar_t.h
 */
34	typedef __darwin_wchar_t wchar_t;
/* /usr/include/i386/_types.h */
70	typedef int					__darwin_ct_rune_t;
103	#if defined(__WCHAR_TYPE__)
104	typedef __WCHAR_TYPE__		__darwin_wchar_t;
105	#else
106	typedef __darwin_ct_rune_t	__darwin_wchar_t;
107	#endif

```

