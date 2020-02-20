# FORK of go-darknet https://github.com/gyonluks/go-darknet applied to FORK of Darknet https://github.com/AlexeyAB/darknet
# go-darknet: Go bindings for Darknet

[![GoDoc](https://godoc.org/github.com/LdDl/go-darknet?status.svg)](https://godoc.org/github.com/LdDl/go-darknet)

go-darknet is a Go package, which uses Cgo to enable Go applications to use
YOLO in [Darknet].

## License

go-darknet follows [Darknet]'s [license].

## Requirements

For proper codebase please use fork of [darknet](https://github.com/AlexeyAB/darknet)
There are instructions for defining GPU/CPU + function for loading image from memory.

In order to use go-darknet, `libdarknet.so` should be available in one of
the following locations:

* /usr/lib
* /usr/local/lib

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
cd $GOPATH/github.com/LdDl/go-darknet/example
#download dataset (coco.names, coco.data, weights and configuration file)
./download_data.sh
#build program
go build main.go
#run it
./main -configFile yolov3.cfg --dataConfigFile coco.data -imageFile sample.jpg -weightsFile yolov3.weights
```

Output should be something like this:
```shell
truck (7): 95.6232% | start point: (78,69) | end point: (222, 291)
truck (7): 81.5451% | start point: (0,114) | end point: (90, 329)
car (2): 99.8129% | start point: (269,192) | end point: (421, 323)
car (2): 99.6615% | start point: (567,188) | end point: (743, 329)
car (2): 99.5795% | start point: (425,196) | end point: (544, 309)
car (2): 96.5765% | start point: (678,185) | end point: (797, 320)
car (2): 91.5156% | start point: (391,209) | end point: (441, 291)
car (2): 88.1737% | start point: (507,193) | end point: (660, 324)
car (2): 83.6209% | start point: (71,199) | end point: (102, 281)
bicycle (1): 59.4000% | start point: (183,276) | end point: (257, 407)
person (0): 96.3393% | start point: (142,119) | end point: (285, 356)
```
## Documentation

See go-darknet's API documentation at [GoDoc].

[Darknet]: https://github.com/pjreddie/darknet
[license]: https://github.com/pjreddie/darknet/blob/master/LICENSE
[darknet.h]: https://github.com/pjreddie/darknet/blob/master/include/darknet.h
[include/darknet.h]: https://github.com/pjreddie/darknet/blob/master/include/darknet.h
[Makefile]: https://github.com/alexeyab/darknet/blob/master/Makefile
[example]: /example
[GoDoc]: https://godoc.org/github.com/LdDl/go-darknet
