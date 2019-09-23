#!/bin/bash

set -e

CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-extldflags "-static"'
THIS_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
cd $THIS_DIR/..

GIT_HASH="$(git rev-list --max-count=1 HEAD | cut -c1-8)"


docker build . -t pwd:$GIT_HASH
docker images
docker push pwd:$GIT_HASH