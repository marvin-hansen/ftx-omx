# Copyright (c) 2022. Marvin Hansen | marvin.hansen@gmail.com

# bin/bash
set -o errexit
set -o nounset
set -o pipefail

echo ""
echo "==============================="
echo " Deploys OMX service:         "
echo "==============================="
echo ""

echo " * Build binary"
command bazel build //:build_cloud

echo " * Build image"
command  bazel run //:image

echo " * Publish image"
command bazel run //:push

# echo " * Deploy to kubernetes"
# command bazel run //:k8s-dev.apply

echo ""
echo "==============================="
echo " OMX Service Deployed!        "
echo "==============================="
echo ""
