#include <gtkmm.h>
#include <iostream>
#include <string>
#include <vector>
#include <fstream>

#define UI_FILE "DreiText.ui"
#define PARTSIZE 1200

std::vector<char> vec;
Gtk::TextView* text1 = 0;
Gtk::TextView* text2 = 0;
Gtk::TextView* text3 = 0;
char * textstr1 = new char[PARTSIZE];
char * textstr2 = new char[PARTSIZE];
char * textstr3 = new char[PARTSIZE];
Gtk::Scale * scroller = 0;

void subs(char * retval, std::vector<char> v, unsigned int from) {
	for (unsigned int i=0; i<PARTSIZE; i++) {
		if ((i+from) < v.size()) {
			retval[i] = v[i+from];
		} else {
			retval[i] = '\0';
		}
	}
}

void scrolling() {
	long val = scroller->get_value();
	val = (long) ( (double) vec.size() * (double) val/1000.0);
	
	subs(textstr1, vec, val);
	text1->get_buffer()->set_text(textstr1);
	subs(textstr2, vec, val + PARTSIZE);
	text2->get_buffer()->set_text(textstr2);
	subs(textstr3, vec, val + 2*PARTSIZE);
	text3->get_buffer()->set_text(textstr3);
}

int main (int argc, char *argv[]) {
	if (argc < 2) {
		std::cerr << "Please add a filename." << std::endl;
		return 2;
	}
	std::string filename = argv[1];
	Gtk::Main kit(argc, argv);

	//Load the Glade file and instiate its widgets:
	Glib::RefPtr<Gtk::Builder> builder;
	try {
		builder = Gtk::Builder::create_from_file(UI_FILE);
	} catch (const Glib::FileError & ex) {
		std::cerr << ex.what() << std::endl;
		return 1;
	}
	Gtk::Window* main_win = 0;
	builder->get_widget("main_window", main_win);

	if (main_win == 0) {
		return 3;
	}
	
	builder->get_widget("text1", text1);
	builder->get_widget("text2", text2);
	builder->get_widget("text3", text3);
	builder->get_widget("scroller", scroller);

	std::ifstream file(filename);
	while (!(file.eof() || file.fail())) {
		char buffer[100];
		file.read(buffer, 100);
		vec.insert(vec.end(), buffer, buffer + file.gcount());
	}
	file.close();

	scrolling();
	scroller->signal_value_changed().connect(
		sigc::ptr_fun(&scrolling)
	);
	
	kit.run(*main_win);
	
	delete[] textstr1;
	delete[] textstr2;
	delete[] textstr3;
	
	return 0;
}
