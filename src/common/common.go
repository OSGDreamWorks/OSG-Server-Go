package common

import (
    "os"
    "fmt"
    "os/signal"
    "encoding/base64"
    "crypto/md5"
    "io"
    "time"
    "sync/atomic"
    "crypto/rc4"
    "encoding/binary"
    "common/logger"
)

func Base64Encode(src []byte) []byte {
    return []byte(base64.URLEncoding.EncodeToString(src))
}

func Base64Decode(src []byte) ([]byte, error) {
    return base64.URLEncoding.DecodeString(string(src))
}

//uuid
var uuid uint32 = 0

const (
    UUIDkey = "MSaUWm6bEUyNTaxcz0DiNICZ4BGmC5AXxOkUMhT7fEsq136lIeaH81kdIYlk3xq4"
    MD5key = "FBz4IeDQPwM7EzyXlPIG71Ud8GD2lxEntx7LxZAPcbxRpHeCfFxyFA0FdcFU0gNR"
)

// UUID() provides unique identifier strings.
func GenUUID(account string) string {
    h1 := md5.New()
    io.WriteString(h1, account)
    io.WriteString(h1, UUIDkey)
    h2 := md5.New()
    io.WriteString(h2, account)
    io.WriteString(h2, MD5key)
    return fmt.Sprintf("%x%x", h1.Sum(nil), h2.Sum(nil))
}

func CheckUUID(uid string,account string) bool {
    h1 := md5.New()
    io.WriteString(h1, account)
    io.WriteString(h1, UUIDkey)
    h2 := md5.New()
    io.WriteString(h2, account)
    io.WriteString(h2, MD5key)
    return uid == fmt.Sprintf("%x%x", h1.Sum(nil), h2.Sum(nil))
}

// UUID() provides unique identifier strings.
func GenPassword(account string, passwd string) string {
    h := md5.New()
    io.WriteString(h, account)
    io.WriteString(h, passwd)
    io.WriteString(h, UUIDkey)
    return fmt.Sprintf("%x", h.Sum(nil))
}

func CheckPassword(hashkey string,account string, passwd string) bool {
    h := md5.New()
    io.WriteString(h, account)
    io.WriteString(h, passwd)
    io.WriteString(h, UUIDkey)
    return hashkey == fmt.Sprintf("%x", h.Sum(nil))
}

// UUID() provides unique identifier strings.
func GenSessionKey() string {

    b := make([]byte, 16)

    t := time.Now().Unix()
    tmpid := uint16(atomic.AddUint32(&uuid, 1))

    b[0] = byte(255)
    b[1] = byte(0)
    b[2] = byte(tmpid)
    b[3] = byte(tmpid >> 8)

    b[4] = byte(t)
    b[5] = byte(t >> 8)
    b[6] = byte(t >> 16)
    b[7] = byte(t >> 24)

    c, _ := rc4.NewCipher([]byte{0x0c, b[2], b[3], b[0]})
    c.XORKeyStream(b[8:], b[:8])

    guid := fmt.Sprintf("%x-%x-%x-%x-%x", b[:4], b[4:6], b[6:8], b[8:12], b[12:])
    h := md5.New()
    io.WriteString(h, guid)
    io.WriteString(h, MD5key)

    return fmt.Sprintf("%x-%x-%x-%x-%x--%x", b[:4], b[4:6], b[6:8], b[8:12], b[12:], h.Sum(nil))
}

func CheckSessionKey(skey string) bool {
    if len(skey) != 70 {
        return false
    }

    b := make([]uint32, 5)
    var s string
    guid := skey[:36]

    _, err := fmt.Sscanf(skey, "%x-%x-%x-%x-%x--%s", &b[0], &b[1], &b[2], &b[3], &b[4], &s)

    if err != nil {
        logger.Debug("err : %v", err.Error())
        return false
    }

    info1 := make([]byte, 4)
    binary.BigEndian.PutUint32(info1, b[0])

    info2 := make([]byte, 4)
    binary.BigEndian.PutUint16(info2[:2], uint16(b[1]))
    binary.BigEndian.PutUint16(info2[2:], uint16(b[2]))

    c, _ := rc4.NewCipher([]byte{0x0c, info1[2], info1[3], info1[0]})

    tmp := make([]byte, 4)

    c.XORKeyStream(tmp, info1)

    if binary.BigEndian.Uint32(tmp) != b[3] {
        return false
    }

    c.XORKeyStream(tmp, info2)

    if binary.BigEndian.Uint32(tmp) != b[4] {
        return false
    }

    h := md5.New()
    io.WriteString(h, guid)
    io.WriteString(h, MD5key)

    if s != fmt.Sprintf("%x", h.Sum(nil)) {
        logger.Debug("%s, %x", guid, h.Sum(nil))
        return false
    }

    return true
}

type SignalHandler func(s os.Signal, arg interface{})

type signalSet struct {
    m map[os.Signal]SignalHandler
}

func signalSetNew() *signalSet {
    ss := new(signalSet)
    ss.m = make(map[os.Signal]SignalHandler)
    return ss
}

func (set *signalSet) register(s os.Signal, handler SignalHandler) {
    if _, found := set.m[s]; !found {
        set.m[s] = handler
    }
}

func (set *signalSet) handle(sig os.Signal, arg interface{}) (err error) {
    if _, found := set.m[sig]; found {
        set.m[sig](sig, arg)
        return nil
    } else {
        return fmt.Errorf("No handler available for signal %v", sig)
    }

    panic("won't reach here")
}

func WatchSystemSignal(watchsingals *[]os.Signal, callbackHandler SignalHandler) {
    ss := signalSetNew()

    for _, wathsingnal := range *watchsingals {
        ss.register(wathsingnal, callbackHandler)
    }

    for {
        c := make(chan os.Signal)
        var sigs []os.Signal
        for sig := range ss.m {
            sigs = append(sigs, sig)
        }
        signal.Notify(c)
        sig := <-c

        err := ss.handle(sig, nil)
        if err != nil {
            fmt.Printf("unknown signal received: %v\n", sig)
        }
    }
}
