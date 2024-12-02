#/bin/sh

if [ "$1" ]
then
	# 进入 protobuf/proto 目录
	cd protobuf/proto

	# 生成 pb 程序
	protoc --go_out=. $1.proto

	# 返回
	cd -

	# 注入tag
	protoc-go-inject-tag -input=protobuf/pb/$1.pb.go
else
	echo "Wrong parameter. It should be like this, ./build_pb.sh common"
fi
