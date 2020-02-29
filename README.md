## GO install
```shell script
sudo add-apt-repository ppa:longsleep/golang-backports
sudo apt-get update
sudo apt-get install golang-go
```
```shell script
sudo nano /lib/systemd/system/vap_helper.service
```

```shell script
[Unit]
Description=vpn helper

[Service]
Type=simple
Restart=always
RestartSec=5s
ExecStart=PATH_TO_HOMEDIR/go/src/ip_country/main

[Install]
WantedBy=multi-user.target
```

```shell script
sudo service vap_helper start
sudo systemctl enable vap_helper
```