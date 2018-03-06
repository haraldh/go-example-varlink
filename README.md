```bash
$ go get -u github.com/haraldh/go-varlink-example
$ cd $GOPATH/src/github.com/haraldh/go-varlink-example
$ make update all
$  varlink call exec:./service/org.example.more.Ping '{"ping" : "test" }' 
{
  "pong": "test"
}
$ ./service unix:@test &
[1] 3126
$ varlink call unix:@test/org.example.more.Ping '{"ping" : "test" }' 
{
  "pong": "test"
}
$ varlink call unix:@test/org.example.more.StopServing
{}
[1]+  Done                    ./service unix:@test
$
```
