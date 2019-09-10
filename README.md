# Server-sent events (SSE) client in Go

[![Build Status](https://travis-ci.org/peterhellberg/sseclient.svg?branch=master)](https://travis-ci.org/peterhellberg/sseclient)
[![Go Report Card](https://goreportcard.com/badge/github.com/peterhellberg/sseclient?style=flat)](https://goreportcard.com/report/github.com/peterhellberg/sseclient)
[![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat)](https://godoc.org/github.com/peterhellberg/sseclient)
[![License MIT](https://img.shields.io/badge/license-MIT-lightgrey.svg?style=flat)](https://github.com/peterhellberg/sseclient#license-mit)

I needed a simple client for the events sent from my Spark Core via the
[Spark API](http://docs.spark.io/api/), unfortunately it includes extra
lines like `:ok` that is not part of any event.


Based on the work by [@cryptix](https://github.com/cryptix/goSSEClient)

<img src="https://data.gopher.se/gopher/viking-gopher.svg" align="right" width="30%" height="300">

## License (MIT)

Copyright (c) 2014-2019 [Peter Hellberg](https://c7.se)

> Permission is hereby granted, free of charge, to any person obtaining
> a copy of this software and associated documentation files (the
> "Software"), to deal in the Software without restriction, including
> without limitation the rights to use, copy, modify, merge, publish,
> distribute, sublicense, and/or sell copies of the Software, and to
> permit persons to whom the Software is furnished to do so, subject to
> the following conditions:
>
> The above copyright notice and this permission notice shall be
> included in all copies or substantial portions of the Software.

> THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
> EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
> MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
> NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE
> LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION
> OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION
> WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
