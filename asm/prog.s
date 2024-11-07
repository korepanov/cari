# ((10 - 20) * (30.5 + 40)) / 4
.data
zero:
.float 0.0
one:
.float 1.0
ten:
.float 10.0
t1:
.float 10
t2:
.float 20
t3:
.float 30.5 
t4:
.float 40 
t5:
.float 4 

.bss
buf:
.skip 8
buf2:
.skip 8
buf3:
.skip 8
buf4:
.skip 8
res1:
.skip 4
res2:
.skip 4
res3:
.skip 4
res4:
.skip 4

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

# transform uint value from %rax register to string 
# %rsi - address of the result
.macro toStr

  push %rax 
  push %rdi 
  push %rdx
  push %rbx 
  push %rcx 
  
  movq $0, (buf2)
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

  mov $buf2, %rbx 
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

# f - float to print 
.macro printFloat f:req
push %rax
push %rbx 
push %r12
push %rsi 

mov \f, %rbx
mov %rbx, (buf4)
mov $buf3, %rbx

# is number negative?
movss (zero), %xmm0 
movss (buf4), %xmm1 
cmpss $1, %xmm0, %xmm1
pextrb $3, %xmm1, %rax
cmp $0, %rax 
jz 1f
# change to positive and save minus  
movb $'-', (%rbx)
inc %rbx
fld (zero) 
fsub (one)
fmul (buf4)
fstp (buf4)
1:
fld (buf4)
movss (buf4), %xmm0 
roundps $3, %xmm0, %xmm0
movss %xmm0, (buf4)
cvtss2si (buf4), %r12
mov %r12, %rax
toStr # integer part 

5:
mov (%rsi), %al
mov %al, (%rbx)
inc %rbx 
inc %rsi
cmp $0, %al 
jnz 5b
dec %rbx 

fsub (buf4)
fstp (buf4)
mov $6, %r10 # number of digits after point 

3:
fld (buf4)
cmp $0, %r10
jz 4f
dec %r10 
movss (ten), %xmm0
movss %xmm0, (buf4)
fmul (buf4)
fstp (buf4)
jmp 3b
4:
cvtss2si (buf4), %rax # here the number after point 
movb $'.', (%rbx)
inc %rbx 
toStr
6:
mov (%rsi), %al
mov %al, (%rbx)
inc %rbx 
inc %rsi
cmp $0, %al 
jnz 6b
dec %rbx 

movb $'\n', (%rbx)
inc %rbx 
movb $0, (%rbx)

mov $buf3, %rsi 
print 

pop %rsi
pop %r12 
pop %rbx 
pop %rax
.endm



.globl _start
_start:
fld (t1)
fsub (t2)
fstp (res1) 

printFloat (res1)

mov $60,  %rax
xor %rdi, %rdi 
syscall
