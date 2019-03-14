set(CMAKE_INCLUDE_CURRENT_DIR ON)
set(CMAKE_AUTOMOC ON)
set(CMAKE_AUTORCC ON)

find_package(Qt5 COMPONENTS Core Quick REQUIRED)

set(GUI_SOURCES thapo-calc-gui.cpp qml.qrc)
add_executable(thapo-calc-gui ${GUI_SOURCES})
target_link_libraries(thapo-calc-gui PRIVATE thapo-calc-lib Qt5::Core Qt5::Quick)

install (TARGETS thapo-calc-gui DESTINATION bin)
