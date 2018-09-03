package Impl

import (
	devtypes "plugin-intermediary/device/types"
	"plugin-intermediary/types"
)

type GPUManager struct{}

func (m *GPUManager) New() error {
	return nil
}

func (m *GPUManager) Start() error {
	return nil
}

func (m *GPUManager) UpdateNodeInfo(*types.NodeInfo) error {
	return nil
}

func (m *GPUManager) Allocate(*types.PodInfo, *types.ContainerInfo) ([]devtypes.Volume, []string, error) {
	v := []devtypes.Volume{}
	names := []string{}
	return v, names, nil
}

func (m *GPUManager) GetName() string {
	return "HelloDevice"
}
