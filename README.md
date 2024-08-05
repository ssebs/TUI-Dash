# TUI Dashboard
A Text User Input Dashboard.

This can be a basic dashboard that shows various information, and support long running TUI animations like from htop or pipe.sh.

## Building / dev setup
- Install golang
- Clone this repo
- `$ go get ...`

## Goals / req'd features
- [ ] display time / date
  - [ ] as asciitext big font
- [ ] display basic htop (top part)
- [ ] display output of any command (pipes.sh, htop, glances, etc.)
  - [ ] i.e. long running command support?
- [ ] display contents of file? Like todo's
- [ ] styles
  - [ ] colored border
  - [ ] color themes
- [ ] config file
  - [ ] theme color
  - [ ] selection of "plugins / things to run" <- decide on what to call these
  - [ ] support for any command

## Stretch goals / wanted features
- [ ] support rendering .gif or .mp4 as ascii?
- [ ] support gui placement customization
- [ ] config editor


