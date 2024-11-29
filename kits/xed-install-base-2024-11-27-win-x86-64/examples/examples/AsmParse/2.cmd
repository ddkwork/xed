ccgo -I . ^
     -I include ^
     --package-name AsmParse ^
     xed-asmparse-main.c -o xed-asmparse-main_gen.go ^
     xed-asmparse.c -o xed-asmparse_gen.go ^
     xed-dot-prep.c -o xed-dot-prep_gen.go ^
     xed-dot.c -o xed-dot_gen.go ^
     xed-examples-util.c -o xed-examples-util_gen.go ^
       -L. ^
         -lxed
