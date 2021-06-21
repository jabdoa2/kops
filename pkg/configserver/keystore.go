/*
Copyright 2020 The Kubernetes Authors.

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

package configserver

import (
	"crypto/x509"
	"fmt"

	"k8s.io/kops/pkg/apis/kops"
	"k8s.io/kops/pkg/apis/nodeup"
	"k8s.io/kops/pkg/pki"
	"k8s.io/kops/upup/pkg/fi"
	"k8s.io/kops/util/pkg/vfs"
)

//configserverKeyStore is a KeyStore backed by the config server.
type configserverKeyStore struct {
	nodeConfig *nodeup.NodeConfig
}

func NewKeyStore(nodeConfig *nodeup.NodeConfig) fi.CAStore {
	return &configserverKeyStore{
		nodeConfig: nodeConfig,
	}
}

// FindPrimaryKeypair implements pki.Keystore
func (s *configserverKeyStore) FindPrimaryKeypair(name string) (*pki.Certificate, *pki.PrivateKey, error) {
	return nil, nil, fmt.Errorf("FindPrimaryKeypair %q not supported by configserverKeyStore", name)
}

// FindKeyset implements fi.Keystore
func (s *configserverKeyStore) FindKeyset(name string) (*fi.Keyset, error) {
	return nil, fmt.Errorf("FindKeyset %q not supported by configserverKeyStore", name)
}

// CreateKeypair implements fi.Keystore
func (s *configserverKeyStore) CreateKeypair(signer string, name string, template *x509.Certificate, privateKey *pki.PrivateKey) (*pki.Certificate, error) {
	return nil, fmt.Errorf("CreateKeypair not supported by configserverKeyStore")
}

// StoreKeyset implements fi.Keystore
func (s *configserverKeyStore) StoreKeyset(name string, keyset *fi.Keyset) error {
	return fmt.Errorf("StoreKeyset not supported by configserverKeyStore")
}

// MirrorTo implements fi.Keystore
func (s *configserverKeyStore) MirrorTo(basedir vfs.Path) error {
	return fmt.Errorf("MirrorTo not supported by configserverKeyStore")
}

// CertificatePool implements fi.CAStore
func (s *configserverKeyStore) CertificatePool(name string, createIfMissing bool) (*fi.CertificatePool, error) {
	return nil, fmt.Errorf("CertificatePool not supported by configserverKeyStore")
}

// FindCertificatePool implements fi.CAStore
func (s *configserverKeyStore) FindCertificatePool(name string) (*fi.CertificatePool, error) {
	return nil, fmt.Errorf("FindCertificatePool not supported by configserverKeyStore")
}

// FindPrivateKey implements fi.CAStore
func (s *configserverKeyStore) FindPrivateKey(name string) (*pki.PrivateKey, error) {
	return nil, fmt.Errorf("FindPrivateKey not supported by configserverKeyStore")
}

// FindCert implements fi.CAStore
func (s *configserverKeyStore) FindCert(name string) (*pki.Certificate, error) {
	for _, cert := range s.nodeConfig.Certificates {
		if cert.Name == name {
			// Special case for the CA certificate
			c, err := pki.ParsePEMCertificate([]byte(cert.Cert))
			if err != nil {
				return nil, fmt.Errorf("error parsing certificate %q: %w", name, err)
			}
			return c, nil
		}
	}

	return nil, fmt.Errorf("FindCert(%q) not supported by configserverKeyStore", name)
}

// ListKeysets implements fi.CAStore
func (s *configserverKeyStore) ListKeysets() (map[string]*fi.Keyset, error) {
	return nil, fmt.Errorf("ListKeysets not supported by configserverKeyStore")
}

// DeleteKeysetItem implements fi.CAStore
func (s *configserverKeyStore) DeleteKeysetItem(item *kops.Keyset, id string) error {
	return fmt.Errorf("DeleteKeysetItem not supported by configserverKeyStore")
}
