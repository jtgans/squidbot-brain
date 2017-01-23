SRCS := $(wildcard *.go brain/*.go)

all: squidbot-brain-rpi squidbot-brain

squidbot-brain-rpi: $(SRCS)
	GOOS=linux GOARCH=arm go build -v
	mv squidbot-brain squidbot-brain-rpi

squidbot-brain: $(SRCS)
	go build

docker: squidbot-brain
	docker build . --force-rm --tag jtgans/squidbot-brain:latest

docker-rpi: squidbot-brain-rpi
	docker build . -f Dockerfile.rpi --force-rm --tag jtgans/squidbot-brain-rpi:latest

docker-all: docker docker-arm

clean:
	rm -f squidbot-brain squidbot-brain-rpi

mrclean: clean
	find -name \*~ -exec rm -f '{}' ';'

.PHONY: all clean mrclean docker docker-arm docker-all
