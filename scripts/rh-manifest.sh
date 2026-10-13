#!/usr/bin/env bash
#
# Generate rh-manifest.txt file.
# Run from repository root.
set -e
set -u

WORKING_FILE="$(mktemp /tmp/rh-manifest.XXXXXXXX)"

atexit() {
	echo "Cleaning up..."
	rm "$WORKING_FILE"
}

trap atexit EXIT

find . -name package-lock.json -execdir npm list --prod --json --all --package-lock-only \; |
	jq -r '..|objects|to_entries|.[]|select(.value.version?) | "\(.key)@\(.value.version)"' \
		>>"${WORKING_FILE}"

sort "${WORKING_FILE}" | uniq >rh-manifest.txt
