PROJECT_DIR=~/go/src/github.com/JackTJC/backend/

cd ${PROJECT_DIR}/proto
protoc --go_out=../pb_gen --go-grpc_out=../pb_gen \
 --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative\
  *.proto