#!/bin/bash
ldd $MESON_INSTALL_DESTDIR_PREFIX/$1/libthapo-calc.dll|grep /mingw| cut -d' ' -f3| while read line; do cp $line $MESON_INSTALL_DESTDIR_PREFIX/$1; done
ldd $MESON_INSTALL_DESTDIR_PREFIX/$1/thapo-calc-cli.exe|grep /mingw| cut -d' ' -f3| while read line; do cp $line $MESON_INSTALL_DESTDIR_PREFIX/$1; done
ldd $MESON_INSTALL_DESTDIR_PREFIX/$1/thapo-calc-gtkmm.exe|grep /mingw| cut -d' ' -f3| while read line; do cp $line $MESON_INSTALL_DESTDIR_PREFIX/$1; done
mkdir -p $MESON_INSTALL_DESTDIR_PREFIX/$2
cp -r /mingw64/lib/gtkmm-3.0 \
	/mingw64/lib/gdk-pixbuf-2.0 \
	$MESON_INSTALL_DESTDIR_PREFIX/$2