// Used in all programms and libraries in order to catch errors

#ifndef ERROR_H
#define ERROR_H

#include <string>

class Myerror {
	public:
		bool error;
		std::string msg;
		Myerror () {
			this->error = false;
			this->msg = "";
		}
		Myerror (std::string msg) {
			error = true;
			this->msg = msg;
		}
};

#endif
