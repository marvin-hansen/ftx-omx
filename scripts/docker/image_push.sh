# Copyright (c) 2021-2022. Marvin Hansen | marvin.hansen@gmail.com

# bin/bash
set -o errexit
set -o nounset
set -o pipefail

# Push an image or a repository to a registry
# https://docs.docker.com/engine/reference/commandline/push/

command docker image tag omx:latest gcr.io/future-309012/omx:latest

command docker image push gcr.io/future-309012/omx:latest