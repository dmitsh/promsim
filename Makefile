all: promsim

promsim:
	CGO_ENABLED=0 go build -a -ldflags '-extldflags "-static"' ./cmd/promsim/

clean:
	rm -f ./promsim

.PHONY: promsim clean
