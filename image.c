#include <darknet.h>

void fill_image_f32(image* im, int w, int h, int c, float* data) {
    int i;
    for (i = 0; i < w*h*c; i++) {
        im->data[i] = data[i];
    }
}

void set_data_f32_val(float* data, int index, float value) {
    data[index] = value;
}

image resize_image_golang(image im, int w, int h) {
    return resize_image(im, w, h);
}


image make_empty_image(int w, int h, int c)
{
    image out;
    out.data = 0;
    out.h = h;
    out.w = w;
    out.c = c;
    return out;
}

image float_to_image(int w, int h, int c, float *data)
{
    image out = make_empty_image(w,h,c);
    fill_image_f32(&out, w, h, c, data);
    // for (i = 0; i < w*h*c; i++) {
    //     out.data[i] = data[i];
    // }
    return out;
}

