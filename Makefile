.ONESHELL:
.PHONY: prepare_cuda prepare_cudnn download_darknet build_darknet build_darknet_gpu clean clean_cuda clean_cudnn sudo_install

# Latest battletested AlexeyAB version of Darknet commit
# LATEST_COMMIT?=f056fc3b6a11528fa0522a468eca1e909b7004b7
LATEST_COMMIT?=9d40b619756be9521bc2ccd81808f502daaa3e9a

# Temporary folder for building Darknet
TMP_DIR?=/tmp/

# Manage cuda version
CUDA_VERSION = 10.2
CUDNN_VERSION = 7.6.5
CUDNN_FULL_VERSION = 7.6.5.32
OS_NAME_LOW_CASE = ubuntu
OS_VERSION_CONCATENATED = 1804
OS_ARCH = x86_64
OS_ALTER_ARCH = linux-x64
OS_FULLNAME = $(OS_NAME_LOW_CASE)$(OS_VERSION_CONCATENATED)
# I guess *.pub is static for most of systems
PUBNAME = 7fa2af80

# Install CUDA
prepare_cuda:
	sudo apt-get install linux-headers-$(uname -r)
	rm -rf $(TMP_DIR)install_cuda
	mkdir $(TMP_DIR)install_cuda
	wget -P $(TMP_DIR)install_cuda https://developer.download.nvidia.com/compute/cuda/repos/$(OS_FULLNAME)/$(OS_ARCH)/cuda-$(OS_FULLNAME).pin
	cd $(TMP_DIR)install_cuda
	sudo mv cuda-$(OS_FULLNAME).pin /etc/apt/preferences.d/cuda-repository-pin-600
	sudo apt-key adv --fetch-keys https://developer.download.nvidia.com/compute/cuda/repos/$(OS_FULLNAME)/$(OS_ARCH)/$(PUBNAME).pub
	sudo add-apt-repository "deb http://developer.download.nvidia.com/compute/cuda/repos/$(OS_FULLNAME)/$(OS_ARCH)/ /"
	sudo apt-get update
	sudo apt-get -y install cuda-$(subst .,-,$(CUDA_VERSION))
	cd -

# Install cuDNN
# Notice: this valid instruction for cuDNN version from v7.2.1 up to 8.1.0.77
prepare_cudnn:
	rm -rf $(TMP_DIR)install_cudnn
	mkdir $(TMP_DIR)install_cudnn
	wget -P $(TMP_DIR)install_cudnn https://developer.download.nvidia.com/compute/redist/cudnn/v${CUDNN_VERSION}/cudnn-${CUDA_VERSION}-${OS_ALTER_ARCH}-v${CUDNN_FULL_VERSION}.tgz
	cd $(TMP_DIR)install_cudnn
	tar -xzvf cudnn-${CUDA_VERSION}-${OS_ALTER_ARCH}-v${CUDNN_FULL_VERSION}.tgz
	sudo cp cuda/include/cudnn*.h /usr/local/cuda/include 
	sudo cp -P cuda/lib64/libcudnn* /usr/local/cuda/lib64 
	sudo chmod a+r /usr/local/cuda/include/cudnn*.h /usr/local/cuda/lib64/libcudnn*
	cd -

# Download AlexeyAB version of Darknet
download_darknet:
	rm -rf $(TMP_DIR)install_darknet
	mkdir $(TMP_DIR)install_darknet
	git clone https://github.com/AlexeyAB/darknet.git $(TMP_DIR)install_darknet
	cd $(TMP_DIR)install_darknet
	git checkout $(LATEST_COMMIT)
	cd -

# Build AlexeyAB version of Darknet for usage with CPU only.
build_darknet:
	cd $(TMP_DIR)install_darknet
	sed -i -e 's/GPU=1/GPU=0/g' Makefile
	sed -i -e 's/CUDNN=1/CUDNN=0/g' Makefile
	sed -i -e 's/LIBSO=0/LIBSO=1/g' Makefile
	$(MAKE) -j $(shell nproc --all)
	$(MAKE) preinstall
	cd -

# Build AlexeyAB version of Darknet for usage with both CPU and GPU (CUDA by NVIDIA).
build_darknet_gpu:
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

clean_cuda:
	rm -rf $(TMP_DIR)install_cuda

clean_cudnn:
	rm -rf $(TMP_DIR)install_cudnn

# Do every step for CPU-based only build.
install_darknet: download_darknet build_darknet sudo_install clean

# Do every step for both CPU and GPU-based build.
install_darknet_gpu: download_darknet build_darknet_gpu sudo_install clean

# Do every step for both CPU and GPU-based build if you haven't installed CUDA.
install_darknet_gpu_cuda: prepare_cuda prepare_cudnn download_darknet build_darknet_gpu sudo_install clean clean_cuda clean_cudnn