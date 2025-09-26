build:
	cmake -G Ninja -B build -S . -D CMAKE_BUILD_TYPE=Release
	cmake --build build
	swig -c++ -go -intgosize 64 webview.i