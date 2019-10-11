set (CMAKE_INCLUDE_CURRENT_DIR ON)
set (CMAKE_AUTOMOC ON)
set (CMAKE_AUTORCC ON)
set (CMAKE_AUTOUIC ON)

find_package (Qt5 COMPONENTS Core Gui Quick QuickControls2 Qml REQUIRED)

set (GUI_SOURCES thapo-calc-gui.cpp qml.qrc)
if(WIN32)
    set (GUI_TYPE WIN32)
endif(WIN32)
# Set some Apple MacOS Specific settings
if (APPLE)
    set (GUI_TYPE MACOSX_BUNDLE)
endif(APPLE)


if(ANDROID)
    add_library(thapo-calc-gui SHARED ${GUI_SOURCES})
else()
    add_executable (thapo-calc-gui ${GUI_TYPE} ${GUI_SOURCES})
endif()

target_compile_definitions(thapo-calc-gui PRIVATE $<$<OR:$<CONFIG:Debug>,$<CONFIG:RelWithDebInfo>>:QT_QML_DEBUG>)
target_link_libraries (thapo-calc-gui PRIVATE thapo-calc-lib Qt5::Core Qt5::Gui Qt5::Quick Qt5::QuickControls2 Qt5::Qml)

install (
    TARGETS thapo-calc-gui
    DESTINATION ${CMAKE_INSTALL_BINDIR}
)

IF (WIN32)
	get_target_property(_qmake_executable Qt5::qmake IMPORTED_LOCATION)
	get_filename_component(_qt_bin_dir "${_qmake_executable}" DIRECTORY)
	find_program(WINDEPLOYQT_EXECUTABLE windeployqt HINTS "${_qt_bin_dir}")
	add_custom_command(TARGET thapo-calc-gui POST_BUILD COMMAND "${CMAKE_COMMAND}" -E env PATH="${_qt_bin_dir}" "${WINDEPLOYQT_EXECUTABLE}" --qmldir "${CMAKE_CURRENT_SOURCE_DIR}" --dir "${CMAKE_INSTALL_PREFIX}/${CMAKE_INSTALL_BINDIR}" "$<TARGET_FILE:thapo-calc-gui>")
ENDIF()
