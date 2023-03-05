run:
	go run main.go
schema:
	@read -p "Enter Schema Name: " name; \
		go run -mod=mod entgo.io/ent/cmd/ent new $$name
generate:
	go generate ./ent
mac:
	env GOOS=darwin GOARCH=amd64 go build -o ./bin/mac/biz
linux:
	env GOOS=linux GOARCH=amd64 go build -o ./bin/linux/biz
