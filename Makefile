all:
	@go generate
	@go build -o service
.PHONY: all

clean:
	rm -f service
.PHONY: clean
