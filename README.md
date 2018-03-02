```bash
$ git get -u github.com/haraldh/go-varlink-example
$ cd $GOPATH/src/github.com/haraldh/go-varlink-example
$ make update all
$  varlink call exec:./service/org.example.more.Ping '{"ping" : "test" }' 
{
  "pong": "test"
}
```
