#!/usr/bin/bash

PROJECT_PATH=/home/pc04/DemoProject/gin_demo/standard_app


printf "\n ====================START GIT==================== \n"

cd $PROJECT_PATH
ssh-agent bash -c "ssh-add  ~/.ssh/vuonghuutri.cit; $1"


printf "====================END GIT===================="
