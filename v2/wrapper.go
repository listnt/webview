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

package webview

import (
	"runtime"
	"sync"
	"unsafe"
	C "github.com/listnt/webview"
)

func init() {
	// Ensure that main.main is called from the main thread
	runtime.LockOSThread()
}

// Hints are used to configure window sizing and resizing
type Hint int

const (
	// Width and height are default size
	HintNone = 0

	// Window size can not be changed by a user
	HintFixed = 1

	// Width and height are minimum bounds
	HintMin = 2

	// Width and height are maximum bounds
	HintMax = 3
)

type WebView interface {

	// Run runs the main loop until it's terminated. After this function exits -
	// you must destroy the webview.
	Run()

	// Terminate stops the main loop. It is safe to call this function from
	// a background thread.
	Terminate()

	// Destroy destroys a webview and closes the native window.
	Destroy()

	// Window returns a native window handle pointer. When using GTK backend the
	// pointer is GtkWindow pointer, when using Cocoa backend the pointer is
	// NSWindow pointer, when using Win32 backend the pointer is HWND pointer.
	Window() unsafe.Pointer

	// SetTitle updates the title of the native window. Must be called from the UI
	// thread.
	SetTitle(title string)

	// SetSize updates native window size. See Hint constants.
	SetSize(w int, h int, hint Hint)

	// Navigate navigates webview to the given URL. URL may be a properly encoded data.
	// URI. Examples:
	// w.Navigate("https://github.com/webview/webview")
	// w.Navigate("data:text/html,%3Ch1%3EHello%3C%2Fh1%3E")
	// w.Navigate("data:text/html;base64,PGgxPkhlbGxvPC9oMT4=")
	Navigate(url string)

	SetUserAgent(userAgent string)

	// SetHtml sets the webview HTML directly.
	// Example: w.SetHtml(w, "<h1>Hello</h1>");
	SetHtml(html string)

	// Init injects JavaScript code at the initialization of the new page. Every
	// time the webview will open a the new page - this initialization code will
	// be executed. It is guaranteed that code is executed before window.onload.
	Init(js string)

	// Eval evaluates arbitrary JavaScript code. Evaluation happens asynchronously,
	// also the result of the expression is ignored. Use RPC bindings if you want
	// to receive notifications about the results of the evaluation.
	Eval(js string)
}

type webview struct {
	w uintptr
}

var (
	m     sync.Mutex
	index uintptr
)

func boolToInt(b bool) int {
	if b {
		return 1
	}
	return 0
}

// New calls NewWindow to create a new window and a new webview instance. If debug
// is non-zero - developer tools will be enabled (if the platform supports them).
func New(debug bool) WebView { return NewWindow(debug, nil) }

// NewWindow creates a new webview instance. If debug is non-zero - developer
// tools will be enabled (if the platform supports them). Window parameter can be
// a pointer to the native window handle. If it's non-null - then child WebView is
// embedded into the given parent window. Otherwise a new window is created.
// Depending on the platform, a GtkWindow, NSWindow or HWND pointer can be passed
// here.
func NewWindow(debug bool, window unsafe.Pointer) WebView {
	w := &webview{}
	w.w = C.Webview_create(boolToInt(debug), uintptr(window))
	return w
}

func (w *webview) Destroy() {
	C.Webview_destroy(w.w)
}

func (w *webview) Run() {
	C.Webview_run(w.w)
}

func (w *webview) Terminate() {
	C.Webview_terminate(w.w)
}

func (w *webview) Window() unsafe.Pointer {
	return unsafe.Pointer(C.Webview_get_window(w.w))
}

func (w *webview) Navigate(url string) {
	C.Webview_navigate(w.w, url)
}

func (w *webview) SetHtml(html string) {
	C.Webview_set_html(w.w, html)
}

func (w *webview) SetTitle(title string) {
	C.Webview_set_title(w.w, title)
}

func (w *webview) SetSize(width int, height int, hint Hint) {
	C.Webview_set_size(w.w, width, height, C.Webview_hint_t(hint))
}

func (w *webview) Init(js string) {
	C.Webview_init(w.w, js)
}

func (w *webview) Eval(js string) {
	C.Webview_eval(w.w, js)
}

func (w *webview) SetUserAgent(userAgent string) {
	C.Webview_set_user_agent(w.w, userAgent)
}
