```bash
$ go get -u github.com/haraldh/go-varlink-example
$ cd $GOPATH/src/github.com/haraldh/go-varlink-example
$ make update all
$ varlink call exec:./service/org.example.more.Ping '{"ping" : "test" }' 
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
$ varlink call -m exec:./service/org.example.more.TestMore '{"n": 10}' 
{
  "state": {
    "start": true
  }
}
{
  "state": {
    "progress": 0
  }
}
{
  "state": {
    "progress": 10
  }
}
{
  "state": {
    "progress": 20
  }
}
{
  "state": {
    "progress": 30
  }
}
{
  "state": {
    "progress": 40
  }
}
{
  "state": {
    "progress": 50
  }
}
{
  "state": {
    "progress": 60
  }
}
{
  "state": {
    "progress": 70
  }
}
{
  "state": {
    "progress": 80
  }
}
{
  "state": {
    "progress": 90
  }
}
{
  "state": {
    "progress": 100
  }
}
{
  "state": {
    "end": true
  }
}

```
