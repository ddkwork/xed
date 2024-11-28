cmake_minimum_required(VERSION 3.30)
set(CMAKE_C_STANDARD 11)
project(AsmParseMain C)

include_directories(../../include)
link_directories(${CMAKE_SOURCE_DIR})
add_executable(AsmParseMain
xed-asmparse-main.c
xed-examples-util.c
)
target_link_libraries(AsmParseMain xed)
