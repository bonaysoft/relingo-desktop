# Relingo Desktop

这是一个非官方的Relingo桌面客户端，为了统计词频，我们采用了拦截官方插件请求的方式进行实现，所以需要使用SwitchyOmega插件将api.relingo.net的请求代理到本软件的代理端口。

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
