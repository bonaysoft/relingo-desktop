# Relingo Desktop

[![main](https://github.com/bonaysoft/relingo-desktop/actions/workflows/main.yml/badge.svg)](https://github.com/bonaysoft/relingo-desktop/actions/workflows/main.yml)
[![](https://img.shields.io/github/downloads/bonaysoft/relingo-desktop/total.svg)](https://github.com/bonaysoft/relingo-desktop/releases)
[![](https://img.shields.io/github/v/release/bonaysoft/relingo-desktop.svg)](https://github.com/bonaysoft/relingo-desktop/releases)
[![](https://img.shields.io/github/license/bonaysoft/relingo-desktop.svg)](https://github.com/bonaysoft/relingo-desktop/blob/master/LICENSE)

这是一个非官方的Relingo桌面客户端，主要用于背单词。

## ✨ 特性

- 生词词频统计
- 每日学习记录
- 快速将生词标记为掌握
- 基于词根词缀的助记功能
- 基于艾宾浩斯遗忘曲线的复习功能
- 更多特性开发中

## 📦️ 安装

访问本项目的 [Releases](https://github.com/bonaysoft/relingo-desktop/releases) 页面下载安装包。

## ⚙️ 使用

1. 首次运行软件会提示下载证书
2. 下载证书后进行信任操作
3. 通过SwitchyOmega插件或者Clash规则将api.relingo.net的请求代理到本软件的代理端口8119
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

<img width="1024" alt="image" src="https://user-images.githubusercontent.com/17308208/226173782-ba597df0-559d-4bba-88a8-d73d99e312ca.png">

