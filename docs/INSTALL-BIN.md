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

