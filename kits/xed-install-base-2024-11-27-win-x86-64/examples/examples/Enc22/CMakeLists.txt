cmake_minimum_required(VERSION 3.30)
set(CMAKE_C_STANDARD 11)
project(Enc22 C)

include_directories(../../include)
link_directories(${CMAKE_SOURCE_DIR})
add_executable(Enc2234)
target_link_libraries(Enc22 xed)
