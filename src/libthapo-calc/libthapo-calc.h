// Defines API to be used by other programms
#ifndef LIBTCALC_H
#define LIBTCALC_H

//Standard libraries which are used by libtcalc and the programs which use it
#include <iostream>
#include <string>

// error.h is used to catch errors in all programs
#include "error.h"

long double calculation(std::string given_string);

#endif
