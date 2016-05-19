#A very simple web framework for practice
[![Build Status](https://travis-ci.org/radaiming/dumpling.svg?branch=master)](https://travis-ci.org/radaiming/dumpling)

## Install

```
> go get -u github.com/radaiming/dumpling
```

## Use
```
package main

import (
	"github.com/radaiming/dumpling"
)

func hello(ctx *dumpling.HTTPContext) {
	ctx.Response("hello world")
}

func main() {
	r := dumpling.New()
	r.Get("/", hello)
	r.Post("/", hello)
	r.Serve("127.0.0.1:9988")
}
```
See more in [samples](https://github.com/radaiming/dumpling/tree/master/samples).

## TODO
* ~~Parse and pass URL parameters and POST content to handler function~~(Done)
* ~~Pass context to handler function?~~(Done)
* ~~Support returning customized HTTP status code~~(Done)
* Write docs/comments/tests
* ~~Support route regex matching~~
* ~~Support serving static file~~(Done)
* ~~Support middleware~~(Done)
* ~~Write middleware for logging and basic auth~~(Done)
* ~~Write a middleware to log request process time~~(Done)