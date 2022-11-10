# duck-blocks
The underlying blockchain to power the Duck Dynasty crypto ecosystem.


## Development

### Go Module Management
Initialize the repository as a go module. Not required any mroe since `go.mod` is preserved in this repository.
```bash
go mod init github.com/duck-dynasty/duck-blocks
```

Clean up packages included in the go module and download any required modules.
```bash
go mod tidy
go get
```

### Build & Run
Run from the repository root to compile and run the program.
```bash
go build .
./duck-blocks
```

### Testing
Run the following from the repository root to run unit tests on all modules and submodules.
```bash
go test --cover ./...
```