.PHONY: proxy fmt lint

proxy:
	@go env -w GO111MODULE="on"
	@go env -w GOPROXY="https://goproxy.cn,direct"

fmt:
	@find . -name '*.go' -not -path "./vendor/*" -not -name "*.pb.go" | xargs gofumpt -w -s -extra
	@find . -name '*.go' -not -path "./vendor/*" -not -name "*.pb.go" | xargs -n 1 -I {} -t goimports-reviser -file-path {} -local "github.com/sliveryou" project-name "github.com/sliveryou/go-tool/" -rm-unused
	@find . -name '*.sh' -not -path "./vendor/*" | xargs shfmt -w -s -i 2 -ci -bn -sr

lint:
	@golangci-lint run ./...
