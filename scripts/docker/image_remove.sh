# Copyright (c) 2021-2022. Marvin Hansen | marvin.hansen@gmail.com

# bin/bash
# set -o errexit # should not end script when docker returns and error i.e. no container found
set -o nounset
set -o pipefail

echo "Remove OMX container "

command docker container stop omx
echo "* OMX Container stopped! "

command docker container rm omx
echo "* OMX Container removed! "