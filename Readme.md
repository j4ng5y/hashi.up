# Hashi.UP
---

This is a little tool to help grab the hashicorp tools more easily, rather than having to go to each web page, find your OS/Arch, and download the tool.

## Usage

```A hashicorp tool downloader/installer/updater/uninstaller.

Usage:
  hashi.up [flags]
  hashi.up [command]

Available Commands:
  help        Help about any command
  install     Install one, some, or all of the hashicorp tools
  remove      Remove one, some, or all of the hashicorp tools
  update      Update one, some, or all of the hashicorp tools

Flags:
  -h, --help      help for hashi.up
  -v, --version   version for hashi.up

Use "hashi.up [command] --help" for more information about a command.
```

### Installation

```Install one, some, or all of the hashicorp tools

Usage:
  hashi.up install [flags]

Flags:
      --all                            download or install all available hashicorp tools at their latest version
      --consul string[="1.8.0"]        download/install consul at a specific version
      --download-only                  only download the hashicorp tools, but don't install them
  -h, --help                           help for install
      --nomad string[="0.11.3"]        the nomad version to install, leave blank for latest
      --packer string[="1.6.0"]        the packer version to install, leave blank for latest
      --terraform string[="0.12.28"]   the terraform version to install, leave blank for latest
      --vagrant string[="2.2.9"]       the vagrant version to install, leave blank for latest
      --vault string[="1.4.2"]         the vault version to install, leave blank for latest
```

Example:

To install all tools at their latest versions: `hashi.up install`

To install one tool at it's latest version: `hashi.up install --terraform`

To install one tool at a specific version: `hashi.up install --terraform=0.11.1`

To install multiple, but not all tools at their latest version: `hashi.up install --terraform --packer --vault`

To install multiple, but not all tools at a specific version: `hashi.up install --terraform=0.11.1 --packer=1.5.6 --vault=1.4.0`

### Updating

```Update one, some, or all of the hashicorp tools

Usage:
  hashi.up update [flags]

Flags:
      --consul string[="1.8.0"]        the consul version to update to, leave blank for latest (default "1.8.0")
  -h, --help                           help for update
      --nomad string[="0.11.3"]        the nomad version to update to, leave blank for latest (default "0.11.3")
      --packer string[="1.6.0"]        the packer version to update to, leave blank for latest (default "1.6.0")
      --terraform string[="0.12.28"]   the terraform version to update to, leave blank for latest (default "0.12.28")
      --vagrant string[="2.2.9"]       the vagrant version to update to, leave blank for latest (default "2.2.9")
      --vault string[="1.4.2"]         the vault version to update to, leave blank for latest (default "1.4.2")        the vault version to install, leave blank for latest
```

Example:

To update all tools at their latest versions: `hashi.up update`

To update one tool at it's latest version: `hashi.up update --terraform`

To update one tool at a specific version: `hashi.up update --terraform=0.11.1`

To update multiple, but not all tools at their latest version: `hashi.up update --terraform --packer --vault`

To update multiple, but not all tools at a specific version: `hashi.up update --terraform=0.11.1 --packer=1.5.6 --vault=1.4.0`

### Uninstalling

```Remove one, some, or all of the hashicorp tools

Usage:
  hashi.up remove [flags]

Flags:
      --consul      uninstall consul
  -h, --help        help for remove
      --nomad       uninstall nomad
      --packer      uninstall packer
      --terraform   uninstall terraform
      --vagrant     uninstall vagrant
      --vault       uninstall vault
```

Example:

To remove all tools: `hashi.up remove`

To remove one tool: `hashi.up remove --terraform`

To remove multiple tools, but not all of them: `hashi.up remove --terraform --packer --vault`

## TODO

* Need to make the latest tool gathering more automatic (rather than a hard-coded version number), it's totally possible, I just haven't gotten around to it yet.
* Whatever else y'all find

## Contributing

Feel free to contribute, I'll set up a whole thing for it if this tool gets used.