// API used by all programs which also includes error.h
#include "thapo-calc-lib.h"
#include "grammar.h"

//Standard libraries which used by libtcalc
#include <iostream>
#include <sstream>

long double calculation(std::string given_string) {
	char c=' ';
	bool success=true;
	std::stringstream input;
	input << given_string;
	try {
		return expression(input, false);
	}
	catch (Myerror myerror) {
		throw myerror;
	}	
}
