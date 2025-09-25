#include "webview/webview.h"

#include <iostream>

#ifdef _WIN32
int WINAPI WinMain(HINSTANCE /*hInst*/, HINSTANCE /*hPrevInst*/,
                   LPSTR /*lpCmdLine*/, int /*nCmdShow*/) {
#else
int main() {
#endif
  try {
    webview::webview w(false, nullptr);
    w.set_title("Basic Example");
    w.set_size(480, 320, WEBVIEW_HINT_NONE);
    w.set_html("Thanks for using webview!");
    w.set_user_agent(
        "Dalvik/2.1.0 (Linux; U; Android 6.0.1; Lenovo YT3-850M Build/MMB29M)");
    w.navigate("https://www.whatismybrowser.com/detect/what-is-my-user-agent/");
    w.run();
  } catch (const webview::exception &e) {
    std::cerr << e.what() << '\n';
    return 1;
  }

  return 0;
}
