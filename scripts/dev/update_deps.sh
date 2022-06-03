# Copyright (c) 2021. Marvin Friedrich Lars Hansen. All Rights Reserved. Contact: marvin.hansen@gmail.com

# bin/bash
set -o errexit
set -o nounset
set -o pipefail

command go get -u ./...

command go mod tidy

# Convert mod dependencies into bazel dependencies
command bazel run //:gazelle -- update-repos -from_file=go.mod

# Update all build files & dependencies
command bazel run //:gazelle

# Regenerate all *.proto.pb file
#bazel run //proto:link

# Build all sources
command bazel build //:build
