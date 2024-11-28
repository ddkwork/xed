cmake_minimum_required(VERSION 3.30)
set(CMAKE_C_STANDARD 11)
project(Ex6 C)

include_directories(../../include)
link_directories(${CMAKE_SOURCE_DIR})
add_executable(Ex6
xed-ex6.c
xed-examples-util.c
)
target_link_libraries(Ex6 xed)
