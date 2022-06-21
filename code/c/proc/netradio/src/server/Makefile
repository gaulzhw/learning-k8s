CFLAGS+=-I../include/ -pthread -Wall
LDFLAGS+=-pthread

all: server

server: server.o thr_channel.o thr_list.o mytbf.o medialib.o
	gcc $^ -o $@ $(CFLAGS) $(LDFLAGS)

clean:
	rm -rf *.o server