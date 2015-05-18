package main

import (
    "app/fightserver"
    "common/logger"
    "common/config"
    "flag"
    "os"
    "common"
    "syscall"
)

var (
    svrConfigFile = flag.String("c", "etc/gameserver.json", "config file name for the fight server")
    fightServerId = flag.Uint64("n", 0, "config id for the fight server")
)

func main() {
    logger.Info("start fight server")

    var cfg config.SvrConfig
    if err := config.ReadConfig(*svrConfigFile, &cfg); err != nil {
        logger.Fatal("load config failed, error is: %v", err)
        return
    }

    fightserver.StartServices(&cfg, fightServerId)

    handler := func(s os.Signal, arg interface{}) {
        logger.Info("handle signal: %v\n", s)
        logger.Info("stop game server")
        os.Exit(0)
    }

    handlerArray := []os.Signal{syscall.SIGINT,
        syscall.SIGILL,
        syscall.SIGFPE,
        syscall.SIGSEGV,
        syscall.SIGTERM,
        syscall.SIGABRT}

    common.WatchSystemSignal(&handlerArray, handler)

    logger.Info("stop fight server")
}
