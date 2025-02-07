#!/usr/bin/env bash

echo "[lint] checking go files"

task lint

# shellcheck disable=SC2181
if [ $? -ne 0 ]; then
  echo "[lint] found issues in go files"
  exit 1
fi

exit 0
