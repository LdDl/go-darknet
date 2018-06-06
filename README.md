# go-darknet: Go bindings for Darknet

[![GoDoc](https://godoc.org/github.com/gyonluks/go-darknet?status.svg)](https://godoc.org/github.com/gyonluks/go-darknet)

go-darknet is a Go package, which uses Cgo to enable Go applications to use
YOLO in [Darknet].

## License

go-darknet follows [Darknet]'s [license].

## Requirements

In order to use go-darknet, `libdarknet.so` should be available in one of
the following locations:

* /usr/lib
* /usr/local/lib

The shared library `libdarknet.so` can be obtained after invoking `make` on
[Darknet]'s codebase.

Also, [darknet.h] should be available in one of the following locations:

* /usr/include
* /usr/local/include

The include file [darknet.h] can be obtained from the `include` directory in
[Darknet]'s codebase. However, some modifications will have to be made.

### Modifying darknet.h for install

Make a copy of [include/darknet.h] and put in the same directory as
`libdarknet.so`.

In Darknet's [Makefile], at the top, there are macros which look like the
following:

```
GPU=0
CUDNN=0
OPENCV=0
OPENMP=0
DEBUG=0
```

If any of the above has the value `1`, they will need to be defined in
[darknet.h].

Do not define the ones with value `0` in [darknet.h]!

For example, if `GPU=1` and `CUDNN=1`, they will need to be defined in
[darknet.h] as follows:

```C
#ifndef DARKNET_API
#define DARKNET_API
#include <stdlib.h>
#include <stdio.h>
#include <string.h>
#include <pthread.h>

#define GPU 1
#define CUDNN 1

#define SECRET_NUM -1234
extern int gpu_index;

// The rest of darknet.h's code...
```

Note the lines `#include GPU 1` and `#include CUDNN 1`. They are added just
after the C standard library's `#include` directives.

It is important to replicate the activated macros (macros with value `1`)
at the top of the [Makefile], with the corresponding `#define` directives
in [darknet.h].

After the changes are made to the copy of [darknet.h], copy it to one of the
locations mentioned above.

## Install

```shell
go get github.com/gyonluks/go-darknet
```

The package name is `darknet`.

## Use

Example Go code/program is provided in the [example] directory. Please
refer to the code on how to use this Go package.

Building and running the example program is easy:

```shell
go install github.com/gyonluks/go-darknet/example

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
[GoDoc]: https://godoc.org/github.com/gyonluks/go-darknet
