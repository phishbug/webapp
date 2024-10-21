#!/bin/bash

echo "Checking if parameters are provided"

# Get the first parameter
PARAM=$1

CURRENT_DATETIME=$(date +"%Y-%m-%d %H:%M:%S")

if [ $# -lt 1 ]; then
    echo "No Parms provided if want to run from params eg: ./template_deploy_to_dev.sh --file=view"
    echo "Prociding with all template, pages, static files"

    echo "Finding Changes"

    git status

    echo "Adding Git Chages"
    git add .

    echo "Git Commit"
    git commit -m "Deploying changes for Go Templates On  $CURRENT_DATETIME"

    echo "Git Push"
    git push origin  main


    scp -r  -i "/home/kunal/pems/phisbug.pem" view  ec2-user@ec2-52-53-78-232.us-west-1.compute.amazonaws.com:~/
    scp -r  -i "/home/kunal/pems/phisbug.pem" pages  ec2-user@ec2-52-53-78-232.us-west-1.compute.amazonaws.com:~/
    scp -r  -i "/home/kunal/pems/phisbug.pem" static  ec2-user@ec2-52-53-78-232.us-west-1.compute.amazonaws.com:~/
elif [ "$PARAM" == "view" ] || [ "$PARAM" == "pages" ] || [ "$PARAM" == "static" ]; then

    echo "Finding Changes"
    git status

    echo "Adding Git Chages"
    git add .

    echo "Git Commit"
    git commit -m "Deploying changes for $PARAM Go Templates On  $CURRENT_DATETIME"

    echo "Git Push"
    git push origin  main

    echo "Param Found: $PARAM"

    scp -r  -i "/home/kunal/pems/phisbug.pem" $PARAM  ec2-user@ec2-52-53-78-232.us-west-1.compute.amazonaws.com:~/
else
    echo "Invalid Command Check doc for deployment"
    exit 1
fi