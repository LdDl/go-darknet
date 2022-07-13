wget --output-document=sample.jpg https://cdn-images-1.medium.com/max/800/1*EYFejGUjvjPcc4PZTwoufw.jpeg
wget --output-document=coco.names https://raw.githubusercontent.com/AlexeyAB/darknet/master/data/coco.names
wget --output-document=yolov7-tiny.cfg https://raw.githubusercontent.com/AlexeyAB/darknet/master/cfg/yolov7-tiny.cfg
sed -i -e "\$anames = coco.names" yolov7-tiny.cfg
wget --output-document=yolov7-tiny.weights https://github.com/AlexeyAB/darknet/releases/download/yolov4/yolov7-tiny.weights