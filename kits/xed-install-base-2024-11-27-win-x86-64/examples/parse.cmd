asmparse -q vaddpd ymm1, ymm2, ymm3
asmparse -q -16 mov ax, dx
asmparse -q vaddpd ymm1{k3}, ymm2, ymm3
asmparse -q vaddpd ymm1{k2}{z}, ymm2, ymm3
asmparse -q vaddpd ymm1{k2}{z}, ymm2, ymmword [ebx]
asmparse -q vaddpd ymm1{k2}{z}, ymm2, ymmword ptr [ebx]
asmparse -q fdiv st(0), st(1)
asmparse -q cvtpi2pd xmm4, mm3
asmparse -q -64 mov cr0, rbx
asmparse -q mov ebx, dr1
asmparse -q call_near eax
asmparse -q call eax
asmparse -q jmp eax
asmparse -q rep cmpsb
asmparse -q repe cmpsb
asmparse -q repne cmpsb
asmparse -q lock  adc  dword [ebx], eax
asmparse -q lock cmpxchg dword ptr [ebx], esi
asmparse -q cmpxchg dword ptr [ebx], esi
asmparse -q lock mov dword ptr [ebx], esi
asmparse -q lock mov edi, dword ptr [ecx]
asmparse -q call far 0x1234:0x10dedead
asmparse -q call 0x1234:0x10dedead

asmparse -q mov dword ptr ds:[edi+5E8],eax
asmparse  mov dword ptr ds:[edi+5E8],eax

asmparse -q call 7FFC0A2B66B0
asmparse -q jmp 7FFC0A36B4C3
asmparse -q mov rax,qword ptr ss:[rsp+40]
asmparse -q or dword ptr ds:[rax+68],r14d
asmparse -q mov rax,qword ptr ss:[rsp+40]
asmparse -q mov rcx,qword ptr ds:[rax+30]
asmparse -q mov qword ptr ds:[7FFC0A4143A8],rcx
asmparse -q call 7FFC0A3093C8
asmparse -q mov ebx,eax
asmparse -q test eax,eax
asmparse -q jns 7FFC0A36B491
asmparse -q mov dword ptr ss:[rsp+28],eax
asmparse -q lea r8,qword ptr ds:[7FFC0A3DD270]
asmparse -q lea rax,qword ptr ds:[7FFC0A3DD148]
asmparse -q xor r9d,r9d
asmparse -q mov edx,DF2
asmparse -q mov qword ptr ss:[rsp+20],rax
asmparse -q lea rcx,qword ptr ds:[7FFC0A3C6C08]
asmparse -q call 7FFC0A2B66B0
asmparse -q jmp 7FFC0A36B4C3
          pause