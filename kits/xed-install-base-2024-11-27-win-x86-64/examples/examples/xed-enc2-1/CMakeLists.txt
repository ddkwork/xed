cmake_minimum_required(VERSION 3.30)
set(CMAKE_C_STANDARD 11)
project(xed-enc2-1 C)

include_directories(../../include)
link_directories(${CMAKE_SOURCE_DIR})
add_executable(xed-enc2-134)
target_link_libraries(xed-enc2-1 xed)
