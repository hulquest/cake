#!/usr/bin/env bash

source ./hack/common.sh

ec=0
echo "Cake dir: ${CAKE_DIR}"
out="$(find ${CAKE_DIR}/ -name '*.go-e')"
if [[ ${out} ]]; then
  echo "FAIL: Found go-e files"
  echo $out
  ec=1
fi
exit ${ec}
