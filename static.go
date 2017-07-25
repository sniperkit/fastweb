package main

import (
	"github.com/valyala/fasthttp"
)

func fsHandler(root string, stripSlashes int) fasthttp.RequestHandler {
	// Setup FS handler
	fs := &fasthttp.FS{
		Root:               root,
		IndexNames:         []string{"index.html"},
		GenerateIndexPages: Configuration.Server.GenerateIndexPages,
		Compress:           Configuration.Server.Compress,
		AcceptByteRange:    Configuration.Server.ByteRange,
	}
	if stripSlashes > 0 {
		fs.PathRewrite = fasthttp.NewVHostPathRewriter(stripSlashes)
	}
	return fs.NewRequestHandler()
}

// design and code by tsingson
