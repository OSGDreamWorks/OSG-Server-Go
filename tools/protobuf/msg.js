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
            "name": "StatusInfo",
            "fields": [
                {
                    "rule": "optional",
                    "type": "string",
                    "name": "name",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "uint32",
                    "name": "level",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "uint32",
                    "name": "experience",
                    "id": 3
                },
                {
                    "rule": "optional",
                    "type": "uint32",
                    "name": "HP",
                    "id": 4
                },
                {
                    "rule": "optional",
                    "type": "uint32",
                    "name": "MP",
                    "id": 5
                },
                {
                    "rule": "optional",
                    "type": "uint32",
                    "name": "Rage",
                    "id": 6
                },
                {
                    "rule": "optional",
                    "type": "uint32",
                    "name": "gender",
                    "id": 7
                },
                {
                    "rule": "optional",
                    "type": "uint32",
                    "name": "modelid",
                    "id": 8
                },
                {
                    "rule": "optional",
                    "type": "Transform",
                    "name": "transform",
                    "id": 9
                }
            ]
        },
        {
            "name": "PropertyBaseInfo",
            "fields": [
                {
                    "rule": "optional",
                    "type": "uint32",
                    "name": "Strenght",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "uint32",
                    "name": "Velocity",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "uint32",
                    "name": "Mana",
                    "id": 3
                },
                {
                    "rule": "optional",
                    "type": "uint32",
                    "name": "Defence",
                    "id": 4
                },
                {
                    "rule": "optional",
                    "type": "uint32",
                    "name": "Stamina",
                    "id": 5
                }
            ]
        },
        {
            "name": "PropertyInfo",
            "fields": [
                {
                    "rule": "optional",
                    "type": "uint32",
                    "name": "maxHP",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "uint32",
                    "name": "maxMP",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "uint32",
                    "name": "ATK",
                    "id": 3
                },
                {
                    "rule": "optional",
                    "type": "uint32",
                    "name": "Armor",
                    "id": 4
                },
                {
                    "rule": "optional",
                    "type": "uint32",
                    "name": "Agility",
                    "id": 5
                },
                {
                    "rule": "optional",
                    "type": "uint32",
                    "name": "Spirit",
                    "id": 6
                },
                {
                    "rule": "optional",
                    "type": "uint32",
                    "name": "Recovery",
                    "id": 7
                }
            ]
        },
        {
            "name": "PropertyReviseInfo",
            "fields": [
                {
                    "rule": "optional",
                    "type": "uint32",
                    "name": "UltimateKill",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "uint32",
                    "name": "Hit",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "uint32",
                    "name": "CounterAttack",
                    "id": 3
                },
                {
                    "rule": "optional",
                    "type": "uint32",
                    "name": "Dodge",
                    "id": 4
                }
            ]
        },
        {
            "name": "PropertyAgainstInfo",
            "fields": [
                {
                    "rule": "optional",
                    "type": "uint32",
                    "name": "MagicAgainst",
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
                    "type": "StatusInfo",
                    "name": "stat",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "PropertyBaseInfo",
                    "name": "bp",
                    "id": 3
                },
                {
                    "rule": "optional",
                    "type": "PropertyInfo",
                    "name": "property",
                    "id": 4
                },
                {
                    "rule": "optional",
                    "type": "PropertyReviseInfo",
                    "name": "revise",
                    "id": 5
                },
                {
                    "rule": "optional",
                    "type": "PropertyAgainstInfo",
                    "name": "against",
                    "id": 6
                }
            ]
        },
        {
            "name": "PlayerInfo",
            "fields": [
                {
                    "rule": "required",
                    "type": "string",
                    "name": "uid",
                    "id": 1
                },
                {
                    "rule": "required",
                    "type": "Login",
                    "name": "account",
                    "id": 2
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
        },
        {
            "name": "CreatureBaseInfo",
            "fields": [
                {
                    "rule": "required",
                    "type": "string",
                    "name": "uid",
                    "id": 1
                },
                {
                    "rule": "required",
                    "type": "StatusInfo",
                    "name": "stat",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "PropertyBaseInfo",
                    "name": "bp",
                    "id": 3
                },
                {
                    "rule": "optional",
                    "type": "PropertyInfo",
                    "name": "property",
                    "id": 4
                },
                {
                    "rule": "optional",
                    "type": "PropertyReviseInfo",
                    "name": "revise",
                    "id": 5
                },
                {
                    "rule": "optional",
                    "type": "PropertyAgainstInfo",
                    "name": "against",
                    "id": 6
                }
            ]
        },
        {
            "name": "AttackInfo",
            "fields": [
                {
                    "rule": "required",
                    "type": "uint32",
                    "name": "droptime",
                    "id": 1
                },
                {
                    "rule": "required",
                    "type": "uint32",
                    "name": "attacker",
                    "id": 2
                },
                {
                    "rule": "required",
                    "type": "uint32",
                    "name": "position",
                    "id": 3
                }
            ]
        },
        {
            "name": "Spell",
            "fields": [
                {
                    "rule": "required",
                    "type": "SpellType",
                    "name": "type",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "uint32",
                    "name": "level",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "uint32",
                    "name": "damage",
                    "id": 3
                },
                {
                    "rule": "optional",
                    "type": "uint32",
                    "name": "range",
                    "id": 4
                }
            ]
        },
        {
            "name": "SpellInfo",
            "fields": [
                {
                    "rule": "required",
                    "type": "uint32",
                    "name": "droptime",
                    "id": 1
                },
                {
                    "rule": "required",
                    "type": "uint32",
                    "name": "attacker",
                    "id": 2
                },
                {
                    "rule": "required",
                    "type": "uint32",
                    "name": "position",
                    "id": 3
                },
                {
                    "rule": "required",
                    "type": "Spell",
                    "name": "spell",
                    "id": 4
                }
            ]
        },
        {
            "name": "BattleTest",
            "fields": [
                {
                    "rule": "optional",
                    "type": "CreatureBaseInfo",
                    "name": "enemy",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "PlayerBaseInfo",
                    "name": "player",
                    "id": 2
                }
            ]
        },
        {
            "name": "BattleInfo",
            "fields": [
                {
                    "rule": "required",
                    "type": "string",
                    "name": "bid",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "CreatureBaseInfo",
                    "name": "partner",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "CreatureBaseInfo",
                    "name": "enemy",
                    "id": 3
                },
                {
                    "rule": "optional",
                    "type": "AttackInfo",
                    "name": "attackunits",
                    "id": 4
                },
                {
                    "rule": "optional",
                    "type": "SpellInfo",
                    "name": "spells",
                    "id": 5
                }
            ]
        },
        {
            "name": "NotifyBattleStart",
            "fields": [
                {
                    "rule": "required",
                    "type": "string",
                    "name": "bid",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "CreatureBaseInfo",
                    "name": "partner",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "CreatureBaseInfo",
                    "name": "enemy",
                    "id": 3
                }
            ]
        },
        {
            "name": "BattleAttackQueue",
            "fields": [
                {
                    "rule": "required",
                    "type": "string",
                    "name": "bid",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "AttackInfo",
                    "name": "attackunits",
                    "id": 2
                },
                {
                    "rule": "optional",
                    "type": "SpellInfo",
                    "name": "spells",
                    "id": 3
                }
            ]
        },
        {
            "name": "NotifyBattleEnd",
            "fields": [
                {
                    "rule": "required",
                    "type": "string",
                    "name": "playerlid",
                    "id": 1
                },
                {
                    "rule": "optional",
                    "type": "uint32",
                    "name": "exp",
                    "id": 2
                }
            ]
        }
    ],
    "enums": [
        {
            "name": "SpellType",
            "values": [
                {
                    "name": "LighningStorm",
                    "id": 1
                },
                {
                    "name": "HealingWave",
                    "id": 2
                }
            ]
        }
    ]
}).build();