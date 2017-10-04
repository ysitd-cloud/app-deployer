package kubernetes

import "k8s.io/api/core/v1"

func (d *deploymentV1) GetSecret() *v1.Secret {
	return &v1.Secret{
		ObjectMeta: d.getObjectMeta(),
		Data:       d.application.GetEnvironment().ToBytesMap(),
		Type:       v1.SecretTypeOpaque,
	}
}
