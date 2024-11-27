cmake_minimum_required(VERSION 3.30)
set(CMAKE_C_STANDARD 11)
project(SymbolTable C)

include_directories(../../include)
link_directories(${CMAKE_SOURCE_DIR})
add_executable(
xed-symbol-table.c
xed-examples-util.c
)
add_executable(SymbolTable34)
target_link_libraries(SymbolTable xed)
