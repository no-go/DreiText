GTKMMFLAG = `pkg-config --cflags --libs gtkmm-3.0`

all:
	g++ -Wall DreiText.cpp -o DreiText $(GTKMMFLAG)

clean:
	rm -rf DreiText
