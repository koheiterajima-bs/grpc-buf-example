#!/bin/bash
set -euo pipefail

go install github.com/bufbuild/buf/cmd/buf@v1.7.0
go install honnef.co/go/tools/cmd/staticcheck@v0.3.3
go install github.com/google/ko@v0.12.0
go install github.com/fullstorydev/grpcurl/cmd/grpcurl@v1.8.7

