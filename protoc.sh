protoc -I proto proto/product_info.proto --go_out=plugins=grpc:./proto
cp proto/product_info.pb.go service/ecommerce
cp proto/product_info.pb.go client/ecommerce
