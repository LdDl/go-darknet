# FORK of go-darknet
# go-darknet: Go bindings for Darknet

[![GoDoc](https://godoc.org/github.com/LdDl/go-darknet?status.svg)](https://godoc.org/github.com/LdDl/go-darknet)

go-darknet is a Go package, which uses Cgo to enable Go applications to use
YOLO in [Darknet].

## License

go-darknet follows [Darknet]'s [license].

## Requirements

For proper codebase please use my fork of [darknet](https://github.com/LdDl/darknet)
There are instructions for defining GPU/CPU + function for loading image from memory.

In order to use go-darknet, `libdarknet.so` should be available in one of
the following locations:

* /usr/lib
* /usr/local/lib

The shared library `libdarknet.so` can be obtained after invoking `make` on
[Darknet]'s codebase.

Also, [darknet.h] should be available in one of the following locations:

* /usr/include
* /usr/local/include

## Install

```shell
go get github.com/LdDl/go-darknet
```

The package name is `darknet`.

## Use

Example Go code/program is provided in the [example] directory. Please
refer to the code on how to use this Go package.

Building and running the example program is easy:

```shell
go install github.com/LdDl/go-darknet/example

# The executable `example` will be available in your $GOPATH/bin
$GOPATH/bin/example
```

## Documentation

See go-darknet's API documentation at [GoDoc].

[Darknet]: https://github.com/pjreddie/darknet
[license]: https://github.com/pjreddie/darknet/blob/master/LICENSE
[darknet.h]: https://github.com/pjreddie/darknet/blob/master/include/darknet.h
[include/darknet.h]: https://github.com/pjreddie/darknet/blob/master/include/darknet.h
[Makefile]: https://github.com/pjreddie/darknet/blob/master/Makefile
[example]: /example
[GoDoc]: https://godoc.org/github.com/LdDl/go-darknet
