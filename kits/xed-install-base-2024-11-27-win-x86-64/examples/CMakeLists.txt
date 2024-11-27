cmake_minimum_required(VERSION 3.30)
set(CMAKE_C_STANDARD 11)
project(asmparse C)

include_directories(../include)
    link_directories(${CMAKE_SOURCE_DIR})
add_executable(asmparse
        xed-asmparse.h
        xed-asmparse.c
        xed-asmparse-main.c
        xed-examples-util.c
)
target_link_libraries(asmparse xed)










# 设置编译选项
#option(USE_ENCODER "Enable encoder" ON)
#option(USE_DECODER "Enable decoder" ON)
#option(BUILD_CPP_EXAMPLES "Build C++ examples" ON)
#option(CLEAN_TARGETS "Clean targets" OFF)
#
## 添加包含目录
#include_directories(../include)
#
## 设置库目录
#
#set(XED_LIB_DIR "./")
#set(XED_INC_DIR "../include")
#file(GLOB XED_ENC2_LIBS "${XED_LIB_DIR}/*xed-enc2-*")
#
## 示例源代码列表
#set(EXAMPLES_SRC
#        xed-disas-raw.c
##        avltree.c
#        xed-disas-hex.c
#        xed-symbol-table.c
#        xed-examples-util.c
#)
#
#if(USE_DECODER)
#    list(APPEND EXAMPLES_SRC
#            xed-dot.c
#            xed-dot-prep.c
#            xed-disas-filter.c
#            xed-nm-symtab.c
#    )
#endif()
#
#if(USE_ENCODER)
#    list(APPEND EXAMPLES_SRC
#            xed-enc-lang.c
#            xed-ex5-enc.c
#    )
#endif()
#
#set(OTHER_EXAMPLES
#        xed-ex1.c
#
#        xed-ex3.c
#        xed-ex4.c
##        xed-ex5-enc.c
#        xed-ex6.c
#        xed-ex7.c
#        xed-ex8.c
#        xed-ex9-patch.c
#        xed-tester.c
#        xed-dec-print.c
#        xed-ex-agen.c
#        xed-ex-ild.c
#        xed-ex-ild2.c
#        xed-min.c
#        xed-reps.c
#        xed-tables.c
#        xed-find-special.c
#        xed-dll-discovery.c
#)
#
## 添加所有其他示例文件
#foreach(src ${OTHER_EXAMPLES})
#    list(APPEND EXAMPLES_SRC ${src})
#endforeach()
#
#
#
## 根据所需示例文件创建可执行文件
#foreach(src ${EXAMPLES_SRC})
#    get_filename_component(exe_name ${src} NAME_WE)
#    add_executable(${exe_name} ${src})
#    # 添加库的搜索路径
#    link_directories(${CMAKE_SOURCE_DIR}../lib)
#    target_link_libraries(${exe_name} xed)
#endforeach()
#
## 清理目标
#if(CLEAN_TARGETS)
#    message(STATUS "Cleaning targets")
#    file(GLOB ALL_EXECS "${CMAKE_BINARY_DIR}/*")
#    file(REMOVE ${ALL_EXECS})
#endif()
#


