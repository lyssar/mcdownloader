<h1 align="center">Welcome to Minecraft Server Download CLI 👋</h1>
<p>
  <img alt="Version" src="https://img.shields.io/badge/version-1.2.0-blue.svg?cacheSeconds=2592000" />
  <a href="./docs" target="_blank">
    <img alt="Documentation" src="https://img.shields.io/badge/documentation-yes-brightgreen.svg" />
  </a>
  <a href="https://opensource.org/licenses/MIT" target="_blank">
    <img alt="License: MIT" src="https://img.shields.io/badge/License-MIT-yellow.svg" />
  </a>
  <a href="https://twitter.com/lyssar\_" target="_blank">
    <img alt="Twitter: lyssar\_" src="https://img.shields.io/twitter/follow/lyssar_.svg?style=social" />
  </a>
</p>

> msdcli is a golang based cli for linux to easily download and setup a minecraft server with one of the existing launcher like spigot, papermc, forge or fabric.

## Install

```sh
export MSDVERSION="1.2.0"
curl -Lo /usr/local/bin/msdcli "https://github.com/lyssar/msdcli/releases/download/${MSDVERSION}/msdcli-amd64"
chmod ugo+x /usr/local/bin/msdcli
```

## Usage

```sh
# To setup a server; in this case mincraft 1.18.1 with forge 39.0.59
msdcli server -mcversion "1.18.1" -type "forge" -serverVersion "39.0.59"

# To get a modpack from curseforge
msdcli modpack -packageId 495683 -serverPackageFileID 3620338
```

## Author

👤 **Sebastian Hens**

* Website: http://lyssar.me/
* Twitter: [@lyssar\_\_](https://twitter.com/lyssar__)
* Github: [@lyssar](https://github.com/lyssar)

## 🤝 Contributing

Contributions, issues and feature requests are welcome!<br />Feel free to check [issues page](https://github.com/lyssar/mcdownloader/issues). 

## Show your support

Give a ⭐️ if this project helped you!

## 📝 License

Copyright © 2022 [Sebastian Hens](https://github.com/lyssar).<br />
This project is [MIT](https://opensource.org/licenses/MIT) licensed.

***
_This README was generated with ❤️ by [readme-md-generator](https://github.com/kefranabg/readme-md-generator)_