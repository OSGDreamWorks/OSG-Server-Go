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
	echo "please run install-require.bat first."
    exit 1
fi

cd "$PRODIR"

export GOPATH="$PRODIR:$PKGDIR"

go run "$PRODIR"/src/runtime/dbserver/main.go

Sleep 3
go run "$PRODIR"/src/runtime/authserver/main.go

Sleep 3
go run "$PRODIR"/src/runtime/loginserver/main.go

Sleep 3
#go run "$PRODIR"/src/runtime/fightserver/main.go

Sleep 3
go run "$PRODIR"/src/runtime/gameserver/main.go

export GOPATH="$OLDGOPATH"

cd "$PRODIR"tools/build/

echo "debug successfully"