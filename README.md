![Repository Top Language](https://img.shields.io/github/languages/top/JavaHutt/coinconv)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/JavaHutt/coinconv)
![Github Repository Size](https://img.shields.io/github/repo-size/JavaHutt/coinconv)
![License](https://img.shields.io/badge/license-MIT-green)
![GitHub last commit](https://img.shields.io/github/last-commit/JavaHutt/coinconv)

<img align="right" width="50%" src="./images/flash.png">

# Coin Converter

## HOWTO

- build the app
```bash
make build
```
And you are ready to go! You can convert currencies like this
```bash
./coinconv 123.45 USD BTC,ETH
```
But it is suggested to use coins IDs instead of codes to avoid collisions
```bash
./coinconv 123.45 USD 1,2
```
Be aware that by default app using Sandbox API, you can add your API key in main file and set **IS_TEST** to false.

## A picture is worth a thousand words

<img src="./images/working-example.png">
