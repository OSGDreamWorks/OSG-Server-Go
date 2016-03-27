#!/usr/bin/env bash

set -e

if [ ! -f install-require.sh ]; then
    echo 'install-require.sh must be run within its container folder' 1>&2
    exit 1
fi

PRODIR=`pwd`/../
PKGDIR=`pwd`/../3rdpkg/

OLDGOPATH="$GOPATH"

if [ ! -d $PKGDIR ]; then
    mkdir $PKGDIR
fi

export GOPATH="$PKGDIR"
cd $PKGDIR
go get github.com/googollee/go-socket.io
go get github.com/garyburd/redigo/redis
go get github.com/go-sql-driver/mysql
go get github.com/yuin/gopher-lua
go get -u github.com/golang/protobuf/proto
go get -u github.com/golang/protobuf/protoc-gen-go
go get github.com/golang/snappy
cd "$PRODIR"

export GOPATH="$OLDGOPATH"

cd "$PRODIR"tools/

echo "install successfully"