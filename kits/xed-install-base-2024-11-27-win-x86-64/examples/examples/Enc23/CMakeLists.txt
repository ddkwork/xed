cmake_minimum_required(VERSION 3.30)
set(CMAKE_C_STANDARD 11)
project(Enc23 C)

include_directories(include)
link_directories(${CMAKE_SOURCE_DIR})
add_executable(Enc23
xed-enc2-3.c
xed-examples-util.c
)
target_link_libraries(Enc23 xed)
