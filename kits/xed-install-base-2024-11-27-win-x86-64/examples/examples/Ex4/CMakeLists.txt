cmake_minimum_required(VERSION 3.30)
set(CMAKE_C_STANDARD 11)
project(Ex4 C)

include_directories(../../include)
link_directories(${CMAKE_SOURCE_DIR})
add_executable(
xed-ex4.c
xed-examples-util.c
)
add_executable(Ex434)
target_link_libraries(Ex4 xed)
