#!/bin/sh

DIR=$1
IMAGE=$2
shift 2

_usage() { echo "usage: $0 DIR IMAGE [ARGS...]"; exit 1; }
test -n "$DIR" || _usage
test -n "$IMAGE" || _usage

# Library path from build container
LIBDIR=/usr/local/lib

set -eux

docker build -t $IMAGE -f Dockerfile.build "$DIR"
TMP=$(docker run -d -v "$PWD/../_mod:/_mod:ro" --entrypoint=/osmosis/scripts/make-build.sh $IMAGE)
docker wait $TMP
docker commit $TMP $IMAGE
docker rm $TMP
docker build -t $IMAGE -f Dockerfile.run "$@" "$DIR"
