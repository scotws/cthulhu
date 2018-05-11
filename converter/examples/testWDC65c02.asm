; Test assembler file for the Cthulhu Assembler
; WDC 65c02 code
; Scot W. Stevenson <scot.stevenson@gmail.com>
; First version: 09. Mai 2018
; This version: 10. Mai 2018

        .mpu "65c02"
        .notation "wdc"
        .origin $8000
        
        .ram $0000-$7FFF
        .rom $8000-$FFFF
        .ram $FF01, $FF02 ; putc, putc_block
        .rom $FF03 ; getc

        .equ start $8000

                lda #$00
                tay
:loop
                sta $2000,y
                dey
                bne loop

                rts ; just of testing

        .word {frog 2 .swap .dup *}
        .byte 1, 2, 3, 4
        .byte 5, 6, 7, 8 ; and that's a wrap!

        .end        
 
