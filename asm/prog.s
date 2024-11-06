# ((10 - 20) * (30.5 + 40)) / 4
.data
zero:
.float 0.0
one:
.float 1.0
ten:
.float 10.0
t1:
.float 10.3
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
.quad
buf2:
.quad 
isNeg:
.byte 
res1:
.float
res2:
.float
res3:
.float
res4:
.float 

.text

# len of the string from %rsi address to the 0 byte 
# %rax - result
 .macro len m:req
 
 push %rdx 
 push \m
 xor %rax, %rax 
 1:
 mov (\m), %dl	
 cmp $0, %dl	
 jz  2f				    
 inc \m		  	
 inc %rax 	    
 jmp 1b   
 2:
 pop \m
 pop %rdx

 .endm

# print everything to 0 byte from %rsi address 
 print:
 push %rax
 push %rdi
 push %rdx
 len %rsi
 
 mov %rax, %rdx 
 mov $1, %rax
 mov $1, %rdi	
 syscall
 pop %rdx
 pop %rdi
 pop %rax		    
ret

/*.macro toStr
 # число в %rax 
 # подготовка преобразования числа в строку
  movq $0, (buf2)
  mov $10, %r8    # делитель
  mov $buf, %rsi  # адрес начала буфера 
  xor %rdi, %rdi  # обнуляем счетчик
# преобразуем путем последовательного деления на 10
toStrlo:
  xor %rdx, %rdx  # число в rdx:rax
  div %r8         # делим rdx:rax на r8
  add $48, %dl    # цифру в символ цифры
  mov %dl, (%rsi) # в буфер
  inc %rsi        # на следующую позицию в буфере
  inc %rdi        # счетчик увеличиваем на 1
  cmp $0, %rax    # проверка конца алгоритма
  jnz toStrlo          # продолжим цикл?
# число записано в обратном порядке,
# вернем правильный, перенеся в другой буфер 
  mov $buf2, %rbx # начало нового буфера
  mov $buf, %rcx  # старый буфер
  add %rdi, %rcx  # в конец
  dec %rcx        # старого буфера
  mov %rdi, %rdx  # длина буфера
# перенос из одного буфера в другой
toStrexc:
  mov (%rcx), %al # из старого буфера
  mov %al, (%rbx) # в новый
  dec %rcx        # в обратном порядке  
  inc %rbx        # продвигаемся в новом буфере
  dec %rdx        # а в старом в обратном порядке
  jnz toStrexc         # проверка конца алгоритма 
  movb $0, (%rbx)
.endm 

.macro printFloat f:req
# is number negative?
movss (zero), %xmm0 
movss \f, %xmm1 
cmpss $1, %xmm0, %xmm1
pextrb $3, %xmm1, %rax
cmp $0, %rax 
jz __floatToStrIsPos
# change to positive
fld (zero) 
fsub (one)
fmul \f
fstp \f
movb $1, (isNeg) 
jmp __floatToStrIsNeg
__floatToStrIsPos:
movb $0, (isNeg)
__floatToStrIsNeg:
fld \f
movss \f, %xmm0 
roundps $3, %xmm0, %xmm0
movss %xmm0, \f
cvtss2si \f, %r12
fsub \f 
fstp \f
mov $6, %r10 

__floatToStrLocal:
fld \f
cmp $0, %r10
jz __floatToStrOk
dec %r10 
movss (ten), %xmm0
movss %xmm0, \f
fmul \f
fstp \f
jmp __floatToStrLocal
__floatToStrOk:
cvtss2si \f, %rax # здесь содержится дробное значение 
toStr
mov $buf2, %rsi 
mark:
print
.endm
*/


.globl _start
_start:
#fld (t2)
#fsub (t1)
#fstp (res1)

mov $message, %rsi

call print
#printFloat (res1)
mov $60,  %rax
xor %rdi, %rdi 
syscall
