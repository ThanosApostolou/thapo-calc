#!/bin/bash
rm -rf builddir
cd ..
meson build-helpers/builddir --prefix=/ --bindir=/ --libdir=/ -Dbundle=true
cd build-helpers
ninja -C builddir
mkdir -p builddir/installation
DESTDIR=installation ninja -C builddir install