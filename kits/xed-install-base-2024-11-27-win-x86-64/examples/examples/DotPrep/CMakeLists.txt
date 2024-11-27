cmake_minimum_required(VERSION 3.30)
set(CMAKE_C_STANDARD 11)
project(DotPrep C)

include_directories(../../include)
link_directories(${CMAKE_SOURCE_DIR})
add_executable(
xed-dot-prep.c
xed-examples-util.c
)
add_executable(DotPrep34)
target_link_libraries(DotPrep xed)
