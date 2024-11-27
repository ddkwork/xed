cmake_minimum_required(VERSION 3.30)
set(CMAKE_C_STANDARD 11)
project(xed-disas-hex C)

include_directories(../../include)
link_directories(${CMAKE_SOURCE_DIR})
add_executable(xed-disas-hex34)
target_link_libraries(xed-disas-hex xed)
