#!/bin/bash



# Function to ensure necessary binaries are on PATH
ensure_binaries_on_path() {
    if ! command -v buf &> /dev/null; then
        if [ -n "$(go env GOBIN)" ]; then
            export PATH="$(go env GOBIN):${PATH}"
        fi
    fi

    if ! command -v grpcurl &> /dev/null; then
        if [ -n "$(go env GOBIN)" ]; then
            export PATH="$(go env GOBIN):${PATH}"
        fi
    fi

    if ! command -v protoc-gen-go &> /dev/null; then
        if [ -n "$(go env GOPATH)" ]; then
            export PATH="$(go env GOPATH)/bin:${PATH}"
        fi
    fi

    if ! command -v protoc-gen-connect-go &> /dev/null; then
        if [ -n "$(go env GOPATH)" ]; then
            export PATH="$(go env GOPATH)/bin:${PATH}"
        fi
    fi

    if ! command -v protoc-gen-es &> /dev/null; then
        if [ -n "$(npm bin)" ]; then
            export PATH="$(pwd)/node_modules/.bin:${PATH}"
        fi
    fi
}


# Function to install dependencies
install_deps() {
    # Install dependencies with specific versions if not already installed
    if ! command -v buf &> /dev/null; then
        go install github.com/bufbuild/buf/cmd/buf@v1.29.0
    fi

    if ! command -v protoc-gen-go &> /dev/null; then
        go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.29.0
    fi

    if ! command -v protoc-gen-connect-es &> /dev/null; then
        go install connectrpc.com/connect/cmd/protoc-gen-connect-go@v1.15.0
    fi

    if ! command -v protoc-gen-es &> /dev/null; then
      npm install --save-dev @bufbuild/buf @connectrpc/protoc-gen-connect-es @bufbuild/protoc-gen-es
      npm install @connectrpc/connect @connectrpc/connect-web @bufbuild/protobuf
    fi
}

# Function to build protocol buffers
build_proto() {
  ensure_binaries_on_path
  install_deps
  cd api/proto
  buf generate
}

# Main script starts here
case "$1" in
    build-proto)
        build_proto
        ;;
    install-deps)
        install_deps
        ;;
    *)
        echo "Usage: $0 {build-proto|install-deps}"
        exit 1
        ;;
esac
