;; Test 02 for GoAsm65816 
;; Scot W. Stevenson <scot.stevenson@gmail.com>
;; First version: 24. Apr 2018
;; This version: 24. Apr 2018
;; Code is in Simpler Assember Notation (SAN)
.mpu 65816 .origin $8000 .native ; it gets serious
start: nop ; this really does nothing
ldx.# 0010 @ lda.# "a" sta.x $00:2FFF  ; note bank byte separator
sta.x 0x30aa dex bne - lda.# %11110000 bra start       
stuff:  .byte 0, $01, 2, 2+1, "4", %0000:1111 ; just a test
moar'stuff: .byte "cats", "dogs", "rats" .end        
