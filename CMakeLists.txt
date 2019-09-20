cmake_minimum_required (VERSION 3.7)
project (thapo-calc LANGUAGES CXX)
set (CMAKE_CXX_STANDARD 14)
set (CMAKE_CXX_STANDARD_REQUIRED ON)
set (thapo-calc_VERSION_MAJOR 0)
set (thapo-calc_VERSION_MINOR 2)
set(CMAKE_WINDOWS_EXPORT_ALL_SYMBOLS ON)

include_directories(src/include)

add_subdirectory (src/thapo-calc-lib)
add_subdirectory (src/thapo-calc-cli)
add_subdirectory (src/thapo-calc-gui)
