// used by libtapocalc.cpp

#ifndef GRAMMAR_H
#define GRAMMAR_H

#include <iostream>
#include <sstream>
#include <string>

long double expression (std::stringstream& input, bool parentheses);
long double term (std::stringstream& input);
long double power (std::stringstream& input);
long double primary (std::stringstream& input);
char read_nonspace_char (std::stringstream& input);
bool isprimary(char c);
/*	
double expression (std::stringstream& input, bool& success, std::string& error);
double term (std::stringstream& input, bool& success, std::string& error);
double primary (std::stringstream& input, bool& success, std::string& error);
*/
#endif
