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

void to_float_and_fill_image(image* im, int w, int h, uint8_t* data) {
    int x, y, idx_source;
    int pixel_count = w * h;
    int idx = 0;

    for (y = 0; y < h; y++) {
        for (x = 0; x < w; x++) {
            idx_source = (y*w + x) * 4;
            im->data[(pixel_count*0) + idx] = (float)data[idx_source] / 255;
            im->data[(pixel_count*1) + idx] = (float)data[idx_source+1] / 255;
            im->data[(pixel_count*2) + idx] = (float)data[idx_source+2] / 255;
            idx++;
        }
    }
}

void nrgba_to_float_and_fill_image(image* im, int w, int h, const uint8_t* pix, int stride) {
    const float scale = 1.0f / 255.0f;
    int pixel_count = w * h;
    float* r_plane = im->data;
    float* g_plane = im->data + pixel_count;
    float* b_plane = im->data + pixel_count * 2;
    int idx = 0;

    for (int y = 0; y < h; y++) {
        const uint8_t* row = pix + y * stride;
        for (int x = 0; x < w; x++) {
            int src = x * 4;
            r_plane[idx] = (float)row[src]     * scale;
            g_plane[idx] = (float)row[src + 1] * scale;
            b_plane[idx] = (float)row[src + 2] * scale;
            idx++;
        }
    }
}
