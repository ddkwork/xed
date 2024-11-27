cmake_minimum_required(VERSION 3.30)
set(CMAKE_C_STANDARD 11)
project(Ex8 C)

include_directories(../../include)
link_directories(${CMAKE_SOURCE_DIR})
add_executable(
xed-ex8.c
xed-examples-util.c
)
add_executable(Ex834)
target_link_libraries(Ex8 xed)
