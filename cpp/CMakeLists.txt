cmake_minimum_required (VERSION 3.7)
project (thapo-calc LANGUAGES CXX)
set (CMAKE_CXX_STANDARD 17)
set (CMAKE_CXX_STANDARD_REQUIRED ON)
set (thapo-calc_VERSION_MAJOR 0)
set (thapo-calc_VERSION_MINOR 2)
set(CMAKE_WINDOWS_EXPORT_ALL_SYMBOLS ON)

if (NOT CMAKE_INSTALL_BINDIR)
    set (CMAKE_INSTALL_BINDIR bin)
endif (NOT CMAKE_INSTALL_BINDIR)
if (NOT CMAKE_INSTALL_LIBDIR)
    set (CMAKE_INSTALL_LIBDIR lib)
endif (NOT CMAKE_INSTALL_LIBDIR)
if (NOT CMAKE_INSTALL_INCLUDEDIR)
    set (CMAKE_INSTALL_INCLUDEDIR include)
endif (NOT CMAKE_INSTALL_INCLUDEDIR)

IF(MSVC)
    SET(CMAKE_CXX_FLAGS "/EHsc")
ENDIF(MSVC)

include_directories(src/include)

add_subdirectory (src/thapo-calc-lib)
#add_subdirectory (src/thapo-calc-gui)

if(NOT ANDROID)
    add_subdirectory (src/thapo-calc-cli)
endif()
