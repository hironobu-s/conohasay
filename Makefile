NAME=conohasay
BINDIR=bin
GOARCH=amd64

all: clean darwin linux

assets:
	go-assets-builder -s /cows/ cows > assets.go

darwin:
	GOOS=$@ GOARCH=$(GOARCH) CGO_ENABLED=0 go build $(GOFLAGS) -o $(BINDIR)/$@/$(NAME)
	cd bin/$@; gzip -c $(NAME) > $(NAME)-osx.$(GOARCH).gz

linux:
	GOOS=$@ GOARCH=$(GOARCH) CGO_ENABLED=0 go build $(GOFLAGS) -o $(BINDIR)/$@/$(NAME)
	cd bin/$@; gzip -c $(NAME) > $(NAME)-linux.$(GOARCH).gz

test:
	$(BUILD_ASSETS)
	go test

clean:
	rm -rf $(BINDIR)
