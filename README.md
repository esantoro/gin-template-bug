Gin-gonic and block in templates
=====================

This repository is a demo for an issue in the `gin-gonic/gin`
repository.

The "bug" prevents templates using the new go1.6 template feature
"block" from being used.

## Bug reproduction

In order to reproduce the bug:

1. run working.go (`go run working.go`): this is the desired output!
2. run not-working.go (`go run not-working.go`)
3. perform a GET request to the not-working go code: `curl
   localhost:6000`

**Note**: for some reason, performing the GET request via the browser
does not work on my machine: an error message is displayed in the
browser (chromium) and no debug information about the HTTP request is
shown by the gin process.

## Reference

+ [How to Use Template Blocks in Go 1.6](http://www.josephspurrier.com/how-to-use-template-blocks-in-go-1-6/)
+ [html/template actions](https://tip.golang.org/pkg/text/template/#hdr-Actions) see "block"
