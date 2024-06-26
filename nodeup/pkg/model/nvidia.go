/*
Copyright 2021 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package model

import (
	"k8s.io/kops/upup/pkg/fi"
	"k8s.io/kops/upup/pkg/fi/nodeup/nodetasks"
)

// NvidiaBuilder installs the Nvidia driver and runtime.
type NvidiaBuilder struct {
	*NodeupModelContext
}

var _ fi.NodeupModelBuilder = &NvidiaBuilder{}

// Build is responsible for installing packages.
func (b *NvidiaBuilder) Build(c *fi.NodeupModelBuilderContext) error {
	if b.InstallNvidiaRuntime() && b.Distribution.IsUbuntu() {
		c.AddTask(&nodetasks.AptSource{
			Name:    "nvidia-container-toolkit",
			Keyring: "https://nvidia.github.io/libnvidia-container/gpgkey",
			Sources: []string{
				"deb https://nvidia.github.io/libnvidia-container/stable/deb/$(ARCH) /",
			},
		})
		c.AddTask(&nodetasks.Package{Name: "nvidia-container-toolkit"})
		c.AddTask(&nodetasks.Package{Name: b.NodeupConfig.NvidiaGPU.DriverPackage})
	}
	return nil
}
