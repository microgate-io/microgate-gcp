module github.com/emicklei/microgate-io/microgate-gcp

go 1.17

require (
	cloud.google.com/go/pubsub v1.3.1
	github.com/emicklei/tre v1.2.0
	github.com/emicklei/xconnect v0.10.1
	github.com/microgate-io/microgate v1.0.0
	github.com/microgate-io/microgate-lib-go v1.0.0
)

require (
	cloud.google.com/go v0.100.2 // indirect
	cloud.google.com/go/compute v1.2.0 // indirect
	cloud.google.com/go/iam v0.2.0 // indirect
	cloud.google.com/go/kms v1.3.0 // indirect
	github.com/blendle/zapdriver v1.3.1 // indirect
	github.com/golang/groupcache v0.0.0-20200121045136-8c9f03a8e57e // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/google/go-cmp v0.5.7 // indirect
	github.com/googleapis/gax-go/v2 v2.1.1 // indirect
	github.com/iancoleman/strcase v0.2.0 // indirect
	github.com/jhump/protoreflect v1.10.1 // indirect
	github.com/processout/grpc-go-pool v1.2.1 // indirect
	github.com/vgough/grpc-proxy v0.0.0-20210913231538-71832b651269 // indirect
	go.opencensus.io v0.23.0 // indirect
	go.uber.org/atomic v1.7.0 // indirect
	go.uber.org/multierr v1.6.0 // indirect
	go.uber.org/zap v1.19.1 // indirect
	golang.org/x/net v0.0.0-20220127200216-cd36cc0744dd // indirect
	golang.org/x/oauth2 v0.0.0-20211104180415-d3ed0bb246c8 // indirect
	golang.org/x/sync v0.0.0-20210220032951-036812b2e83c // indirect
	golang.org/x/sys v0.0.0-20220209214540-3681064d5158 // indirect
	golang.org/x/text v0.3.7 // indirect
	google.golang.org/api v0.69.0 // indirect
	google.golang.org/appengine v1.6.7 // indirect
	google.golang.org/genproto v0.0.0-20220211171837-173942840c17 // indirect
	google.golang.org/grpc v1.44.0 // indirect
	google.golang.org/protobuf v1.27.1 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
)

replace (
	github.com/microgate-io/microgate v1.0.0 => ../microgate
	github.com/microgate-io/microgate-lib-go v1.0.0 => ../microgate-lib-go
)
