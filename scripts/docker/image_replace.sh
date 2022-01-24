# Copyright (c) 2021-2022. Marvin Hansen | marvin.hansen@gmail.com

# bin/bash
set -o errexit
set -o nounset
set -o pipefail

echo "Replacing OMX with latest version "

command docker container stop omx
echo "* OMX Container stopped! "

command docker container rm omx
echo "* OMX Container removed! "

command docker pull gcr.io/future-309012/omx:latest
echo "* latest OMX Container pulled! "

command docker run -d --name omx --network my-net -p 80:80 gcr.io/future-309012/omx:latest
echo "* latest OMX Container started! "

echo " "

echo "* Currently running containers: "
command docker container ls
echo " "

echo "OMX container replace complete "
