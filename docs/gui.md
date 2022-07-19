## Installation of GUI dependencies

The integrated OpenGL GUI support is enabled by default. Debugging
code will execute the built-in Emulator and GUI by default.

To select the GUI mode to use, set the following build flags:

* `nogui`: disables all GUI modules
* `noopengl` `sdl` enables the SDL GUI

The following libraries need to be installed, depending on the operating system:

### **macOS**

Xcode or Command Line Tools for Xcode:

```
xcode-select --install
```

### **Ubuntu/Debian-like**

For OpenGL support:

```
apt install build-essential libgl1-mesa-dev xorg-dev
```

For SDL support:

```
apt install libsdl2{,-image,-mixer,-ttf,-gfx}-dev
```

### **CentOS/Fedora-like**

For OpenGL support:

```
yum install @development-tools libX11-devel libXcursor-devel libXrandr-devel libXinerama-devel mesa-libGL-devel libXi-devel libXxf86vm-devel
```

For SDL support:

```
yum install SDL2{,_image,_mixer,_ttf,_gfx}-devel
```

### Windows

For SDL support:

1. Install [msys2](http://www.msys2.org/)
2. Start msys2 and execute:
```
pacman -S --needed base-devel mingw-w64-i686-toolchain mingw-w64-x86_64-toolchain mingw64/mingw-w64-x86_64-SDL2
```
3. Add `c:\tools\msys64\mingw64\bin\` to the user path environment variable
