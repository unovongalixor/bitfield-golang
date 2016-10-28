package bitfield

type Bitfield struct {
    index      int64 // reading index
    autosizing bool
    bytes      []byte
}

func NewBitfield(autosizing bool, size int) *Bitfield {
    bf := Bitfield{}
    bf.autosizing = autosizing
    bf.bytes = make([]byte, size)

    return &bf
}

func (bf *Bitfield) Size() int64 { 
    return int64(len(bf.bytes)) 
}

func (bf *Bitfield) Bytes() []byte {
    return bf.bytes
}

func (bf *Bitfield) Grow(n int64) {
    if n < 0 {
        panic("bitfield.Grow: negative count")
    }

    new_bytes := make([]byte, int64(len(bf.bytes) + 1) + n)
    copy(bf.bytes, new_bytes)
    bf.bytes = new_bytes
}

func (bf *Bitfield) SetBit(n int) {
    index := int64(n / 4)
    bit := uint16(n % 4)

    if index > bf.Size() {
        if bf.autosizing == true {
            bf.Grow(index - bf.Size())
        } else {
            panic("INVALID INDEX")
        }
    }

    var flag byte = 1   // flag = 0001
    flag = flag << bit // shift flag to bit we want to set
    bf.bytes[index] = bf.bytes[index] | flag
}

func (bf *Bitfield) ClearBit(n int) {
    index := int64(n / 4)
    bit := uint16(n % 4)

    if index > bf.Size() {
        panic("INVALID INDEX")
    }

    var flag byte = 1   // flag = 0001
    flag = flag << bit // shift flag to bit we want to set
    flag = ^flag
    bf.bytes[index] = bf.bytes[index] & flag
}

func (bf *Bitfield) GetBit(n int) bool {
    index := int64(n / 4)
    bit := uint16(n % 4)

    if index > bf.Size() {
        return false
    }

    var flag byte = 1   // flag = 0001
    flag = flag << bit // shift flag to bit we want to set
    return (bf.bytes[index] & flag) > 0;
}