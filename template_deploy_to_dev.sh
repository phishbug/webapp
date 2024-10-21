#!/bin/bash

# Configuration
REMOTE_USER="ec2-user"  # Change this to your remote username
REMOTE_HOST="ec2-52-53-78-232.us-west-1.compute.amazonaws.com"  # Change this to your remote host (IP or domain)
REMOTE_DIR="/home/ec2-user"  # Change this to the desired remote directory
SSH_KEY="/home/kunal/pems/phisbug.pem"

# Check if we are in a git repository
if ! git rev-parse --is-inside-work-tree &> /dev/null; then
    echo "Not a git repository."
    exit 1
fi

# Get the output of git status
git_status_output=$(git status)

echo "Deploying To git"

CURRENT_DATETIME=$(date +"%Y-%m-%d %H:%M:%S")

echo "Adding Git Chages"
git add .

echo "Git Commit"
git commit -m "Deploying changes for Go On  $CURRENT_DATETIME"

echo "Git Push"
git push origin  main


# Folders to check
folders=("view" "static" "pages")

# Function to copy files to the remote server
copy_files() {
    local files=("$@")
    for file in "${files[@]}"; do
        echo "$REMOTE_USER@$REMOTE_HOST:$REMOTE_DIR/$file"
        scp -i "$SSH_KEY" "$file" "$REMOTE_USER@$REMOTE_HOST:$REMOTE_DIR/$file"
    done
}

# Arrays to hold files
files_to_copy=()

# Check for modified, untracked, and deleted files
for folder in "${folders[@]}"; do
    # Modified files
    while IFS= read -r line; do
        if [[ "$line" == *"$folder/"* ]]; then
            files_to_copy+=("$line")
        fi
    done < <(echo "$git_status_output" | grep 'modified:' | awk '{print $2}')

    # Untracked files
    while IFS= read -r line; do
        if [[ "$line" == *"$folder/"* ]]; then
            files_to_copy+=("$line")
        fi
    done < <(echo "$git_status_output" | grep 'Untracked files:' -A 100 | grep -v 'use' | grep -v 'Untracked files:' | sed 's/^ *//')

    # Deleted files
    while IFS= read -r line; do
        if [[ "$line" == *"$folder/"* ]]; then
            files_to_copy+=("$line")
        fi
    done < <(echo "$git_status_output" | grep 'deleted:' | awk '{print $2}')
done

# Copy files to remote server
if [ ${#files_to_copy[@]} -gt 0 ]; then
    echo "Copying files to remote server..."
    copy_files "${files_to_copy[@]}"
    echo "Files copied successfully."
else
    echo "No files to copy."
fi
