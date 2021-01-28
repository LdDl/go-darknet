.ONESHELL:
.PHONY: download build clean

# Latest battletested AlexeyAB version of Darknet commit
LATEST_COMMIT?=f056fc3b6a11528fa0522a468eca1e909b7004b7

# Temporary folder for building Darknet
TMP_DIR?=/tmp/

# Download AlexeyAB version of Darknet
download:
	rm -rf $(TMP_DIR)install_darknet
	mkdir $(TMP_DIR)install_darknet
	git clone https://github.com/AlexeyAB/darknet.git $(TMP_DIR)install_darknet
	cd $(TMP_DIR)install_darknet
	git checkout $(LATEST_COMMIT)
	cd -

# Build AlexeyAB version of Darknet for usage with CPU only.
build:
	cd $(TMP_DIR)install_darknet
	sed -i -e 's/GPU=1/GPU=0/g' Makefile
	sed -i -e 's/CUDNN=1/CUDNN=0/g' Makefile
	sed -i -e 's/LIBSO=0/LIBSO=1/g' Makefile
	$(MAKE) -j $(shell nproc --all)
	$(MAKE) preinstall
	cd -

# Build AlexeyAB version of Darknet for usage with both CPU and GPU (CUDA by NVIDIA).
build_gpu:
	cd $(TMP_DIR)install_darknet
	sed -i -e 's/GPU=0/GPU=1/g' Makefile
	sed -i -e 's/CUDNN=0/CUDNN=1/g' Makefile
	sed -i -e 's/LIBSO=0/LIBSO=1/g' Makefile
	$(MAKE) -j $(shell nproc --all)
	$(MAKE) preinstall
	cd -

# Install system wide.
sudo_install:
	cd $(TMP_DIR)install_darknet
	sudo cp libdarknet.so /usr/lib/libdarknet.so
	sudo cp include/darknet.h /usr/include/darknet.h
	sudo ldconfig
	cd -

# Cleanup temporary files for building process
clean:
	rm -rf $(TMP_DIR)install_darknet

# Do every step for CPU-based only build.
install: download build sudo_install clean

# Do every step for both CPU and GPU-based build.
install_gpu: download build_gpu sudo_install clean