# go-cachebust

Package **cachebust** provides tools for cache-busting, for the Go programming language.

Cache-Busting codes are useful when trying to prevent a resource on the Web (such an image, a CSS file, or anything else) from being served from a cache.
Whether that be a web-browser cache, or some cache on (or behind) the Internet.

## Documention

Online documentation, which includes examples, can be found at: http://godoc.org/sourcecode.social/reiver/go-cachebust

[![GoDoc](https://godoc.org/sourcecode.social/reiver/go-cachebust?status.svg)](https://godoc.org/sourcecode.social/reiver/go-cachebust)

## Example

You can get a cache-busting code with code like:

```golang
cacheBustingCode := cachebust.Code()
```
And you might use it to do something like:

```golang
var url string = "https://example.com/images/cover.png"

url = url + "?cb=" cacheBustingCode
```

Which would make `url` look something like this:
```golang
"https://example.com/images/cover.png?cb=gtpqqoq141-3bkhjtdwagrqq"
```

In general, the cache-busting codes look something like these:

* `gtpqqoq141-3bkhjtdwagrqq`
* `gtps0bieqy-h3xw4swfsab5`
* `h4w7d3moqo-omd4wbb3z1iq`
* `hg2mtkhalg-3m28oenx1m9n1`
* `idlx5kggsz-dx4f7vny6aok`
* `jxi2czgiwp-p9xlmsgl8gzm`
* `n1acr8u6kh-akc4bqx3rff5`

They only contains lower-case English letters (i.e., "a" to "z"), (western) arabic digits (i.e., "0" to "9"), and a hyphen (i.e., "-").

And thus they are safe to use without escaping in a URL and in HTML.

## Import

To import package **cachebust** use `import` code like the follownig:
```
import "sourcecode.social/reiver/go-cachebust"
```

## Installation

To install package **cachebust** do the following:
```
GOPROXY=direct go get https://sourcecode.social/reiver/go-cachebust
```

## Author

Package **cachebust** was written by [Charles Iliya Krempeaux](http://changelog.ca)
