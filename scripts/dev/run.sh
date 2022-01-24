# Copyright (c) 2021-2022. Marvin Hansen | marvin.hansen@gmail.com

# bin/bash
set -o errexit
set -o nounset
set -o pipefail

 bazel run //:server_noarch
