igo
===

igo is a collection of interfaces around Golang's standard packages - making it easy to mock and make assertions about calls to functions within these packages.

A fake (generated using [counterfeiter](https://github.com/maxbrunsfeld/counterfeiter)) is also provided for each interface along with helper functions that makes more complicated faking easy.

Structs that expose public attributes have been given corresponding get and set methods.

This is by no means a complete collection of interfaces and fakes and the focus is on the most commonly used packages.

Packages
========
- [ios](os/)
- [iexec](os/exec/)

Installation
============

See package specific documentation
