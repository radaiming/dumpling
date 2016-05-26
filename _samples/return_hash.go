/*
> curl -F 1=@1.jpg -F 2=@2.jpg 127.0.0.1:9988
1.jpg -> 5063f79713f2a54813f47ed7c5184b0be3c0ae49
2.jpg -> 7f20a6d2ab38a9b55161bd4c40e4a64165f19238
*/
package main

import (
	"crypto/sha1"
	"fmt"
	"io"

	"github.com/radaiming/dumpling"
	"github.com/radaiming/dumpling/middlewares"
)

func hashAndReturn(ctx *dumpling.HTTPContext) {
	returnContent := ""
	for _, fileHeaders := range ctx.GetMultipartForm().File {
		for _, fileHeader := range fileHeaders {
			fileName := fileHeader.Filename
			f, err := fileHeader.Open()
			if err != nil {
				continue
			}
			defer f.Close()
			h := sha1.New()
			_, err = io.Copy(h, f)
			if err != nil {
				continue
			}
			returnContent += fmt.Sprintf("%s -> %x\n", fileName, h.Sum(nil))
		}
	}
	ctx.GetMultipartForm().RemoveAll()
	ctx.Response(returnContent)
}

func main() {
	r := dumpling.New()
	r.Plug(middlewares.ProcessTimeLogger(r))
	r.Post("/", hashAndReturn)
	r.Serve("127.0.0.1:9988")
}
