cmake_minimum_required(VERSION 3.30)
set(CMAKE_C_STANDARD 11)
project(FindSpecial C)

include_directories(../../include)
link_directories(${CMAKE_SOURCE_DIR})
add_executable(
xed-find-special.c
xed-examples-util.c
)
add_executable(FindSpecial34)
target_link_libraries(FindSpecial xed)
