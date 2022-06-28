CFLAGS+=-pthread
LDFLAGS+=-pthread

all: relayer

relayer: relayer.o main.o
	gcc $^ -o $@ $(CFLAGS) $(LDFLAGS)

clean:
	rm -rf *.o relayer