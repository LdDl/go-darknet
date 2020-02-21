package darknet

// #include <stdlib.h>
// #include "classes.h"
import "C"

func makeClassNames(names **C.char, classes int) []string {
	out := make([]string, classes)
	for i := 0; i < classes; i++ {
		n := C.get_class_name(names, C.int(i), C.int(classes))
		out[i] = C.GoString(n)
	}
	return out
}
