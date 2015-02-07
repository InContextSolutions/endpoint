# Endpoint Example

To run the example, start the server

```
go run main.go
```

and [get foo](http://0.0.0.0:8080/foo).

In the terminal, you should see two log statements detailing the request flow.

## Get

```
$ curl http://0.0.0.0:8080/foo
the middleware told me the answer is 42
```

## Post

```
$ echo '{"text": "Hello **world**!"}' | curl -H "Content-Type:application/json" -d @- http://0.0.0.0:8080/bar
You posted map[text:Hello **world**!]
```
