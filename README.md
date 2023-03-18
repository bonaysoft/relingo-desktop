# Relingo Desktop

这是一个非官方的Relingo桌面客户端，为了统计词频，我们采用了拦截官方插件请求的方式进行实现，所以需要使用SwitchyOmega插件将api.relingo.net的请求代理到本软件的代理端口。

每日任务：记忆10个单词。

点击单词出现一个记忆面板，采用词根记忆法帮助记忆单词。每个单词需要造一个句子，完成本次记忆。


按照艾宾浩斯记忆曲线，将记忆过的单词周期性的出现在当天待记忆列表中。

记忆面板展示内容

1. 单词
2. Tag
3. 音标 朗读
4. 词性，含义
5. 构词
6. 例句


## About

## 安装

### MacOS

TODO

### Windows

TODO

## Live Development

To run in live development mode, run `wails dev` in the project directory. This will run a Vite development
server that will provide very fast hot reload of your frontend changes. If you want to develop in a browser
and have access to your Go methods, there is also a dev server that runs on http://localhost:34115. Connect
to this in your browser, and you can call your Go code from devtools.

## Building

To build a redistributable, production mode package, use `wails build`.
