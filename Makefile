SOURCES=server.go

all:
	go build $(SOURCES)
	./server

build:
	go build $(SOURCES)

clean:
	rm server
