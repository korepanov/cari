# ((10 - 20) * (30.5 + 40)) / 4
.data
t1:
.double 10
t2:
.double 20
t3:
.double 30.5 
t4:
.double 40 
t5:
.double 4 

.bss
res1:
.double
res2:
.double
res3:
.double
res4:
.double 

.text

printFloat:



.globl _start
_start:
fld (t2)
fsub (t1)
fstp (res1)

mov $60,  %rax
xor %rdi, %rdi 
syscall
