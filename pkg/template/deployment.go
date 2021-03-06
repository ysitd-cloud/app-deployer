package template

import (
	"fmt"

	"k8s.io/api/apps/v1beta2"
	"k8s.io/api/core/v1"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func GenerateDeployment(id, image, tag string, env map[string]string) *v1beta2.Deployment {
	meta := getObjectMeta(id)
	container := getContainer(id, image, tag, env)

	template := v1.PodTemplateSpec{
		ObjectMeta: meta,
		Spec: v1.PodSpec{
			Containers: []v1.Container{container},
		},
	}

	selector := &metav1.LabelSelector{
		MatchLabels: getLabels(id),
	}

	spec := v1beta2.DeploymentSpec{
		Replicas: int32Ptr(1),
		Template: template,
		Selector: selector,
	}

	return &v1beta2.Deployment{
		ObjectMeta: meta,
		Spec:       spec,
	}
}

func getContainer(id, image, tag string, env map[string]string) v1.Container {
	image = fmt.Sprintf("%s:%s", image, tag)

	port := v1.ContainerPort{
		Name:          "http",
		ContainerPort: 80,
		Protocol:      v1.ProtocolTCP,
	}

	container := v1.Container{
		Name:  GetName(id),
		Image: image,
		Ports: []v1.ContainerPort{port},
		Env:   getEnv(id, env),
	}

	return container
}

func getEnv(id string, env map[string]string) []v1.EnvVar {
	vars := make([]v1.EnvVar, 0)

	selector := v1.LocalObjectReference{
		Name: GetName(id),
	}

	for k := range env {
		source := v1.SecretKeySelector{
			LocalObjectReference: selector,
			Key:                  k,
		}

		env := v1.EnvVar{
			Name: k,
			ValueFrom: &v1.EnvVarSource{
				SecretKeyRef: &source,
			},
		}

		vars = append(vars, env)
	}

	return vars
}

func int32Ptr(i int32) *int32 { return &i }
