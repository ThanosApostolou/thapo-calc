# tapocalc
TApoCalc (Thanos Apostolou' Calculator) is a simple calculator written in c++. Currently there is the application **tapocalc-cli** which is using the library **libtapocalc**. In the future a Gui will be added, most likely **tapocalc-qt** which will be using the same library.

**BUILD INSTRUCTIONS**
```
meson --prefix=/usr --buildtype=release builddir
cd builddir
ninja
DESTDIR="DESTINATION" ninja install
```
**BUILD BINARIES FOR WINDOWS**

Use build-win64.sh (or build-win32.sh) and distribute the tapocalc-win64 (or tapocalc-win32) folder in a zip archive.
