#!/bin/sh
echo "Rebuilding..."
go build -o nordvpn-cycle
chmod +x nordvpn-cycle
echo "Restarting..."
./nordvpn-cycle