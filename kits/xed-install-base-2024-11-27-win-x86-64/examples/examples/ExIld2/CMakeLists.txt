cmake_minimum_required(VERSION 3.30)
set(CMAKE_C_STANDARD 11)
project(ExIld2 C)

include_directories(../../include)
link_directories(${CMAKE_SOURCE_DIR})
add_executable(ExIld234)
target_link_libraries(ExIld2 xed)
