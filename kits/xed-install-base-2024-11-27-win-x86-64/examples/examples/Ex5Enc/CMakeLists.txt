cmake_minimum_required(VERSION 3.30)
set(CMAKE_C_STANDARD 11)
project(Ex5Enc C)

include_directories(../../include)
link_directories(${CMAKE_SOURCE_DIR})
add_executable(Ex5Enc34)
target_link_libraries(Ex5Enc xed)
