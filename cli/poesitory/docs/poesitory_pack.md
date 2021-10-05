## poesitory pack

Creates a tarball of a plugin

```
poesitory pack [flags]
```

### Options

```
  -h, --help            help for pack
      --no-npm          disables running "npm run build" before packing the plugin
      --output string   sets the output tarball file (default is ./plugin-name.tar.gz)
      --path string     the path to the plugin (default ".")
```

### Options inherited from parent commands

```
  -p, --upload string   token for upload authentication
  -u, --user string     token for user authentication
  -w, --web             use web authentication
```

### SEE ALSO

* [poesitory](poesitory.md)	 - Poesitory CLI allows you to push and pull Nevermore Plugins to or from Poesitory.

