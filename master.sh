#!/bin/bash

# check permission
if [ "$EUID" -ne 0 ]; then
    echo "please run as root"
    exit 1
fi


echo "starting build project..."
cd transport || exit
go build .
echo "build project successfully"
echo

echo "starting register service..."

CWD=$(pwd)

cat > /etc/systemd/system/mahironode.service <<EOL
[Unit]
Description=MahiroCD Master Service
After=network.target

[Service]
ExecStart=$CWD/mahironode
WorkingDirectory=$CWD
Restart=always
User=root

[Install]
WantedBy=multi-user.target
EOL

systemctl daemon-reload

echo "register service successfully"

systemctl enable mahironode
systemctl start mahironode

echo "service has been started"

