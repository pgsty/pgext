---
title: "CMD: sty"
description: How to manage pigsty with pig sty subcommand
icon: SquareTerminal
weight: 630
---

**pig** can also be used as a CLI tool for Pigsty, a batteries-included free PostgreSQL RDS solution.
It brings HA, PITR, monitoring, infrastructure as code (IaC), and rich extension support to your PostgreSQL clusters.

```bash
pig sty - Init (Download), Bootstrap, Configure, and Deploy Pigsty

  pig sty init    [-mpfvd]        # install pigsty (~/pigsty by default)
  pig sty boot    [-rmpk]         # install ansible and prepare offline pkg
  pig sty conf    [-cvmrsoxnpg --raw] # configure pigsty and generate config
  pig sty deploy                  # use pigsty to deploy everything (CAUTION!)
  pig sty get                     # download pigsty source tarball
  pig sty list                    # list available pigsty versions

Examples:
  pig sty init                 # extract and init ~/pigsty
  pig sty boot                 # install ansible & other deps
  pig sty conf                 # generate pigsty.yml config file
  pig sty deploy               # run the deploy.yml playbook
```

| Command | Description | Notes |
|:---|:---|:---|
| `sty init` | Install Pigsty | |
| `sty boot` | Install Ansible prerequisites | Requires sudo or root |
| `sty conf` | Generate configuration | |
| `sty deploy` | Run deployment playbook | |
| `sty list` | List available Pigsty versions | |
| `sty get` | Download Pigsty source tarball | |
{.full-width}


## Quick Start

Use `pig sty` to bootstrap and deploy Pigsty on the current node.

```bash
pig sty init                     # install Pigsty to ~/pigsty
pig sty boot                     # install Ansible prerequisites
pig sty conf                     # generate configuration
pig sty deploy                   # run deployment playbook
```

See the detailed setup guide: <https://pigsty.io/docs/setup/install/>


## sty init

Download and install the Pigsty distribution into `~/pigsty`.

```bash
pig sty init                   # install latest version to ~/pigsty
pig sty init -f                # install and overwrite existing pigsty directory
pig sty init -m                # install from the pigsty.cc mirror source
pig sty init -p /tmp/pigsty    # install to selected directory /tmp/pigsty
pig sty init -v 3.4            # fetch and install selected version v3.4.1
pig sty init 3                 # fetch and install latest v3 major version
```

**Options:**

- `-p|--path`: target installation directory, default `~/pigsty`
- `-f|--force`: force overwrite of existing pigsty directory
- `-m|--mirror`: prefer the `pigsty.cc` mirror as the primary source
- `-v|--version`: Pigsty version
- `-d|--dir`: download directory, default `/tmp`


## sty boot

Install Ansible and its dependencies.

```bash
pig sty boot                     # install Ansible
pig sty boot -r china            # use China region mirror
pig sty boot -m                  # same as --region china
pig sty boot -k                  # keep existing repositories
pig sty boot -p /path/to/pkg     # selected offline package path
```

**Options:**

- `-r|--region`: region, such as default, china, europe
- `-m|--mirror`: same as `--region china`
- `-p|--path`: offline package path
- `-k|--keep`: keep existing repositories

See: <https://pigsty.io/docs/setup/offline/#bootstrap>


## sty conf

Configure Pigsty with `./configure` and generate the configuration file.

```bash
pig sty conf                       # use default meta.yml config
pig sty conf -g                    # generate random passwords (recommended!)
pig sty conf -c rich               # use conf/rich.yml template with more extensions
pig sty conf -c ha/full            # use conf/ha/full.yml 4-node HA template
pig sty conf -c slim               # use conf/slim.yml minimal template
pig sty conf -c supabase           # use conf/supabase.yml self-hosting template
pig sty conf -v 18 -c rich         # use conf/rich.yml template with PostgreSQL 18
pig sty conf -r china -s           # use China mirror and skip IP probing
pig sty conf -m -s                 # mirror mode, same as --region china
pig sty conf -x                    # write proxy settings from environment variables
pig sty conf -c full -g -O ha.yml  # full HA template, random passwords, output to ha.yml
pig sty conf --raw                 # use legacy shell configure workflow
```

**Options:**

- `-c|--conf`: configuration template name, such as meta/rich/slim/full/supabase
- `--ip`: primary node IP address
- `-v|--version`: PostgreSQL major version, 18/17/16/15/14; 19 beta can be specified explicitly
- `-r|--region`: upstream repository region, such as default/china/europe
- `-m|--mirror`: same as `--region china`
- `-O|--output-file`: output config file path, default `pigsty.yml`
- `-s|--skip`: skip IP probing
- `-p|--port`: SSH port
- `-x|--proxy`: write proxy settings from environment variables
- `-n|--non-interactive`: non-interactive mode
- `-g|--generate`: generate random default passwords, recommended
- `--raw`: use the legacy shell configure workflow

See: <https://pigsty.io/docs/setup/install/#configure>


## sty deploy

Deploy Pigsty with the `deploy.yml` playbook.

```bash
pig sty deploy       # run deploy.yml, falling back to install.yml if not found
pig sty install      # same as deploy, for backward compatibility
pig sty d            # short alias
pig sty de           # short alias
pig sty ins          # short alias
```

This command runs the `deploy.yml` playbook from your Pigsty installation directory. For backward compatibility, if `deploy.yml` does not exist but `install.yml` exists, `install.yml` is used instead.

> **Warning**: This operation modifies your system. Use it carefully.


## sty list

List available Pigsty versions.

```bash
pig sty list                     # list available versions
pig sty list -m                  # list versions from the pigsty.cc mirror
```


## sty get

Download the Pigsty source tarball.

```bash
pig sty get                      # download latest version
pig sty get v3.4.0               # download selected version
pig sty get -m                   # download from the pigsty.cc mirror
```
