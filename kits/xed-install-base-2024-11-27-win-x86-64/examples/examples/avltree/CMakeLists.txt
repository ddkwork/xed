cmake_minimum_required(VERSION 3.30)
set(CMAKE_C_STANDARD 11)
project(Avltree C)

include_directories(../../include)
link_directories(${CMAKE_SOURCE_DIR})
add_executable(
avltree.c
xed-examples-util.c
)
add_executable(Avltree34)
target_link_libraries(Avltree xed)
