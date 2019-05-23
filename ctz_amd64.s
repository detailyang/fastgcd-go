// +build amd64

#include "textflag.h"

// func ctz(x uint64) uint64
TEXT ·ctz(SB),NOSPLIT,$0-16
    MOVQ x+0(FP), BX
    BSFQ BX, BX
    MOVQ BX, ret+8(FP)
    RET
