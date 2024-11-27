cmake_minimum_required(VERSION 3.30)
set(CMAKE_C_STANDARD 11)
project(DisasFilter C)

include_directories(../../include)
link_directories(${CMAKE_SOURCE_DIR})
add_executable(
xed-disas-filter.c
xed-examples-util.c
)
add_executable(DisasFilter34)
target_link_libraries(DisasFilter xed)
