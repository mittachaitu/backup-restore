
# Get the currently used golang install path (in GOPATH/bin, unless GOBIN is set)
GOBIN := $(or $(shell go env GOBIN 2>/dev/null), $(shell go env GOPATH 2>/dev/null)/bin)

# find or download controller-gen
controller-gen:
ifneq ($(shell controller-gen --version 2> /dev/null), Version: v0.2.9)
	@(cd /tmp; GO111MODULE=on go get sigs.k8s.io/controller-tools/cmd/controller-gen@v0.2.9)
CONTROLLER_GEN=$(GOBIN)/controller-gen
else
CONTROLLER_GEN=$(shell which controller-gen)
endif

.PHONY: kubegen
# code generation for custom resources
kubegen:
	./hack/update-codegen.sh

.PHONY: generated_files
generated_files: kubegen protobuf
