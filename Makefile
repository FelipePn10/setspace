# Makefile

.PHONY: run-orders run-kitchen gen

run-orders:
	@go run github.com/FelipePn10/setspace/services/orders

run-kitchen:
	@go run github.com/FelipePn10/setspace/services/kitchen

gen:
	@protoc \
		--proto_path=protobuf "protobuf/orders.proto" \
		--go_out=services/common/genproto/orders --go_opt=paths=source_relative \
		--go-grpc_out=services/common/genproto/orders --go-grpc_opt=paths=source_relative
