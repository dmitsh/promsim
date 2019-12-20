all: promsim

promsim:
	go build ./cmd/promsim

clean:
	rm -f ./promsim

.PHONY: promsim clean
