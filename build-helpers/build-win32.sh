#!/bin/bash
rm -rf builddir-win32
cd ..
meson --cross-file meson-win32.config --prefix=/ --bindir=/ --libdir=/ build-helpers/builddir-win32
cd build-helpers
ninja -C builddir-win32
mkdir -p builddir-win32/installation
DESTDIR=installation ninja -C builddir-win32 install
cp /usr/i686-w64-mingw32/bin/libstdc++-6.dll builddir-win32/installation/
cp /usr/i686-w64-mingw32/bin/libgcc_s_sjlj-1.dll builddir-win32/installation/
cp /usr/i686-w64-mingw32/bin/libwinpthread-1.dll builddir-win32/installation/