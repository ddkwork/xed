cmake_minimum_required(VERSION 3.30)
set(CMAKE_C_STANDARD 11)
project(Ex7 C)

include_directories(include)
link_directories(${CMAKE_SOURCE_DIR})
add_executable(Ex7
xed-ex7.c
xed-examples-util.c
)
target_link_libraries(Ex7 xed)
