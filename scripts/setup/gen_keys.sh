# Copyright (c) 2021-2022. Marvin Hansen | marvin.hansen@gmail.com

# bin/bash
set -o errexit
set -o nounset
set -o pipefail

command go run src/cmd/*.go

command go run src/cmd/*.go

command go run src/cmd/*.go