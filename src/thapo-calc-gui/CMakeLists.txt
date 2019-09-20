set (CMAKE_INCLUDE_CURRENT_DIR ON)
set (CMAKE_AUTOMOC ON)
set (CMAKE_AUTORCC ON)
set (CMAKE_AUTOUIC ON)

find_package (Qt5 COMPONENTS Core Gui Quick QuickControls2 Qml REQUIRED)

set (GUI_SOURCES thapo-calc-gui.cpp qml.qrc)
add_executable (thapo-calc-gui ${GUI_SOURCES})
target_link_libraries (thapo-calc-gui PRIVATE thapo-calc-lib Qt5::Core Qt5::Gui Qt5::Quick Qt5::QuickControls2 Qt5::Qml)

install (TARGETS thapo-calc-gui DESTINATION bin)

IF (WIN32)
	get_target_property(_qmake_executable Qt5::qmake IMPORTED_LOCATION)
	get_filename_component(_qt_bin_dir "${_qmake_executable}" DIRECTORY)
	find_program(WINDEPLOYQT_EXECUTABLE windeployqt HINTS "${_qt_bin_dir}")
	add_custom_command(TARGET thapo-calc-gui POST_BUILD COMMAND "${CMAKE_COMMAND}" -E env PATH="${_qt_bin_dir}" "${WINDEPLOYQT_EXECUTABLE}" --qmldir "${CMAKE_CURRENT_SOURCE_DIR}" "$<TARGET_FILE:thapo-calc-gui>")
	install (CODE "execute_process(COMMAND \"${CMAKE_COMMAND}\" -E env PATH=\"${_qt_bin_dir}\" \"${WINDEPLOYQT_EXECUTABLE}\" --qmldir \"${CMAKE_CURRENT_SOURCE_DIR}\" --release \"${CMAKE_INSTALL_PREFIX}/bin/thapo-calc-gui.exe\")")
ENDIF()