#!/bin/bash
rm -rf builddir-win64
cd ..
meson --cross-file meson-win64.config --prefix=/ --bindir=/ --libdir=/ build-helpers/builddir-win64
cd build-helpers
ninja -C builddir-win64
mkdir -p builddir-win64/installation
DESTDIR=installation ninja -C builddir-win64 install
cp /usr/x86_64-w64-mingw32/bin/libstdc++-6.dll builddir-win64/installation/
cp /usr/x86_64-w64-mingw32/bin/libgcc_s_seh-1.dll builddir-win64/installation/
cp /usr/x86_64-w64-mingw32/bin/libwinpthread-1.dll builddir-win64/installation/