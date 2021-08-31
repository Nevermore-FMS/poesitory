## poesitory pull

Pulls a plugin from Poesitory

### Synopsis

Pulls a plugin from Poesitory
	- The identifier parameter comes in the following form:

		pluginname[#channel][@version]

		If channel is omitted, it will default to STABLE
		If version is omitted, it will default to the latest version for the given channel
	

```
poesitory pull [identifier] [flags]
```

### Options

```
  -h, --help          help for pull
      --path string   path to place the plugin
```

### Options inherited from parent commands

```
  -p, --upload string   token for upload authentication
  -u, --user string     token for user authentication
  -w, --web             use web authentication
```

### SEE ALSO

* [poesitory](poesitory.md)	 - Poesitory CLI allows you to push and pull Nevermore Plugins to or from Poesitory.

