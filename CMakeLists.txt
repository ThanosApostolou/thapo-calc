cmake_minimum_required (VERSION 3.7)
project (thapo-calc)
set (thapo-calc_VERSION_MAJOR 0)
set (thapo-calc_VERSION_MINOR 2)

include_directories(src/include)

add_subdirectory(src/thapo-calc-lib)
add_subdirectory(src/thapo-calc-cli)


#set(GTKMM_SOURCES src/thapo-calc-gtkmm/thapo-calc-gtkmm.cpp src/thapo-calc-gtkmm/buttons.cpp)
#add_executable(thapo-calc-gtkmm ${GTKMM_SOURCES})
#target_link_libraries(thapo-calc-gtkmm thapo-calc-lib)
