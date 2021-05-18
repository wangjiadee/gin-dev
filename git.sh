#!/bin/bash

set +x

GIT_BRANCH=`git branch|awk -F ' ' '{print $2}'`

echo "The commit branch this time is: ${GIT_BRANCH}"

git add .

git commit -m "$1"

git push origin ${GIT_BRANCH}

if [ $? -ne 0 ];then
    echo "git push failed! Pls check!"
else
    GIT_COMMIT=`git reflog |head -1|awk -F ' ' '{print $1}'`
    echo "The commit id: ${GIT_COMMIT}"
fi
