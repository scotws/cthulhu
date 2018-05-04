# Reformat mnemonic list for data.go 
# Scot W. Stevenson <scot.stevenson@gmail.com>
# First version: 04. Mai 2018
# This version: 04. Mai 2018

template = "{0}: Opcode{{{0}, {1}, X, {2}, false}},"

with open('opcodes.txt', 'r') as f:
    raw_list = f.readlines()
 
for l in raw_list:
    ws = l.split()
    print(template.format(ws[1], ws[2], ws[0]))
