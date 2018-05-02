;; Test file for GoAsm65816 
;; Scot W. Stevenson <scot.stevenson@gmail.com>
;; First version: 21. Apr 2018
;; This version: 21. Apr 2018

;; Code is in Simpler Assember Notation

        .mpu 65816
        .origin $8000
        .native         ; it gets serious

        .equ target1 $00:2fff
        .equ target2 $0030aa
start:
                nop             ; this really does nothing

                ldx.# 0010
@
                lda.# "a"
                sta.x target1   ; note bank byte separator
                sta.x target2
                dex
                bne -

                lda.# %11110000

                jmp start       

;; Silly subroutine. Call with char value in A
.scope
got_a?:
        .axy8
                ldy.# 00        ; false
                ldx.# 00
        _loop:
                lda.x check_a
                beq _done
                cmp.# "a"
                beq _found!
                inx
                bra _loop

_found!:
                dey             ; to $ff
_done:
        .axy16
                rts
.scend

stuff:  .byte 0, $01, 2, 2+1, "4", %0000:1111 ; just a test

check_a:
        .byte "This is a string", 0
        .end        
