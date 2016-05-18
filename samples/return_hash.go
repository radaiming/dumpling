/*
> curl -F 1=@1.jpg -F 2=@2.jpg 127.0.0.1:9988
1.jpg -> 5063f79713f2a54813f47ed7c5184b0be3c0ae49
2.jpg -> 7f20a6d2ab38a9b55161bd4c40e4a64165f19238
*/
package main

import (
	"crypto/sha1"
	"fmt"

	"github.com/radaiming/dumpling"
)

func hashAndReturn(ctx *dumpling.HTTPContext) {
	returnContent := ""
	multipartStreamReader := ctx.GetMultipartStreamReader()
	for {
		part, err := multipartStreamReader.NextPart()
		if err != nil {
			break
		}
		h := sha1.New()
		buffer := make([]byte, 1024)
		for {
			n, err := part.Read(buffer)
			h.Write(buffer[0:n])
			if err != nil {
				break
			}
		}
		returnContent += fmt.Sprintf("%s -> %x\n", part.FileName(), h.Sum(nil))
	}
	ctx.Response(returnContent)
}

func main() {
	r := dumpling.New()
	r.Post("/", hashAndReturn)
	r.Serve("127.0.0.1:9988")
}
