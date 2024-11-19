# Compiler of integer arithmetic expressions into assembler code of AT&T syntax

## Get started 
1. To make AST:
```
./cari -i example.cari -ast
```
```
start
   ├─ *
   │  ├─ +
   │  │  ├─ *
   │  │  │  ├─ -25
   │  │  │  └─ 10
   │  │  └─ *
   │  │     ├─ -1
   │  │     └─ -7
   │  └─ -100
   ├─ -
   │  ├─ 10
   │  └─ 7
   └─ +
      ├─ /
      │  ├─ -
      │  │  ├─ +
      │  │  │  ├─ 512
      │  │  │  └─ 421
      │  │  └─ +
      │  │     ├─ 3321
      │  │     └─ 70
      │  └─ -
      │     ├─ 900
      │     └─ +
      │        ├─ 30
      │        └─ 78
      └─ -
         ├─ +
         │  ├─ 32
         │  └─ 7
         └─ 90
```
2. To compile to the assembler code:
```
./cari -i example.cari -o example.s
```
3. Now you can build your program:
```
as example.s -o example.o
ld example.o -o example
```
4. Run the program.
```
./example
```
```
24300
3
-54
```
## Build 
From `cmd` folder:
```
go build -o cari main.go
```

## Limits
**minumum** number: -9223372036854775808  
**maximum** number: 9223372036854775807
