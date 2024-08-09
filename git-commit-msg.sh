#!/bin/bash

# Function to determine the commit message type
function get_commit_message {
    local file=$1

    if [[ $file == *.html ]]; then
        echo "add: update HTML file $file"
    elif [[ $file == *.css ]]; then
        echo "feat: update CSS styles in $file"
    elif [[ $file == *.js ]]; then
        echo "feat: update JavaScript file $file"
    elif [[ $file == *.go ]]; then
        if git diff --cached --name-only | grep -q "$file"; then
            echo "refactor: Refactor Go code in $file"
        else
            echo "add: Add new Go file $file"
        fi
    elif [[ $file == *.md ]]; then
        echo "docs: Add or update documentation in $file"
    else
        echo "chore: Update $file"
    fi
}

# Add and commit each file individually
for file in $(git ls-files --modified --others --exclude-standard); do
    git add "$file"
    commit_message=$(get_commit_message "$file")
    git commit -m "$commit_message"
done
