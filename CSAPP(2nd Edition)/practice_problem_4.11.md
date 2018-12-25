Fill in the right-hand column of the following table to describe the processing of the irmovl instruction on line 4 of the object code in Figure 4.17

Figure 4.17:

<pre>
0x000: 30f209000000 |    irmovl $9,  %edx
0x006: 30f315000000 |    irmovl $21, %ebx
0x00c: 6123         |    subl %edx, %ebx
0x00e: 30f480000000 |    irmovl $128,%esp
0x014: 404364000000 |    rmmovl %esp, 100(%ebx)
0x01a: a02f         |    pushl %edx
0x01c: b00f         |    popl  %eax
0x01e: 7328000000   |    je done
0x023: 8029000000   |    call proc
0x028:              | done:
0x028: 00           |    halt
0x029:              | proc:
0x029: 90           |    ret
</pre>



<pre>
             Geneeric               Specific
Stage        irmovl V, rB           irmovl $128, %esp
-----------------------------------------------------------------
Fetch        icode:ifun <- M₁[PC]   icode:ifun <- M₁[0x00e] = 3:0
             rA:rB <- M₁[PC+1]      rA:rB <- M₁[0x00f] = f:4
             valC <- M₄[PC+2]       valC <- M₄[0x010] = 128
             valP <- PC+6           valP <- 0x00e + 6 = 0x014
Decode
Excute       valE <- 0 + valC       valE <- 0 + 128
Memory
Write back   R[rB] <- valE          R[%esp] <- valE = 128
PC update    PC <- valP             PC <- valP = 0x014
</pre>



How does this instruction execution modify the registers and the PC?

首先是Fetch，从内存中读取指令，得知是 irmovl 指令，目的寄存器为 esp，下一个指令地址为当前 PC + 6，即 0x014；Execute 阶段计算出结果 128；Write back 阶段将结果 128 写入寄存器 esp；PC update 更新 PC 的值为 0x014