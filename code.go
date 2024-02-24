package cachebust

import (
	"math/rand"
	"strconv"
	"time"
)

var (
	randomness = rand.New( rand.NewSource(time.Now().Unix()) )
)

// Code return a cache-busting code.
//
// Cache-Busting codes are useful when trying to prevent a resource on the Web
// (such an image, a CSS file, or anything else) from being served from a cache.
// Whether that be a web-browser cache, or some cache on (or behind) the Internet.
//
// You can get a cache-busting code with code like:
//
//	cacheBustingCode := cachebust.Code()
//
// And you might use it to do something like:
//
//	var url string = "https://example.com/images/cover.png"
//
//	url = url + "?cb=" cacheBustingCode
//
// The cache-busting code that Code returns looks like this:
//
//	gtpqqoq141-3bkhjtdwagrqq
//
// And only contains lower-case English letters (i.e., "a" to "z"), (western) arabic digits (i.e., "0" to "9"), and a hyphen (i.e., "-").
//
// It is thus safe to use without escaping in a URL and in HTML.
func Code() string {
	var buffer [256]byte
	var p []byte = buffer[0:0]

	{
		var n int64 = time.Now().UnixMicro()
		var s string = strconv.FormatInt(n, 36)

		p = append(p, s...)
	}
	{
		p = append(p, "-"...)
	}
	{
		var n uint64 = randomness.Uint64()
		var s string = strconv.FormatUint(n, 36)

		p = append(p, s...)
	}

	return string(p)
}
