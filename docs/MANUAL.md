# Manual for the Cthulhu Assembler
Scot W. Stevenson <scot.stevenson@gmail.com>
First version: 23. Apr 2018
This version: 19. May 2018

**THIS IS JUST A COLLECTION OF NOTES AT THE MOMENT**

## Introduction 

There are various assemblers for the 6502 and 65c02, but very few for the 65816.



## Philosophy

Cthulhu does not try to rewrite your code, even for optimization. It tries to
give you as much control as possible -- if you wanted the machine to do it for
you, you probably wouldn't be programming in assembler in the first place.  It
will make suggestions about code optimization if asked, however. 

Instead, it tries to check for as many errors as possible. 



## Forget all of that. What's with the name?

The assembler was originally called "go65816". But then I opened my first book
on compilers, FEHLT, and what I found was a horror beyond human comprehension.
The sentiment instantly reminded me of the works of H.P. Lovecraft, and so there
was a name change. Since then, I have found more accessible books on compilers
and feel much better.


## Command Line Options

- **-d** "debug" Debugging mode.
- **-f** "format" Format output. 
- **-fo** "format only" Only format the source code
- **-ff <FILE>** "format file" Name of the file the formatted source code from
  `-f` is saved as. If not included, output goes to standard output
- **-i <FILE>** "input" Input file (required).
- **-l** "listing" Generate listing file.
- **-m <STRING>** "MPU". Target processor. Currently supported are `6502`, `65c02`, and `65816`,
  default is `65c02`. 
- **-s** "symbol" Generate symbol table file.
- **-v** "verbose" Verbose mode. 

## The source code file

### Assembler Syntax

The syntax for Cthulhu should be quick to write and easy to parse to allow
automatic formatting. Mostly, the type of a word can be determined by its 
**first character**. The following list has the informal definitions, a formal
grammar can be found in /docs/grammar.txt

- **Comments** start with a semi-colon (`;`) and run to the end of the line

- **Directives** start with a dot (`.`) and consist of lower-case letters,
  numbers, the start (`*`), and the exclamation mark `!`. They can have
  parameters. 

- **Hex numbers** start with the traditional dollar sign (`$`). It can contain
  dots and colons as separators for easier reading.

- **Binary numbers** start with the traditional percent sign (`%`). It can contain
  dots and colons as separators for easier reading.

- **Decimal numbers** are, well, decimal numbers.

- Octal number are not supported. 

- **Strings** and **single characters** are enclosed by quotation marks (`\"`).
  This means that single characters are not enclosed in single quotes.

- **Mnemonics** belong to a fixed set of 256 words. Cthulhu accepts
  Simpler Assembler Notation (SAN), but includes a conversion program to move
  traditional WDC notation to SAN. 

- **Symbols** start with upper- or lowercase (unicode) letter, not a number or a
  special character. They can then contain futher upper- or lowercase letters,
  numbers, or special charaters out of the list `#?!_.'@~^&=|`. Note that
  the square braces `[` and `]`, as well as curly brances `{` `}`, the comma
  `,`, semi-colon `;` and parens `(` and `)` are not legal characters for
  symbols.  They cannot be the same as mnemonics.

- **Labels** are basically symbols with are defined a special way: They mark
  their position on the far beginning of the line with a colon (`:`) if they
  are global, and an underscore (`_`) if they are local. Local labels are only
  valid inside a scope marked by  `.scope` and `.scend`. Apart from that, they
  follow the same rules as symbols.

- **Math terms** are handled inside curly braces and follow reverse polish
  notation (RPN). See below for a further description.


### Indentation

The assembler can provide automatically formatted code following the model of
the `gofmt` tool. Simplified, **labels** should start at the beginning of the
line and by on a line by themselves, **directives** should be indented by eight
spaces (one tab), and **mnenomics** by 16 spaces (two tabs). Though there is no
artificial limit to the line length, lines beyond 80 characters are discouraged.

### Example Code

```
        .ram $0000 ... $7FFF
        .rom $8000 ... $FFFF

        .mpu "65816"
        .org $00:8000

        .equ target $2000

        .native
        .a16
        .xy8

;; This are getting serious
start:
        .scope
                lda.# 0000
                ldx.# $FF
_loop:
                sta.x target
                dex             ; remember two bytes per A at 16 bit
                dex
                bne loop        ; note no underscore
        .scend

                stp
```

Note that `yes!`, `wtf?`, and `No!No!No!` are all legal symbols and (with colons
at the end) legal labels. 

### Labels

Labels must start with a letter or an underscore `_`, that is, they follow the same
rules for symbols. When defined, they take a colon `:` as last character. 


```
global_label: ; accessable from everywhere

        .scope
_local_Label: ; only accessable in this scope
        .scend
```

During formatting, label definitions (the ones with the colons) are moved to the
very front of the line.


### RPN Math Syntax

Cthulhu allows a form of Reverse Polish Notation (RPN) for math, which is
appropriate because it was first built for a Forth system called Liara Forth. It
is introduced by a curly brace and followed by a stack notation with numbers --
binary, hex, and decimal -- and mathematical operation, along with some special
functions. Examples:

```
        .equ target $2000

                lda.# {40 10 +}
                sta.zi target            ; STA (target)
                sta.zi {target 1 +}      ; STA (target+1)
```

It is an error if there is more than one element on the math stack when the
operation is finished. 

### Error handling

Cthulhu follows the philosophy that each pass should find as many problems as
possible. If an error is found in a module such as the lexer or parser, that
module attempts to continue as far as possible and then stops the program.

> Internally, each module handles errors with a `reportErr()` function that
> prints an error message and increases a counter. The function `recover()`
> attempts to find a way to continue.


## List of Directives

This is a complete list of available and planned directives. A "(n/a)" signals
that this word is not yet available.

- **.a8** (n/a)
- **.a16** (n/a) No parameters.
- **.!a8** (n/a)
- **.advance** (n/a) 
- **.and**
- **.assert** (n/a) Takes one of the following options: **a8 a16 xy8 xy16 native emulated**. Checks during
  assembly to make sure that the given parameter is true. Aborts with an error 
  message if not. (65816 only)

- **.!axy16** (n/a) 
- **.axy16** (n/a) 
- **.axy8** (n/a) 
- **.bank** ADDRESS (n/a) 
- **.byte** ADDRESS (n/a) 
- **.drop** (RPN only)
- **.dup**
- **.emulate** (n/a) 

- **.end** (n/a) No parameters. Marks end of assembly program.

- **.equ** (n/a) Required paramters: **<SYMBOL> <NUMBER>**. Defines a symbol.

- **.here** Inserts current Program Counter (PC) address
- **.include** STRING (n/a) 
- **.long** (n/a) 
- **.lsb** (n/a) 
- **.lshift**
- **.msb** (n/a) 

- **.mpu** Takes a string of **"6502"**, **"65c02"**, **"65816"**

- **.native** (n/a) 

- **.or**
- **.origin** (n/a) 
- **.ram** ADDRESS ADDRESS
- **.rshift**
- **.rom** ADDRESS ADDRESS  
- **.status** (n/a) 
- **.swap**
- **.word** (n/a) 
- **.xor**
- **.xy16** (n/a) 
- **.xy8** (n/a) 

### Reserved for future use

- **.if** (n/a) 
- **.then** (n/a) 
- **.else** (n/a) 
- **.invoke** (n/a) 
- **.loop** ("x" | "y") <RANGE> 
- **.lend**
- **.macro** (n/a) 
- **.mend** (n/a) 
- **.print** Takes a string and prints it turning compilation (useful for
  debugging)
- **.scope** (n/a) 
- **.scend** (n/a) 

## Pseudoinstructions

- **move** <NUMBER> <SOURCE> <DESTINATION> For non-overlapping moves, this
  will generate the MVP/MVN code

- **.loop** (X|Y) <value1> ... <value2>

## Error handling

Cthulhu tries to report as many errors as possible at once.

## Literature and Websites


[1] Cooper, Keith D. and Torczon, Linda *Engineering a Compiler*, 2nd edition
Elsevier Inc 2012.

For literature on the Cthulhu Mythos and H.P. Lovecraft, 

The Super Tiny Compiler
https://github.com/hazbo/the-super-tiny-compiler/blob/master/compiler.go

## And finally

*Ph'nglui mglw'nafh Cthulhu R'lyeh wgah'nagl fhtagn*



