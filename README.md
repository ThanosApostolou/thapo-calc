# thapo-calc
ThApo-Calc (Thanos Apostolou' Calculator) is a simple calculator written in c++. Currently there is the command line app **thapo-calc-cli** which is using the library **libthapo-calc**. In the future a Gui will be added, most likely **thapo-calc-qt** which will be using the same library.


**BUILD DEPENDENCIES**
  - Qt5Qml
  - Qt5Core
  - Qt5Gui
  - Qt5Quick
  - Qt5QuickControls2

**BUILD INSTRUCTIONS**
  - LINUX
    ```
    mkdir build && cd build
    cmake -GNinja -DCMAKE_BUILD_TYPE=Release -DCMAKE_PREFIX_PATH="${HOME}/Qt/5.13.1/gcc_64 ..
    ninja
    DESTDIR=install ninja install
    ```
  - LINUX BUNDLE
    ```
    mkdir build && cd build
    cmake -GNinja -DCMAKE_BUILD_TYPE=Release " ..
    ninja
    DESTDIR=install ninja install
    ```
  - ANDROID from Linux

    open project with qtcreator (install the android kits using QT MaintenanceTool)

  - WINDOWS

    Configure vs studio and qt creator and use the cmake install target.
