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

After server started check if REST-requests works. We provide cURL-based example
```shell
curl -F 'image=@sample.jpg' 'http://localhost:8090/detect_objects'
```

Servers response should be something like this:
```json
{
    "net_time": "43.269289ms",
    "overall_time": "43.551604ms",
    "num_detections": 44,
    "detections": [
        {
            "class_id": 7,
            "class_name": "truck",
            "probability": 49.51231,
            "start_point": {
                "x": 0,
                "y": 136
            },
            "end_point": {
                "x": 85,
                "y": 311
            }
        },
        {
            "class_id": 2,
            "class_name": "car",
            "probability": 36.36933,
            "start_point": {
                "x": 95,
                "y": 152
            },
            "end_point": {
                "x": 186,
                "y": 283
            }
        },
        {
            "class_id": 7,
            "class_name": "truck",
            "probability": 48.417683,
            "start_point": {
                "x": 95,
                "y": 152
            },
            "end_point": {
                "x": 186,
                "y": 283
            }
        },
        {
            "class_id": 7,
            "class_name": "truck",
            "probability": 45.652023,
            "start_point": {
                "x": 694,
                "y": 178
            },
            "end_point": {
                "x": 798,
                "y": 310
            }
        },
        {
            "class_id": 2,
            "class_name": "car",
            "probability": 76.8402,
            "start_point": {
                "x": 1,
                "y": 145
            },
            "end_point": {
                "x": 84,
                "y": 324
            }
        },
        {
            "class_id": 7,
            "class_name": "truck",
            "probability": 25.592052,
            "start_point": {
                "x": 107,
                "y": 89
            },
            "end_point": {
                "x": 215,
                "y": 263
            }
        },
        {
            "class_id": 2,
            "class_name": "car",
            "probability": 99.87823,
            "start_point": {
                "x": 511,
                "y": 185
            },
            "end_point": {
                "x": 748,
                "y": 328
            }
        },
        {
            "class_id": 2,
            "class_name": "car",
            "probability": 99.819336,
            "start_point": {
                "x": 261,
                "y": 189
            },
            "end_point": {
                "x": 427,
                "y": 322
            }
        },
        {
            "class_id": 2,
            "class_name": "car",
            "probability": 99.64055,
            "start_point": {
                "x": 426,
                "y": 197
            },
            "end_point": {
                "x": 539,
                "y": 311
            }
        },
        {
            "class_id": 2,
            "class_name": "car",
            "probability": 74.56263,
            "start_point": {
                "x": 692,
                "y": 186
            },
            "end_point": {
                "x": 796,
                "y": 316
            }
        },
        {
            "class_id": 2,
            "class_name": "car",
            "probability": 72.79756,
            "start_point": {
                "x": 388,
                "y": 206
            },
            "end_point": {
                "x": 437,
                "y": 276
            }
        },
        {
            "class_id": 1,
            "class_name": "bicycle",
            "probability": 72.27595,
            "start_point": {
                "x": 178,
                "y": 270
            },
            "end_point": {
                "x": 268,
                "y": 406
            }
        },
        {
            "class_id": 0,
            "class_name": "person",
            "probability": 97.30075,
            "start_point": {
                "x": 143,
                "y": 135
            },
            "end_point": {
                "x": 268,
                "y": 343
            }
        }
    ]
}
```