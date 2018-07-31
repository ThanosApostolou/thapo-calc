#include <iostream>
#include <gtkmm.h>
#include "buttons.h"

int main (int argc, char* argv[]) {
	std::cout << "hello gui\n";
	auto app = Gtk::Application::create(argc, argv, "org.gtkmm.examples.base");
	Gtk::Window window;
	window.set_default_size(200, 200);
	Buttons button1;
	return app->run(button1);
}