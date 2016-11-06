go get -u github.com/golang/protobuf/{proto,protoc-gen-go}
protoc -I api/ api/api.proto --go_out=plugins=grpc:api

protoc -I /usr/local/include -I ~/workspace/vendoring/src/vendor/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis -I api/ api/api.proto --go_out=plugins=grpc:api
protoc -I /usr/local/include -I ~/workspace/vendoring/src/vendor/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis -I api/ api/api.proto --grpc-gateway_out=logtostderr=true:api 
