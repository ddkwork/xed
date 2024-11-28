cmake_minimum_required(VERSION 3.30)
set(CMAKE_C_STANDARD 11)
project(Tester C)

include_directories(../../include)
link_directories(${CMAKE_SOURCE_DIR})
add_executable(Tester
xed-tester.c
xed-examples-util.c
)
target_link_libraries(Tester xed)
