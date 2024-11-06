# ((10 - 20) * (30.5 + 40)) / 4
.data
enter:
.ascii "\n"
.space 1, 0
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
message:
.ascii "Mymessage"
.space 1,0

.bss
buf:
.skip 8
buf2:
.skip 8
buf3:
.skip 8
isNeg:
.skip 1
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
mov $buf3, %rbx

# is number negative?
movss (zero), %xmm0 
movss \f, %xmm1 
cmpss $1, %xmm0, %xmm1
pextrb $3, %xmm1, %rax
cmp $0, %rax 
jz 1f
# change to positive and save minus  
mark:
movb $'-', (%rbx)
inc %rbx
fld (zero) 
fsub (one)
fmul \f
fstp \f
movb $1, (isNeg) 
jmp 2f
1:
movb $0, (isNeg)
2:
fld \f
movss \f, %xmm0 
roundps $3, %xmm0, %xmm0
movss %xmm0, \f
cvtss2si \f, %r12
mov %r12, %rax
toStr # integer part 

/*5:
mov (%rsi), %al
mov %al, (%rbx)
inc %rbx 
inc %rsi
cmp $0, %al 
jnz 5b*/

fsub \f 
fstp \f
mov $6, %r10 # number of digits after point 

3:
fld \f
cmp $0, %r10
jz 4f
dec %r10 
movss (ten), %xmm0
movss %xmm0, \f
fmul \f
fstp \f
jmp 3b
4:
cvtss2si \f, %rax # here the number after point 
mov $buf3, %rsi 
print 

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
