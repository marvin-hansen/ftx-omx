# Copyright (c) 2021-2022. Marvin Hansen | marvin.hansen@gmail.com

# bin/bash
set -o errexit
set -o nounset
set -o pipefail

docker run --name omx --network my-net -p 80:80 omx:latest
