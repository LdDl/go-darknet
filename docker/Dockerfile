# Build phase
FROM ubuntu:18.04 as builder

ENV darknet_commit=a234a5022333c930de08f2470184ef4e0c68356e

WORKDIR /root/build
COPY Makefile.cpu .
RUN apt-get -y update && \
	apt-get -y install --no-install-recommends git build-essential ca-certificates && \
	git clone https://github.com/AlexeyAB/darknet && \
	cd darknet && \
	git checkout $darknet_commit && \
	cp -f /root/build/Makefile.cpu Makefile && \
	make

# Final Image
FROM golang:1.14

RUN apt-get -y update && \
        apt-get -y install --no-install-recommends nano bash jq

WORKDIR /root
COPY --from=builder /root/build/darknet/darknet \
	/root/build/darknet/libdarknet.so \
	/root/build/darknet/include/darknet.h \
	./staging/

RUN mv staging/darknet /usr/local/bin && \
	mv staging/darknet.h /usr/include && \
	mv staging/libdarknet.so /usr/lib && \
	rm -rf staging

RUN go get -u github.com/LdDl/go-darknet \
    && go get -u github.com/disintegration/imaging

WORKDIR /darknet

COPY download_data.sh .

CMD ["/bin/bash"]
