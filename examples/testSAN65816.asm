; Test assembler file for the Cthulhu Assembler
; Scot W. Stevenson <scot.stevenson@gmail.com>
; First version: 09. Mai 2018
; This version: 10. Mai 2018

        .mpu "65816"
        .origin $00:8000

        .ram $0000 ... $7fff
        .ram $FF01, $FF02   

        .rom $8000 ... $ffff
        .rom $FF03, $FF04 ... $FF06

        .equ dog %0000.1111
        .equ cat { dog 1000 + }

        .native
        .axy16

        .include "somefile.asm"

                lda.# 0000
                tay
myloop:
                sta.y $2000
                sta.y $2000+$100
                dey
                bne myloop
stop:           
                stp ; it's all over, baby!

        .word {frog 2 .swap .dup *}, $ff, "frog"
        .byte 1 ... 4
        .byte 5, 6, 7, 8 ; and that's a wrap!

        .end        
 
