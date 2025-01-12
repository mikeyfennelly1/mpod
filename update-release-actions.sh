#!/bin/bash

# add all files to staging area
git add --all

# commit files with commit message
git commit -m "$1"

# push existing local commits
git push origin main

# remove tag locally
git tag -d v0.0.7
# remove tag on remote
git push origin --delete v0.0.7

# create tag again
git tag -a v0.0.7 -m "this tag was created via script to test actions"
git push origin v0.0.7
