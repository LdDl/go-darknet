#include <darknet.h>

image new_darknet_image() {
    image img;
    img.w = 0;
    img.h = 0;
    img.c = 0;
    img.data = NULL;
    return img;
}

image prepare_image(image img, int w, int h, int c){
    img.w=w;
    img.h=h;
    img.c=c;
    return img;
}