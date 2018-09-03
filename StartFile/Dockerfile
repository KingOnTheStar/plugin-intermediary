FROM ubuntu:16.04 as build

RUN apt-get update && apt-get install -y --no-install-recommends \
        g++ \
        ca-certificates \
	git \
        wget && \
    rm -rf /var/lib/apt/lists/*

ENV GOLANG_VERSION 1.10.3
RUN wget -nv -O - https://storage.googleapis.com/golang/go${GOLANG_VERSION}.linux-amd64.tar.gz \
    | tar -C /usr/local -xz
ENV GOPATH /go
ENV GOBIN /go/bin
ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH

RUN go get -u google.golang.org/grpc
RUN go get -u github.com/golang/protobuf/protoc-gen-go

WORKDIR /go/src/plugin-intermediary
COPY plugin-intermediary .
RUN go build -o /go/bin/kubelet_client

WORKDIR /go/src/nvidia-device-plugin
COPY k8s-device-plugin .

RUN export CGO_LDFLAGS_ALLOW='-Wl,--unresolved-symbols=ignore-in-object-files' && \
    go install -ldflags="-s -w" -v nvidia-device-plugin

FROM debian:stretch-slim

ENV NVIDIA_VISIBLE_DEVICES=all
ENV NVIDIA_DRIVER_CAPABILITIES=utility

WORKDIR /go/src/ca
COPY kubernetes ./kubernetes
RUN mv kubernetes /etc

COPY --from=build /go/bin/nvidia-device-plugin /usr/bin/nvidia-device-plugin
COPY --from=build /go/bin/kubelet_client /usr/bin/kubelet_client

WORKDIR /go/src/
COPY start.sh .
RUN chmod 777 start.sh

ENTRYPOINT ./start.sh
