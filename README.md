# TUI Dashboard
A Text User Input Dashboard. This can be a basic dashboard that shows various information, and maybe have a couple light animations.

## Building / dev setup
- Install golang
- Clone this repo
- `$ go mod tidy`

## Installing
- `$ go install github.com/ssebs/tui-dash`

## Bugs / TODO
- [ ] https://www.inngest.com/blog/interactive-clis-with-bubbletea
- [ ] Decide what I want this to be
- [ ] support resizing
- [ ] learn how to layout stuff

## Goals
- [x] display time / date
  - [ ] as asciitext big font
- [ ] display system info
  - [ ] RAM usage
  - [ ] CPU usage
  - [ ] DISK / Network IO / usage?
  - [ ] hostname, IP addr
- [ ] MOTD (curl an API?)
- [ ] display contents of file? Like todo's
- [ ] Create an animation
- [ ] Create breakout?
- [x] styles
  - [x] colored border
  - [ ] color themes



