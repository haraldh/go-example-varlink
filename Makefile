all:
	@go get -u github.com/varlink/go-varlink
	@go get -u github.com/varlink/go-varlink/varlink-generator
	@go generate
	@go build -o service
.PHONY: all

clean:
	rm -f service
.PHONY: clean
