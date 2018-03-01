all:
	@(cd orgexamplemore && go generate)
	@go build -o service
.PHONY: all

update:
	@go get -u github.com/varlink/go-varlink
	@go get -u github.com/varlink/go-varlink/varlink-generator
.PHONY: update

clean:
	rm -f service
.PHONY: clean
