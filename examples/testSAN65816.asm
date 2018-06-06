; Test assembler file for the Cthulhu Assembler
; Scot W. Stevenson <scot.stevenson@gmail.com>
; First version: 09. Mai 2018
; This version: 10. Mai 2018

        .mpu "65816"
        .origin $00:8000

        .ram $FF00
        .ram $FF01, $FF02   
        .ram $0000 ... $7fff

        .rom $FF03, $FF04 ... $FF05
        .rom $8000 ... $EFFF, $FF06

        .equ mouse 1 ; simple, right?
        .equ rat mouse + 1
        .equ dog %0000.1111
        .equ cat { dog .dup - 1000 + }

        .native
        .axy16

        .include "include_01.asm"

                lda.# 0000
                tay
myloop:
                sta.y $2000 ; not enough
                sta.y $2000+$100
                dey
                bne myloop
stop:           
                stp ; it's all over, baby!

        .byte 1 ... 4
        .byte 5 ... 6
        .byte "a" ... "z"
        .byte 5, 6, 7, 8 ; and that's a wrap!
        .word {frog 2 .swap *}, $ff , "frog"

        .end        
 
