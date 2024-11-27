cmake_minimum_required(VERSION 3.30)
set(CMAKE_C_STANDARD 11)
project(xed-find-special C)

include_directories(../../include)
link_directories(${CMAKE_SOURCE_DIR})
add_executable(xed-find-special34)
target_link_libraries(xed-find-special xed)
