#!/bin/bash

# Define colors
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[0;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

echo -e "$GREEN Building Go Build $NC"

go build

echo -e "Zip build $NC"

tar -cvzf webapp.tar.gz webapp

echo -e "$YELLOW Zip Complete $NC"

echo -e "$RED Copy From Local to Remote $NC"

echo "Deploying To git"

CURRENT_DATETIME=$(date +"%Y-%m-%d %H:%M:%S")

echo "Adding Git Chages"
git add .

echo "Git Commit"
git commit -m "Deploying changes for Go On  $CURRENT_DATETIME"

echo "Git Push"
git push origin  main

# Get the first parameter
PARAM=$1

if [ $# -lt 1 ]; then
    
    echo "Only Tar will Deploy"

    scp -r  -i "/home/kunal/pems/phisbug.pem" webapp.tar.gz  ec2-user@ec2-52-53-78-232.us-west-1.compute.amazonaws.com:~/
else
    scp -r  -i "/home/kunal/pems/phisbug.pem" webapp.tar.gz  ec2-user@ec2-52-53-78-232.us-west-1.compute.amazonaws.com:~/
    scp -r  -i "/home/kunal/pems/phisbug.pem" view  ec2-user@ec2-52-53-78-232.us-west-1.compute.amazonaws.com:~/
    scp -r  -i "/home/kunal/pems/phisbug.pem" pages  ec2-user@ec2-52-53-78-232.us-west-1.compute.amazonaws.com:~/
    scp -r  -i "/home/kunal/pems/phisbug.pem" static  ec2-user@ec2-52-53-78-232.us-west-1.compute.amazonaws.com:~/        
fi



echo "Copy Complete"

# Variables
INSTANCE_USER="ec2-user"  # Change this to your instance's user (e.g., ubuntu, ec2-user)
INSTANCE_IP="ec2-52-53-78-232.us-west-1.compute.amazonaws.com"  # Replace with your instance's public IP or DNS
PRIVATE_KEY_PATH="/home/kunal/pems/phisbug.pem"  # Replace with the path to your private key

# SSH into the instance and run commands
ssh -i "$PRIVATE_KEY_PATH" "$INSTANCE_USER@$INSTANCE_IP" << 'EOF'
    # Navigate to the home repository
    cd /home/ec2-user  # Replace with your repository path

    # Run your commands
    echo "Running commands in the repository"
    # Example commands:
    
    tar -xzvf webapp.tar.gz

    echo "Untar Finish"

    sudo systemctl restart myapp

    echo "App restart Complete"
    
    # Add any other commands you want to run here

EOF
