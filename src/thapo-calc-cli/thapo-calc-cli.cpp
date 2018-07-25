#include "libthapo-calc.h"
#include "error.h"
#include "limits"
int main(int argc, char* argv[]) {
	long double result=0;
	bool success=true;
	std::string given_string="", error="";
	if (argc == 1) {
		while (true) {
			result=0;
			given_string="";
			std::cout << "Enter expression for calculation or exit to quit (Supported Operations +, -, *, /, ^):\n> ";
			getline(std::cin,given_string);
			if (given_string == "exit") {
				break;
			}
			else {
				try {
					result = calculation(given_string);
					std::cout << "= " << result << "\n\n";
				}
				catch (Myerror myerror) {
					std::cout << myerror.msg << "\n\n";	
				}
			}
		}
	}
	else if (argc == 3) {
		std::string argument=argv[1];
		if (argument == "-t" || argument == "--text") {
			given_string = argv[2];
			try {
				result = calculation(given_string);
				std::cout << "= " << result << "\n\n";
			}
			catch (Myerror myerror) {
				std::cout << myerror.msg << "\n\n";	
			}
		}
	}
	else {
		std::cout << "Type tcalc to use the program or pas the \"-t\" option to get fast result\n";
	}
	return 0;
}
