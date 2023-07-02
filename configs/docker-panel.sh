#!/bin/bash

# This script generates a MiniBoard Panel from Docker containers.
# It can be added to board.yaml under 'panels:' section.

HNAME=`hostname`
docker ps -a --format "{{.Names}}">/tmp/miniboard-docker.txt

echo $HNAME':'
echo '      name: '$HNAME
echo '      scan: false'
echo '      hosts:'

i=0
while read NAME; do
    echo '           '$i:
    let "i++"
    echo '               name: '$NAME
    ADDR=`docker inspect $NAME | grep HostIp | sed '1!d;s/"HostIp": //;s/,//'`
    echo '               addr: '$ADDR
    PORT=`docker inspect $NAME | grep HostPort | sed '1!d;s/"HostPort": //;s/,//'`
    echo '               port: '$PORT
done </tmp/miniboard-docker.txt