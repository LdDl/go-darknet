#include <stdlib.h>

#include <darknet.h>

#include "network.h"

#include "detection.h"

struct network_box_result perform_network_detect(network *n, image *img, int classes, float thresh, float hier_thresh, float nms) {
    image sized = letterbox_image(*img, n->w, n->h);
    struct network_box_result result = { NULL };
    float *X = sized.data;
    int letter_box = 0;
    network_predict_ptr(n, X);
    result.detections = get_network_boxes(n, img->w, img->h, thresh, hier_thresh, 0, 1, &result.detections_len, letter_box);
    if (nms) {
        do_nms_sort(result.detections, result.detections_len, classes, nms);
    }
    free_image(sized);
    return result;
}

