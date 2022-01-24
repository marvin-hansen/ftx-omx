# Copyright (c) 2021-2022. Marvin Hansen | marvin.hansen@gmail.com

# bin/bash
# set -o errexit
set -o nounset
set -o pipefail

echo "## (!)(!) ## Hard Reset: Delete, rebuild DB & replace OMX with latest version! ## (!)(!) ##"

command docker container stop omx
echo "* OMX Container stopped! "

command docker container rm omx
echo "* OMX Container removed! "

command /bin/bash scripts/db/db_configure.sh
echo "* New DB Container configured! "

# Uncomment & edit the section below when pulling from registry & deploying pulled image
# command docker pull gcr.io/future-309012/omx:latest
# echo "* Latest OMX Container pulled! "

# command docker run -d --name omx --network my-net -p 80:80 gcr.io/future-309012/omx:latest
# echo "* Latest OMX Container started! "

command docker build -t omx:latest .
echo "* Latest OMX Container build! "

command docker run -d --name omx --network my-net -p 80:80 omx:latest
echo "* Latest OMX Container started! "

echo " "

echo "* Currently running containers: "
command docker container ls
echo " "

echo "Hard reset complete "
