# Example Go application using go-darknet and REST

This is an example Go server application (in terms of REST) which uses go-darknet.

## Run

Navigate to example folder:

```shell
cd $GOPATH/github.com/LdDl/go-darknet/example/rest_example
```

Download dataset (sample of image, coco.names, yolov3.cfg, yolov3.weights).
```shell
./download_data_v3.sh
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
go build main.go && ./main --configFile=yolov3.cfg --weightsFile=yolov3.weights --port 8090
```

@todo