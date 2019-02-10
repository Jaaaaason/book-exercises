/* undefine S_ISLNK first */
#ifdef S_ISLNK
#undef S_ISLNK
#endif

#include "ourhdr.h"

int
main(int argc, char *argv[])
{
#ifdef S_ISLNK
	printf("S_ISLNK exists\n");
	if(S_ISLNK(S_IFLNK))
		printf("S_ISLNK works\n");
	else
		printf("S_ISLNK not works\n");
#endif

	return 0;
}
