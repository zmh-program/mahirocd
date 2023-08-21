#!/bin/bash

# check permission
if [ "$EUID" -ne 0 ]; then
    echo "please run as root"
    exit 1
fi


echo "starting build project..."
go build .
echo "build project successfully"
echo

echo "starting register service..."

CWD=$(pwd)

cat > /etc/systemd/system/mahirocd.service <<EOL
[Unit]
Description=MahiroCD Node Service
After=network.target

[Service]
ExecStart=$CWD/mahirocd
WorkingDirectory=$CWD
Restart=always
User=root

[Install]
WantedBy=multi-user.target
EOL

systemctl daemon-reload

echo "register service successfully"

systemctl enable mahirocd
systemctl start mahirocd

echo "service has been started"
