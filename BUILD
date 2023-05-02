load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")
load("@bazel_gazelle//:def.bzl", "gazelle")

# gazelle:prefix github.com/flakrimjusufi/grpc-with-rest
gazelle(name = "gazelle")

go_library(
    name = "grpc-with-rest_lib",
    srcs = ["main.go"],
    importpath = "github.com/flakrimjusufi/grpc-with-rest",
    visibility = ["//visibility:private"],
    deps = [
        "//client:client_lib",
        "//helper",
        "//proto",
        "@com_github_grpc_ecosystem_grpc_gateway_v2//runtime",
        "@com_github_joho_godotenv//:godotenv",
        "@org_golang_google_grpc//:go_default_library",
    ],
)

go_binary(
    name = "grpc-with-rest",
    embed = [":grpc-with-rest_lib"],
    visibility = ["//visibility:public"],
)

gazelle(
    name = "gazelle-update-repos",
    args = [
        "-from_file=go.mod",
        "-to_macro=deps.bzl%go_dependencies",
        "-prune",
    ],
    command = "update-repos",
)
