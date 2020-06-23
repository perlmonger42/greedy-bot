# How to use this script:
#   fswatch -o **/*.go | xargs -n1 -I{} test_all.sh
# For some reason, I can't just put that into this script.
# fswatch never seems to trigger when I do.
clear;
go test


### #!/usr/bin/env bash
### SCRIPT_PATH="$( cd "$(dirname "$0")" >/dev/null 2>&1 ; pwd -P )"
### echo invoked: $SCRIPT_PATH "$@"
### if [[ $1 == "test" ]]; then
###   clear;
###   go test
### else
###   echo "fswatch -o **/*.go | xargs -n1 -I{} $SCRIPT_PATH test"
###   fswatch -o **/*.go | xargs -n1 -I{} $SCRIPT_PATH test
### fi
