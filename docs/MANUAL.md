# Manual for the GoAsm65816 Assembler
Scot W. Stevenson <scot.stevenson@gmail.com>
First version: 23. Apr 2018
This version: 02. May 2018



The assembler is built with the stengths and special characteristics of the
65816 MPU in mind. 



## Assembler Syntax

The syntax for GoAsm65816 should be quick to write and easy to parse to allow
automatic formatting. Mostly, the type of a word can be determined by its 
**first character**. The following list has the informal definitions, a formal
grammar will be included in the doc folder at a later point.

- **Comments** start with a semi-colon (`;`) and run to the end of the line. A
  double semi-colon (`;;`) forces a whole-lime comment regardless of where it
  begins. 

- **Directives** start with a dot (`.`) and consist of lower-case letters and
  numbers. The can have parameters. 

- **Hex numbers** start with the traditional dollar sign (`$`). It can contain
  dots and colons as separators for easier reading.

- **Binary numbers** start with the traditional percent sign (`%`). It can contain
  dots and colons as separators for easier reading.

- **Decimal numbers** are, well, decimal numbers.

- Octal number are not supported. 

- **Strings** and **single characters** are enclosed by quotation marks (`\"`).
  This means that single characters are not enclosed in single quotes.

- **Mnemonics** belong to a fixed set of 256 words. They consist of lower case
  letters and dots, with a few special characters such as `#` for immediate
  values.

- **Symbols** start with upper- or lowercase (unicode) letter, not a number or a
  special character. They can then contain futher upper- or lowercase letters,
  numbers, or special charaters out of the list `#?!_.+-*/\\'@~^&=|`. Note that
  the square braces `[` and `]`, as well as curly brances `{` `}`, the comma
  `,`, semi-colon `;` and parens `(` and `)` are not legal characters for
  symbols.  They cannot be the same as mnemonics.

- **Labels** start with a colon (`:`) if the are global, and an underscore (`_`)
  if they are local must be inside a local scope, enclosed by `.scope` and
  `.scend`. After that, they follow the same rules as symbols.

- **Math terms** are handled inside square braces and follow reverse polish
  notation (RPN). See below for a further description.


## Indentation

The assembler can provide automatically formatted code following the model of
the `gofmt` tool. Simplified, **labels** should start at the beginning of the
line and by on a line by themselves, **directives** should be indented by eight
spaces (one tab), and **mnenomics** by 16 spaces (two tabs). Though there is no
artificial limit to the line length, lines beyond 80 characters are discouraged.

### Example Code

```
        .ram $0000 $7FFF
        .rom $8000 $FFFF

        .org $00:8000

        .equ target $2000

        .native
        .a16
        .xy8

;; This are getting serious
:start
        .scope
                lda.# 0000
                ldx.# $FF
_loop
                sta.x target
                dex             ; remember two bytes per A at 16 bit
                dex
                bne _loop
        .scend

                stp
```

Note that `yes!`, `wtf?`, and `oh-no-no` are all legal symbols and (with colons
and underscores at the beginning) legal labels. 


