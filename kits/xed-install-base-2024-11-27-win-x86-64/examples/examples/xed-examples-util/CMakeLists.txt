cmake_minimum_required(VERSION 3.30)
set(CMAKE_C_STANDARD 11)
project(xed-examples-util C)

include_directories(../../include)
link_directories(${CMAKE_SOURCE_DIR})
add_executable(xed-examples-util34)
target_link_libraries(xed-examples-util xed)
