#!/usr/bin/env bash

MPOD="../mpod/mpod"
MPOD_VERSION="0.0.1"

@test "mpod version logs version" {
  run $MPOD version
  [ "$status" -eq 0 ]
  [[ "$output" =~ $MPOD_VERSION ]]   # Check the output contains the version string
}