var _root = dcodeIO.ProtoBuf.newBuilder({})['import']({
    "package": "protobuf",
    "messages": [
        {
            "name": "Request",
            "fields": [
                {
                    "rule": "required",
                    "type": "uint64",
                    "name": "id",
                    "id": 1
                },
                {
                    "rule": "required",
                    "type": "string",
                    "name": "method",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "bytes",
                    "name": "serialized_request",
                    "id": 3
                },
                {
                    "rule": "optional",
                    "type": "string",
                    "name": "timer",
                    "id": 4
                },
                {
                    "rule": "optional",
                    "type": "string",
                    "name": "desc",
                    "id": 5
                }
            ],
            "extensions": [
                100,
                536870911
            ]
        },
        {
            "name": "RpcErrorResponse",
            "fields": [
                {
                    "rule": "required",
                    "type": "string",
                    "name": "method",
                    "id": 1
                },
                {
                    "rule": "required",
                    "type": "string",
                    "name": "text",
                    "id": 2
                }
            ]
        },
        {
            "name": "ConnectorInfo",
            "fields": [
                {
                    "rule": "required",
                    "type": "uint32",
                    "name": "ServerId",
                    "id": 1
                },
                {
                    "rule": "required",
                    "type": "uint32",
                    "name": "PlayerCount",
                    "id": 2
                },
                {
                    "rule": "required",
                    "type": "string",
                    "name": "TcpServerIp",
                    "id": 3
                },
                {
                    "rule": "required",
                    "type": "string",
                    "name": "HttpServerIp",
                    "id": 4
                }
            ]
        },
        {
            "name": "ConnectorInfoResult",
            "fields": [
                {
                    "rule": "required",
                    "type": "Result",
                    "name": "result",
                    "id": 1,
                    "options": {
                        "default": "OK"
                    }
                },
                {
                    "rule": "required",
                    "type": "uint32",
                    "name": "server_time",
                    "id": 2
                }
            ],
            "enums": [
                {
                    "name": "Result",
                    "values": [
                        {
                            "name": "OK",
                            "id": 0
                        },
                        {
                            "name": "ERROR",
                            "id": 1
                        }
                    ]
                }
            ]
        },
        {
            "name": "LoginInfo",
            "fields": [
                {
                    "rule": "required",
                    "type": "string",
                    "name": "serverIp",
                    "id": 1
                },
                {
                    "rule": "required",
                    "type": "string",
                    "name": "gsInfo",
                    "id": 2
                }
            ]
        },
        {
            "name": "Login",
            "fields": [
                {
                    "rule": "optional",
                    "type": "string",
                    "name": "uid",
                    "id": 1
                },
                {
                    "rule": "required",
                    "type": "string",
                    "name": "account",
                    "id": 2
                },
                {
                    "rule": "required",
                    "type": "string",
                    "name": "password",
                    "id": 3
                },
                {
                    "rule": "optional",
                    "type": "uint32",
                    "name": "create_time",
                    "id": 4
                },
                {
                    "rule": "optional",
                    "type": "string",
                    "name": "option",
                    "id": 5
                },
                {
                    "rule": "optional",
                    "type": "uint32",
                    "name": "language",
                    "id": 6
                },
                {
                    "rule": "optional",
                    "type": "string",
                    "name": "udid",
                    "id": 7
                },
                {
                    "rule": "optional",
                    "type": "string",
                    "name": "sessionKey",
                    "id": 8
                }
            ]
        },
        {
            "name": "LoginResult",
            "fields": [
                {
                    "rule": "required",
                    "type": "Result",
                    "name": "result",
                    "id": 1,
                    "options": {
                        "default": "OK"
                    }
                },
                {
                    "rule": "required",
                    "type": "uint32",
                    "name": "server_time",
                    "id": 2
                },
                {
                    "rule": "required",
                    "type": "string",
                    "name": "sessionKey",
                    "id": 3
                },
                {
                    "rule": "required",
                    "type": "string",
                    "name": "uid",
                    "id": 4
                },
                {
                    "rule": "optional",
                    "type": "string",
                    "name": "errmsg",
                    "id": 5
                }
            ],
            "enums": [
                {
                    "name": "Result",
                    "values": [
                        {
                            "name": "OK",
                            "id": 0
                        },
                        {
                            "name": "SERVERERROR",
                            "id": 1
                        },
                        {
                            "name": "USERNOTFOUND",
                            "id": 2
                        },
                        {
                            "name": "AUTH_FAILED",
                            "id": 3
                        },
                        {
                            "name": "ISONFIRE",
                            "id": 4
                        }
                    ]
                }
            ]
        },
        {
            "name": "Ping",
            "fields": []
        },
        {
            "name": "PingResult",
            "fields": [
                {
                    "rule": "required",
                    "type": "uint32",
                    "name": "server_time",
                    "id": 1
                }
            ]
        },
        {
            "name": "Chat",
            "fields": [
                {
                    "rule": "required",
                    "type": "string",
                    "name": "msg",
                    "id": 1
                }
            ]
        },
        {
            "name": "ChatResult",
            "fields": [
                {
                    "rule": "required",
                    "type": "string",
                    "name": "msg",
                    "id": 1
                }
            ]
        },
        {
            "name": "PlayerBaseInfo",
            "fields": [
                {
                    "rule": "required",
                    "type": "string",
                    "name": "uid",
                    "id": 1
                },
                {
                    "rule": "required",
                    "type": "string",
                    "name": "name",
                    "id": 2
                },
                {
                    "rule": "required",
                    "type": "uint32",
                    "name": "level",
                    "id": 3
                },
                {
                    "rule": "required",
                    "type": "uint32",
                    "name": "experience",
                    "id": 4
                },
                {
                    "rule": "required",
                    "type": "uint32",
                    "name": "HP",
                    "id": 6
                },
                {
                    "rule": "required",
                    "type": "uint32",
                    "name": "MP",
                    "id": 7
                },
                {
                    "rule": "required",
                    "type": "uint32",
                    "name": "Rage",
                    "id": 8
                },
                {
                    "rule": "required",
                    "type": "uint32",
                    "name": "maxHP",
                    "id": 9
                },
                {
                    "rule": "required",
                    "type": "uint32",
                    "name": "maxMP",
                    "id": 10
                },
                {
                    "rule": "required",
                    "type": "uint32",
                    "name": "maxRage",
                    "id": 11
                },
                {
                    "rule": "optional",
                    "type": "uint32",
                    "name": "gender",
                    "id": 12
                },
                {
                    "rule": "optional",
                    "type": "uint32",
                    "name": "modelid",
                    "id": 13
                },
                {
                    "rule": "optional",
                    "type": "Transform",
                    "name": "transform",
                    "id": 14
                },
                {
                    "rule": "optional",
                    "type": "uint32",
                    "name": "Strenght",
                    "id": 15
                },
                {
                    "rule": "optional",
                    "type": "uint32",
                    "name": "Velocity",
                    "id": 16
                },
                {
                    "rule": "optional",
                    "type": "uint32",
                    "name": "Mana",
                    "id": 17
                },
                {
                    "rule": "optional",
                    "type": "uint32",
                    "name": "Defence",
                    "id": 18
                },
                {
                    "rule": "optional",
                    "type": "uint32",
                    "name": "Stamina",
                    "id": 19
                },
                {
                    "rule": "optional",
                    "type": "uint32",
                    "name": "ATK",
                    "id": 20
                },
                {
                    "rule": "optional",
                    "type": "uint32",
                    "name": "Armor",
                    "id": 21
                },
                {
                    "rule": "optional",
                    "type": "uint32",
                    "name": "Agility",
                    "id": 22
                },
                {
                    "rule": "optional",
                    "type": "uint32",
                    "name": "Spirit",
                    "id": 23
                },
                {
                    "rule": "optional",
                    "type": "uint32",
                    "name": "Recovery",
                    "id": 24
                }
            ]
        },
        {
            "name": "PlayerInfo",
            "fields": [
                {
                    "rule": "required",
                    "type": "PlayerBaseInfo",
                    "name": "base",
                    "id": 1
                }
            ]
        },
        {
            "name": "Vector3",
            "fields": [
                {
                    "rule": "required",
                    "type": "float",
                    "name": "X",
                    "id": 1
                },
                {
                    "rule": "required",
                    "type": "float",
                    "name": "Y",
                    "id": 2
                },
                {
                    "rule": "required",
                    "type": "float",
                    "name": "Z",
                    "id": 3
                }
            ]
        },
        {
            "name": "Quaternion",
            "fields": [
                {
                    "rule": "required",
                    "type": "float",
                    "name": "X",
                    "id": 1
                },
                {
                    "rule": "required",
                    "type": "float",
                    "name": "Y",
                    "id": 2
                },
                {
                    "rule": "required",
                    "type": "float",
                    "name": "Z",
                    "id": 3
                },
                {
                    "rule": "required",
                    "type": "float",
                    "name": "W",
                    "id": 4
                }
            ]
        },
        {
            "name": "Transform",
            "fields": [
                {
                    "rule": "required",
                    "type": "Vector3",
                    "name": "position",
                    "id": 1
                },
                {
                    "rule": "required",
                    "type": "Quaternion",
                    "name": "rotation",
                    "id": 2
                },
                {
                    "rule": "required",
                    "type": "Vector3",
                    "name": "scale",
                    "id": 3
                }
            ]
        }
    ]
}).build();