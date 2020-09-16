#!/usr/bin/env bash
# A simple script as example

set -o errexit
set -o nounset
set -o pipefail

FIRST_ARGUMENT="$1"

function dummy() {
  echo $FIRST_ARGUMENT
}

dummy

function handle_exit() {
  # Add cleanup code here
  echo "cleanup block"
  exit 0
}
trap handle_exit 0 SIGHUP SIGINT SIGQUIT SIGABRT SIGTERM