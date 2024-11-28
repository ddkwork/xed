cmake_minimum_required(VERSION 3.30)
set(CMAKE_C_STANDARD 11)
project(Ex3 C)

include_directories(../../include)
link_directories(${CMAKE_SOURCE_DIR})
add_executable(Ex3
xed-ex3.c
xed-examples-util.c
)
target_link_libraries(Ex3 xed)
