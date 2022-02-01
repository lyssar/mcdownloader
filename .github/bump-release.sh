#!/bin/bash

set -euo pipefail

if [ -f "README.md" ]
then
  echo "[VERSION BUMP] Found README.md and start bumping"
  sed -i "s/^\(.*\/version\-\)\([0-9]*\.[0-9]*\.[0-9]*\)\(\-.*\)/\1${NEXTVERSION}\3/" README.md
  sed -i "s/^\(.*MSDVERSION\=\"\)\([0-9]*\.[0-9]*\.[0-9]*\)\(\".*\)/\1${NEXTVERSION}\3/" README.md
fi
