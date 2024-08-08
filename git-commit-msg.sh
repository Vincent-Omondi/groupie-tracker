#!/bin/bash

changes=$(git status --porcelain)
untracked=$(git status --porcelain | grep '^??')
modified=$(git status --porcelain | grep '^ M')

if [ -n "$untracked" ]; then
  commit_message="feat: add new files and directories"
elif [ -n "$modified" ]; then
  commit_message="fix: resolve issues with modified files"
else
  commit_message="docs: update README.md with new information"
fi

git add .
git commit -m "$commit_message"