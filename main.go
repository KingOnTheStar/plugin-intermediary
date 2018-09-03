package main

import (
	"github.com/golang/glog"
	"github.com/spf13/pflag"
	"k8s.io/apiserver/pkg/util/flag"
	kubeletapp "k8s.io/kubernetes/cmd/kubelet/app"
	"k8s.io/kubernetes/cmd/kubelet/app/options"
	"plugin-intermediary/kubeadvertise"
	"plugin-intermediary/kubelet_client"
)

func main() {
	kubeletFlags := options.NewKubeletFlags()
	kubeletFlags.AddFlags(pflag.CommandLine)

	// construct KubeletConfiguration object and register command line flags mapping
	defaultConfig, err := options.NewKubeletConfiguration()
	if err != nil {
		glog.Errorf("NewKubeletConfiguration Error: %v", err)
		return
	}
	options.AddKubeletConfigFlags(pflag.CommandLine, defaultConfig)

	flag.InitFlags()

	// validate the initial KubeletFlags, to make sure the dynamic-config-related flags aren't used unless the feature gate is on
	if err := options.ValidateKubeletFlags(kubeletFlags); err != nil {
		glog.Errorf("ValidateKubeletFlags: %v", err)
		return
	}
	// bootstrap the kubelet config controller, app.BootstrapKubeletConfigController will check
	// feature gates and only turn on relevant parts of the controller
	kubeletConfig, _, err := kubeletapp.BootstrapKubeletConfigController(
		defaultConfig, kubeletFlags.InitConfigDir, kubeletFlags.DynamicConfigDir)
	if err != nil {
		glog.Errorf("BootstrapKubeletConfigController Error: %v", err)
		return
	}

	// construct a KubeletServer from kubeletFlags and kubeletConfig
	kubeletServer := &options.KubeletServer{
		KubeletFlags:         *kubeletFlags,
		KubeletConfiguration: *kubeletConfig,
	}

	glog.Info("Hello World!")
	done := make(chan bool)
	// start the device advertiser
	if _, err := kubeadvertise.StartDeviceAdvertiser(kubeletServer, done); err != nil {
		glog.Errorf("StartDeviceAdvertiser Error: %v", err)
		return
	}

	kubelet_client.RpcService()

	<-done // wait forever
	done <- true
}
