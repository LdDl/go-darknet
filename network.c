#include <stdlib.h>

#include <darknet.h>

#include "network.h"

#include "detection.h"

struct network_box_result perform_network_detect(network *n, image *img, int classes, float thresh, float hier_thresh, float nms, int letter_box) {
    image sized;
    if (letter_box) {
        sized = letterbox_image(*img, n->w, n->h);
    } else {
        sized = resize_image(*img, n->w, n->h);
    }
    struct network_box_result result = { NULL };
    float *X = sized.data;
    network_predict_ptr(n, X);
    int nboxes = 0;
    detection *dets = get_network_boxes(n, img->w, img->h, thresh, hier_thresh, 0, 1, &nboxes, letter_box);
    result.detections = get_network_boxes(n, img->w, img->h, thresh, hier_thresh, 0, 1, &result.detections_len, letter_box);
    if (nms) {
        do_nms_sort(result.detections, result.detections_len, classes, nms);
    }
    free_image(sized);
    return result;
}

