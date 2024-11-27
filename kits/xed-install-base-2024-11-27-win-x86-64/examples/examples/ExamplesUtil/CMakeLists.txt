cmake_minimum_required(VERSION 3.30)
set(CMAKE_C_STANDARD 11)
project(ExamplesUtil C)

include_directories(../../include)
link_directories(${CMAKE_SOURCE_DIR})
add_executable(
xed-examples-util.c
xed-examples-util.c
)
add_executable(ExamplesUtil34)
target_link_libraries(ExamplesUtil xed)
