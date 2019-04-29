package main

import (
	pb "consignment-service/proto/consignment"
	"context"
	"google.golang.org/grpc"
	"log"
	"net"
)

const (
	port = ":50051"
)

type IRepository interface {
	Create(*pb.Consignment) (*pb.Consignment, error)
	GetAll() []*pb.Consignment
}

// Repository - 模拟一个数据库，我们会在此后使用真正的数据库替代他
type Repository struct {
	consignments []*pb.Consignment
}

func (repo *Repository) Create(consignment *pb.Consignment) (*pb.Consignment, error) {
	updated := append(repo.consignments, consignment)
	repo.consignments = updated
	return consignment, nil
}

func (repo *Repository) GetAll() []*pb.Consignment {
	return repo.consignments
}
// service要实现在proto中定义的所有方法。当你不确定时
// 可以去对应的*.pb.go文件里查看需要实现的方法及其定义
type service struct {
	repo IRepository
}

// CreateConsignment - 在proto中，我们只给这个微服务定一个了一个方法
// 就是这个CreateConsignment方法，它接受一个context以及proto中定义的
// Consignment消息，这个Consignment是由gRPC的服务器处理后提供给你的
func (s *service) CreateConsignment(ctx context.Context, req *pb.Consignment) (*pb.Response, error) {
	// 保存我们的consignment
	consignment, err := s.repo.Create(req)
	if err != nil {
		return nil, err
	}
	// 返回的数据也要符合proto中定义的数据结构
	return &pb.Response{Created: true, Consignment: consignment}, nil
}

func (s *service) GetConsignments(ctx context.Context, req *pb.GetRequest) (*pb.Response, error) {
	consignments := s.repo.GetAll()
	return &pb.Response{Consignments: consignments}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalln(err)
	}

	repo := &Repository{}
	s := grpc.NewServer()
	pb.RegisterShippingServiceServer(s, &service{repo: repo})
	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
