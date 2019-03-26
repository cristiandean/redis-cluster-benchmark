.PHONY: setup-single-node setup-cluster run 


setup-single-node:
	@echo "initializing redis single node..."
	redis-server --port 5555

setup-cluster:
	@echo "initializing redis-cluster..."
	docker run -e 'IP=0.0.0.0' -p 5000-5002:5000-5002/tcp -p 6379:6479 -p 7000-7007:7000-7007/tcp grokzen/redis-cluster:latest

run:
	go run main.go
