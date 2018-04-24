# Manual for the Simpler Assembler Notation (SAN)
Scot W. Stevenson <scot.stevenson@gmail.com> 
First version: 23. April 2018
This version: 24. April 2018

Simpler Assembler Notation (SAN) is a, well, not so complicated way of
writing assembler code for the 65816 MPU. 

## History




## Design Goals


(The following is copied from TAN and still not done)




Directives always come first in their line (except for the Current Line Symbol,
see below). Assignments for example are `.equ mynumber 13.`


### Mnemonics 

TinkAsm uses a different format for the actual MPU instructions, the Typist's
Assembler Notation (TAN). See [the
introduction](https://docs.google.com/document/d/16Sv3Y-3rHPXyxT1J3zLBVq4reSPYtY2G6OSojNTm4SQ/)
to TAN for details.

TinkAsm and TAN try to ensure that all source files have the same formatting, a
philosophy it takes from the Go programming language. An equivalent tool to Go's
`gofmt`, `tinkfmt.py`, is included. 


### Definitions

TinkAsm requires two definitions at the beginning of the source file: Where
assembly is to start (`.origin`) and which processor the code is to be
generated for (`.mpu`). Failure to provide either one will abort the assembly
process with a FATAL error. Supported MPUs are `6502`, `65c02` (upper and
lowercase "c" are accepted), and `65816`. 


### Assignments 

To assign a value to a variable, use `.equ` followed by the symbol and the value.

```
        .equ a_bore 4
```
Modifications and math terms are allowed (see below for details).

```
        .equ less 8001
        .equ some .lsb less
        .equ more less + 1 
        .equ other .msb less + 1
```

### Labels

A label must begin with a upper- or lowercase letter and end with a colon. The
characters inbetween those, if any, can be letters, numbers, or special
characters from the list `!?&_`. 

```
ice&fire:       nop
                nop
                jmp ice&fire
```

During formatted output, labels are given their own line.

### Anonymous Labels 

Anonymous labels are used for trivial uses such as loops that don't warrant a
full label. It consists of `@` as the generic label, used as a normal label, and
either a `+` or a `-` after the jump or branch instruction. 

```
@               nop
                nop
                jmp -           ; anonymous reference 
```

The `-` or `+` always refers to the next or previous `@`. These directives
cannot be modified. 


### Current Line Symbol

To reference the current address, by default the directive `.*` is used 
instead of an operand. It can be modified and subjected to mathematical 
operations.
```
                jmp .*
```

### Numbers 

The common hex prefixes `$` and `0x` are recognized, with `0x`
being the recommended format. For binary numbers, use `%`. Octal numbers are
not supported.

(Numbers may contain the `:` character for better readability.)

```
        .equ address $00:fffa
```

This is especially useful for the bank bytes of the 65816.


### Single Characters and Strings

Single characters that are to be converted to ASCII values are in double quotes
just like strings (`lda.# "a"`). Note that though the code can be in Unicode
(labels, etc), strings for the 65816 themselves are currently still only ASCII.


### Modifiers and Math

Normal references to labels (but not anonymous labels) and symbols can be
modified by "modifiers" such as `.lsb` and simple mathematical terms such as `{
label + 2 }` .  White space is significant, so `label+2` is not legal (and will
be identified as a symbol). You can use anything that is a simple Python 3 math
instruction (including `**`) because  the term between the brackets is santized
and then sent to EVAL. Yes, EVAL is evil. Modifiers and math terms can be used
in data lines, such as `.byte .lsb { 1 + 1 } .msb { 2 + 2 }`


### Other 

It is assumed that branches will always be given a label, not the relative
offset. There is in fact currently no way to pass on such offset. 


## Macros

The macro system of TinkAsm is in its first stage. Macros do not accept
parameters and cannot reference other macros. Including parameters is a high
priority for the next version.

To define a macro, use the directive `.macro` followed by the name string of the
macro. No label may precede the directive. In a future version, parameters will
follow the name string. The actual macro contents follow in the subsequent
lines in the usual format. The macro definition is terminated by `.endmacro` in
its own line.

To expand a macro, use the `.invoke` directive followed by the macro's name. In
future versions, this will be followed optional parameters. 

Currently, there are no system macros such as `.if`, `.then`, `.else` or loop
constructs. These are to be added in a future version.


## List of Directives

### Directives for all MPUs

`@` - The default anonymous label symbol. Used at the very beginning of a line and
referenced by `+` and `-` for jumps and branches.

`+` - As an operand to a branch or jump instruction: Refer to the next anonymous 
label. 

`-` - As an operand to a branch or jump instruction: Refer to previous anonymous
label. 

`.*` - As an operand in the first position after the mnemonic: Marks current
address (eg `jmp { .* + 2 }`). 

`.advance` - Jump ahead to the address given as parameter, filling the space
in between with zeros.

`.bank` - Isolate bank byte - the highest byte - of following number. This is a
modifier. Thought it pretty much only makes sense for the 65816 MPU, it is
supported for other formats as well. 

`.byte` - Store the following list of comma-delimited bytes. Parameters
can be in any supported number base or symbols. 
Example: `.byte 'a', 2, { size + 1 }, "Kaylee", %11001001`

`.end` - Marks the end of the assembly source code. Must be last line in
original source file. Required. 

`.endmacro` - End definition of the macro that was last defined by `.macro`. 

`.include` - Inserts code from an external file, the name of which is given as a
parameter. 

`.invoke` - Inserts the macro given as parameter. 

`.long` - Store the following list of comma-delimited 24-bit as bytes.
The assembler handles the conversion to little-endian format. Parameters can be
in any supported number base or symbols.

`.lsb` - Isolate least significant byte of following number. This is a 
modifier.

`.macro` - Start definition of the macros, the name of which follows
immediately as the next string. Parameters are not supported in this version.
The definition ends with the first `.endmacro`. Macros cannot be nested.

`.msb` - Isolate most significant byte of following number. This is a 
modifier.

`.mpu` - Provides the target MPU as the parameter. Required. Legal values are
`6502`, `65c02`, or `65816`. 

`.origin` - Start assembly at the address provided as a parameter.
Required for the program to run. Example: `.origin 8000`

`.save` - Given a symbol and a number, save the current address during assembly 
as the symbol and skip over the number of bytes. Used to reserve a certain
number of bytes at a certain location. Example: `.save counter 2`

`.scope` - 

`.scend` - 

`.skip` - Jump head by the number of bytes given as a parameter, filling the
space in between with zeros. Example: `.skip 100`

`.word` - Store the following list of comma-delimited 16-bit words as
bytes. The assembler handles the conversion to little-endian format. Parameters
can be in any supported number base or symbols. Note that WDC uses "double 
byte" for 16-bit values, but the rest of the world uses "word". 


### Directives for 65816 only

`.a8` and `.a16` - Switch the size of the A register to 8 or 16 bit. The switch
to 16 bit only works in native mode. These insert the required instructions
as well as the internal control sequences (see below) and should be used instead
of directly coding the `rep`/`sep` instructions.

`.xy8` and `.xy16` - Switch the size of the X and Y registers to 8 or 16 bit.
The switch to 16 bit only works in native mode. These insert the required instructions
as well as the internal control sequences (see below) and should be used instead
of directly coding the `rep`/`sep` instructions.

`.axy8` and `.axy16` - Switch the size of the A, X, and Y registers to 8 or 16
bit. The switch to 16 bit only works in native mode. These insert the required instructions
as well as the internal control sequences (see below) and should be used instead
of directly coding the `rep`/`sep` instructions.

`.emulated` - Switch the MPU to emulated mode, inserting the required
instructions and control sequences. Use this directive instead of directly
coding `sec`/`xce`. 

`.native` - Switch the MPU to native mode, inserting the required
instructions and control sequences. Use this directive instead of directly
coding `clc`/`xce`. 


### Direct use of 65816 control directives

Internally, the mode and register size switches are handled by inserting
"control directives" into the source code. Though the above directives such as
`.native` or `.a16` should be enough, you can insert the control sequences
directly to ensure that the assembler handles the sizes correctly. These do not
encode any instructions.

`.!a8`, `.!a16` - Tell the assembler to interpret the A register as 8 or 16 bit.
Note the switch to 16 bit only works in native mode.

`.!xy8`, `.!xy16` - Tell the assembler to interpret the X and Y register as 8 or
16 bit.  Note the switch to 16 bit only works in native mode.

`.!emulated` - Tell the assembler to ensure emulated mode. Note this does not
insert the control sequences `.a8!` and `.xy8!` as the full directive
`.emulated` does.

`.!native` - Tell the assembler to ensure we're in native mode. 


### Notes on various opcodes

TinkAsm enforces the signature byte of `brk` for all processors.

The `mvp` and `mvn` move instructions are annoying even in the Typist's
Assembler Notation because they break the rule that every instruction has only
one operand. Also, the source and destination parameters are reversed in machine
code from their positions in assembler. The format used here is 

```
                mvp <src>,<dest>
```

The source and destination parameters can be modified (such as `.lsb 1000`) or
consist of math terms (`{ home + 2 }`).  For longer terms, you can use the
`.bank` directive.

```
        .equ source 01:0000
        .equ target 02:0000

                mvp .bank source , .bank target
```


## Internals 





### Tinkfmt Source Code Formatter

TinkAsm is currently shipped with one tool, the formatter Tinkfmt. See the
README file in the directory `tinkfmt` for details. 


## SOURCES AND THANKS 

TinkAsm was inspired by Samuel A. Falvo II, who pointed me to a paper by Sarkar
*et al*, ["A Nanopass Framework for Compiler Education"](
http://www.cs.indiana.edu/~dyb/pubs/nano-jfp.pdf)
The authors discuss using compilers with multiple small passes for educational
purposes. 

David Salomon's [*Assemblers And
Loaders*](http://www.davidsalomon.name/assem.advertis/AssemAd.html) was
invaluable and is highly recommended for anybody who wants to write their own. 


