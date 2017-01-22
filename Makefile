SRCS := $(wildcard *.go brain/*.go)

all: squidbot-brain-rpi squidbot-brain

squidbot-brain-rpi: $(SRCS)
	GOOS=linux GOARCH=arm go build -v
	mv squidbot-brain squidbot-brain-rpi

squidbot-brain: $(SRCS)
	go build

clean:
	-rm squidbot-brain squidbot-brain-rpi

mrclean: clean
	-find -name \*~ -exec rm -f '{}' ';'

.PHONY: all clean mrclean
