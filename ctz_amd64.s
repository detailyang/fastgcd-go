// +build amd64

// func ctz(x uint64) uint64
TEXT ·ctz(SB),$0-16
    MOVQ x+0(FP), BX
    BSFQ BX, BX
    MOVQ BX, ret+16(FP)
    RET
