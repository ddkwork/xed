cmake_minimum_required(VERSION 3.30)
set(CMAKE_C_STANDARD 11)
project(Ex9Patch C)

include_directories(../../include)
link_directories(${CMAKE_SOURCE_DIR})
add_executable(
xed-ex9-patch.c
xed-examples-util.c
)
add_executable(Ex9Patch34)
target_link_libraries(Ex9Patch xed)
