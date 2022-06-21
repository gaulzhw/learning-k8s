all: queue

queue: main.o queue.o
	$(CC) $^ -o $@

clean:
	rm -rf queue *.o