module github.com/hzxiao/microservice-tutorial/consignment-service

go 1.12

require (
	github.com/golang/protobuf v1.3.1
	github.com/hzxiao/microservice-tutorial/vessel-service v0.0.0-20190501060753-f254ba09a622
	github.com/micro/go-micro v1.1.0
	github.com/tidwall/pretty v0.0.0-20190325153808-1166b9ac2b65 // indirect
	github.com/xdg/scram v0.0.0-20180814205039-7eeb5667e42c // indirect
	github.com/xdg/stringprep v1.0.0 // indirect
	go.mongodb.org/mongo-driver v1.0.1
	golang.org/x/net v0.0.0-20190501004415-9ce7a6920f09
)

replace github.com/hzxiao/microservice-tutorial => ../
