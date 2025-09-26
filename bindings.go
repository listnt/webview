/*
* Copyright (c) 2025 Aidar Kudabaev.
* All rights reserved.
* This code is provided solely for the purpose of evaluating my candidacy
* for employment. It may not be used, copied, modified, or distributed
* without explicit permission from the author.
 */

package webview

// #cgo pkg-config: gtk4 webkitgtk-6.0
// #cgo LDFLAGS: -L${SRCDIR}/build/core -lwebview
// #include "core/include/webview.h"
import "C"
