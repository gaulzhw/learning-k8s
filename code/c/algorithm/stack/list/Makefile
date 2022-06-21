all: stack

stack: main.o stack.o llist.o
	$(CC) $^ -o $@

clean:
	rm -rf stack *.o