; Test file for GoAsm65816 
; Scot W. Stevenson <scot.stevenson@gmail.com>
; First version: 06. May 2018
; This version: 06. May 2018

        .ram $0000-$7FFF
        .rom $8000-$FFFF
        .ram $FFF0, $FF01

        .notation wdc
        .origin $8000
        .mpu 65c02

        .equ target1 $2fff
        .equ target2 $30aa
:start
                nop             ; this really does nothing

                ldx #10
@
                lda #"a"
                sta target1,x   ; note bank byte separator
                sta target2,x
                sta target1+1,x
                dex
                bne -

                lda #%0000.0000.1111.0000

                jmp start       

; Silly subroutine. Call with char value in A
        .scope
:got_a?
                ldy #00        ; false
                tyx
_loop
                lda check_a,x
                beq done
                cmp #"a"
                beq found!
                inx
                bra loop

_found!
                dey             ; to $ff
_done
                rts
        .scend

:stuff
        .byte 0, $0A, 2, 2+1, "4", %0000:1111 ; just a test

:check_a
        .byte "This is a string", 0
        .end        
