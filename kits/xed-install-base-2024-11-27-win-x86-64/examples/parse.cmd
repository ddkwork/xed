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
pause