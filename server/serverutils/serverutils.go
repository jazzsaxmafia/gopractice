package serverutils

import (
	"context"
	"errors"
	"log"
	pb "productinfo/proto/productinfopb"

	wrapper "github.com/golang/protobuf/ptypes/wrappers"
	"github.com/google/uuid"
)

// server is used to implement ecommerce/product_info.
type Server struct {
	productMap map[string]*pb.Product
}

// AddProduct implements ecommerce.AddProduct
func (s *Server) AddProduct(ctx context.Context, in *pb.Product) (*wrapper.StringValue, error) {
	out, err := uuid.NewUUID()
	if err != nil {
		log.Fatal(err)
	}
	in.Id = out.String()
	if s.productMap == nil {
		s.productMap = make(map[string]*pb.Product)
	}
	s.productMap[in.Id] = in
	log.Printf("New product added - ID : %s, Name : %s", in.Id, in.Name)
	return &wrapper.StringValue{Value: in.Id}, nil
}

// GetProduct implements ecommerce.GetProduct
func (s *Server) GetProduct(ctx context.Context, in *wrapper.StringValue) (*pb.Product, error) {
	value, exists := s.productMap[in.Value]
	if exists {
		log.Printf("New product retrieved - ID : %s", in)
		return value, nil
	}

	return nil, errors.New("Product does not exist for the ID" + in.Value)
}
