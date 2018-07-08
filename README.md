# gterm

gterm is intended to provide a straightforward terminal library for Go with a clear
separation between low level terminal management, basic screen output functions,
primitive graphics functions, and higher level text user interface framework.

It features:
* terminal initialization and management
* an event channel for receiving keyboard and mouse events
* basic primitive functions for querying and updating the display screen
* basic drawing routines
* basic collection of user interface widgets
* TUI (text user interface) framework

Ncurses is the gold standard for text user interface and this is *not* an attempt
to duplicate that functionality in Go. There are other excellent packages available
for Go that attempt to wrap, duplicate, or provide similar functionality in native Go.
The main difference between this and other packages is that the goal here is to
provide a useful set of layered abstractions that can be adopted as desired, from
simple screen, cursor, keyboard and mouse support, up to a full TUI framework, for
Go developers.

Internally, gterm uses a provider model that in theory allows different terminal
libraries to be used as the underlying tty driver wrapper. The current version uses
[termbox-go](https://github.com/nsf/termbox-go).

