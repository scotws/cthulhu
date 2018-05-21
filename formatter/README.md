# Formatter for the Cthulhu Assembler
Scot W. Stevenson <scot.stevenson@gmail.com> 
First version: 11. May 2018 
This version: 21. May 2018 

The formatter is called with the `-f` option and produces a (well) perfectly
formatted version of the source code based on the Immediate Representation (IR)
of the assembler. The philosophy is based on the `gofmt` tool included with the Go
(golang) programming language: There is one and only one format, and it's
handled by a machine, not a specification. 

## Formatting rules

The formatter is based on Simpler Assembler Notation (SAN), see the
specification for a complete set of rule. Briefly:

- Indentation is in steps of eight characters (one classic tabulator)
- **Labels** start at the first column of a line and have the line to
  themselves. This is true for local labels and anonymous labels as well
- **Directives** are indented by one step, including the `.scope` directive
- **Instructions** are indented by two steps
- All directives and instructions are lower case
- Inline comments have one space after the line they comment
- There are only single empty lines in sequence

## Cthulhu options

- **-f** "format" Run the formatter as part of Cthulhu
- **-fo** "format only" Stop the assembler after formatting the source code
- **-ff <FILE>** "format file" Output to file named. If not present, it is sent
  to the standard output
