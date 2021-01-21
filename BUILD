load("@bazel_gazelle//:def.bzl", "gazelle")

# gazelle:proto disable_global
# gazelle:exclude tools/deb-tools/
gazelle(
    name = "gazelle",
    external = "vendored",
    prefix = "kope.io/etcd-manager",
)

load("//images:etcd.bzl", "supported_etcd_arch_and_version")

[
    genrule(
        name = "etcd-v{version}-linux-{arch}_{bin}".format(
            version = version,
            arch = arch,
            bin = bin,
        ),
        srcs = ["@etcd_{version}_{arch}_tar//file".format(
            version = version,
            arch = arch,
            bin = bin,
        )],
        outs = ["etcd-v{version}-linux-{arch}/{bin}".format(
            version = version,
            arch = arch,
            bin = bin,
        )],
        cmd = "tar -x -z --no-same-owner -f ./$(location @etcd_{version}_{arch}_tar//file) etcd-v{version}-linux-{arch}/{bin} && mv etcd-v{version}-linux-{arch}/{bin} \"$@\"".format(
            version = version,
            arch = arch,
            bin = bin,
        ),
        visibility = ["//visibility:public"],
    )
    for (arch, version) in supported_etcd_arch_and_version()
    for (bin) in [
        "etcd",
        "etcdctl",
    ]
]

load("@io_bazel_rules_go//go:def.bzl", "nogo")

nogo(
    name = "nogo",
    config = "analysis-config.json",
    visibility = ["//visibility:public"],
    deps = [
        "@org_golang_x_tools//go/analysis/passes/assign:go_tool_library",
        "@org_golang_x_tools//go/analysis/passes/atomic:go_tool_library",
        "@org_golang_x_tools//go/analysis/passes/bools:go_tool_library",
        "@org_golang_x_tools//go/analysis/passes/buildtag:go_tool_library",
        "@org_golang_x_tools//go/analysis/passes/loopclosure:go_tool_library",
        "@org_golang_x_tools//go/analysis/passes/lostcancel:go_tool_library",
        "@org_golang_x_tools//go/analysis/passes/nilfunc:go_tool_library",
        "@org_golang_x_tools//go/analysis/passes/nilness:go_tool_library",
        "@org_golang_x_tools//go/analysis/passes/printf:go_tool_library",
        # shadow is well-known to have too many false positives
        #"@org_golang_x_tools//go/analysis/passes/shadow:go_tool_library",
        "@org_golang_x_tools//go/analysis/passes/structtag:go_tool_library",
        "@org_golang_x_tools//go/analysis/passes/unreachable:go_tool_library",
        "@org_golang_x_tools//go/analysis/passes/unusedresult:go_tool_library",
    ],
)

