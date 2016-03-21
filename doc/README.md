# OSG-Server-Go 文档

- [框架介绍](#框架介绍)
	- [核心组件](#核心组件)
	- [多服示例](#多服示例)
	- [Protobuf](#protobuf)
- [Go代码风格](#go代码风格)
	- [登陆流程](#登陆流程)
- [Lua代码风格](#lua代码风格)
	- [战斗流程](#战斗流程)
- [Js代码风格](#js代码风格)

## 框架介绍
Framework introduction

### 核心组件
The Core Component

### 多服示例
The Servers Demo

### Protobuf
The Protobuf Style

## Go代码风格
Go coding Style

### 登陆流程
Login Process
[登陆流程](http://git.oschina.net/3dseals/OSG-Server-Go/blob/master/doc/Login_Process.md "登陆流程")
首先客户端发送10002[CL_CheckAccount]传给登陆服账号密码 登陆服收到账号密码后传给认证服，通过认证后返回给客户端是否成功的结果 然后客户端拿着成功登陆的uid和sessionKey再到游戏服调用20002[CS_CheckSession]二次确认 游戏服查找缓存中的sessionKey来决定允许客户端的访问

## Lua代码风格
Lua coding Style

### 战斗流程
Fighting Process
[战斗流程](http://git.oschina.net/3dseals/OSG-Server-Go/blob/master/doc/Fighting_Process.md "战斗流程")

## Js代码风格
Js coding Style
(Need to complete)