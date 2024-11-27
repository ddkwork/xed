cmake_minimum_required(VERSION 3.30)
set(CMAKE_C_STANDARD 11)
project(Enc23 C)

include_directories(../../include)
link_directories(${CMAKE_SOURCE_DIR})
add_executable(Enc2334)
target_link_libraries(Enc23 xed)
