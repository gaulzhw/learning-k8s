CFLAGS+=-pthread
LDFLAGS+=-pthread

all: mysem

mysem: main.o mysem.o
	$(CC) $^ -o $@

clean:
	rm -rf mysem *.o