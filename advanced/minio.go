package advanced

type BytePoolCap struct {
    c    chan []byte
    w    int
    wcap int
}

func (bp *BytePoolCap) Get() (b []byte) {
    select {
    case b = <-bp.c:
    // reuse existing buffer
    default:
        // create new buffer
        if bp.wcap > 0 {
            b = make([]byte, bp.w, bp.wcap)
        } else {
            b = make([]byte, bp.w)
        }
    }
    return
}

func (bp *BytePoolCap) Put(b []byte) {
    select {
    case bp.c <- b:
        // buffer went back into pool
    default:
        // buffer didn't go back into pool, just discard
    }
}

func NewBytePoolCap(maxSize int, width int, capwidth int) (bp *BytePoolCap) {
	// 工厂函数，用于生成*BytePoolCap
    return &BytePoolCap{
        c:    make(chan []byte, maxSize),
        w:    width,
        wcap: capwidth,
    }
}