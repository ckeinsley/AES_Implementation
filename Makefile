BINARY_NAME = aes_encrypt

all: build build-linux build-windows build-windows_64

build:
	go get -t -v
	go build -o $(BINARY_NAME) -v

run:
	go get -t -v
	go build -o $(BINARY_NAME) -v ./...
	./$(BINARY_NAME)

build-linux:
	go get -t -v
	GOOSCGO_ENABLED=0 GOOS=linux GOARCH=386 go build -a -installsuffix cgo -o $(BINARY_NAME)_linux -v

build-windows:
	go get -t -v
	CGO_ENABLED=0 GOOS=windows GOARCH=386 go build -a -installsuffix cgo -o $(BINARY_NAME)_windows_386.exe -v

build-windows_64:
	go get -t -v
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -a -installsuffix cgo -o $(BINARY_NAME)_windows_64.exe -v

clean:
	rm $(BINARY_NAME)*