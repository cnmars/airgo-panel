<img width="200px" src="https://telegraph-image.pages.dev/file/c48a2f45ebf102dd66131.png" align="left"/>

# AirGo Front and rear separation, multi-user, multi protocol proxy service management system, simple and easy to use

![License](https://img.shields.io/badge/License-GPL_v3.0-red)
![Go](https://img.shields.io/badge/Golang-orange?logo=Go&logoColor=white)
![Gorm](https://img.shields.io/badge/Gorm-yellow&logo=gorm)
![Gin](https://img.shields.io/badge/Gin-green?logo=)
![Vue](https://img.shields.io/badge/Vue.js-00b6ff?logo=vuedotjs&logoColor=white)
![TypeScript](https://img.shields.io/badge/TypeScript-blue?logo=TypeScript&logoColor=white)
![Linux](https://img.shields.io/badge/Linux-purple?logo=linux&logoColor=white)

<hr/>

Support：vless，vmess，shadowsocks，hysteria2

<hr/>

<h3><a href='./README.md'>中文</a> | English</h3>

<h3>⚠️⚠️⚠️Attention: The English document comes from machine translation</h3>

<hr/>

<!-- TOC -->
* [AirGo Front and rear separation, multi-user, multi protocol proxy service management system, simple and easy to use](#airgo-front-and-rear-separation-multi-user-multi-protocol-proxy-service-management-system-simple-and-easy-to-use)
* [Display of panel functions](#display-of-panel-functions)
* [Catalogue：](#catalogue)
* [1 Deployment - front-end and back-end non separation](#1-deployment---front-end-and-back-end-non-separation)
  * [1-1 Direct installation](#1-1-direct-installation)
  * [1-2 Installing using Docker](#1-2-installing-using-docker)
* [2 Deployment - front-end and back-end separation](#2-deployment---front-end-and-back-end-separation)
  * [2-1 back-end](#2-1-back-end)
    * [2-1-1 Direct installation](#2-1-1-direct-installation)
    * [2-1-2 Installing through Docker](#2-1-2-installing-through-docker)
  * [2-2 front end](#2-2-front-end)
    * [2-2-1 Deploy to cloud platforms such as Vercel](#2-2-1-deploy-to-cloud-platforms-such-as-vercel)
    * [2-2-2 Deploy to web application servers such as Nginx, Caddy, OpenResty, etc](#2-2-2-deploy-to-web-application-servers-such-as-nginx-caddy-openresty-etc)
    * [2-2-3 Set partial personalized data](#2-2-3-set-partial-personalized-data)
* [3 Configure SSL (optional)](#3-configure-ssl--optional-)
  * [3-1 Set up SSL certificates for the front-end](#3-1-set-up-ssl-certificates-for-the-front-end)
  * [3-2 Set up SSL certificates for the backend](#3-2-set-up-ssl-certificates-for-the-backend)
* [4 Configuration file description](#4-configuration-file-description)
* [5 Docking nodes](#5-docking-nodes)
  * [5-1 V2bX](#5-1-v2bx)
    * [5-1-1 Directly install V2bX](#5-1-1-directly-install-v2bx)
    * [5-1-2 docker install V2bX](#5-1-2-docker-install-v2bx)
  * [5-2 XrayR](#5-2-xrayr)
    * [5-2-1 Install XrayR directly](#5-2-1-install-xrayr-directly)
    * [5-2-2 docker install XrayR](#5-2-2-docker-install-xrayr)
* [6 Update Panel](#6-update-panel)
  * [6-1 Update backend](#6-1-update-backend)
  * [6-2 Update front-end](#6-2-update-front-end)
* [7 命令行](#7-命令行)
* [8 开发](#8-开发)
<!-- TOC -->

TG Channel：[https://t.me/Air_Go](https://t.me/Air_Go)

TG group：[https://t.me/AirGo_Group](https://t.me/AirGo_Group)

Last update date of document：2024.4.8

# Display of panel functions

<div style="color: darkgray" >Display of panel functions</div>

<table>
<tr>
    <td> <img src="https://github.com/ppoonk/AirGo/raw/main/assets/image/1.png">
    <td> <img src="https://github.com/ppoonk/AirGo/raw/main/assets/image/2.png">
    <td> <img src="https://github.com/ppoonk/AirGo/raw/main/assets/image/3.png">
<tr>
    <td> <img src="https://github.com/ppoonk/AirGo/raw/main/assets/image/4.png">
    <td> <img src="https://github.com/ppoonk/AirGo/raw/main/assets/image/5.png">
    <td> <img src="https://github.com/ppoonk/AirGo/raw/main/assets/image/6.png">
</table>
<table>
<tr>
    <td> <img src="https://github.com/ppoonk/AirGo/raw/main/assets/image/7.png">
</table>

# Catalogue：

# 1 Deployment - front-end and back-end non separation

## 1-1 Direct installation

- Install the core, use Linux systems such as Ubuntu, Debian, Centos, etc., execute the following command, and then follow the prompts to install

```
bash <(curl -Ls https://raw.githubusercontent.com/ppoonk/AirGo/main/server/scripts/install.sh)
```

- Modify the configuration file to the directory '/usr/local/AirGo/config. yaml'. For the first installation, data will be automatically initialized based on the configuration file config.yaml. Please make sure to change the administrator account and password
- Start Core, ` systemictl start AirGo`
- Browser access:` http://ip:port `, where the port is the value set in the configuration file

## 1-2 Installing using Docker

- Create a new configuration file in the appropriate directory, for example:/$PWD/air/config.yaml. The content of the configuration file is as follows. For the first installation, data will be automatically initialized based on the configuration file config.yaml. Please make sure to modify the administrator account and password

```
system:
  admin-email: admin@oicq.com
  admin-password: adminadmin
  http-port: 80
  https-port: 443
  db-type: sqlite
mysql:
  address: mysql.sql.com
  port: 3306
  config: charset=utf8mb4&parseTime=True&loc=Local
  db-name: imdemo
  username: imdemo
  password: xxxxxx
  max-idle-conns: 10
  max-open-conns: 100
sqlite:
  path: ./air.db
```

- The starting Docker command reference is as follows：

```
docker run -tid \
  -v $PWD/air/config.yaml:/air/config.yaml \
  -p 80:80 \
  -p 443:443 \
  --name airgo \
  --restart always \
  --privileged=true \
  ppoiuty/airgo:latest
```

- docker compose is as follows：

```
version: '3'

services:
  airgo:
    container_name: airgo
    image: ppoiuty/airgo:latest
    ports:
      - "80:80"
      - "443:443"
    restart: "always"
    privileged: true
    volumes:
      - ./config.yaml:/air/config.yaml
```

- Browser access:` http://ip:port `, where the port is the value set in the configuration file


# 2 Deployment - front-end and back-end separation

## 2-1 back-end

### 2-1-1 Direct installation
```
bash <(curl -Ls https://raw.githubusercontent.com/ppoonk/AirGo/main/server/scripts/install.sh)
```

- Modify the configuration file to the directory `/usr/local/AirGo/config. yaml`. For the first installation, data will be automatically initialized based on the configuration file config.yaml. Please make sure to change the administrator account and password
- Start Core, ` systemictl start AirGo`

### 2-1-2 Installing through Docker

- Prepare the configuration file config.yaml in advance, refer to [config.yaml](https://github.com/ppoonk/AirGo/blob/main/server/config.yaml) For the first installation, data will be automatically initialized based on the configuration file config.yaml. Please make sure to modify the administrator account and password
- The starting Docker command reference is as follows:

```
docker run -tid \
  -v $PWD/air/config.yaml:/air/config.yaml \
  -p 80:80 \
  -p 443:443 \
  --name airgo \
  --restart always \
  --privileged=true \
  ppoiuty/airgo:latest
```

- docker compose is as follows：

```
version: '3'
services:
  airgo:
    container_name: airgo
    image: ppoiuty/airgo:latest
    ports:
      - "80:80"
      - "443:443"
    restart: "always"
    privileged: true
    volumes:
      - ./config.yaml:/air/config.yaml
```
## 2-2 front end

### 2-2-1 Deploy to cloud platforms such as Vercel

- Fork this project, modify The `window.httpurl` field of `web/index.HTML` to your own backend address. Due to Vercel restrictions, please fill in the HTTPS interface address
- Login [Vercel](https://vercel.com)，Add New Project，Refer to the configuration diagram below and pay attention to the settings in the red circle!
  ![image](https://telegraph-image.pages.dev/file/afe97f45857b988ebd005.png)
- After successful deployment, customize the domain name (domain name resolution to 76.76.21.21)

### 2-2-2 Deploy to web application servers such as Nginx, Caddy, OpenResty, etc
It is recommended to use `github codespaces` compilation as it will not install additional dependencies on your computer

- Fork this project, modify The `window. httpurl` field of `web/index.HTML` to your own backend address
- Under `project/web/`, execute `npm i&&npm run build`
- The packaged static resource folder is web, upload the web folder to the appropriate location on the server. Create a new website (purely static), select the web folder for the website location

### 2-2-3 Set partial personalized data

modify  `keywords`，`description`，`title`，`favicon.ico` in `./web/index.html`





# 3 Configure SSL (optional)

## 3-1 Set up SSL certificates for the front-end

Certificates can be directly applied for and imported through platforms such as `Pagoda Panel (bt. cn)` and `1Panel (1Panel. cn)`

## 3-2 Set up SSL certificates for the backend

- Apply for or import a certificate through the `Pagoda Panel (bt. cn)`, `1Panel (1Panel. cn)`, and then enable reverse proxy
- If you already have the certificate, simply copy it to the installation directory (/usr/local/AirGo/), rename it to `air. cer`, `air. key`, and then restart AirGo


# 4 Configuration file description

```
system:
  mode: release                //Mode, default to release. If it is dev, that is, development mode. The console will output more information
  admin-email: admin@oicq.com  //Administrator account needs to be modified before initialization!
  admin-password: adminadmin   //Administrator password needs to be modified before initialization!
  http-port: 8899              //Core listening port
  https-port: 443              //Core listening port
  db-type: sqlite              //Database type, optional values: MySQL, Mariadb, sqlite
mysql:
  address: xxx.com             //MySQL database address
  port: 3306                   //MySQL database port
  db-name: xxx                 //MySQL database name
  username: xxx                //MySQL database username
  password: xxx                //MySQL database password
  config: charset=utf8mb4&parseTime=True&loc=Local //Just keep it as default
  max-idle-conns: 10
  max-open-conns: 100
sqlite:
  path: ./air.db               //Sqlite database file name
```

# 5 Docking nodes

**V2bx and XrayR are currently supported, but the official version is not currently supported. Please use the following version：**

## 5-1 V2bX

### 5-1-1 Directly install V2bX

```
bash <(curl -Ls https://raw.githubusercontent.com/ppoonk/V2bX/main/scripts/install.sh)
```
- After installation is completed, please modify the configuration file as needed in `/usr/local/AV/config. json`
- Start: Use administrative scripts `AV` or directly start AV using `systemictl start AV`

### 5-1-2 docker install V2bX
- Prepare the configuration file config.json in advance, refer to [config.json](https://github.com/ppoonk/V2bX/blob/main/config.json)

- The starting Docker command reference is as follows:

```
docker run -tid \
  -v $PWD/av/config.json:/etc/V2bX/config.json \
  --name av \
  --restart always \
  --net=host \
  --privileged=true \
  ppoiuty/av:latest
```

- docker compose is as follows：

```
version: '3'
services:
  AV:
    container_name: AV
    image: ppoiuty/av:latest
    network_mode: "host"
    restart: "always"
    privileged: true
    volumes:
      - ./config.json:/etc/V2bX/config.json
```

## 5-2 XrayR



### 5-2-1 Install XrayR directly

```
bash <(curl -Ls https://raw.githubusercontent.com/ppoonk/XrayR-for-AirGo/main/scripts/manage.sh)
```

- After installation is completed, please modify the configuration file as needed in `/usr/local/XrayR/config. yml`
- Start: Use management script `XrayR` or directly `systemictl start XrayR`

### 5-2-2 docker install XrayR

-Prepare the configuration file config.yml in advance, refer to [config.yml](https://github.com/ppoonk/XrayR-for-AirGo/blob/main/config.yml)

-The starting Docker command reference is as follows:

```
docker run -tid \
  -v $PWD/xrayr/config.yml:/etc/XrayR/config.yml \
  --name xrayr \
  --restart always \
  --net=host \
  --privileged=true \
  ppoiuty/xrayr:latest
```

- docker compose is as follows：

```
version: '3'
services:
  xrayr:
    container_name: xrayr
    image: ppoiuty/xrayr:latest
    network_mode: "host"
    restart: "always"
    privileged: true
    volumes:
      - ./config.yml:/etc/XrayR/config.yml
```

# 6 Update Panel
When updating, please check the `front-end version` and `back-end core version`, which are in different positions and have the same version number, as shown in the figure:
![](https://github.com/ppoonk/AirGo/raw/main/assets/image/8.png)

## 6-1 Update backend

- Method 1: Download a new binary file, replace the old one, and then execute `./AirGo update`
- Method 2: After version `v0.2.5`, the update can be completed by clicking the `Upgrade button` through `Panel - Administrator - System`
- Explanation: After updating the core, the menu and casbin permissions (API permissions) bound to the role will be set to the default values of the current core


## 6-2 Update front-end
Redeploy according to [2-2 front end](#2-2-front-end)


# 7 command line

```
./AirGo help                    Help about any command
./AirGo reset --resetAdmin      reset admin password
./AirGo start                   start AirGo
./AirGo update                  update AirGo
./AirGo version                 show the version of AirGo
```

# 8 develop

[Click to view more](https://github.com/ppoonk/AirGo/wiki)



