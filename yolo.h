#include <stdint.h>
#include <stdlib.h>

#define C_SHARP_MAX_OBJECTS 1000

typedef struct bbbox_t {
    unsigned int x, y, w, h;       // (x,y) - top-left corner, (w, h) - width & height of bounded box
    float prob;                    // confidence - probability that the object was found correctly
    unsigned int obj_id;           // class of object - from range [0, classes-1]
    unsigned int track_id;         // tracking id for video (0 - untracked, 1 - inf - tracked object)
    unsigned int frames_counter;   // counter of frames on which the object was detected
    float x_3d, y_3d, z_3d;        // center of object (in Meters) if ZED 3D Camera is used
} bbbox_t;

typedef struct bbbox_t_container {
    bbbox_t candidates[C_SHARP_MAX_OBJECTS];
} bbbox_t_container;


extern  int init(const char *configurationFilename, const char *weightsFilename, int gpu);
extern  int detect_image(const char *filename, bbbox_t_container *container);
extern  int detect_image_mat(const uint8_t* data, const size_t data_length, bbbox_t_container *container);
extern  int detect_mat(const void *mat, bbbox_t_container *container);
extern  int dispose();
extern  int get_device_count();
extern  int get_device_name(int gpu, char* deviceName);