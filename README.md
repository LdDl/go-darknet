# go-darknet: Go bindings for Darknet

### This is fork of go-darknet https://github.com/gyonluks/go-darknet applied to FORK of Darknet https://github.com/AlexeyAB/darknet

[![GoDoc](https://godoc.org/github.com/LdDl/go-darknet?status.svg)](https://godoc.org/github.com/LdDl/go-darknet)

go-darknet is a Go package, which uses Cgo to enable Go applications to use YOLO in [Darknet].

## Table of Contents

- [Requirements](#requirements)
- [Installation](#installation)
- [Usage](#usage)
- [Documentation](#documentation)
- [License](#license)

## Requirements

For proper codebase please use fork of [darknet](https://github.com/AlexeyAB/darknet). Latest commit I've tested [here](https://github.com/AlexeyAB/darknet/commit/64fb042c63637038671ae9d53c06165599b28912)

In order to use go-darknet, `libdarknet.so` should be available in one of
the following locations:

* /usr/lib
* /usr/local/lib

Also, [darknet.h] should be available in one of the following locations:

* /usr/include
* /usr/local/include

To achieve it, after Darknet compilation (via make) execute following command:
```shell
sudo cp libdarknet.so /usr/lib/libdarknet.so && sudo cp include/darknet.h /usr/local/include/darknet.h
```
Note: do not forget to set LIBSO=1 in Makefile before executing make:
```Makefile
LIBSO=1
```
## Installation

```shell
go get github.com/LdDl/go-darknet
```

## Usage

Example Go program is provided in the [example] directory. Please refer to the code on how to use this Go package.

Building and running program:

Navigate to [example] folder
```shell
cd $GOPATH/github.com/LdDl/go-darknet/example
```

Download dataset (sample of image, coco.names, yolov3.cfg, yolov3.weights).
```shell
./download_data.sh
```
Note: you don't need *coco.data* file anymore, because script below does insert *coco.names* into 'names' filed in *yolov3.cfg* file (so AlexeyAB's fork can deal with it properly)
So last rows in yolov3.cfg file will look like:
```bash
......
[yolo]
mask = 0,1,2
anchors = 10,13,  16,30,  33,23,  30,61,  62,45,  59,119,  116,90,  156,198,  373,326
classes=80
num=9
jitter=.3
ignore_thresh = .7
truth_thresh = 1
random=1
names = coco.names # this is path to coco.names file
```

Build and run program
```
go build main.go && ./main --configFile=yolov3.cfg --weightsFile=yolov3.weights --imageFile=sample.jpg
```

Output should be something like this:
```shell
truck (7): 49.5197% | start point: (0,136) | end point: (85, 311)
car (2): 36.3747% | start point: (95,152) | end point: (186, 283)
truck (7): 48.4384% | start point: (95,152) | end point: (186, 283)
truck (7): 45.6590% | start point: (694,178) | end point: (798, 310)
car (2): 76.8379% | start point: (1,145) | end point: (84, 324)
truck (7): 25.5731% | start point: (107,89) | end point: (215, 263)
car (2): 99.8783% | start point: (511,185) | end point: (748, 328)
car (2): 99.8194% | start point: (261,189) | end point: (427, 322)
car (2): 99.6408% | start point: (426,197) | end point: (539, 311)
car (2): 74.5610% | start point: (692,186) | end point: (796, 316)
car (2): 72.8053% | start point: (388,206) | end point: (437, 276)
bicycle (1): 72.2932% | start point: (178,270) | end point: (268, 406)
person (0): 97.3026% | start point: (143,135) | end point: (268, 343)
```
## Documentation

See go-darknet's API documentation at [GoDoc].

## License

go-darknet follows [Darknet]'s [license].


[Darknet]: https://github.com/pjreddie/darknet
[license]: https://github.com/pjreddie/darknet/blob/master/LICENSE
[darknet.h]: https://github.com/AlexeyAB/darknet/blob/master/include/darknet.h
[include/darknet.h]: https://github.com/AlexeyAB/darknet/blob/master/include/darknet.h
[Makefile]: https://github.com/alexeyab/darknet/blob/master/Makefile
[example]: /example
[GoDoc]: https://godoc.org/github.com/LdDl/go-darknet
