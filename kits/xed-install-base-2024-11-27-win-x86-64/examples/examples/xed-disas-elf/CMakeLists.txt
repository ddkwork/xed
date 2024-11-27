cmake_minimum_required(VERSION 3.30)
set(CMAKE_C_STANDARD 11)
project(xed-disas-elf C)

include_directories(../../include)
link_directories(${CMAKE_SOURCE_DIR})
add_executable(xed-disas-elf34)
target_link_libraries(xed-disas-elf xed)
