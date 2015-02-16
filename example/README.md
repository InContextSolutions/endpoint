# Endpoint Example

To run the example, start the server

```
go run main.go
```

## Get

```
$ curl http://localhost:8080/foo/something
the answer is 42
```

## Post

```
$ echo '{"text": "Hello **world**!"}' | curl -H "Content-Type:application/json" -d @- http://localhost:8080/bar/baz
You posted {"text": "Hello **world**!"}
```
