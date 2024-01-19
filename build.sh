MODULE_NAME="github.com/stevenmaharaj/gateway_draft1"

protoc --go_out=. --go_opt=module=$MODULE_NAME \
--go-grpc_out=. --go-grpc_opt=module=$MODULE_NAME \
proto/*.proto
