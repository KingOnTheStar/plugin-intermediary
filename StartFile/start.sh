#!/bin/bash

kubelet_client --address=10.130.21.133 --hostname-override=10.130.21.133 --pod-infra-container-image=registry.cn-hangzhou.aliyuncs.com/imooc/pause-amd64:3.0 --kubeconfig=/etc/kubernetes/kubelet.kubeconfig --experimental-bootstrap-kubeconfig=/etc/kubernetes/bootstrap.kubeconfig --cert-dir=/etc/kubernetes/ca --hairpin-mode hairpin-veth --network-plugin=cni --cni-conf-dir=/etc/cni/net.d --cni-bin-dir=/home/tusimple/kubernetes/calico_bin --cluster-dns=10.68.0.2 --cluster-domain=cluster.local. --allow-privileged=true --fail-swap-on=false --logtostderr=true --v=2 --runtime-cgroups=/systemd/system.slice --kubelet-cgroups=/systemd/system.slice --feature-gates=DevicePlugins=true &

sleep 5

nvidia-device-plugin
