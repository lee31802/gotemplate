# 获取pb文件夹下所有的proto文件
PROTO_FILES := $(shell find pb -name '*.proto')

# 定义protoc命令
PROTOC_CMD := protoc --go_out=paths=source_relative:. --go-grpc_out=paths=source_relative:.

# 生成每个proto文件对应的目标文件
PROTO_GEN_TARGETS := $(PROTO_FILES:.proto=.pb.go)
GRPC_GEN_TARGETS := $(PROTO_FILES:.proto=.pb.grpc.go)

# 定义proto目标，执行protoc命令生成pb.go和pb.grpc.go文件
proto: $(PROTO_GEN_TARGETS) $(GRPC_GEN_TARGETS)

# 规则：从proto文件生成pb.go文件
%.pb.go: %.proto
	$(PROTOC_CMD) --go_opt=M$<=$(dir $<) $<

# 规则：从proto文件生成pb.grpc.go文件
%.pb.grpc.go: %.proto
	$(PROTOC_CMD) --go-grpc_opt=M$<=$(dir $<) $<

# 清理生成的pb.go和pb.grpc.go文件
cleanproto:
	rm -f $(PROTO_GEN_TARGETS) $(GRPC_GEN_TARGETS)