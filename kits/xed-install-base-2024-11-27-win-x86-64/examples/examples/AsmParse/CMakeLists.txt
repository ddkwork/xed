cmake_minimum_required(VERSION 3.30)
set(CMAKE_C_STANDARD 11)
project(AsmParse C)

include_directories(include)
link_directories(${CMAKE_SOURCE_DIR})
add_executable(AsmParse
xed-asmparse.c
xed-examples-util.c
)
target_link_libraries(AsmParse xed)
