#!/usr/bin/env bash

BUCKET='pub.astranetwork.com'
OS="$(uname -s)"

usage () {
    cat << EOT
Usage: $0 [option] command

Options:
   -d          download all the binaries
   -h          print this help
Note: Arguments must be passed at the end for ./astra to work correctly.
For instance: ./astra.sh balances <one-address> --node=https://rpc.s0.p.astranetwork.com/

EOT
}

set_download () {
    local rel='mainnet'
    case "$OS" in
	Darwin)
	    FOLDER=release/darwin-x86_64/${rel}/
	    BIN=( astra libbls384_256.dylib libcrypto.1.0.0.dylib libgmp.10.dylib libgmpxx.4.dylib libmcl.dylib )
	    ;;
	Linux)
	    FOLDER=release/linux-x86_64/${rel}/
	    BIN=( astra )
	    ;;
	*)
	    echo "${OS} not supported."
	    exit 2
	    ;;
    esac
}

do_download () {
    # download all the binaries
    for bin in "${BIN[@]}"; do
	rm -f ${bin}
	curl http://${BUCKET}.s3.amazonaws.com/${FOLDER}${bin} -o ${bin}
    done
    chmod +x astra
}

while getopts "dh" opt; do
    case ${opt} in
        d)
            set_download
            do_download
            exit 0
            ;;
        h|*)
            usage
            exit 1
            ;;
    esac
done

shift $((OPTIND-1))

if [ "$OS" = "Linux" ]; then
    ./astra "$@"
else
    DYLD_FALLBACK_LIBRARY_PATH="$(pwd)" ./astra "$@"
fi
