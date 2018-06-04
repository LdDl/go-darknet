package darknet

// #include "class_name.h"
import "C"

func freeClassNames(names **C.char) {
	C.free_class_names(names)
}

func loadClassNames(dataConfigFile string) **C.char {
	return C.read_class_names(C._GoStringPtr(dataConfigFile))
}

func makeClassNames(names **C.char, classes int) []string {
	out := make([]string, classes)
	for i := 0; i < classes; i++ {
		n := C.get_class_name(names, C.int(i), C.int(classes))
		s := C.GoString(n)
		out[i] = s
	}

	return out
}
