# thapo-calc
ThApo-Calc (Thanos Apostolou' Calculator) is a simple calculator written in c++. Currently there is the command line app **thapo-calc-cli** which is using the library **libthapo-calc**. In the future a Gui will be added, most likely **thapo-calc-qt** which will be using the same library.


**BUILD DEPENDENCIES**
  - cmake & ninja (or make)
  - Qt5Qml
  - Qt5Core
  - Qt5Gui
  - Qt5Quick
  - Qt5QuickControls2

**BUILD INSTRUCTIONS**
  - LINUX
    ```
    cmake -GNinja -DCMAKE_INSTALL_PREFIX=/usr -DCMAKE_BUILD_TYPE=Release . -B build
    ninja -C build
    DESTDIR=/  ninja -C build install
    ```

  - ANDROID from Linux

    Setup QtCreator with all android kits (using QT MaintenanceTool)
    Add all kits to the project and build the apk files.

  - WINDOWS

    Install VS Studio with msvc
    Install cmake and make and add them to path
    Setup Qt Creator with latest qt msvc builds.
    Configure project and change CMAKE_INSTALL_PREFIX to PROJECT\LOCATION\out\install
    Add build step `cmake.exe --build . --target install` to Release configuration.
    Or open "x64 native tools command prompt for VS 2019" and run:
    ```
    cmake -G Ninja -DCMAKE_BUILD_TYPE=Release -DCMAKE_PREFIX_PATH=C:\Qt\5.13.1\msvc2017_64 -DCMAKE_INSTALL_PREFIX=build/install . -B build
    ninja -C build
    ninja -C build install
    ```