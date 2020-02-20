#pragma once

#include <darknet.h>

extern void fill_image_f32(image *im, int w, int h, int c, float* data);
extern void set_data_f32_val(float* data, int index, float value);
extern image resize_image_golang(image im, int w, int h);
extern image make_empty_image(int w, int h, int c);
extern image float_to_image(int w, int h, int c, float *data);