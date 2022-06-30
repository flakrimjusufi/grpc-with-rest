# This is used only in ci builds to run tests.

CLUSTER=$1
NAMESPACE=$2

bazelisk test \
    --platform_suffix="bazel-test" \
    --@io_bazel_rules_go//go/config:race \
    --define cluster=$CLUSTER \
    --define namespace=$NAMESPACE \
    --test_tag_filters=fast \
    --build_tag_filters=fast \
    --test_output=errors \
    --nocache_test_results \
    //...
