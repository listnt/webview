/*
 * MIT License
 *
 * Copyright (c) 2017 Serge Zaitsev
 * Copyright (c) 2022 Steffen André Langnes
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in
 * all copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 * 
 * Additional Notice:
 * Copyright (c) 2025 Aidar Kudabaev
 * All rights reserved.
 * This code is provided solely for the purpose of evaluating my
 * candidacy for employment. It may not be used, copied, modified,
 * or distributed without explicit permission from the author.
 */

%module webview

%{
#include "core/include/webview.h"
%}

typedef void *webview_t;

/// Window size hints
typedef enum {
  /// Width and height are default size.
  WEBVIEW_HINT_NONE,
  /// Width and height are minimum bounds.
  WEBVIEW_HINT_MIN,
  /// Width and height are maximum bounds.
  WEBVIEW_HINT_MAX,
  /// Window size can not be changed by a user.
  WEBVIEW_HINT_FIXED
} webview_hint_t;

extern webview_t webview_create(int debug, void *window);
extern void webview_destroy(webview_t w);
extern void webview_run(webview_t w);
extern void webview_terminate(webview_t w);
extern void webview_dispatch(webview_t w, void (*fn)(webview_t w, void *arg),
                             void *arg);
extern void *webview_get_window(webview_t w);
extern void *webview_get_native_handle(webview_t w,
                                       webview_native_handle_kind_t kind);
extern void webview_set_title(webview_t w, const char *title);
extern void webview_set_size(webview_t w, int width, int height,
                             webview_hint_t hints);
extern void webview_navigate(webview_t w, const char *url);
extern void webview_set_user_agent(webview_t w, const char *user_agent);
extern void webview_set_html(webview_t w, const char *html);
extern void webview_init(webview_t w, const char *js);
extern void webview_eval(webview_t w, const char *js);
extern void webview_bind(webview_t w, const char *name,
                         void (*fn)(const char *seq, const char *req,
                                    void *arg),
                         void *arg);
extern void webview_unbind(webview_t w, const char *name);
extern void webview_return(webview_t w, const char *seq, int status,
                           const char *result);
