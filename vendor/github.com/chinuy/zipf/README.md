# zipf [![Build Status](https://travis-ci.org/chinuy/zipf.svg?branch=master)](https://travis-ci.org/chinuy/zipf)
Random number generator following Zipf's law

# Motivation
The [built-in package](https://golang.org/pkg/math/rand/#Zipf) implementation doesn't support *alpha* < 1.0.

Referring to the answer in [here](https://stackoverflow.com/questions/1366984/generate-random-numbers-distributed-by-zipf/),
here is a go version zipf generator.

The generator will randomly produce a number in the range of [0, n-1].

# Usage

```go
// import "github.com/chinuy/zipf"
r := rand.New(rand.NewSource(time.Now().UTC().UnixNano()))
z := zipf.NewZipf(r, 1.0, 10) // n = 10, alpha = 1.0
for i := 0; i < 10; i++ {
  randNumber := z.Uint64()
 // do something you want
}
```
