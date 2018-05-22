#!/bin/bash

#Variables - START
TARBALLURL="https://github.com/bulwark-crypto/Bulwark/releases/download/1.2.4/bulwark-1.2.4.0-linux64.tar.gz"
TARBALLNAME="bulwark-1.2.4.0-linux64.tar.gz"
BOOTSTRAPURL="https://github.com/bulwark-crypto/Bulwark/releases/download/1.2.4/bootstrap.dat.zip"
BOOTSTRAPARCHIVE="bootstrap.dat.zip"
BWKVERSION="1.2.4.0"
#Variables - END

sudo adduser bulwark

#Bulwark Service - START
sudo cat > /etc/systemd/system/bulwarkd.service << EOL
[Unit]
Description=Bulwarks's distributed currency daemon
After=network.target
[Service]
User=bulwark
Group=bulwark
WorkingDirectory=/home/bulwark
Type=forking
ExecStart=/usr/local/bin/bulwarkd -datadir=/home/bulwark/.bulwark -conf=/home/bulwark/.bulwark/bulwark.conf -daemon
ExecStop=/usr/local/bin/bulwark-cli -datadir=/home/bulwark/.bulwark -conf=/home/bulwark/.bulwark/bulwark.conf stop
#KillMode=process
Restart=always
TimeoutSec=120
RestartSec=30
[Install]
WantedBy=multi-user.target
EOL
sleep 1
#Bulwark Service - END

#Bulwark Config - START
sudo mkdir /home/bulwark/.bulwark
wget $BOOTSTRAPURL && unzip $BOOTSTRAPARCHIVE -d /home/bulwark/.bulwark/ && rm $BOOTSTRAPARCHIVE
sudo touch /home/bulwark/.bulwark/bulwark.conf
sudo chown -R bulwark:bulwark /home/bulwark/.bulwark
RPCUSER=$(cat /dev/urandom | tr -dc 'a-zA-Z0-9' | fold -w 12 | head -n 1)
RPCPASSWORD=$(cat /dev/urandom | tr -dc 'a-zA-Z0-9' | fold -w 32 | head -n 1)
sudo cat > /home/bulwark/.bulwark/bulwark.conf << EOL
rpcusername=${RPCUSER}
rpcpassword=${RPCPASSWORD}
daemon=1
EOL
#Bulwark Config - END

#Golang Setup - START
sudo wget https://dl.google.com/go/go1.10.2.linux-amd64.tar.gz
sudo tar -C /usr/local -xzf go1.10.2.linux-amd64.tar.gz
sudo rm go1.10.2.linux-amd64.tar.gz
sudo mkdir -p /home/bulwark/go/bin
sudo chown -R bulwark:bulwark /home/bulwark/go
sleep 1
# put into global /etc/profile
export PATH=$PATH:/usr/local/go/bin
sudo echo "PATH=/usr/local/go/bin:$PATH" >> /etc/profile
source /etc/profile
sleep 1
# put into user's ~/.profile
export GOPATH=/home/bulwark/go
export PATH=$PATH:$GOPATH/bin
sudo echo "" >> /home/bulwark/.profile
sudo echo "# Bulwark settings" >> /home/bulwark/.profile
sudo echo "GOPATH=/home/bulwark/go" >> /home/bulwark/.profile
sudo echo "PATH=$PATH:$GOPATH/bin" >> /home/bulwark/.profile
sleep 1
#Golang Setup - END

#BWK-Dash Setup - START
sudo mkdir -p /home/bulwark/go/src/github.com/dustinengle
sudo git clone https://github.com/dustinengle/bwk-dash /home/bulwark/go/src/github.com/dustinengle/bwk-dash
sudo GOOS=linux GOARCH=amd64 go build -o /usr/local/bin/bwk-cron /home/bulwark/go/src/github.com/dustinengle/bwk-dash/cmd/bwk-cron/*.go
sudo GOOS=linux GOARCH=amd64 go build -o /usr/local/bin/bwk-dash /home/bulwark/go/src/github.com/dustinengle/bwk-dash/cmd/bwk-dash/*.go
# Setup systemd service and start.
sudo cat > /etc/systemd/system/bwk-dash.service << EOL
[Unit]
Description=Bulwark Home Node Dashboard
After=network.target
[Service]
User=bulwark
Group=bulwark
WorkingDirectory=/home/bulwark/dash
ExecStart=/usr/local/bin/bwk-dash
Restart=always
TimeoutSec=5
RestartSec=5
[Install]
WantedBy=multi-user.target
EOL
sleep 1
# Create .env file.
# Uses same RPCUSER and RPCPASSWORD as above configuration.
sudo mkdir -p /home/bulwark/dash
sudo cat > /home/bulwark/dash/.env << EOL
DASH_DONATION_ADDRESS=TESTADDRESSHERE
DASH_PORT=8080
DASH_RPC_ADDR=localhost
DASH_RPC_PORT=52541
DASH_RPC_USER=${RPCUSER}
DASH_RPC_PASS=${RPCPASSWORD}
DASH_WEBSITE=/home/bulwark/dash
DASH_DB=/home/bulwark/dash/bwk-dash.db
EOL
sleep 1
# Copy the html files to the dash folder.
sudo cp -R /home/bulwark/go/src/github.com/dustinengle/bwk-dash/client/build/* /home/bulwark/dash/
# Cleanup/enforce ownership.
sudo chown -R bulwark:bulwark /home/bulwark/dash
# Setup cron job for bwk-cron.
sudo crontab -u bulwark -l > mycron
echo '*/5 * * * * cd /home/bulwark/dash && /usr/local/bin/bwk-cron' >> mycron
sudo crontab -u bulwark mycron
sleep 1
sudo rm -f mycron
# Enable dashboard service.
sudo systemctl enable bwk-dash
#BWK-Dash Setup - END

#Bulwark Node - START
sudo wget $TARBALLURL
sleep 2
sudo tar -xzf $TARBALLNAME -C /usr/local
sudo rm $TARBALLNAME
sleep 3
cd ~
sudo chown -R bulwark:bulwark /home/bulwark/bulwark/
sleep 1
sudo systemctl enable bulwarkd.service
sleep 1
#Bulwark Node - END

sudo reboot
