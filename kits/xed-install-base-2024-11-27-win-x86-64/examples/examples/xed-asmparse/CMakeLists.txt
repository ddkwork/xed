cmake_minimum_required(VERSION 3.30)
set(CMAKE_C_STANDARD 11)
project(xed-asmparse C)

include_directories(../../include)
link_directories(${CMAKE_SOURCE_DIR})
add_executable(xed-asmparse34)
target_link_libraries(xed-asmparse xed)
