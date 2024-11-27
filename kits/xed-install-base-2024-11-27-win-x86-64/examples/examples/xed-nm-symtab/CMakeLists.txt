cmake_minimum_required(VERSION 3.30)
set(CMAKE_C_STANDARD 11)
project(xed-nm-symtab C)

include_directories(../../include)
link_directories(${CMAKE_SOURCE_DIR})
add_executable(xed-nm-symtab34)
target_link_libraries(xed-nm-symtab xed)
