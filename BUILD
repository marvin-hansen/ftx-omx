load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

package(default_visibility = ["//visibility:public"])

load("@bazel_gazelle//:def.bzl", "DEFAULT_LANGUAGES", "gazelle", "gazelle_binary")

gazelle_binary(
    name = "gazelle_binary",
    languages = DEFAULT_LANGUAGES,
    visibility = ["//visibility:public"],
)

# gazelle:prefix ftx-omx/
gazelle(
    name = "gazelle",
    gazelle = "//:gazelle_binary",
)

filegroup(
    name = "build_cloud",
    srcs = [
        "//:server",
    ],
    visibility = ["//visibility:public"],
)

filegroup(
    name = "build_local",
    srcs = [
        "//:server_noarch",
    ],
    visibility = ["//visibility:public"],
)

go_library(
    name = "lib",
    srcs = ["main.go"],
    importpath = "ftx-omx/",
    deps = [
        "//src/cfg",
        "//src/service",
    ],
)

go_binary(
    name = "server",
    embed = [":lib"],
    gc_linkopts = [
        "-linkmode",
        "external",
        "-extldflags",
        "-static",
        "-w",
        "-s",
    ],
    goarch = "amd64",
    goos = "linux",
    pure = "off",
    static = "on",
    visibility = ["//visibility:public"],
)

go_binary(
    name = "server_noarch",
    embed = [":lib"],
)

## ==== Container image ==== ##
# Define the container image to build (similar to the content in a Dockerfile)
load("@io_bazel_rules_docker//container:container.bzl", "container_image")

container_image(
    # bazel run //:image
    name = "image",
    base = "@scratch//image",
    cmd = ["./server"],
    files = [":server"],
    ports = [
        "80/tcp",  # Service endpoint.
    ],
    visibility = ["//visibility:public"],
)

# Define the push commando to publish container to registry
load("@io_bazel_rules_docker//container:container.bzl", "container_push")

# https://github.com/bazelbuild/rules_docker#container_push-1
container_push(
    # bazel run //:push
    name = "push",
    format = "OCI",
    image = ":image",
    registry = "gcr.io",
    repository = "future-309012/omx",
    tag = "latest",
    visibility = ["//visibility:public"],
)

## ==== Kubernetes deployment ==== ##
# Template for each environment in WORKSPACE@Line:140-190
#
