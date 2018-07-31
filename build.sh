#!/bin/bash
rm -rf builddir
meson builddir
ninja -C builddir
if [[ `uname` == "MINGW64_NT-10.0" ]]
	then
		cd builddir
		mkdir -p installation
		DESTDIR=installation ninja install
		echo "build for windows finished"
	else echo "build for linux finished"
fi