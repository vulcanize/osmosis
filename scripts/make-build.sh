#!/bin/sh

# Library path from build container
LIBDIR=/usr/local/lib

exec env BUILD_TAGS=linux \
     env CGO_LDFLAGS="-L${LIBDIR} -Wl,-rpath,${LIBDIR}" \
     make build
