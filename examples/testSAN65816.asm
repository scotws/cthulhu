; Test assembler file for the Cthulhu Assembler
; Scot W. Stevenson <scot.stevenson@gmail.com>
; First version: 09. Mai 2018
; This version: 10. Mai 2018

        .mpu "65816"
        .origin $00:8000

        .equ frog %0000.1111

        .native
        .axy16

                lda.# 0000
                tay
_loop:
@
                sta.y $2000
                sta.y $2000+$100
                dey
                bne _loop
                bne -

                stp

        .word {frog 2 .swap .dup *}, $ff, "frog"
        .byte 1, 2, 3, 4
        .byte 5, 6, 7, 8 ; and that's a wrap!

        .end        
 
