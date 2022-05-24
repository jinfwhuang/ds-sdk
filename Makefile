clean-proto:
	rm -rf build/github.com

proto:
	rm -rf build/github.com/jinfwhuang/ds-toolkit/proto/
	protoc -I=./third_party/googleapis -I=./proto --go_out=build --go-grpc_out=build proto/**/*.proto
	cp -r build/github.com/jinfwhuang/ds-toolkit/proto/  proto

#	rm -rf build/github.com/jinfwhuang/ds-toolkit/proto
#	protoc -I=./third_party/googleapis -I=./proto/identity --go_out=build --go-grpc_out=build proto/identity/login.proto
#	cp build/github.com/jinfwhuang/ds-toolkit/proto/identity/*  proto/identity/

generate-eth-code:
	cp -r npm-pkg/contracts-identity/artifacts/abi go-pkg/user
	go run github.com/ethereum/go-ethereum/cmd/abigen --abi "go-pkg/user/abi/UserRegistry.json" --pkg user --type UserRegistry --out "go-pkg/user/user-registry.go"

hardhat-setup:
	cd npm-pkg/contracts-identity
	NETWORK=hardhat npx hardhat run scripts/hardhat-setup.ts

# Targets that are not associated with explicit filename or filedir
.PHONY: clean proto generate-eth-code hardhat-setup
