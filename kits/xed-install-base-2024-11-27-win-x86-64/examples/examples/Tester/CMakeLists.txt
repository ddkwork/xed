cmake_minimum_required(VERSION 3.30)
set(CMAKE_C_STANDARD 11)
project(Tester C)

include_directories(../../include)
link_directories(${CMAKE_SOURCE_DIR})
add_executable(
xed-tester.c
xed-examples-util.c
)
add_executable(Tester34)
target_link_libraries(Tester xed)
