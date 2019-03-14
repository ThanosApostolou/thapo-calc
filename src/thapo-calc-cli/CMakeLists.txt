set(CLI_SOURCES thapo-calc-cli.cpp)
add_executable(thapo-calc ${CLI_SOURCES})
target_link_libraries(thapo-calc thapo-calc-lib)

install (TARGETS thapo-calc DESTINATION bin)
