#!/bin/sh
 
STAGED_GO_FILES=$(git diff --cached --name-only -- '*.go')     
if [[ $STAGED_GO_FILES == "" ]]; then                          
    echo "no go files updated"
else
    for file in $STAGED_GO_FILES; do
        go fmt $file                                           
        git add $file
    done
fi
 
golangci-lint run            