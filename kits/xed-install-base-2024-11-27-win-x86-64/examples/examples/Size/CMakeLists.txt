cmake_minimum_required(VERSION 3.30)
set(CMAKE_C_STANDARD 11)
project(Size C)

include_directories(../../include)
link_directories(${CMAKE_SOURCE_DIR})
add_executable(
xed-size.c
xed-examples-util.c
)
add_executable(Size34)
target_link_libraries(Size xed)
