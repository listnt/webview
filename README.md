# Requirements

A webview fork with additional ability to set user-agent. 

```
Right now supports setting user-agent only on linux 
```
## Tools

- Cmake
- Ninja

## Packages

* Debian:
  * WebKitGTK 6.0, GTK 4:
    * Development: `apt install libgtk-4-dev libwebkitgtk-6.0-dev`
    * Production: `apt install libgtk-4-1 libwebkitgtk-6.0-4`
  * WebKitGTK 4.1, GTK 3, libsoup 3:
    * Development: `apt install libgtk-3-dev libwebkit2gtk-4.1-dev`
    * Production: `apt install libgtk-3-0 libwebkit2gtk-4.1-0`
  * WebKitGTK 4.0, GTK 3, libsoup 2:
    * Development: `apt install libgtk-3-dev libwebkit2gtk-4.0-dev`
    * Production: `apt install libgtk-3-0 libwebkit2gtk-4.0-37`
* Fedora:
  * WebKitGTK 6.0, GTK 4:
    * Development: `dnf install gtk4-devel webkitgtk6.0-devel`
    * Production: `dnf install gtk4 webkitgtk6.0`
  * WebKitGTK 4.1, GTK 3, libsoup 3:
    * Development: `dnf install gtk3-devel webkit2gtk4.1-devel`
    * Production: `dnf install gtk3 webkit2gtk4.1`
  * WebKitGTK 4.0, GTK 3, libsoup 2:
    * Development: `dnf install gtk3-devel webkit2gtk4.0-devel`
    * Production: `dnf install gtk3 webkit2gtk4.0`
* FreeBSD:
  * GTK 4: `pkg install webkit2-gtk4`
  * GTK 3: `pkg install webkit2-gtk3`

## Development Dependencies

In addition to the dependencies mentioned earlier in this document for developing *with* the webview library, the following are used during development *of* the webview library.

* Amalgamation:
  * Python >= 3.9
* Checks:
  * `clang-format`
  * `clang-tidy`
* Documentation:
  * Doxygen
  * Graphvis

# how to compile 

```
make build
```