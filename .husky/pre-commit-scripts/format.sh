#!/usr/bin/env bash

echo "[format] formatting code"

task format

# shellcheck disable=SC2181
if [ $? -ne 0 ]; then
  echo "[format] found issues formatting go files"
  exit 1
fi

git add --all
exit 0
