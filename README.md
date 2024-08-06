# TUI Dashboard
A Text User Input Dashboard.

This can be a basic dashboard that shows various information, and support long running TUI animations like from htop or pipe.sh.

## Building / dev setup
- Install golang
- Clone this repo
- `$ go get ...`

## Bugs / TODO
- [ ] debugger https://stackoverflow.com/a/39062734
- [ ] border on right is off by a pixel
- [ ] support resizing
- [ ] learn how to layout stuff
- [ ] 

## Goals / req'd features
- [x] display time / date
  - [ ] as asciitext big font
- [ ] display basic htop (top part)
- [ ] display output of any command (pipes.sh, htop, glances, etc.)
  - [ ] i.e. long running command support?
- [ ] display contents of file? Like todo's
- [x] styles
  - [x] colored border
  - [ ] color themes
- [ ] config file
  - [ ] theme color
  - [ ] selection of "plugins / things to run" <- decide on what to call these
  - [ ] support for any command
- [ ] Run up to 2 cmds at startup
  - [ ] e.g. htop and pipes.sh
  - [ ] Allow quitting each and editing the command on shortcut
  - [ ] Allow entering context of program on enter

## Stretch goals / wanted features
- [ ] support rendering .gif or .mp4 as ascii?
- [ ] support gui placement customization
- [ ] config editor


