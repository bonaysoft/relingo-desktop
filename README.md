# Relingo Desktop

[![main](https://github.com/bonaysoft/relingo-desktop/actions/workflows/main.yml/badge.svg)](https://github.com/bonaysoft/relingo-desktop/actions/workflows/main.yml)
[![](https://img.shields.io/github/downloads/bonaysoft/relingo-desktop/total.svg)](https://github.com/bonaysoft/relingo-desktop/releases)
[![](https://img.shields.io/github/v/release/bonaysoft/relingo-desktop.svg)](https://github.com/bonaysoft/relingo-desktop/releases)
[![](https://img.shields.io/github/license/bonaysoft/relingo-desktop.svg)](https://github.com/bonaysoft/relingo-desktop/blob/master/LICENSE)

这是一个非官方的[Relingo](https://relingo.net/)桌面客户端，主要用于背单词。

<a href="#%EF%B8%8F-安装" target="blank"><strong>📦️ 下载安装包</strong></a>&nbsp;&nbsp;|&nbsp;&nbsp;
<a href="https://t.me/relingodesktop" target="_blank"><strong>💬 加入交流群</strong></a>

## ✨ 特性

- 生词词频统计
- 每日学习记录
- 快速将生词标记为掌握
- 基于词根词缀的助记功能
- 基于艾宾浩斯遗忘曲线的复习功能
- 更多特性开发中

## 📦️ 安装

访问本项目的 [Releases](https://github.com/bonaysoft/relingo-desktop/releases) 页面下载安装包。

- macOS 用户可以通过 Homebrew 来安装：`brew install --cask saltbo/bin/relingo`

## 🆙 更新

访问本项目的 [Releases](https://github.com/bonaysoft/relingo-desktop/releases) 页面下载安装包覆盖安装即可。

- macOS 用户可以通过 Homebrew 来安装：`brew update && brew upgrade --cask saltbo/bin/relingo`


## ⚙️ 使用

1. 首次运行软件会提示下载证书
2. 下载证书后，打开证书进行信任操作
3. 通过 [SwitchyOmega](https://chrome.google.com/webstore/detail/proxy-switchyomega/padekgcemlokbadohgkifijomclgjgif?hl=en)
插件或者Clash规则将api.relingo.net的请求代理到本软件的代理端口8119
4. 刷新任意页面触发relingo插件的加载

## FAQ

### 1. 为什么要拦截api.relingo.net的请求？

答：为了统计每日词频，官方的接口没有提供这个数据，我们通过拦截官方接口方式对单词数据进行统计并存储在您本地

### 2. 为什么要信任证书？

答：官方域名使用https提供服务，拦截请求需要获得您的授权。

### 3. 信任证书后如何保证没有拦截我的其他数据？

答：项目代码是开源的，您可以在pkg/proxy/proxy.go中看到所有逻辑。
如果还不放心，建议您通过SwitchyOmega插件或者Clash规则只把api.relingo.net的请求转发到8119.

## 📜 开源许可

本项目仅供个人学习使用，禁止用于商业用途。

基于 Apache-2.0 license 许可进行开源。

## 🖼️ 截图

<img width="1200" alt="image" src="https://user-images.githubusercontent.com/17308208/227981740-f7c3739e-22fd-485b-9ac7-d1a7a557e45c.png">
<img width="1200" alt="image" src="https://user-images.githubusercontent.com/17308208/227981900-e9e96e43-fb95-4366-b345-87fa71893e34.png">


