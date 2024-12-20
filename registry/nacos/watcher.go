package nacos

import (
	"context"
	"fmt"

	"github.com/go-kratos/kratos/v2/registry"
	"github.com/nacos-group/nacos-sdk-go/v2/clients/naming_client"
	"github.com/nacos-group/nacos-sdk-go/v2/model"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
)

var _ registry.Watcher = (*watcher)(nil)

type watcher struct {
	serviceName    string
	clusters       []string
	groupName      string
	ctx            context.Context
	cancel         context.CancelFunc
	subscribeChan  chan []model.Instance
	cli            naming_client.INamingClient
	kind           string
	subscribeParam *vo.SubscribeParam
}

func newWatcher(ctx context.Context, cli naming_client.INamingClient, serviceName, groupName, kind string, clusters []string) (*watcher, error) {
	w := &watcher{
		serviceName:   serviceName,
		clusters:      clusters,
		groupName:     groupName,
		cli:           cli,
		kind:          kind,
		subscribeChan: make(chan []model.Instance, 1),
	}
	w.ctx, w.cancel = context.WithCancel(ctx)

	w.subscribeParam = &vo.SubscribeParam{
		ServiceName: serviceName,
		Clusters:    clusters,
		GroupName:   groupName,
		SubscribeCallback: func(instances []model.Instance, err error) {
			w.subscribeChan <- instances
		},
	}
	e := w.cli.Subscribe(w.subscribeParam)
	return w, e
}

func (w *watcher) Next() ([]*registry.ServiceInstance, error) {
	select {
	case <-w.ctx.Done():
		return nil, w.ctx.Err()
	case instances := <-w.subscribeChan:
		items := make([]*registry.ServiceInstance, 0, len(instances))
		for _, in := range instances {
			kind := w.kind
			if k, ok := in.Metadata["kind"]; ok {
				kind = k
			}
			items = append(items, &registry.ServiceInstance{
				ID:        in.InstanceId,
				Name:      in.ServiceName,
				Version:   in.Metadata["version"],
				Metadata:  in.Metadata,
				Endpoints: []string{fmt.Sprintf("%s://%s:%d", kind, in.Ip, in.Port)},
			})
		}
		return items, nil
	}
}

func (w *watcher) Stop() error {
	err := w.cli.Unsubscribe(w.subscribeParam)
	w.cancel()
	return err
}