server:
	go build -o server ./greeter_server
run: server
	./server
run-client:
	go run greeter_client/*
test-client:
	go test greeter_client/*
test-server:
	go test greeter_server/*
genmock:
	mockgen -source=helloworld/helloworld.pb.go -destination mock/helloworld.pb.go
