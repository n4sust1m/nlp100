# nlp100/golang

```
$ docker image build -t nlp100-go .
prepare execution environment

$ docker container run -it -v $(pwd):/app nlp100-go go test ./${ test number, ex: 00, 01, 10 }
run test

$ go run ./xx/xxx.go
execute this code
```
