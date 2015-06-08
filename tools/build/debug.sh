#!/usr/bin/env bash

set -e

if [ ! -f debug.sh ]; then
    echo 'debug.sh must be run within its container folder' 1>&2
    exit 1
fi

PRODIR=`pwd`/../../
PKGDIR=`pwd`/../../3rdpkg/

OLDGOPATH="$GOPATH"

if [ ! -d $PKGDIR ]; then
    mkdir $PKGDIR
fi

export GOPATH="$PKGDIR"
cd $PKGDIR
go get github.com/googollee/go-socket.io
go get github.com/garyburd/redigo/redis
go get github.com/go-sql-driver/mysql
go get github.com/Shopify/go-lua
go get code.google.com/p/goprotobuf/proto
go get code.google.com/p/snappy-go/snappy
cd "$PRODIR"

export GOPATH="$PRODIR:$PKGDIR"

go run "$PRODIR"/src/runtime/dbserver/main.go

Sleep 3
go run "$PRODIR"/src/runtime/authserver/main.go

Sleep 3
go run "$PRODIR"/src/runtime/gateserver/main.go

Sleep 3
go run "$PRODIR"/src/runtime/fightserver/main.go

Sleep 3
go run "$PRODIR"/src/runtime/gameserver/main.go

export GOPATH="$OLDGOPATH"

cd "$PRODIR"tools/build/

echo "debug successfully"