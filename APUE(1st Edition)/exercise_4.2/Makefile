CC = gcc
CFLAG = -g -Wall
INCLUDES = -I.

all: a.out

a.out: exercise_4.2.o
	$(CC) $(CFLAG) -o a.out exercise_4.2.o

exercise_4.2.o: exercise_4.2.c
	$(CC) $(CFLAG) $(INCLUDES) -c exercise_4.2.c

clean:
	rm *.o a.out
