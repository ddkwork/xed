cmake_minimum_required(VERSION 3.30)
set(CMAKE_C_STANDARD 11)
project(Enc21 C)

include_directories(include)
link_directories(${CMAKE_SOURCE_DIR})
add_executable(Enc21
xed-enc2-1.c
xed-examples-util.c
)
target_link_libraries(Enc21 xed)
