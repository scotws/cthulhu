# Manual for the GoAsm65816 Assembler
Scot W. Stevenson <scot.stevenson@gmail.com>
First version: 23. Apr 2018
This version: 23. Apr 2018






## Internals

Without going into the formal grammar, the rules for elements are as follows:

**COMMENTS** - Comments start with a ';' semicolon. If the formatted output
should recognize a comment as a whole-line comment, use ';;' two semicolons. In
both cases, the whole rest of the line is ignored.

**NUMBERS** - SAN distingishes three types of numbers:

- **Decimal** - The usual, built of the digits 0 to 9
- **Hexadecimal** - Prefixed with either `$` (traditional) or `0x` (preferred)
  and followed by a hex number. The digits a to f can be either upper or
  lowercase, though upper case is recommended to make it easier to distingish
  them from symbols.
- **Binary** - Prefixed by `%` and either `0` or `1`

Note that octals are not supported.


**LABELS** - A label must start with an upper or lower case letter, or the `_`
(underscore) sign for local labels inside a `.scope` region. They must end with
a colon `:`. Between the first and last letter, they can contain letters,
numbers or any of the characters `_?!&#`. These are legal labels:
```
fire&ice:
wtf?:
_keep_it_local:
b:
```
The colon *must* appear at the end of the label, but that's the only place it is
allowed to appear. There must be whitespace after the colon or an end of the
line.

**DIRECTIVES** - Directives start with a period `.` and are followed by letters,
numbers, and special characters. They always appear before their parameters.

```
        .equ company 3
        .equ bore 3+1           ; 3+1 is calculated
```

Directives are indented one tab's worth of spaces. There must be whitespace or
the end of the line after a directive.

**SYMBOLS** - A symbol must start with an upper or lowercase letter, or the `_`
underscore sign for local labels inside a `.scope` region. For obvious reasons,
they can contain the same special characters as labels.

```
        lda.# company
        jmp wtf?
        jmp fire&ice+2
        
```
Internally, a string is defined as a symbol when all other
alternatives are exhausted -- that is, the string is not a number, not a label,
and not a directive. There must be whitespace, a math operation (see below) 
or the end of the line after a symbol.after a symbol.


```
```
