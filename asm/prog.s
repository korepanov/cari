# This code was made by cari
# ((10 - 20) * (30 + 40)) / 4
.data
enter:
.ascii "\n"
.space 1, 0

t1:
.quad 10
t2:
.quad 20
t3:
.quad 30 
t4:
.quad 40 
t5:
.quad 4

t6:
.quad -9223372036854775808
#.quad 9223372036854775807



.bss
buf:
.skip 21
buf2:
.skip 21
res1:
.skip 8
res2:
.skip 8
res3:
.skip 8
res4:
.skip 8

.text

# len of the string from %rsi address to the 0 byte 
# %rax - result
 .macro len
 
 push %rdx  
 push %rsi
 
 xor %rax, %rax 
 1:
 mov (%rsi), %dl	
 cmp $0, %dl	
 jz  2f				    
 inc %rsi		  	
 inc %rax 	    
 jmp 1b   
 2:
 
 pop %rsi 
 pop %rdx

 .endm

# print everything to 0 byte from %rsi register address 
 .macro print
 push %rax
 push %rdi
 push %rdx
 
 len 
 mov %rax, %rdx
 
 mov $1, %rax
 mov $1, %rdi	
 syscall
 pop %rdx
 pop %rdi
 pop %rax		    
 .endm

# clear 21 byte from %rsi
.macro clear
 push %rax
 mov $21, %al
 1:
 movb $0, (%rsi)
 inc %rsi
 dec %al 
 cmp $0, %al 
 jnz 1b
 pop %rax
.endm

# transform uint value from %rax register to string 
# %rsi - address of the result
.macro toStr

  push %rax 
  push %rdi 
  push %rdx
  push %rbx 
  push %rcx 
  
  mov $buf, %rsi 
  clear
  mov $buf2, %rsi 
  clear
  mov $buf2, %rbx
  
  cmp $0, %rax 
  jge 3f
  mov $0, %rdx 
  sub $1, %rdx 
  imul %rdx  
  movb $'-', (%rbx)
  inc %rbx 
  3:
  
  mov $10, %r8    
  mov $buf, %rsi 
  xor %rdi, %rdi  

1:
  xor %rdx, %rdx  
  div %r8         
  add $48, %dl    
  mov %dl, (%rsi) 
  inc %rsi        
  inc %rdi        
  cmp $0, %rax    
  jnz 1b         

  mov $buf, %rcx  
  add %rdi, %rcx  
  dec %rcx        
  mov %rdi, %rdx  

2:
  mov (%rcx), %al 
  mov %al, (%rbx) 
  dec %rcx          
  inc %rbx        
  dec %rdx       
 
  jnz 2b          
  mov $buf2, %rsi 
  
  pop %rcx
  pop %rbx 
  pop %rdx
  pop %rdi 
  pop %rax
.endm


.globl _start
_start:

mov (t1), %rax 
mov (t2), %rbx 
sub %rbx, %rax
mov %rax, (res1)

toStr 
print 
mov $enter, %rsi 
print 

mov (t3), %rax 
mov (t4), %rbx
add %rbx, %rax 
mov %rax, (res2)

toStr 
print 
mov $enter, %rsi 
print 

mov (res1), %rax 
mov (res2), %rbx 
imul %rbx, %rax 
mov %rax, (res3)
toStr 
print 
mov $enter, %rsi 
print 

mov (res3), %rax
mov (t5), %rbx 
cqo
idiv %rbx
mov %rax, (res4)

toStr
print 
mov $enter, %rsi 
print 

mov $60,  %rax
xor %rdi, %rdi 
syscall
