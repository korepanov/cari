package program

const dataBegin = `
.data
enter:
.ascii "\n"
.space 1, 0
`
const bssBegin = `
.bss
buf:
.skip 21
buf2:
.skip 21
`

const textBegin = `
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
`
const textEnd = `
mov $60,  %rax
xor %rdi, %rdi 
syscall
`
