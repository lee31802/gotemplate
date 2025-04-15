# 获取pb文件夹下所有的proto文件
PROTO_FILES := $(shell find pb -name '*.proto')

# 定义protoc命令
PROTOC_CMD := protoc --go_out=./

# 生成每个proto文件对应的目标文件
PROTO_GEN_TARGETS := $(PROTO_FILES:.proto=.pb.go)

# 定义proto目标，执行protoc命令生成pb.go文件
proto: $(PROTO_GEN_TARGETS)

# 规则：从proto文件生成pb.go文件
%.pb.go: %.proto
	$(PROTOC_CMD) --go_opt=M$<=$(dir $<) $<

# 清理生成的pb.go文件
cleanproto:
	rm -f $(PROTO_GEN_TARGETS)