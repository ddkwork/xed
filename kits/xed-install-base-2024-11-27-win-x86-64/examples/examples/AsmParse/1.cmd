gcc -o AsmParse ^
    xed-asmparse.c ^
    xed-examples-util.c ^
    xed-asmparse-main.c ^
    xed-dot.c ^
    xed-dot-prep.c ^
    -I. ^
    -Iinclude ^
    -L. ^
    -lxed
