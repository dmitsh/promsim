DOMAIN=promsim

all: bin

bin:
	CGO_ENABLED=0 go build -a -ldflags '-extldflags "-static"' ./cmd/promsim/

cert: $(DOMAIN).key $(DOMAIN).crt

$(DOMAIN).key:
	openssl genrsa -out $(DOMAIN).key 2048

$(DOMAIN).crt:
	openssl req -new -x509 -sha256 -key $(DOMAIN).key -out $(DOMAIN).crt -subj "/C=US/ST=CA/L=SF/O=Sysdig/OU=Promsim/CN=none.com/emailAddress=none@none.com"

clean:
	rm -f ./promsim $(DOMAIN).key $(DOMAIN).crt

.PHONY: bin cert clean
