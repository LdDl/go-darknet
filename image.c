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
