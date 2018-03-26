[![GoDoc](https://godoc.org/github.com/sedx/svg?status.svg)](https://godoc.org/github.com/sedx/svg)

Since SVG based on XML this library uses all power of Go `encoding/xml` to build SVG documents.

**WORK IN PROGRESS**
Library is not stable yet. Big changes of package structure, API and other things can be occurred.

### Features

- [ ] Full SVG1.2 support ()
    - [ ] All elements
    - [ ] All attributes
    - [ ] Constants for attributes values
- [ ] Unmarshal SVG to SVG struct
    - [ ] Custom types should implement `Unmarshaler` interface
    - [ ] Add test (for example from  https://github.com/w3c/web-platform-tests/tree/master/svg)
- [ ] SVG elements lookup
- [ ] Modifying exists SVG elements
- [ ] Insert SVG elements
    - [ ] In Specified position
    - [ ] After some elements
    - [ ] Before some element
    - [ ] Instead some element
    
