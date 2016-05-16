#A very simple web framework for practice.

## Install

```
> go get -u github.com/radaiming/dumpling
```

## Use
```
package main

import (
    "dumpling"
)

func hello() string {
    return "hello world"
}

func main() {
    dumpling.AddRoute("/", "GET", hello)
    dumpling.Serve("127.0.0.1:9988")
}
```

## TODO
* Parse and pass URL parameters and POST content to handler function
* Support returning customized HTTP status code
* Write docs/comments
* Support route regex matching
* Support serving static file