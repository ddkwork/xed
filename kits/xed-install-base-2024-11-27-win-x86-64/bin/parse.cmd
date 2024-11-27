xed-asmparse-main -q vaddpd ymm1, ymm2, ymm3
xed-asmparse-main -q xchg eax, ebx, ecx, edx, eax, edx, edx, edx, eax, r12, r13
xed-asmparse-main -q -16 mov ax, dx
xed-asmparse-main -q vaddpd ymm1{k3}, ymm2, ymm3
xed-asmparse-main -q vaddpd ymm1{k2}{z}, ymm2, ymm3
xed-asmparse-main -q vaddpd ymm1{k2}{z}, ymm2, ymmword [ebx]
xed-asmparse-main -q vaddpd ymm1{k2}{z}, ymm2, ymmword ptr [ebx]
xed-asmparse-main -q fdiv st(0), st(1)
xed-asmparse-main -q cvtpi2pd xmm4, mm3
xed-asmparse-main -q -64 mov cr0, rbx
xed-asmparse-main -q mov ebx, dr1
xed-asmparse-main -q call_near eax
xed-asmparse-main -q call eax
xed-asmparse-main -q jmp eax
xed-asmparse-main -q rep cmpsb
xed-asmparse-main -q repe cmpsb
xed-asmparse-main -q repne cmpsb
xed-asmparse-main -q lock  adc  dword [ebx], eax
xed-asmparse-main -q lock cmpxchg dword ptr [ebx], esi
xed-asmparse-main -q cmpxchg dword ptr [ebx], esi
xed-asmparse-main -q lock mov dword ptr [ebx], esi
xed-asmparse-main -q lock mov edi, dword ptr [ecx]
xed-asmparse-main -q call far 0x1234:0x10dedead
xed-asmparse-main -q call 0x1234:0x10dedead
pause
