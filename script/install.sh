#!/bin/bash

#Variables - START
CHARS="/-\|"
TARBALLURL="https://github.com/bulwark-crypto/Bulwark/releases/download/1.2.4/bulwark-1.2.4.0-linux64.tar.gz"
TARBALLNAME="bulwark-1.2.4.0-linux64.tar.gz"
BOOTSTRAPURL="https://github.com/bulwark-crypto/Bulwark/releases/download/1.2.4/bootstrap.dat.zip"
BOOTSTRAPARCHIVE="bootstrap.dat.zip"
BWKVERSION="1.2.4.0"
#Variables - END

if [ "$(id -u)" != "0" ]; then
    echo "Sorry, this script needs to be run as root. Do \"sudo bash run.sh\""
    exit 1
fi

#System Setup - START
sudo echo "Preparing installation..."
sudo apt-get -y update
sleep 2
sudo apt-get -y upgrade
sleep 2
sudo apt-get -y dist-upgrade
sleep 2
sudo apt-get update -y
sleep 2
sudo apt-get install htop -y
sleep 3
sudo apt-get install nano -y
sleep 3
sudo apt-get install ufw -y
sleep 3
sudo apt-get install fail2ban -y
sleep 3
sudo apt-get install git -y
sleep 3
sudo apt install unzip -y
sleep 3
sudo wget --directory-prefix=/etc/fail2ban/ https://raw.githubusercontent.com/whywefight/Bulwark-MN-Install/master/jail.local
#System Setup - END

#Bulwark User Setup - START
sudo adduser --gecos "" bulwark --disabled-password > /dev/null
sleep 1
#Bulwark User Setup - END

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
sudo su -c "echo 'PATH=/usr/local/go/bin:$PATH' >> /etc/profile"
source /etc/profile
sleep 1
# put into user's ~/.profile
export GOPATH=/home/bulwark/go
export PATH=$PATH:$GOPATH/bin
echo "" >> /home/bulwark/.profile
echo "# Bulwark settings" >> /home/bulwark/.profile
sudo sh -c "echo 'GOPATH=/home/bulwark/go' >> /home/bulwark/.profile"
sudo sh -c "echo 'PATH=$PATH:$GOPATH/bin' >> /home/bulwark/.profile"
sudo sh -c "source /home/bulwark/.profile" bulwark
sleep 1
#Golang Setup - END

#BWK-Dash Setup - START
sudo su -c "go get -u github.com/dustinengle/bwk-dash" bulwark
sudo su -c "GOOS=linux GOARCH=amd64 go build -o /home/bulwark/go/bin/bwk-cron /home/bulwark/go/src/github.com/dustinengle/bwk-dash/cmd/bwk-cron/*.go" bulwark
sudo su -c "GOOS=linux GOARCH=amd64 go build -o /home/bulwark/go/bin/bwk-dash /home/bulwark/go/src/github.com/dustinengle/bwk-dash/cmd/bwk-dash/*.go" bulwark
# Setup systemd service and start.
sudo cat > /etc/systemd/system/bwk-dash.service << EOL
[Unit]
Description=Bulwark Home Node Dashboard
After=network.target
[Service]
User=bulwark
Group=bulwark
WorkingDirectory=/home/bulwark/dash
ExecStart=/home/bulwark/go/bin/bwk-dash
Restart=always
TimeoutSec=5
RestartSec=5
[Install]
WantedBy=multi-user.target
EOL
sleep 1
# Create .env file.
# Uses same RPCUSER and RPCPASSWORD as above configuration.
mkdir -p /home/bulwark/dash
cat > /home/bulwark/dash/.env << EOL
DASH_DONATION_ADDRESS=TESTADDRESSHERE
DASH_PORT=8080
DASH_RPC_ADDR=localhost
DASH_RPC_PORT=52541
DASH_RPC_USER=${$RPCUSER}
DASH_RPC_PASS=${$RPCPASSWORD}
DASH_WEBSITE=/home/bulwark/dash
DASH_DB=/home/bulwark/dash/bwk-dash.db
EOL
sleep 1
# Copy the html files to the dash folder.
sudo su -c "cp -R /home/bulwark/go/src/github.com/dustinengle/bwk-dash/client/build/* /home/bulwark/dash/" bulwark
# Cleanup/enforce ownership.
sudo chown -R bulwark:bulwark /home/bulwark/dash
# Run cron job for first time manually.
sudo su -c "bwk-cron" bulwark
sleep 1
# Setup cron job for bwk-cron.
crontab -u bulwark -l > mycron
echo '*/5 * * * * cd /home/bulwark/dash && bwk-cron' >> mycron
crontab -u bulwark mycron
sleep 1
rm -f mycron
# Start and enable dashboard service.
sudo systemctl start bwk-dash
sudo systemctl enable bwk-dash
#BWK-Dash Setup - END

#Firewall Setup - START
sudo ufw allow 8080
sleep 2
sudo ufw allow 9050
sleep 2
sudo ufw allow 52543
sleep 2
sudo ufw allow from 127.0.0.1 to 127.0.0.1 port 52541
sleep 2
sudo ufw allow from `ip addr | grep 'state UP' -A2 | tail -n1 | awk '{print $2}' | cut -f1  -d'/' | awk -F"." '{print $1"."$2"."$3".0/24"}'` to any port 22
yes | sudo ufw enable
sleep 2
#Firewall Setup - END

#Bulwark Node - START
sudo wget $TARBALLURL
sleep 2
sudo tar -xzf $TARBALLNAME
sudo mv bin bulwark
sudo rm $TARBALLNAME
cd bulwark
sudo cp bulwark* /usr/local/bin
sleep 3
cd ~
sudo mv /home/pi/bulwark /home/bulwark/
sudo chown -R bulwark:bulwark /home/bulwark/bulwark/
sleep 1
sudo systemctl enable bulwarkd.service
sleep 1
sudo systemctl start bulwarkd.service
sudo echo "Starting up bulwarkd, please wait"

# Wait for bulwark to finish starting to prevent errors in line 158
until sudo su -c "bulwark-cli getinfo 2>/dev/null | grep 'balance' > /dev/null" bulwark; do
  for (( i=0; i<${#CHARS}; i++ )); do
    sleep 2
    echo -en "${CHARS:$i:1}" "\r"
  done
done

sudo echo ""
sudo echo "I will open the getinfo screen for you in watch mode now, close it with CTRL + C once we are fully synced."
sleep 20
watch bulwark-cli -datadir=/home/bulwark/.bulwark -conf=/home/bulwark/.bulwark/bulwark.conf getinfo
sudo echo "Daemon Status:"
sudo systemctl status bulwarkd.service | sed -n -e 's/^.*Active: //p'
sudo echo ""
sudo echo "Show Active Peers:"
bulwark-cli -datadir=/home/bulwark/.bulwark -conf=/home/bulwark/.bulwark/bulwark.conf getpeerinfo | sed -n -e 's/^.*"addr" : //p'
sudo echo ""
sudo echo "Firewall Rules:"
sudo ufw status
sudo echo ""
sudo echo "Fail2Ban:"
sudo systemctl status fail2ban.service | sed -n -e 's/^.*Active: //p'
sudo echo ""
sudo echo "Important Other Infos:"
sudo echo ""
sudo echo "Bulwark bin dir: /home/bulwark/bulwark"
sudo echo "bulwark.conf: /home/bulwark/.bulwark/bulwark.conf"
sudo echo "Start daemon: sudo systemctl start bulwarkd.service"
sudo echo "Restart daemon: sudo systemctl restart bulwarkd.service"
sudo echo "Status of daemon: sudo systemctl status bulwarkd.service"
sudo echo "Stop daemon: sudo systemctl stop bulwarkd.service"
sudo echo "Check bulwarkd status: bulwark-cli getinfo"
sudo echo "Check masternode status: bulwark-cli masternode status"
sleep 5
sudo echo ""
sudo echo "Adding bulwark-cli shortcut to ~/.profile"
echo "alias bulwark-cli='sudo bulwark-cli -config=/home/bulwark/.bulwark/bulwark.conf -datadir=/home/bulwark/.bulwark'" >> /home/pi/.profile
sudo echo "Installation finished."
read -p "Press Enter to continue, the system will reboot."
#Bulwark Node - END

sudo reboot
