# checkr

### A cli tool to check whether a package name is available on different package managers.

![Screen Recording](https://raw.githubusercontent.com/shopnilsazal/checkr/master/recording.gif)


## Install

```
$ brew install shopnilsazal/tap/checkr
```

Or you can download binary from [releases](https://github.com/shopnilsazal/checkr/releases) page.


## Usage

```
$ checkr -n packageName -m packageManager
$ checkr -n django -m pip

To search from multiple package managers, write them using comma
$ checkr -n django -m pip,gem,npm

To search from all supported package managers, write 'all' keyword
$ checkr -n django -m all


$ checkr --help
  -n string
    Provide a package name you want to search. (default "react")
  -m string
    Provide package managers where you want to search. (default "npm")

Examples
  $ checkr -n django -m pip
  ✖ django is unavailable on 'pip'

  $ checkr -n rails -m pip,gem,npm
  ✖ rails is unavailable on 'pip'
  ✖ rails is unavailable on 'gem'
  ✖ rails is unavailable on 'npm'
  
  $ checkr -n mynewpackage007 -m all
  ✔ mynewpackage007 is available on 'gem'
  ✔ mynewpackage007 is available on 'hex'
  ✔ mynewpackage007 is available on 'pip'
  ✔ mynewpackage007 is available on 'npm'
  ✔ mynewpackage007 is available on 'crates'

```

### Available Package Manager to Search
```
pip - Python Package Index
npm - NodeJS Package Manager
gem - Ruby Gem
hex - Package manager for the Erlang ecosystem
crates - The Rust crate Registry
```