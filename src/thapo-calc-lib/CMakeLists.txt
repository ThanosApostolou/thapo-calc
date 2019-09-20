set (LIB_SOURCES grammar.cpp thapo-calc-lib.cpp)
add_library (thapo-calc-lib SHARED ${LIB_SOURCES})
set_target_properties (thapo-calc-lib PROPERTIES PUBLIC_HEADER "../include/thapo-calc-lib.h")

install (
    TARGETS thapo-calc-lib
    PUBLIC_HEADER DESTINATION include
)
