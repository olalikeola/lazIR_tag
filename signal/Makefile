# Makefile for Raspberry Pi Zero W - IR reading

# Compiler and flags
CC = gcc
CFLAGS = -O2 -Wall

# Target executable
TARGET = signal

# Source files
SRCS = src/main.c src/gpio.c src/timer.c

# Object files
OBJS = $(SRCS:.c=.o)

# Include directories
INCLUDES = -Iinclude

# Libraries
LIBS = -lwiringPi

# Build rules
all: $(TARGET)

$(TARGET): $(OBJS)
	$(CC) $(CFLAGS) $(OBJS) -o $(TARGET) $(LIBS)

%.o: %.c
	$(CC) $(CFLAGS) $(INCLUDES) -c $< -o $@

clean:
	rm -f $(OBJS) $(TARGET)

.PHONY: all clean
