#!/bin/bash

set -euo pipefail

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null && pwd )"

# Ensure that `bashttpd` can run the `pung` bin
export PATH="$PATH:$DIR/bin"

# Need to cd into it, because bashttpd searches relative
cd "$DIR/server"
socat TCP4-LISTEN:8080,fork "EXEC:$DIR/server/bashttpd"