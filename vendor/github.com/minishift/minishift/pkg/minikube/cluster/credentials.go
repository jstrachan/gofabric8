/*
Copyright 2016 The Kubernetes Authors All rights reserved.

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

package cluster

import (
	"net"

	"github.com/minishift/minishift/pkg/minishift/registration"
	"github.com/minishift/minishift/pkg/util"
)

var (
	// This is the internalIP the the API server and other components communicate on.
	internalIP             = net.ParseIP(DefaultServiceClusterIP)
	RegistrationParameters = new(registration.RegistrationParameters)
)

func GenerateCerts(pub, priv string, ip net.IP) error {
	ips := []net.IP{ip, internalIP}
	if err := util.GenerateSelfSignedCert(pub, priv, ips, GetAlternateDNS(DefaultDNSDomain)); err != nil {
		return err
	}
	return nil
}
