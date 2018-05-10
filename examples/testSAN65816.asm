; Test assembler file for the Cthulhu Assembler
; Scot W. Stevenson <scot.stevenson@gmail.com>
; First version: 09. Mai 2018
; This version: 09. Mai 2018

        .mpu "65816"
        .notation "san"
        .origin $00:8000

        .native
        .axy16

                nop
                nop
                nop

        .byte 1, 2, 3, 4

        .end        
 
