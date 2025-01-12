#!/bin/bash

# remove tag locally
git tag -d v0.0.7
# remove tag on remote
git push origin --delete v0.0.7

# create tag again
git tag -a v0.0.7 -m "this tag was created via script to test actions"
git push origin v0.0.7
