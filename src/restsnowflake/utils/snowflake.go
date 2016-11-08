package utils

import (
        "sync"
        "time"
        "fmt"
)

var snowflakeList map[uint32]*SnowFlake = make(map[uint32]*SnowFlake)

const (
        nano = 1000 * 1000
)

const (
        WorkerIdBits = 10                                          // worker id
        MaxWorkerId  = -1 ^ (-1 << WorkerIdBits)                   // worker id mask
        SequenceBits = 12                                          // sequence
        MaxSequence  = -1 ^ (-1 << SequenceBits)                   // sequence mask
)

var (
        Since int64 = time.Date(2010, 1, 0, 0, 0, 0, 0, time.UTC).UnixNano() / nano
)

type SnowFlake struct {
        lastTimestamp uint64
        workerId      uint32
        sequence      uint32
        lock          sync.Mutex
}

func (sf *SnowFlake) uint64() uint64 {
        return (sf.lastTimestamp << (WorkerIdBits + SequenceBits)) |
                (uint64(sf.workerId) << SequenceBits) |
                (uint64(sf.sequence))
}

func (sf *SnowFlake) Next() (uint64, error) {
        sf.lock.Lock()
        defer sf.lock.Unlock()

        ts := timestamp()
        if ts == sf.lastTimestamp {
                sf.sequence = (sf.sequence + 1) & MaxSequence
                if sf.sequence == 0 {
                        ts = tilNextMillis(ts)
                }
        } else {
                sf.sequence = 0
        }

        if ts < sf.lastTimestamp {
                return 0, fmt.Errorf("Invalid timestamp: %v - precedes %v", ts, sf)
        }
        sf.lastTimestamp = ts
        return sf.uint64(),  nil
}

func NewSnowFlake(workerId uint32) (*SnowFlake, error) {
        if workerId < 0 || workerId > MaxWorkerId {
                return nil, fmt.Errorf("Worker id %v is invalid", workerId)
        }
        return &SnowFlake{workerId: workerId}, nil
}

func timestamp() uint64 {
        return uint64(time.Now().UnixNano()/nano - Since)
}

func tilNextMillis(ts uint64) uint64 {
        i := timestamp()
        for i < ts {
                i = timestamp()
        }
        return i
}

func GetSnowFlakeObj(workerid uint32)(sf *SnowFlake,err error){
        if v, ok := snowflakeList[workerid]; ok {
                fmt.Println("find obj workerid:",workerid)
                return v,nil
        } else {
                sf,_:=NewSnowFlake(workerid)
                snowflakeList[workerid]=sf
                fmt.Println("create obj workerid:",workerid)
                return sf,nil
        }
}
