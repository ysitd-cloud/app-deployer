package service

import (
	"code.ysitd.cloud/component/deployer/pkg/deployer"
	"code.ysitd.cloud/component/deployer/pkg/manager"
	pb "code.ysitd.cloud/grpc/schema/deployer"
	"github.com/tonyhhyip/go-di-container"
)

func NewServiceProvider(app container.Container) container.ServiceProvider {
	sp := &grpcServiceProvider{
		AbstractServiceProvider: container.NewAbstractServiceProvider(true),
	}
	sp.SetContainer(app)
	return sp
}

func newService(manager manager.Manager, deployer deployer.Controller) pb.DeployerServer {
	return &service{
		manager:  manager,
		deployer: deployer,
	}
}
