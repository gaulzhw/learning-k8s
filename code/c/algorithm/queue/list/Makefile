all: queue

queue: main.o queue.o llist.o
	$(CC) $^ -o $@

clean:
	rm -rf queue *.o