all:
	@go get -u github.com/varlink/go-varlink/varlink-go-generator
	@go get -d -u github.com/varlink/go-varlink
	@go generate
	@go build -o service
.PHONY: all

clean:
	rm -f service
.PHONY: clean
