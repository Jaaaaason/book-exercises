CC = gcc
CFLAG = -g -Wall
INCLUDES = -I../include

all: a.out

a.out: exercise_4.1.o error.o
	$(CC) $(CFLAG) -o a.out exercise_4.1.o error.o

exercise_4.1.o: exercise_4.1.c
	$(CC) $(CFLAG) $(INCLUDES) -c exercise_4.1.c

error.o: ../include/error.c
	$(CC) $(CFLAG) $(INCLUDES) -c ../include/error.c

clean:
	rm *.o a.out
