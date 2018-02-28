all:
	@go generate github.com/varlink/go-varlink
	@go build -o service
.PHONY: all

clean:
	rm -f service
.PHONY: clean
