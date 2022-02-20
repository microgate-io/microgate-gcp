package queue

import (
	"context"
	"sync"

	"cloud.google.com/go/pubsub"
	"github.com/emicklei/tre"
	"github.com/emicklei/xconnect"
	apiqueue "github.com/microgate-io/microgate-lib-go/v1/queue"
	mlog "github.com/microgate-io/microgate/v1/log"
)

// QueueingServiceImpl is used to publish or receive messages via GCP Pub/Sub
// https://pkg.go.dev/cloud.google.com/go/pubsub#Topic
type QueueingServiceImpl struct {
	apiqueue.UnimplementedQueueingServiceServer
	topics        map[string]*pubsub.Topic        // one service can publish to multiple topics
	subscriptions map[string]*pubsub.Subscription // one service can receive from multiple subscriptions
	pubsubClient  *pubsub.Client
	mutex         *sync.RWMutex
}

func NewQueueingServiceImpl(config xconnect.Document) (*QueueingServiceImpl, error) {
	p, err := config.FindString("project_id")
	if err != nil {
		return nil, tre.New(err, "missing config paramater [project_id]")
	}
	ctx := context.Background()
	pc, err := pubsub.NewClient(ctx, p)
	if err != nil {
		return nil, tre.New(err, "cannot create pubsub client", "project_id", p)
	}
	return &QueueingServiceImpl{
		topics:        map[string]*pubsub.Topic{},
		subscriptions: map[string]*pubsub.Subscription{},
		pubsubClient:  pc,
		mutex:         new(sync.RWMutex)}, nil
}

func (s *QueueingServiceImpl) Publish(ctx context.Context, req *apiqueue.PublishRequest) (*apiqueue.PublishResponse, error) {
	mlog.Debugw(ctx, "Publish", "req", req)
	s.mutex.RLock()
	top, ok := s.topics[req.TopicId]
	s.mutex.RUnlock()
	if !ok {
		// first time, create the topic (reference to the GCP resource)
		s.mutex.Lock()
		top = s.pubsubClient.Topic(req.TopicId)
		s.topics[req.TopicId] = top
		s.mutex.Unlock()
	}
	_ = top.Publish(ctx, &pubsub.Message{Data: req.Message})
	return new(apiqueue.PublishResponse), nil
}
func (s *QueueingServiceImpl) Subscribe(ctx context.Context, req *apiqueue.SubscribeRequest) (*apiqueue.SubscribeResponse, error) {
	mlog.Debugw(ctx, "Subscribe", "req", req)
	s.mutex.RLock()
	_, ok := s.subscriptions[req.SubscriptionId]
	s.mutex.RUnlock()
	if !ok {
		// first time, create the subscription (reference to the GCP resource)
		s.mutex.Lock()
		sub := s.pubsubClient.Subscription(req.SubscriptionId)
		s.subscriptions[req.SubscriptionId] = sub
		s.mutex.Unlock()
		// start receiving
		err := sub.Receive(context.Background(), func(ctx context.Context, m *pubsub.Message) {
			mlog.Debugw(ctx, "received message", "subscription", req.SubscriptionId, "data", m.Data)
			m.Ack()
		})
		if err != nil {
			return nil, tre.New(err, "cannot create receive from subscription", "subscription", req.SubscriptionId)
		}
	}
	return new(apiqueue.SubscribeResponse), nil
}

func (s *QueueingServiceImpl) Shutdown() error { return nil }
