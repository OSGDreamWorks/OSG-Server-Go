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

go build -o "$PRODIR"/bin/dbserver		"$PRODIR"/src/runtime/dbserver/main.go

go build -o "$PRODIR"/bin/authserver		"$PRODIR"/src/runtime/authserver/main.go

go build -o "$PRODIR"/bin/loginserver		"$PRODIR"/src/runtime/loginserver/main.go

#go build -o "$PRODIR"/bin/fightserver		"$PRODIR"/src/runtime/fightserver/main.go

go build -o "$PRODIR"/bin/gameserver		"$PRODIR"/src/runtime/gameserver/main.go

export GOPATH="$OLDGOPATH"

cd "$PRODIR"tools/build/

cp -rf "$PRODIR"etc "$PRODIR"bin/etc

echo "build successfully"