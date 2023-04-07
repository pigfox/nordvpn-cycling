#!/bin/sh
while true
do
  echo "Last status check:" $(date '+%Y-%m-%d %H:%M:%S')
  nordvpn status
  sleep 60
  clear
done