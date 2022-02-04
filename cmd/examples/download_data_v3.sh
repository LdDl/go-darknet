wget --output-document=sample.jpg https://cdn-images-1.medium.com/max/800/1*EYFejGUjvjPcc4PZTwoufw.jpeg
wget --output-document=coco.names https://raw.githubusercontent.com/AlexeyAB/darknet/master/data/coco.names
wget --output-document=yolov3.cfg https://raw.githubusercontent.com/AlexeyAB/darknet/master/cfg/yolov3.cfg
sed -i -e "\$anames = coco.names" yolov3.cfg
wget --output-document=yolov3.weights https://pjreddie.com/media/files/yolov3.weights