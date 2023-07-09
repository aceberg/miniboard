[![Main-Docker](https://github.com/aceberg/miniboard/actions/workflows/main-docker.yml/badge.svg)](https://github.com/aceberg/miniboard/actions/workflows/main-docker.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/aceberg/miniboard)](https://goreportcard.com/report/github.com/aceberg/miniboard)
[![Maintainability](https://api.codeclimate.com/v1/badges/064657fe6ff0adb0d3ba/maintainability)](https://codeclimate.com/github/aceberg/miniboard/maintainability)
![Docker Image Size (latest semver)](https://img.shields.io/docker/image-size/aceberg/miniboard)

<h1><a href="https://github.com/aceberg/miniboard">
    <img src="https://raw.githubusercontent.com/aceberg/miniboard/main/assets/logo.png" width="20" />
</a>miniboard</h1>
Lightweight dashboard with tabs, uptime monitoring and notifications.<br>
Can be configured through GUI or yaml file.     

- [Quick start](https://github.com/aceberg/miniboard#quick-start)
- [Install binary](docs/INSTALL-BIN.md)
- [Usage](https://github.com/aceberg/miniboard#usage)
- [Config](https://github.com/aceberg/miniboard#config)
- [Options](https://github.com/aceberg/miniboard#options)
- [Local network only](https://github.com/aceberg/miniboard#local-network-only) 
- [Thanks](https://github.com/aceberg/miniboard#thanks)

![screenshot](https://raw.githubusercontent.com/aceberg/miniboard/main/assets/Screenshot%202023-06-25%20at%2022-55-05%20MiniBoard%20-%20Docker.png)
<details>
  <summary>More Themes</summary>
  <img src="https://raw.githubusercontent.com/aceberg/miniboard/main/assets/MiniBoard%20-%20Themes.png">
</details>     

## Quick start

```sh
docker run --name miniboard \
-e "TZ=$YOURTIMEZONE" \
-v ~/.dockerdata/miniboard:/data/miniboard \
-v /var/run/docker.sock:/var/run/docker.sock \ # Only needed to create panel from Docker containers (On Edit panels page)
-p 8849:8849 \
aceberg/miniboard
```
Or use [docker-compose.yml](docker-compose.yml)   
There is also [binary installation](docs/INSTALL-BIN.md) available.


## Usage
All configuration can be done both with GUI or config file. To start with GUI first create new panel, then create new tab with this panel. Then you can add as many hosts (to panels), panels and tabs as you want.   
<br>
Also, you can start with example [board.yaml](configs/board.yaml) file. All fields are commented there.
  

## Config
Configuration can be done through config file or environment variables

| Variable  | Description | Default |
| --------  | ----------- | ------- |
| HOST | Listen address | 0.0.0.0 |
| PORT   | Port for web GUI | 8849 |
| THEME | Any theme name from https://bootswatch.com in lowcase | flatly |
| COLOR | Background color: light or dark | dark |
| COLORON | Online host color | #89ff89 |
| COLOROFF | Offline host color | #ff3232 |
| BTNWIDTH | Adjust buttons to theme | 180px |
| WEBREFRESH | Refresh interval for Tabs and Uptime pages (seconds) | 60 |
| TZ | Time zone (for uptime monitor) | "" |

## Options
| Key  | Description | Default | 
| --------  | ----------- | ------- | 
| -b | Path to board file | /data/miniboard/board.yaml |
| -c | Path to config file | /data/miniboard/config.yaml |
| -n | Path to node modules (see below) | "" |

## Local network only
By default, this app pulls themes, icons and fonts from the internet. But, in some cases, it may be useful to have an independent from global network setup. I created a separate [image](https://github.com/aceberg/my-dockerfiles/tree/main/node-bootstrap) with all necessary modules and fonts.
Run with Docker:
```sh
docker run --name node-bootstrap          \
    -p 8850:8850                          \
    aceberg/node-bootstrap
```
```sh
docker run --name miniboard \
    -v ~/.dockerdata/miniboard:/data/miniboard \
    -p 8849:8849 \
    aceberg/miniboard -n "http://127.0.0.1:8850"
```
Or use [docker-compose](docker-compose-local.yml)

## Thanks
- All go packages listed in [dependencies](https://github.com/aceberg/miniboard/network/dependencies)
- Favicon and logo: [Skateboard icons created by Freepik - Flaticon](https://www.flaticon.com/free-icons/skateboard)
- [Bootstrap](https://getbootstrap.com/)
- Themes: [Free themes for Bootstrap](https://bootswatch.com)