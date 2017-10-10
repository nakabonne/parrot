# parrot
A cli tool to display the log of requests

# Usage

```
$ go build -o parrot
$ parrot -p 8080

// Throw any request
$ curl -d key=value http://localhost:8080

// displayed your request
POST / HTTP/1.1
Host: localhost:8080
Accept: */*
Content-Length: 9
Content-Type: application/x-www-form-urlencoded
User-Agent: curl/7.51.0

key=value
```

# Installation

```
$ go get github.com/ryonakao/parrot
```
