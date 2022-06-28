all: sqstack

sqstack: main.o sqstack.o
	$(CC) $^ -o $@

clean:
	rm -rf sqstack *.o