---
title: "CMD: sty"
description: How to manage pigsty with pig sty subcommand
icon: SquareTerminal
weight: 630
---


The **pig** can also be used as a cli tool for Pigsty — the battery-include free PostgreSQL RDS.
Which brings HA, PITR, Monitoring, IaC, and all the extensions to your PostgreSQL cluster.

```bash
pig sty -Init (Download), Bootstrap, Configure, and Deploy Pigsty

  pig sty init    [-pfvd]         # install pigsty (~/pigsty by default)
  pig sty boot    [-rpk]          # install ansible and prepare offline pkg
  pig sty conf    [-cvrsoxnpg]   # configure pigsty and generate config
  pig sty deploy                  # use pigsty to deploy everything (CAUTION!)
  pig sty get                     # download pigsty source tarball
  pig sty list                    # list available pigsty versions

Usage:
  pig sty [command]

Aliases:
  sty, s, pigsty

Examples:
  Get Started: https://pigsty.io/docs/setup/install/
  pig sty init                 # extract and init ~/pigsty
  pig sty boot                 # install ansible & other deps
  pig sty conf                 # generate pigsty.yml config file
  pig sty deploy               # run the deploy.yml playbook

Available Commands:
  boot        Bootstrap Pigsty
  conf        Configure Pigsty
  deploy      run pigsty deploy.yml playbook
  get         download pigsty available versions
  init        Install Pigsty
  list        list pigsty available versions

Flags:
  -h, --help   help for sty
```

You can use the `pig sty` subcommand to bootstrap pigsty on current node.



------

## `sty init`

```bash
pig sty init
  -p | --path    : where to install, ~/pigsty by default
  -f | --force   : force overwrite existing pigsty dir
  -v | --version : pigsty version, latest by default
  -d | --dir     : download directory, /tmp by default

Usage:
  pig sty init [flags]

Aliases:
  init, i

Examples:

  pig sty init                   # install to ~/pigsty with the latest version
  pig sty init -f                # install and OVERWRITE existing pigsty dir
  pig sty init -p /tmp/pigsty    # install to another location /tmp/pigsty
  pig sty init -v 3.4            # get & install specific version v3.4.1
  pig sty init 3                 # get & install specific version v3 latest


Flags:
  -d, --dir string       pigsty download directory (default "/tmp")
  -f, --force            overwrite existing pigsty (false by default)
  -h, --help             help for init
  -p, --path string      target directory (default "~/pigsty")
  -v, --version string   pigsty version string
```

------

## `sty boot`

```bash
pig sty boot
  [-r|--region <region]   [default,china,europe]
  [-p|--path <path>]      specify another offline pkg path
  [-k|--keep]             keep existing upstream repo during bootstrap

Check https://pigsty.io/docs/setup/offline/#bootstrap for details

Usage:
  pig sty boot [flags]

Aliases:
  boot, b, bootstrap

Flags:
  -h, --help            help for boot
  -k, --keep            keep existing repo
  -p, --path string     offline package path
  -r, --region string   default,china,europe,...
```

------

## `sty conf`

```bash
Configure pigsty with ./configure

pig sty conf
  [-c|--conf <name>]      # config template: [meta|rich|slim|full|supabase|...]
  [--ip <ip>]             # primary IP address (skip with -s)
  [-v|--version <pgver>]  # postgres major version: [18|17|16|15|14|13]
  [-r|--region <region>]  # upstream repo region: [default|china|europe]
  [-o|--output <file>]    # output config file path (default: pigsty.yml)
  [-s|--skip]             # skip IP address probing
  [-p|--port <port>]      # specify SSH port (for ssh accessibility check)
  [-x|--proxy]            # write proxy env from environment
  [-n|--non-interactive]  # non-interactive mode

  [-g|--generate]         # generate random default passwords (RECOMMENDED!)

Check https://pigsty.io/docs/setup/install/#configure for details

Usage:
  pig sty conf [flags]

Aliases:
  conf, c, configure

Examples:

  pig sty conf                       # use default meta.yml config
  pig sty conf -g                    # generate with random passwords (RECOMMENDED!)
  pig sty conf -c rich               # use conf/rich.yml template (with more extensions)
  pig sty conf -c ha/full            # use conf/ha/full.yml 4-node ha template
  pig sty conf -c slim               # use conf/slim.yml template (minimal install)
  pig sty conf -c supabase           # use conf/supabase.yml template (self-hosting)
  pig sty conf -v 17 -c rich         # use conf/rich.yml template with PostgreSQL 17
  pig sty conf -r china -s           # use china region mirrors, skip IP probe
  pig sty conf -x                    # write proxy env from environment to config
  pig sty conf -c full -g -o ha.yml  # full HA template with random passwords to ha.yml


Flags:
  -c, --conf string       config template name
  -g, --generate          generate random passwords
  -h, --help              help for conf
      --ip string         primary ip address
  -n, --non-interactive   non-interactive mode
  -o, --output string     output config file path
  -p, --port string       SSH port (only used if set)
  -x, --proxy             write proxy env from environment
  -r, --region string     upstream repo region
  -s, --skip              skip ip probe
  -v, --version string    postgres major version
```

------

## `sty deploy`

```bash
Deploy Pigsty using the deploy.yml playbook

This command runs the deploy.yml playbook from your Pigsty installation.
For backward compatibility, if deploy.yml doesn't exist but install.yml does,
it will use install.yml instead.

Aliases: deploy, d, de, install, ins (for backward compatibility)

WARNING: This operation makes changes to your system. Use with caution!

Usage:
  pig sty deploy [flags]

Aliases:
  deploy, d, de, install, ins

Examples:
  pig sty deploy       # run deploy.yml (or install.yml if deploy.yml not found)
  pig sty install      # same as deploy (backward compatible)
  pig sty d            # short alias
  pig sty de           # short alias
  pig sty ins          # short alias

Flags:
  -h, --help   help for deploy
```
