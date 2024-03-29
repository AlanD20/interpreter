# Interpreter

This tool allows you to write scripts, rendering tools, and many more tools
without sacrificing `Speed` and `Memory`.

This implementation uses Go as a base language to take advantage of features
that Go provides.

ZoomLang (.zlm):

Implies a language that allows users to 'zoom' through scripting. The ".zlm"
extension is short and easy to remember.

## Roadmap

- [x] REPL Support
- [x] Script File Support with `.iad` extension
- [ ] Error Handling (internal): error detection and reporting within the lexer,
      parser, and interpreter.
- [ ] Testing (internal): Lexer, parser, interpreter, and built-in functions.
- [ ] Tree-walk interpreter
  - [x] Grammar
    - [x] Keywords
    - [x] Data types
    - [x] Variables
    - [x] Functions
    - [x] Control Structure
    - [x] Operators
  - [x] Lexer/Scanner
    - [ ] Error Detection
  - [ ] Parser
    - [ ] Abstract Syntax Tree (AST)
    - [ ] Probably different AST structures and algorithms to find the best one?
  - [ ] Interpreter/Evaluate Expressions
    - [ ] Variables
    - [ ] Functions
    - [ ] Operators
    - [ ] Control Flow
    - [ ] Resolvers and Bindings
    - [ ] Classes
    - [ ] Inheritance
  - [ ] Build-in Functions/States and Statements
    - [ ] print
- [ ] Bytecode
  - [ ] Compiling Expressions
  - [ ] Types
  - [ ] Strings
  - [ ] Variables
  - [ ] Closures
  - [ ] Memory Management
    - [ ] Garbage Collection
  - [ ] Classes and Instances
  - [ ] Superclasses
  - [ ] Optimization: Include constant folding, loop unrolling, and register
        allocation
- [ ] Modules Support
- [ ] Syntax Review/Grammar efficiency
- [ ] Documentation
  - [ ] API
  - [ ] Descriptions
  - [ ] Examples
  - [ ] Best Practices

## Notes/Features

- Extend the `token` struct to add necessary information where the token is. For
  example, another property to store what is the column number of the token.

## Grammar

- Here is the grammar:

  ```plaintext
  expression     → literal
                | unary
                | binary
                | grouping ;

  literal        → NUMBER | STRING | "true" | "false" | "null" ;
  grouping       → "(" expression ")" ;
  unary          → ( "-" | "!" ) expression ;
  binary         → expression operator expression ;
  operator       → "==" | "!=" | "<" | "<=" | ">" | ">="
                | "+"  | "-"  | "*" | "/" ;
  ```

## Why?

This project is purely a hobby project to explore and learn more about how to
build an interpreter tool from scratch and learn more about optimization, AST,
garbage collection, and many related topics that helps me understand how
interpeters work. <br> My intention is to create a simple, compiled scripting
language that helps beginners to learn quickly to write scripts that is
efficient in memory and speed. <br> I try to omit boilerplates as much as
possible to help people start writing their script without wasting time on
boilerplates.

## Contribution

Feel free to contribute, give suggestions, or recommendations for the design
decisions.

## License

This project is under [MIT License](/LICENSE)
