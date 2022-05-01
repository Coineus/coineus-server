#!/bin/bash
set -e
echo "***** Pulling Latest API Image ******"
docker image pull safderun/coineus-server:latest
echo ""
echo "***** Removing Old Container *****"
START=$(date +%s)
docker container rm -f coineus-server
echo ""
echo "***** Docker-Compose Up Runnig  *****"
docker-compose -f ./docker/server-compose.yaml up -d
echo ""
END=$(date +%s)
DIFF=$(( $END - $START ))
echo "API Updated Succesfully, Interruption for $DIFF Seconds"