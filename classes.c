#include <stdlib.h>

char *get_class_name(char **names, int index, int names_len) {
    if (index >= names_len) {
        return NULL;
    }
    return names[index];
}