#!/usr/bin/env bash

set -euo pipefail

if [[ "$#" -ne 1 ]]; then
  echo "Usage: $0 FILE" >&2
  exit 1
fi

xml_file="$1"

parsed=$(./parser "$xml_file" | xmllint --c14n --pretty 1 -)
original=$(xmllint --c14n --pretty 1 "$xml_file")

# diff and ignore white space
diff -w <(echo ${original}) <(echo ${parsed})
