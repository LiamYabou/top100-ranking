
PROJECTNAME=$(shell basename "${PWD}")

## test: run all the test of the project.
test: go-test

## compile: compile the instructions located the `./cmd` directory into the `./bin` directory.
compile: go-tidy go-compile-rpc-server-cmd

go-test:
	@echo "  > Testing..."
	@gotest -v ./...

go-tidy:
	@echo "  > Tidying the dependencies from the \`go.mod\` file..."
	@go mod tidy
	@echo "  > Done."

go-compile-rpc-server-cmd:
	@echo "  > Compiling the instruction of the rpc server..."
	@go build -o ./bin/ ./cmd/launch_rpc_server
	@echo "  > Done."

.PHONY: help
help: Makefile
	@echo
	@echo "  Choose a command to run in "${PROJECTNAME}": "
	@echo
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'
	@echo
