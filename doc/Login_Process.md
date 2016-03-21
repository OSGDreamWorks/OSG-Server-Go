# 登陆流程

## 数据结构

### 客户端

```cpp
message CL_CheckAccount{
	optional string uid = 1;			//可选玩家uid 全服唯一，通过服务端返回记录，第一次登陆可以传空
	required string account = 2;		//玩家账号
	required string password = 3;		//玩家密码
	optional string option = 4;			//登陆参数可选项，用于第三方认证特殊参数，json格式字符串
	optional uint32 language = 5;		//客户端语言
	optional string udid = 6;			//客户端设备id
}


message LC_CheckAccountResult{
	enum Result {
		OK = 1;								//登陆成功
		SERVERERROR = 2;					//服务端内部错误
		USERNOTFOUND = 3;					//用户没有找到
		AUTH_FAILED = 4;					//认证失败
		ISONFIRE = 5;						//用户已经在线
	}
	required Result result = 1;				//登录返回结果
	required uint32 server_time = 2;		//服务器时间
	required string sessionKey = 3;			//登陆后服务端返回的会话key， 需要在游戏服验证正确性
	required string uid = 4;				//服务端返回的玩家uid
	optional string gameServerIp = 5;		//游戏服的链接地址，客户端会被自动分配一个低负载的游戏服
}


message CS_CheckSession{
	required string uid = 1;				//服务端返回的玩家uid
	required string sessionKey = 2;			//登陆后服务端返回的会话key， 需要在游戏服验证正确性
	optional uint32 timestamp = 3;			//服务器时间
}

message SC_CheckSessionResult{
	enum Result {
		OK = 1;
		SERVERERROR = 2;
		USERNOTFOUND = 3;
		AUTH_FAILED = 4;
		ISONFIRE = 5;
	}
	required Result result = 1;
	optional uint32 server_time = 2;
	optional string errmsg = 3;
}
```

### 服务端

```cpp
message LA_CheckAccount{
	optional string uid = 1;
	required string account = 2;
	required string password = 3;
	optional string option = 4;
	optional uint32 language = 5;
	optional string udid = 6;
}

message AL_CheckAccountResult{
	enum Result {
		OK = 1;
		SERVERERROR = 2;
		USERNOTFOUND = 3;
		AUTH_FAILED = 4;
		ISONFIRE = 5;
	}
	required Result result = 1;
	required string sessionKey = 2;
	required string uid = 3;
	optional string errmsg = 4;
}

message SL_UpdatePlayerCount{
	required uint32 ServerId = 1;				//游戏服id
	required uint32 PlayerCount = 2;			//当前连接玩家数
	required string TcpServerIp = 3;			//tcp 链接地址
	required string HttpServerIp = 4;			//http链接地址
}

message LS_UpdatePlayerCountResult{
	enum Result {
		OK = 1;
		SERVERERROR = 2;
	}
	required Result result = 1;
	required uint32 server_time = 2;
}
```

## 接口

### 客户端

```cpp
enum CL_Protocol {
	eCL_CheckAccount 		= 10002;	//客户端发给服务端登陆接口
}

enum LC_Protocol {
	eLC_CheckAccountResult	= 10502;	//服务端返回给客户端的回调
}


enum CS_Protocol {
	eCS_CheckSession 		= 20002;		//检测会话key
}

enum SC_Protocol {
	eSC_CheckSessionResult 	= 30002;		//返回检测会话的结果
}
```

### 服务端

```cpp
enum LA_Protocol {
	eLA_CheckAccount 		= 11002;	//服务端发给认证服接口
}

enum AL_Protocol {
	eAL_CheckAccountResult  = 12002;	//认证服回调的结果
}

enum LS_Protocol {
	eLS_UpdatePlayerCountResult	= 13002;
}

enum SL_Protocol {
	eSL_UpdatePlayerCount	= 14002;	//游戏服发给登陆服的在线玩家连接数
}
```

## 流程

首先客户端发送10002[CL_CheckAccount]传给登陆服账号密码
登陆服收到账号密码后传给认证服，通过认证后返回给客户端是否成功的结果
然后客户端拿着成功登陆的uid和sessionKey再到游戏服调用20002[CS_CheckSession]二次确认
游戏服查找缓存中的sessionKey来决定允许客户端的访问