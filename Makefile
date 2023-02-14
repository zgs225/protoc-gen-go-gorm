GO_FILES = $(shell find . -name '*.go' -and -not -name '*.pb.go' -and -not -name '*.gorm.go')
GOMOD_FILES = go.mod go.sum
PROTOBUF_FILES = $(shell find ./example -name '*.proto')
ANNOTATION_PROTOBUF_FILES = $(shell find ./proto -name '*.proto')

install: $(GOMOD_FILES) $(GO_FILES) gen-annotation
	@echo "Installing..."
	go mod tidy
	go install

gen-annotation: $(ANNOTATION_PROTOBUF_FILES)
	protoc --go_out=. --go_opt=paths=source_relative $(ANNOTATION_PROTOBUF_FILES)

example: install $(PROTOBUF_FILES)
	protoc --go_out=. --go_opt=paths=source_relative --go-gorm_out=. --go-gorm_opt=paths=source_relative -I ./proto/ -I . $(PROTOBUF_FILES)

clean:
	find . \( -name '*.pb.go' -or -name '*.gorm.go' \) -exec rm -f {} \;
