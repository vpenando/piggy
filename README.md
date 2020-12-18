# Piggy

[![Build Status](https://travis-ci.org/vpenando/piggy.svg)](https://travis-ci.org/vpenando/piggy)
[![Go Report Card](https://goreportcard.com/badge/github.com/vpenando/piggy)](https://goreportcard.com/report/github.com/vpenando/piggy)

Piggy is a free, open source software designed to help you track and classify your spendings.

It provides features such as:
* Adding, editing, deleting, sorting & filtering expenses
* Adding custom categories
* Navigating between months by pressing the Page Up / Page Down keys
* Generating Excel reports
* And much more!

It also embeds some basic security features to prevent HTML injections (into the category names or expense descriptions) and malware uploads.

New features (statistics, annual reports, ...) are already under development.

## Installation

Start the server (or add it to your startup programs), then open `http://localhost:8081` in your favorite browser. That's it!
The application has no third party dependency.

**Tip**: If you want to set Piggy as a desktop-like app, it's possible with Google Chrome!
* Open the Chrome menu
* Click "More tools"
* Click "Create shortcut"
* Check "Open as window"

Et voilà! 

## Custom build

If you want to build the application by yourself, run the following command: `go build -o <APPLICATION_NAME>`.
