[![GoDoc](https://godoc.org/github.com/LdDl/go-darknet?status.svg)](https://godoc.org/github.com/LdDl/go-darknet)
[![Sourcegraph](https://sourcegraph.com/github.com/LdDl/go-darknet/-/badge.svg)](https://sourcegraph.com/github.com/LdDl/go-darknet?badge)
[![Go Report Card](https://goreportcard.com/badge/github.com/LdDl/go-darknet)](https://goreportcard.com/report/github.com/LdDl/go-darknet)
[![GitHub tag](https://img.shields.io/github/tag/LdDl/go-darknet.svg)](https://github.com/LdDl/go-darknet/releases)

# go-darknet: Go bindings for Darknet (Yolo V4, Yolo V3)
### go-darknet is a Go package, which uses Cgo to enable Go applications to use YOLO V4/V3 in [Darknet].

#### Since this repository https://github.com/gyonluks/go-darknet  is no longer maintained I decided to move on and make little different bindings for Darknet.
#### This bindings aren't for [official implementation](https://github.com/pjreddie/darknet) but for [AlexeyAB's fork](https://github.com/AlexeyAB/darknet).

#### Paper Yolo v4: https://arxiv.org/abs/2004.10934
#### Paper Yolo v3: https://arxiv.org/abs/1804.02767

## Table of Contents

- [Why](#why)
- [Requirements](#requirements)
- [Installation](#installation)
- [Usage](#usage)
- [Documentation](#documentation)
- [License](#license)

## Why
**Why does this repository exist?**

Because this repository https://github.com/gyonluks/go-darknet is no longer maintained.

**What is purpose of this bindings when you can have [GoCV](https://github.com/hybridgroup/gocv#gocv) (bindings to OpenCV) and it handle Darknet YOLO perfectly?**

Well, you don't need bunch of OpenCV dependencies and OpenCV itself sometimes.

Example of such project here: https://github.com/LdDl/license_plate_recognition#license-plate-recognition-with-go-darknet---- .


## Requirements

You need to install fork of [darknet](https://github.com/AlexeyAB/darknet). Latest commit I've tested is [here](https://github.com/AlexeyAB/darknet/commit/9d40b619756be9521bc2ccd81808f502daaa3e9a). It corresponds last official [YOLOv4 release](https://github.com/AlexeyAB/darknet/releases/tag/yolov4)

Use provided [Makefile](Makefile).

* For CPU-based instalattion:
    ```shell
    make install_darknet
    ```
* For both CPU and GPU-based instalattion if you HAVE CUDA installed:
    ```shell
    make install_darknet_gpu
    ```
    Note: I've tested CUDA [10.2](https://developer.nvidia.com/cuda-10.2-download-archive) and cuDNN is [7.6.5](https://developer.nvidia.com/rdp/cudnn-archive#a-collapse765-102))

* For both CPU and GPU-based instalattion if you HAVE NOT CUDA installed:
    ```shell
    make install_darknet_gpu_cuda
    ```
    Note: There is some struggle in Makefile for cuDNN, but I hope it works in Ubuntu atleast. Do not forget provide proper CUDA and cuDNN versions.


## Installation

```shell
go get github.com/LdDl/go-darknet
```

## Usage

Example Go program is provided in the [examples] directory. Please refer to the code on how to use this Go package.

Building and running program:

* Navigate to [examples] folder
    ```shell
    cd ${YOUR PATH}/github.com/LdDl/go-darknet/cmd/examples
    ```

* Download dataset (sample of image, coco.names, yolov4.cfg (or v3), yolov4.weights(or v3)).
    ```shell
    #for yolo v4
    ./download_data.sh
    #for yolo v3
    ./download_data_v3.sh
    ```
* Note: you don't need *coco.data* file anymore, because sh-script above does insert *coco.names* into 'names' field in *yolov4.cfg* file (so AlexeyAB's fork can deal with it properly)
    So last rows in yolov4.cfg file will look like:
    ```bash
    ......
    [yolo]
    .....
    iou_loss=ciou
    nms_kind=greedynms
    beta_nms=0.6

    names = coco.names # this is path to coco.names file
    ```

* Also do not forget change batch and subdivisions sizes from:
    ```shell
    batch=64
    subdivisions=8
    ```
    to
    ```shell
    batch=1
    subdivisions=1
    ```
    It will reduce amount of VRAM used for detector test.


* Build and run example program
    Yolo V4:
    ```shell
    go build -o base_example/main base_example/main.go && ./base_example/main --configFile=yolov4.cfg --weightsFile=yolov4.weights --imageFile=sample.jpg
    ```

    Output should be something like this:
    ```shell
    traffic light (9): 73.5040% | start point: (238,73) | end point: (251, 106)
    truck (7): 96.6401% | start point: (95,79) | end point: (233, 287)
    truck (7): 96.4774% | start point: (662,158) | end point: (800, 321)
    truck (7): 96.1841% | start point: (0,77) | end point: (86, 333)
    truck (7): 46.8694% | start point: (434,173) | end point: (559, 216)
    car (2): 99.7370% | start point: (512,188) | end point: (741, 329)
    car (2): 99.2532% | start point: (260,191) | end point: (422, 322)
    car (2): 99.0333% | start point: (425,201) | end point: (547, 309)
    car (2): 83.3920% | start point: (386,210) | end point: (437, 287)
    car (2): 75.8621% | start point: (73,199) | end point: (102, 274)
    car (2): 39.1925% | start point: (386,206) | end point: (442, 240)
    bicycle (1): 76.3121% | start point: (189,298) | end point: (253, 402)
    person (0): 97.7213% | start point: (141,129) | end point: (283, 362)
    ```

    Yolo V3:
    ```
    go build main.go && ./main --configFile=yolov3.cfg --weightsFile=yolov3.weights --imageFile=sample.jpg
    ```

    Output should be something like this:
    ```shell
    truck (7): 49.5123% | start point: (0,136) | end point: (85, 311)
    car (2): 36.3694% | start point: (95,152) | end point: (186, 283)
    truck (7): 48.4177% | start point: (95,152) | end point: (186, 283)
    truck (7): 45.6520% | start point: (694,178) | end point: (798, 310)
    car (2): 76.8402% | start point: (1,145) | end point: (84, 324)
    truck (7): 25.5920% | start point: (107,89) | end point: (215, 263)
    car (2): 99.8782% | start point: (511,185) | end point: (748, 328)
    car (2): 99.8193% | start point: (261,189) | end point: (427, 322)
    car (2): 99.6405% | start point: (426,197) | end point: (539, 311)
    car (2): 74.5627% | start point: (692,186) | end point: (796, 316)
    car (2): 72.7975% | start point: (388,206) | end point: (437, 276)
    bicycle (1): 72.2760% | start point: (178,270) | end point: (268, 406)
    person (0): 97.3007% | start point: (143,135) | end point: (268, 343)
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
[examples]: cmd/examples/base_example
[GoDoc]: https://godoc.org/github.com/LdDl/go-darknet
