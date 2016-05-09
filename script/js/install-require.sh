#!/usr/bin/env bash

set -e

if [ ! -f install-require.sh ]; then
    echo 'install-require.sh must be run within its container folder' 1>&2
    exit 1
fi



echo "install successfully"