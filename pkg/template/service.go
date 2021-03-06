package template

import (
	"k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
)

func getPortData(port int32) intstr.IntOrString {
	return intstr.IntOrString{
		Type:   intstr.Int,
		IntVal: port,
	}
}

func GenerateService(id string) *v1.Service {
	selector := getLabels(id)

	ports := []v1.ServicePort{
		{
			Name:       "http",
			Port:       80,
			TargetPort: getPortData(80),
			Protocol:   v1.ProtocolTCP,
		},
	}

	spec := v1.ServiceSpec{
		Type:     v1.ServiceTypeClusterIP,
		Selector: selector,
		Ports:    ports,
	}

	return &v1.Service{
		ObjectMeta: getObjectMeta(id),
		Spec:       spec,
	}
}
