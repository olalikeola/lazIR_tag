CC = gcc
CFLAGS = -Wall -Wextra

all: signal

signal: signal.o
	$(CC) $(CFLAGS) -o signal signal.o

signal.o: signal.c
	$(CC) $(CFLAGS) -c signal.c

clean:
	rm -f signal signal.o
