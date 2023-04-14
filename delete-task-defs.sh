#!/bin/bash

## ./delete-task-defs.sh <family>
## Will discover the latest task defintion, and begin deleting it recursivly
set -e
FAMILY=$1

# Get latest revision
REVISION=$(aws ecs describe-task-definition --task-definition "${FAMILY}" | jq '.taskDefinition.revision')
## Loop through delete with backoff
./build/ecs-task-mgr delete -t "${FAMILY}:${REVISION}"