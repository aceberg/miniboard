# Install binary

## From .deb repository (recommended)
```sh
curl -s --compressed "https://aceberg.github.io/ppa/KEY.gpg" | gpg --dearmor | sudo tee /etc/apt/trusted.gpg.d/aceberg.gpg
```
```sh
sudo curl -s --compressed -o /etc/apt/sources.list.d/aceberg.list "https://aceberg.github.io/ppa/aceberg.list"
```
```sh
sudo apt update && sudo apt install miniboard
```
## From .deb file
Download [latest](https://github.com/aceberg/miniboard/releases/latest) release, install with your package maneger

## From .tar.gz
Download [latest](https://github.com/aceberg/miniboard/releases/latest) release, then
```sh
tar xvzf miniboard-*.tar.gz
cd miniboard
sudo ./install.sh
```
<br>

# Usage
## Systemd as user (recommended)
Enable and start service, replace `MYUSER` with your username
```sh
sudo systemctl enable miniboard@MYUSER.service
sudo systemctl start miniboard@MYUSER.service
```
Web GUI will be available at [http://0.0.0.0:8849](http://0.0.0.0:8849)   
Config files are stored at `/home/MYUSER/.config/miniboard/`   


## Systemd as root
Enable and start service
```sh
sudo systemctl enable miniboard.service
sudo systemctl start miniboard.service
```
Web GUI will be available at [http://0.0.0.0:8849](http://0.0.0.0:8849)   
Config files are stored at `/etc/miniboard/`

## From command line
Just run `miniboard`. Be mindful of the config files paths listed in [options](https://github.com/aceberg/miniboard#options) section.

