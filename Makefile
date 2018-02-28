all:
	@( cd $$GOPATH/src/github.com/varlink/go-varlink; go generate )
	@go build
.PHONY: all

clean:
	rm -f go-example
.PHONY: clean
