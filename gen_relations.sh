#!/bin/bash

# Check if the correct number of arguments are passed
if [ "$#" -ne 2 ]; then
    echo "Usage: $0 <number_of_times> <user>"
    exit 1
fi

n=$1
user=$2

binding="${user}_binding"
zed relationship create role_binding:"$binding" subject user:"$user"
zed relationship create role_binding:"$binding" granted role:patch_viewer

for ((i=1; i<=n; i++))
do
    workspace="${i}_test_workspace"
    zed relationship create workspace:"$workspace" user_grant role_binding:"$binding"
done


