# checkr

### A cli tool to check whether a package name is available on different package managers.

![Screen Recording](https://raw.githubusercontent.com/shopnilsazal/checkr/master/recording.gif)


## Usage

```
$ checkr --help

  Usage
    $ checkr
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