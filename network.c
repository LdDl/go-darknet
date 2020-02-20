#include <stdlib.h>

#include <darknet.h>

#include "network.h"

#include "detection.h"

struct network_box_result perform_network_detect(network *n, image img, int classes, float thresh, float hier_thresh, float nms, int letter_box) {
    image sized;
    int ww = img.w;
    int hh =  img.h;
    int cc =  img.c;
    if (letter_box) {
        printf("using letter %d %d %d %d %d %d\n", letter_box, n->w, n->h, img.w, img.h, img.c);
        sized = letterbox_image(img, n->w, n->h);
    } else {
        printf("not using letter: %d %d %d %d %d %d\n", letter_box, n->w, n->h, img.w, img.h, img.c);
        sized = resize_image(img, n->w, n->h);
    }

    // printf("\n>>>>>>>>>>>>>>Fourth Print (Golang)\n");
    // for (int i = 0; i< 50; i++) {
    //     printf("%f\n", sized.data[i]);
    // }
    // printf("\n<<<<<<<<<<<<<<Done\n");

    // sized.data = float*{0.815686, 0.592875, 0.645386, 0.731689, 0.659976};
    // newImg := resize_image_golang(imgDarknet.image, 416, 416)
    
    struct network_box_result result = { NULL };
    float *X = sized.data;
    float *outnwt = network_predict_ptr(n, X);

    // printf("\n>>>>>>>>>>>>>>Golang out\n");
    // for (int i = 0; i< 100; i++) {
    //     printf("%f ", outnwt[i]);
    // }
    // printf("\n<<<<<<<<<<<<<<Done\n");

    int nboxes = 0;
    detection *dets = get_network_boxes(n, img.w, img.h, thresh, hier_thresh, 0, 1, &nboxes, letter_box);
    printf("Clang number of detections: %d %d %d %f %f %d\n", nboxes, img.w, img.h, thresh, hier_thresh, letter_box);

    result.detections = get_network_boxes(n, img.w, img.h, thresh, hier_thresh, 0, 1, &result.detections_len, letter_box);
    printf("Golang number of detections: %d\n", result.detections_len);

    // printf("Clang number of detections: %d\n", result.detections_len);
    if (nms) {
        do_nms_sort(result.detections, result.detections_len, classes, nms);
    }

    free_image(sized);

    return result;
}

