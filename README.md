# nuklear-fridge

__Warning: This library is work in progress! Not all nuklear elements are migrated yet, and the documentation is non existent. Use at own caution!__

nuklear-fridge is a small wrapper library around [golang-ui/nuklear](https://github.com/golang-ui/nuklear) with the aim to ease usage of the library.

The main feature of nuklear-fridge is the ability to assemble a so called _fridge_ - a collective description of multiple nuklear windows and their contents.

See the examples folder for some inspiration. The desktop example is taken directly from [golang-ui/nuklear](https://github.com/golang-ui/nuklear) and adapted to a _fridge_.

## Future

The eventual goal ofnuklear-fridge is to have fridges read and saved to the file system (probably xml or json). This enables to build reusable UIs, which can be easily shared. Also, this keeps the definition of content (the ui) code (and thus implementation) agnostic.

## Name

The name is a reference to the [fridge scene](https://www.youtube.com/watch?v=jn4Vhkmb4Lw) of the otherwise medicore Indiana Jones 4.
