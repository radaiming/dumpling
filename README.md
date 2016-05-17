#A very simple web framework for practice

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

func hello() (int, map[string]string, string) {
	return 200, nil, "hello world"
}

func main() {
	r := dumpling.New()
	r.Get("/", hello)
	r.Post("/", hello)
	r.Serve("127.0.0.1:9988")
}
```

## TODO
* Parse and pass URL parameters and POST content to handler function
* Pass context to handler function?
* ~~Support returning customized HTTP status code~~(Done)
* Write docs/comments
* Support route regex matching
* Support serving static file
* ~~Support middleware~~(Done)