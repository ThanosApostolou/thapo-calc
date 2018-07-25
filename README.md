# thapo-calc
ThApo-Calc (Thanos Apostolou' Calculator) is a simple calculator written in c++. Currently there is the command line app **thapo-calc-cli** which is using the library **libthapo-calc**. In the future a Gui will be added, most likely **thapo-calc-qt** which will be using the same library.

**BUILD INSTRUCTIONS**
```
meson --prefix=/usr --buildtype=release builddir
cd builddir
ninja
ninja install
```
**BUILD BINARIES FOR WINDOWS**

Use the scripts in **build-helpers** folder and distribute the **installation** folders in zip archives.
