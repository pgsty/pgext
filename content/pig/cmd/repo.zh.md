---
title: "命令：repo"
description: 如何使用 pig repo 子命令管理软件仓库？
icon: SquareTerminal
weight: 610
---

`pig repo` 命令是一个综合性的软件包仓库管理工具。它提供了添加、移除、创建和管理软件仓库的功能，支持 RPM 系统（RHEL/CentOS/Rocky/Alma）和 Debian 系统（Debian/Ubuntu）。该命令是安装 PostgreSQL 及其扩展所需仓库设置的关键工具。

```bash
pig repo - 管理 Linux 软件仓库（apt/dnf）

Usage: pig repo <command>

Commands:
  add       添加新仓库
  set       清空、覆盖并更新仓库
  rm        移除仓库
  list      打印可用仓库与模块列表
  info      获取仓库详细信息
  status    显示当前仓库状态
  update    更新仓库缓存
  create    创建本地 YUM/APT 仓库
  cache     从本地仓库创建离线包
  boot      从离线包引导仓库
  reload    刷新仓库元数据

Flags:
  -h, --help   获取帮助信息

Global Flags:
      --debug              启用调试模式
  -H, --home string        pigsty 主目录路径
  -i, --inventory string   配置清单路径
      --log-level string   日志级别: debug, info, warn, error, fatal, panic (默认 "info")
      --log-path string    日志文件路径，默认为终端输出

使用 "pig repo [command] --help" 获取命令的详细信息。
```

| 命令                            | 描述              | 备注                |
|-------------------------------|-----------------|-------------------|
| [`repo list`](#repo-list)     | 打印可用仓库与模块列表     |                   |
| [`repo info`](#repo-info)     | 获取仓库详细信息        |                   |
| [`repo status`](#repo-status) | 显示当前仓库状态        |                   |
| [`repo add`](#repo-add)       | 添加新仓库           | 需要 sudo 或 root 权限 |
| [`repo set`](#repo-set)       | 清空、覆盖并更新仓库      | 需要 sudo 或 root 权限 |
| [`repo rm`](#repo-rm)         | 移除仓库            | 需要 sudo 或 root 权限 |
| [`repo update`](#repo-update) | 更新仓库缓存          | 需要 sudo 或 root 权限 |
| [`repo create`](#repo-create) | 创建本地 YUM/APT 仓库 | 需要 sudo 或 root 权限 |
| [`repo cache`](#repo-cache)   | 从本地仓库创建离线包      | 需要 sudo 或 root 权限 |
| [`repo boot`](#repo-boot)     | 从离线包引导仓库        | 需要 sudo 或 root 权限 |



## 快速入门

```bash
# 方法 1：清理干净现有仓库，添加所有必要仓库并更新缓存（推荐）
pig repo add all --remove --update    # 移除旧仓库，添加所有必要仓库，更新缓存

# 方法 1 变体：一步到位
pig repo set                          # = pig repo add all --remove --update 

# 方法 2：温和方式 - 仅添加所需仓库，保留你目前的仓库配置
pig repo add pgsql                    # 添加 PGDG 和 Pigsty 仓库并更新缓存
pig repo add pigsty --region=china    # 添加 Pigsty 仓库，指定使用中国区域
pig repo add pgdg   --region=default  # 添加 PGDG ，指定使用默认区域
pig repo add infra  --region=europe   # 添加 INFRA 仓库 ，指定使用欧洲区域

# 如果上面没有-u|--update 选项一步到位，请额外执行此命令
pig repo update                       # 更新系统包缓存 
```



## 全局概览

```bash
pig repo - 管理 Linux APT/YUM 仓库

  pig repo list                    # 可用仓库列表            (info)
  pig repo info   [repo|module...] # 显示仓库信息            (info)
  pig repo status                  # 显示当前仓库状态        (info)
  pig repo add    [repo|module...] # 添加仓库和模块          (root)
  pig repo rm     [repo|module...] # 移除仓库和模块          (root)
  pig repo update                  # 更新仓库包缓存          (root)
  pig repo create                  # 在当前系统创建仓库      (root)
  pig repo boot                    # 从离线包引导仓库        (root)
  pig repo cache                   # 将仓库缓存为离线包      (root)

Usage:
  pig repo [command]

别名:
  repo, r

可用命令:
  add         添加新仓库
  boot        从离线包引导仓库
  cache       从本地仓库创建离线包
  create      创建本地 YUM/APT 仓库
  info        获取仓库详细信息
  list        打印可用仓库列表
  rm          移除仓库
  set         清空、覆盖并更新仓库
  status      显示当前仓库状态
  update      更新仓库缓存

标志:
  -h, --help   获取帮助信息

全局标志:
      --debug              启用调试模式
  -H, --home string        pigsty 主目录路径
  -i, --inventory string   配置清单路径
      --log-level string   日志级别: debug, info, warn, error, fatal, panic (默认 "info")
      --log-path string    日志文件路径，默认为终端输出
```




## 模块

在 pig 中，APT/YUM 仓库被组织为 **模块** —— 服务于特定目的的一组仓库。

|    模块     | 说明                      | 仓库列表                                                                                                              |
|:---------:|-------------------------|:------------------------------------------------------------------------------------------------------------------|
|   `all`   | 安装 PG 所需的全部核心模块         | `node` + `infra` + `pgsql`                                                                                        |
|  `pgsql`  |                         | `pigsty-pgsql` + `pgdg`                                                                                           |
| `pigsty`  | Pigsty Infra + PGSQL 仓库 | pigsty-infra, pigsty-pgsql                                                                                        |
|  `pgdg`   | PGDG 官方仓库               | pigsty-pgsql, pgdg-common, pgdg-el8fix, pgdg-el9fix, pgdg-el10fix, pgdg13, pgdg14, pgdg15, pgdg16, pgdg17, pgdg   |
|  `node`   | Linux 系统仓库              | base, updates, extras, epel, centos-sclo, centos-sclo-rh, baseos, appstream, powertools, crb, security, backports |
|  `infra`  | 基础设施组件仓库                | pigsty-infra, nginx, docker-ce                                                                                    |
|  `beta`   | PostgreSQL Beta 版本      | pgdg19-beta, pgdg-beta                                                                                            |
|  `extra`  | PGDG Non-Free 与三方扩展     | pgdg-extras, pgdg13-nonfree, pgdg14-nonfree, pgdg15-nonfree, pgdg16-nonfree, pgdg17-nonfree, timescaledb, citus   |
| `groonga` | PGroonga 仓库             | groonga                                                                                                           |
|  `mssql`  | WiltonDB 仓库             | wiltondb                                                                                                          |
| `percona` | Percona PG + PG_TDE     | percona                                                                                                           |

除此之外，pig 还自带了一些其他数据库的 APT/DNF 仓库：`redis`, `kubernetes`, `grafana`, `clickhouse`, `gitlab`, `haproxy`, `mongodb`, `mysql`，在此不再展开。

通常来说，为了安装 PostgreSQL `node` （Linux 系统仓库） 和 `pgsql`（PGDG + Pigsty）是必选项，`infra` 仓库是可选项（包含了一些工具，IvorySQL Kernel 等）。
您可以使用特殊的 `all` 模块，一次性添加所有需要的仓库到系统中，对绝大多数用户来说，这是合适的起点。

```bash
pig repo add all      # 添加 node,pgsql,infra 三个仓库到系统中
pig repo add          # 不添加任何参数时，默认使用 all 模块
pig repo set          # 使用 set 替代 add 时，将清理备份现有仓库定义并覆盖式更新
```

## 仓库

Pigsty 中可用仓库的完整定义位于 [`cli/repo/assets/repo.yml`](https://github.com/pgsty/pig/blob/main/cli/repo/assets/repo.yml)。

```yaml
# EL 7/8/9 REPOS
- { name: pigsty-local   ,description: 'Pigsty Local'       ,module: local   ,releases: [7,8,9,10]       ,arch: [x86_64, aarch64] ,baseurl: { default: 'file:///www/pigsty'  }, meta: {skip_if_unavailable: 1}} # used by intranet nodes
- { name: pigsty-infra   ,description: 'Pigsty INFRA'       ,module: infra   ,releases: [7,8,9,10]       ,arch: [x86_64, aarch64] ,baseurl: { default: 'https://repo.pigsty.io/yum/infra/$basearch' ,china: 'https://repo.pigsty.cc/yum/infra/$basearch' }}
- { name: pigsty-pgsql   ,description: 'Pigsty PGSQL'       ,module: pgsql   ,releases: [7,8,9,10]       ,arch: [x86_64, aarch64] ,baseurl: { default: 'https://repo.pigsty.io/yum/pgsql/el$releasever.$basearch' ,china: 'https://repo.pigsty.cc/yum/pgsql/el$releasever.$basearch' }}
- { name: nginx          ,description: 'Nginx Repo'         ,module: infra   ,releases: [7,8,9,10]       ,arch: [x86_64, aarch64] ,baseurl: { default: 'https://nginx.org/packages/rhel/$releasever/$basearch/' }}
- { name: docker-ce      ,description: 'Docker CE'          ,module: infra   ,releases: [7,8,9,10]       ,arch: [x86_64, aarch64] ,baseurl: { default: 'https://download.docker.com/linux/centos/$releasever/$basearch/stable'    ,china: 'https://mirrors.aliyun.com/docker-ce/linux/centos/$releasever/$basearch/stable'  ,europe: 'https://mirrors.xtom.de/docker-ce/linux/centos/$releasever/$basearch/stable' }}
- { name: base           ,description: 'EL 7 Base'          ,module: node    ,releases: [7       ]       ,arch: [x86_64         ] ,baseurl: { default: 'http://vault.centos.org/centos/$releasever/os/$basearch/'                 ,china: 'https://mirrors.aliyun.com/centos/$releasever/os/$basearch/'        ,europe: 'https://mirrors.xtom.de/centos/$releasever/os/$basearch/'           }}
- { name: updates        ,description: 'EL 7 Updates'       ,module: node    ,releases: [7       ]       ,arch: [x86_64         ] ,baseurl: { default: 'http://vault.centos.org/centos/$releasever/updates/$basearch/'            ,china: 'https://mirrors.aliyun.com/centos/$releasever/updates/$basearch/'   ,europe: 'https://mirrors.xtom.de/centos/$releasever/updates/$basearch/'      }}
- { name: extras         ,description: 'EL 7 Extras'        ,module: node    ,releases: [7       ]       ,arch: [x86_64         ] ,baseurl: { default: 'http://vault.centos.org/centos/$releasever/extras/$basearch/'             ,china: 'https://mirrors.aliyun.com/centos/$releasever/extras/$basearch/'    ,europe: 'https://mirrors.xtom.de/centos/$releasever/extras/$basearch/'       }}
- { name: epel           ,description: 'EL 7 EPEL'          ,module: node    ,releases: [7       ]       ,arch: [x86_64         ] ,baseurl: { default: 'http://archives.fedoraproject.org/pub/archive/epel/$releasever/$basearch/',china: 'https://mirrors.aliyun.com/epel/$releasever/$basearch/'             ,europe: 'https://mirrors.xtom.de/epel/$releasever/$basearch/'                }}
- { name: centos-sclo    ,description: 'EL 7 SCLo'          ,module: node    ,releases: [7       ]       ,arch: [x86_64         ] ,baseurl: { default: 'http://vault.centos.org/centos/$releasever/sclo/$basearch/sclo/'          ,china: 'https://mirrors.aliyun.com/centos/$releasever/sclo/$basearch/sclo/' ,europe: 'https://mirrors.xtom.de/centos/$releasever/sclo/$basearch/sclo/'    }}
- { name: centos-sclo-rh ,description: 'EL 7 SCLo rh'       ,module: node    ,releases: [7       ]       ,arch: [x86_64         ] ,baseurl: { default: 'http://vault.centos.org/centos/$releasever/sclo/$basearch/rh/'            ,china: 'https://mirrors.aliyun.com/centos/$releasever/sclo/$basearch/rh/'   ,europe: 'https://mirrors.xtom.de/centos/$releasever/sclo/$basearch/rh/'      }}
- { name: baseos         ,description: 'EL 8+ BaseOS'       ,module: node    ,releases: [  8,9,10]       ,arch: [x86_64, aarch64] ,baseurl: { default: 'https://dl.rockylinux.org/pub/rocky/$releasever/BaseOS/$basearch/os/'     ,china: 'https://mirrors.aliyun.com/rockylinux/$releasever/BaseOS/$basearch/os/'     ,europe: 'https://mirrors.xtom.de/rocky/$releasever/BaseOS/$basearch/os/'     }}
- { name: appstream      ,description: 'EL 8+ AppStream'    ,module: node    ,releases: [  8,9,10]       ,arch: [x86_64, aarch64] ,baseurl: { default: 'https://dl.rockylinux.org/pub/rocky/$releasever/AppStream/$basearch/os/'  ,china: 'https://mirrors.aliyun.com/rockylinux/$releasever/AppStream/$basearch/os/'  ,europe: 'https://mirrors.xtom.de/rocky/$releasever/AppStream/$basearch/os/'  }}
- { name: extras         ,description: 'EL 8+ Extras'       ,module: node    ,releases: [  8,9,10]       ,arch: [x86_64, aarch64] ,baseurl: { default: 'https://dl.rockylinux.org/pub/rocky/$releasever/extras/$basearch/os/'     ,china: 'https://mirrors.aliyun.com/rockylinux/$releasever/extras/$basearch/os/'     ,europe: 'https://mirrors.xtom.de/rocky/$releasever/extras/$basearch/os/'     }}
- { name: powertools     ,description: 'EL 8 PowerTools'    ,module: node    ,releases: [  8     ]       ,arch: [x86_64, aarch64] ,baseurl: { default: 'https://dl.rockylinux.org/pub/rocky/$releasever/PowerTools/$basearch/os/' ,china: 'https://mirrors.aliyun.com/rockylinux/$releasever/PowerTools/$basearch/os/' ,europe: 'https://mirrors.xtom.de/rocky/$releasever/PowerTools/$basearch/os/' }}
- { name: crb            ,description: 'EL 9 CRB'           ,module: node    ,releases: [    9,10]       ,arch: [x86_64, aarch64] ,baseurl: { default: 'https://dl.rockylinux.org/pub/rocky/$releasever/CRB/$basearch/os/'        ,china: 'https://mirrors.aliyun.com/rockylinux/$releasever/CRB/$basearch/os/'        ,europe: 'https://mirrors.xtom.de/rocky/$releasever/CRB/$basearch/os/'        }}
- { name: epel           ,description: 'EL 8/9 EPEL'        ,module: node    ,releases: [  8,9   ]       ,arch: [x86_64, aarch64] ,baseurl: { default: 'https://mirrors.edge.kernel.org/fedora-epel/$releasever/Everything/$basearch/' ,china: 'https://mirrors.aliyun.com/epel/$releasever/Everything/$basearch/'     ,europe: 'https://mirrors.xtom.de/epel/$releasever/Everything/$basearch/'     }}
- { name: epel           ,description: 'EL 10 EPEL'         ,module: node    ,releases: [      10]       ,arch: [x86_64, aarch64] ,baseurl: { default: 'https://mirrors.edge.kernel.org/fedora-epel/$releasever.0/Everything/$basearch/' ,china: 'https://mirrors.aliyun.com/epel/$releasever.0/Everything/$basearch/' ,europe: 'https://mirrors.xtom.de/epel/$releasever.0/Everything/$basearch/' } }
- { name: pgdg-common    ,description: 'PostgreSQL Common'  ,module: pgsql   ,releases: [7,8,9,10]       ,arch: [x86_64, aarch64] ,baseurl: { default: 'https://download.postgresql.org/pub/repos/yum/common/redhat/rhel-$releasever-$basearch' ,china: 'https://repo.pigsty.cc/yum/pgdg/common/redhat/rhel-$releasever-$basearch' , europe: 'https://mirrors.xtom.de/postgresql/repos/yum/common/redhat/rhel-$releasever-$basearch' }}
- { name: pgdg-el8fix    ,description: 'PostgreSQL EL8FIX'  ,module: pgsql   ,releases: [  8     ]       ,arch: [x86_64, aarch64] ,baseurl: { default: 'https://download.postgresql.org/pub/repos/yum/common/pgdg-centos8-sysupdates/redhat/rhel-8-$basearch/' ,china: 'https://repo.pigsty.cc/yum/pgdg/common/pgdg-centos8-sysupdates/redhat/rhel-8-x86_64/' , europe: 'https://mirrors.xtom.de/postgresql/repos/yum/common/pgdg-centos8-sysupdates/redhat/rhel-8-x86_64/' } }
- { name: pgdg-el9fix    ,description: 'PostgreSQL EL9FIX'  ,module: pgsql   ,releases: [    9   ]       ,arch: [x86_64, aarch64] ,baseurl: { default: 'https://download.postgresql.org/pub/repos/yum/common/pgdg-rocky9-sysupdates/redhat/rhel-9-$basearch/'  ,china: 'https://repo.pigsty.cc/yum/pgdg/common/pgdg-rocky9-sysupdates/redhat/rhel-9-x86_64/' , europe: 'https://mirrors.xtom.de/postgresql/repos/yum/common/pgdg-rocky9-sysupdates/redhat/rhel-9-x86_64/' }}
- { name: pgdg13         ,description: 'PostgreSQL 13'      ,module: pgsql   ,releases: [7,8,9,10]       ,arch: [x86_64, aarch64] ,baseurl: { default: 'https://download.postgresql.org/pub/repos/yum/13/redhat/rhel-$releasever-$basearch' ,china: 'https://repo.pigsty.cc/yum/pgdg/13/redhat/rhel-$releasever-$basearch' ,europe: 'https://mirrors.xtom.de/postgresql/repos/yum/13/redhat/rhel-$releasever-$basearch' }}
- { name: pgdg14         ,description: 'PostgreSQL 14'      ,module: pgsql   ,releases: [7,8,9,10]       ,arch: [x86_64, aarch64] ,baseurl: { default: 'https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-$releasever-$basearch' ,china: 'https://repo.pigsty.cc/yum/pgdg/14/redhat/rhel-$releasever-$basearch' ,europe: 'https://mirrors.xtom.de/postgresql/repos/yum/14/redhat/rhel-$releasever-$basearch' }}
- { name: pgdg15         ,description: 'PostgreSQL 15'      ,module: pgsql   ,releases: [7,8,9,10]       ,arch: [x86_64, aarch64] ,baseurl: { default: 'https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-$releasever-$basearch' ,china: 'https://repo.pigsty.cc/yum/pgdg/15/redhat/rhel-$releasever-$basearch' ,europe: 'https://mirrors.xtom.de/postgresql/repos/yum/15/redhat/rhel-$releasever-$basearch' }}
- { name: pgdg16         ,description: 'PostgreSQL 16'      ,module: pgsql   ,releases: [  8,9,10]       ,arch: [x86_64, aarch64] ,baseurl: { default: 'https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-$releasever-$basearch' ,china: 'https://repo.pigsty.cc/yum/pgdg/16/redhat/rhel-$releasever-$basearch' ,europe: 'https://mirrors.xtom.de/postgresql/repos/yum/16/redhat/rhel-$releasever-$basearch' }}
- { name: pgdg17         ,description: 'PostgreSQL 17'      ,module: pgsql   ,releases: [  8,9,10]       ,arch: [x86_64, aarch64] ,baseurl: { default: 'https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-$releasever-$basearch' ,china: 'https://repo.pigsty.cc/yum/pgdg/17/redhat/rhel-$releasever-$basearch' ,europe: 'https://mirrors.xtom.de/postgresql/repos/yum/17/redhat/rhel-$releasever-$basearch' }}
- { name: pgdg18         ,description: 'PostgreSQL 18'      ,module: pgsql   ,releases: [  8,9,10]       ,arch: [x86_64, aarch64] ,baseurl: { default: 'https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-$releasever-$basearch' ,china: 'https://repo.pigsty.cc/yum/pgdg/18/redhat/rhel-$releasever-$basearch' ,europe: 'https://mirrors.xtom.de/postgresql/repos/yum/18/redhat/rhel-$releasever-$basearch' }}
- { name: pgdg19-beta    ,description: 'PostgreSQL 19 Beta' ,module: beta    ,releases: [  8,9,10]       ,arch: [x86_64, aarch64] ,baseurl: { default: 'https://download.postgresql.org/pub/repos/yum/testing/19/redhat/rhel-$releasever-$basearch' ,china: 'https://repo.pigsty.cc/yum/pgdg/testing/19/redhat/rhel-$releasever-$basearch' ,europe: 'https://mirrors.xtom.de/postgresql/repos/yum/testing/19/redhat/rhel-$releasever-$basearch' }}
- { name: pgdg-extras    ,description: 'PostgreSQL Extra'   ,module: extra   ,releases: [7,8,9,10]       ,arch: [x86_64, aarch64] ,baseurl: { default: 'https://download.postgresql.org/pub/repos/yum/common/pgdg-rhel$releasever-extras/redhat/rhel-$releasever-$basearch' ,china: 'https://repo.pigsty.cc/yum/pgdg/common/pgdg-rhel$releasever-extras/redhat/rhel-$releasever-$basearch' , europe: 'https://mirrors.xtom.de/postgresql/repos/yum/common/pgdg-rhel$releasever-extras/redhat/rhel-$releasever-$basearch' }}
- { name: pgdg13-nonfree ,description: 'PostgreSQL 13+'     ,module: extra   ,releases: [7,8,9,10]       ,arch: [x86_64         ] ,baseurl: { default: 'https://download.postgresql.org/pub/repos/yum/non-free/13/redhat/rhel-$releasever-$basearch' ,china: 'https://repo.pigsty.cc/yum/pgdg/non-free/13/redhat/rhel-$releasever-$basearch' ,europe: 'https://mirrors.xtom.de/postgresql/repos/yum/non-free/13/redhat/rhel-$releasever-$basearch' }}
- { name: pgdg14-nonfree ,description: 'PostgreSQL 14+'     ,module: extra   ,releases: [7,8,9,10]       ,arch: [x86_64         ] ,baseurl: { default: 'https://download.postgresql.org/pub/repos/yum/non-free/14/redhat/rhel-$releasever-$basearch' ,china: 'https://repo.pigsty.cc/yum/pgdg/non-free/14/redhat/rhel-$releasever-$basearch' ,europe: 'https://mirrors.xtom.de/postgresql/repos/yum/non-free/14/redhat/rhel-$releasever-$basearch' }}
- { name: pgdg15-nonfree ,description: 'PostgreSQL 15+'     ,module: extra   ,releases: [7,8,9,10]       ,arch: [x86_64         ] ,baseurl: { default: 'https://download.postgresql.org/pub/repos/yum/non-free/15/redhat/rhel-$releasever-$basearch' ,china: 'https://repo.pigsty.cc/yum/pgdg/non-free/15/redhat/rhel-$releasever-$basearch' ,europe: 'https://mirrors.xtom.de/postgresql/repos/yum/non-free/15/redhat/rhel-$releasever-$basearch' }}
- { name: pgdg16-nonfree ,description: 'PostgreSQL 16+'     ,module: extra   ,releases: [  8,9,10]       ,arch: [x86_64         ] ,baseurl: { default: 'https://download.postgresql.org/pub/repos/yum/non-free/16/redhat/rhel-$releasever-$basearch' ,china: 'https://repo.pigsty.cc/yum/pgdg/non-free/16/redhat/rhel-$releasever-$basearch' ,europe: 'https://mirrors.xtom.de/postgresql/repos/yum/non-free/16/redhat/rhel-$releasever-$basearch' }}
- { name: pgdg17-nonfree ,description: 'PostgreSQL 17+'     ,module: extra   ,releases: [  8,9,10]       ,arch: [x86_64         ] ,baseurl: { default: 'https://download.postgresql.org/pub/repos/yum/non-free/17/redhat/rhel-$releasever-$basearch' ,china: 'https://repo.pigsty.cc/yum/pgdg/non-free/17/redhat/rhel-$releasever-$basearch' ,europe: 'https://mirrors.xtom.de/postgresql/repos/yum/non-free/17/redhat/rhel-$releasever-$basearch' }}
- { name: timescaledb    ,description: 'TimescaleDB'        ,module: extra   ,releases: [7,8,9   ]       ,arch: [x86_64, aarch64] ,baseurl: { default: 'https://packagecloud.io/timescale/timescaledb/el/$releasever/$basearch'  }}
- { name: percona        ,description: 'Percona TDE'        ,module: percona ,releases: [  8,9   ]       ,arch: [x86_64, aarch64] ,baseurl: { default: 'https://repo.pigsty.io/yum/percona/el$releasever.$basearch' ,china: 'https://repo.pigsty.cc/yum/percona/el$releasever.$basearch' ,origin: 'http://repo.percona.com/ppg-17.5/yum/release/$releasever/RPMS/$basearch'  }}
- { name: wiltondb       ,description: 'WiltonDB'           ,module: mssql   ,releases: [7,8,9   ]       ,arch: [x86_64, aarch64] ,baseurl: { default: 'https://repo.pigsty.io/yum/mssql/el$releasever.$basearch', china: 'https://repo.pigsty.cc/yum/mssql/el$releasever.$basearch' , origin: 'https://download.copr.fedorainfracloud.org/results/wiltondb/wiltondb/epel-$releasever-$basearch/' }}
- { name: groonga        ,description: 'Groonga'            ,module: groonga ,releases: [  8,9,10]       ,arch: [x86_64, aarch64] ,baseurl: { default: 'https://packages.groonga.org/almalinux/$releasever/$basearch/' }}
- { name: mysql          ,description: 'MySQL'              ,module: mysql   ,releases: [7,8,9   ]       ,arch: [x86_64, aarch64] ,baseurl: { default: 'https://repo.mysql.com/yum/mysql-8.0-community/el/$releasever/$basearch/', china: 'https://mirrors.tuna.tsinghua.edu.cn/mysql/yum/mysql-8.0-community-el7-$basearch/'}}
- { name: mongo          ,description: 'MongoDB'            ,module: mongo   ,releases: [7,8,9   ]       ,arch: [x86_64, aarch64] ,baseurl: { default: 'https://repo.mongodb.org/yum/redhat/$releasever/mongodb-org/8.0/$basearch/' ,china: 'https://mirrors.aliyun.com/mongodb/yum/redhat/$releasever/mongodb-org/8.0/$basearch/' }}
- { name: redis          ,description: 'Redis'              ,module: redis   ,releases: [7       ]       ,arch: [x86_64, aarch64] ,baseurl: { default: 'https://rpmfind.net/linux/remi/enterprise/$releasever/remi/$basearch/' }}
- { name: redis          ,description: 'Redis'              ,module: redis   ,releases: [  8,9,10]       ,arch: [x86_64, aarch64] ,baseurl: { default: 'https://rpmfind.net/linux/remi/enterprise/$releasever/redis72/$basearch/' }}
- { name: grafana        ,description: 'Grafana'            ,module: grafana ,releases: [7,8,9,10]       ,arch: [x86_64, aarch64] ,baseurl: { default: 'https://rpm.grafana.com', china: 'https://mirrors.aliyun.com/grafana/yum/' }}
- { name: kubernetes     ,description: 'Kubernetes'         ,module: kube    ,releases: [7,8,9,10]       ,arch: [x86_64, aarch64] ,baseurl: { default: 'https://pkgs.k8s.io/core:/stable:/v1.33/rpm/', china: 'https://mirrors.aliyun.com/kubernetes-new/core/stable/v1.33/rpm/' }}
- { name: gitlab-ee      ,description: 'Gitlab EE'          ,module: gitlab  ,releases: [  8,9   ]       ,arch: [x86_64, aarch64] ,baseurl: { default: 'https://packages.gitlab.com/gitlab/gitlab-ee/el/$releasever/$basearch' }}
- { name: gitlab-ce      ,description: 'Gitlab CE'          ,module: gitlab  ,releases: [  8,9   ]       ,arch: [x86_64, aarch64] ,baseurl: { default: 'https://packages.gitlab.com/gitlab/gitlab-ce/el/$releasever/$basearch' }}
- { name: clickhouse     ,description: 'ClickHouse'         ,module: click   ,releases: [  8,9,10]       ,arch: [x86_64, aarch64] ,baseurl: { default: 'https://packages.clickhouse.com/rpm/stable/', china: 'https://mirrors.aliyun.com/clickhouse/rpm/stable/' }}
# DEB 11/12/13 Ubuntu 20/22/24 REPOS
- { name: pigsty-local   ,description: 'Pigsty Local'       ,module: local   ,releases: [11,12,13,20,22,24] ,arch: [x86_64, aarch64] ,baseurl: { default: 'http://${admin_ip}/pigsty ./' }}
- { name: pigsty-pgsql   ,description: 'Pigsty PgSQL'       ,module: pgsql   ,releases: [11,12,13,20,22,24] ,arch: [x86_64, aarch64] ,baseurl: { default: 'https://repo.pigsty.io/apt/pgsql/${distro_codename} ${distro_codename} main', china: 'https://repo.pigsty.cc/apt/pgsql/${distro_codename} ${distro_codename} main' }}
- { name: pigsty-infra   ,description: 'Pigsty Infra'       ,module: infra   ,releases: [11,12,13,20,22,24] ,arch: [x86_64, aarch64] ,baseurl: { default: 'https://repo.pigsty.io/apt/infra/ generic main' ,china: 'https://repo.pigsty.cc/apt/infra/ generic main' }}
- { name: nginx          ,description: 'Nginx'              ,module: infra   ,releases: [11,12,   20,22,24] ,arch: [x86_64, aarch64] ,baseurl: { default: 'http://nginx.org/packages/${distro_name} ${distro_codename} nginx' }}
- { name: docker-ce      ,description: 'Docker'             ,module: infra   ,releases: [11,12,13,20,22,24] ,arch: [x86_64, aarch64] ,baseurl: { default: 'https://download.docker.com/linux/${distro_name} ${distro_codename} stable'                               ,china: 'https://mirrors.aliyun.com/docker-ce/linux/${distro_name} ${distro_codename} stable' }}
- { name: base           ,description: 'Debian Basic'       ,module: node    ,releases: [11,12,13         ] ,arch: [x86_64, aarch64] ,baseurl: { default: 'http://deb.debian.org/debian/ ${distro_codename} main non-free-firmware'                                  ,china: 'https://mirrors.aliyun.com/debian/ ${distro_codename} main restricted universe multiverse' }}
- { name: updates        ,description: 'Debian Updates'     ,module: node    ,releases: [11,12,13         ] ,arch: [x86_64, aarch64] ,baseurl: { default: 'http://deb.debian.org/debian/ ${distro_codename}-updates main non-free-firmware'                          ,china: 'https://mirrors.aliyun.com/debian/ ${distro_codename}-updates main restricted universe multiverse' }}
- { name: security       ,description: 'Debian Security'    ,module: node    ,releases: [11,12,13         ] ,arch: [x86_64, aarch64] ,baseurl: { default: 'http://security.debian.org/debian-security ${distro_codename}-security main non-free-firmware'            ,china: 'https://mirrors.aliyun.com/debian-security/ ${distro_codename}-security main non-free-firmware' }}
- { name: base           ,description: 'Ubuntu Basic'       ,module: node    ,releases: [         20,22,24] ,arch: [x86_64         ] ,baseurl: { default: 'https://mirrors.edge.kernel.org/ubuntu/ ${distro_codename}           main universe multiverse restricted' ,china: 'https://mirrors.aliyun.com/ubuntu/ ${distro_codename}           main restricted universe multiverse' }}
- { name: updates        ,description: 'Ubuntu Updates'     ,module: node    ,releases: [         20,22,24] ,arch: [x86_64         ] ,baseurl: { default: 'https://mirrors.edge.kernel.org/ubuntu/ ${distro_codename}-backports main restricted universe multiverse' ,china: 'https://mirrors.aliyun.com/ubuntu/ ${distro_codename}-updates   main restricted universe multiverse' }}
- { name: backports      ,description: 'Ubuntu Backports'   ,module: node    ,releases: [         20,22,24] ,arch: [x86_64         ] ,baseurl: { default: 'https://mirrors.edge.kernel.org/ubuntu/ ${distro_codename}-security  main restricted universe multiverse' ,china: 'https://mirrors.aliyun.com/ubuntu/ ${distro_codename}-backports main restricted universe multiverse' }}
- { name: security       ,description: 'Ubuntu Security'    ,module: node    ,releases: [         20,22,24] ,arch: [x86_64         ] ,baseurl: { default: 'https://mirrors.edge.kernel.org/ubuntu/ ${distro_codename}-updates   main restricted universe multiverse' ,china: 'https://mirrors.aliyun.com/ubuntu/ ${distro_codename}-security  main restricted universe multiverse' }}
- { name: base           ,description: 'Ubuntu Basic'       ,module: node    ,releases: [         20,22,24] ,arch: [        aarch64] ,baseurl: { default: 'http://ports.ubuntu.com/ubuntu-ports/ ${distro_codename}             main universe multiverse restricted' ,china: 'https://mirrors.aliyun.com/ubuntu-ports/ ${distro_codename}           main restricted universe multiverse' }}
- { name: updates        ,description: 'Ubuntu Updates'     ,module: node    ,releases: [         20,22,24] ,arch: [        aarch64] ,baseurl: { default: 'http://ports.ubuntu.com/ubuntu-ports/ ${distro_codename}-backports   main restricted universe multiverse' ,china: 'https://mirrors.aliyun.com/ubuntu-ports/ ${distro_codename}-updates   main restricted universe multiverse' }}
- { name: backports      ,description: 'Ubuntu Backports'   ,module: node    ,releases: [         20,22,24] ,arch: [        aarch64] ,baseurl: { default: 'http://ports.ubuntu.com/ubuntu-ports/ ${distro_codename}-security    main restricted universe multiverse' ,china: 'https://mirrors.aliyun.com/ubuntu-ports/ ${distro_codename}-backports main restricted universe multiverse' }}
- { name: security       ,description: 'Ubuntu Security'    ,module: node    ,releases: [         20,22,24] ,arch: [        aarch64] ,baseurl: { default: 'http://ports.ubuntu.com/ubuntu-ports/ ${distro_codename}-updates     main restricted universe multiverse' ,china: 'https://mirrors.aliyun.com/ubuntu-ports/ ${distro_codename}-security  main restricted universe multiverse' }}
- { name: pgdg           ,description: 'PGDG'               ,module: pgsql   ,releases: [11,12,13,   22,24] ,arch: [x86_64, aarch64] ,baseurl: { default: 'http://apt.postgresql.org/pub/repos/apt/ ${distro_codename}-pgdg main'            ,china: 'https://repo.pigsty.cc/apt/pgdg/ ${distro_codename}-pgdg main' }}
- { name: pgdg-beta      ,description: 'PGDG Beta'          ,module: beta    ,releases: [11,12,13,   22,24] ,arch: [x86_64, aarch64] ,baseurl: { default: 'http://apt.postgresql.org/pub/repos/apt/ ${distro_codename}-pgdg-testing main 19' ,china: 'https://repo.pigsty.cc/apt/pgdg/ ${distro_codename}-pgdg-testing main 19' }}
- { name: timescaledb    ,description: 'TimescaleDB'        ,module: extra   ,releases: [11,12,   20,22,24] ,arch: [x86_64, aarch64] ,baseurl: { default: 'https://packagecloud.io/timescale/timescaledb/${distro_name}/ ${distro_codename} main' }}
- { name: citus          ,description: 'Citus'              ,module: extra   ,releases: [11,12,   20,22   ] ,arch: [x86_64, aarch64] ,baseurl: { default: 'https://packagecloud.io/citusdata/community/${distro_name}/ ${distro_codename} main' } }
- { name: percona        ,description: 'Percona TDE'        ,module: percona ,releases: [11,12,   20,22,24] ,arch: [x86_64, aarch64] ,baseurl: { default: 'https://repo.pigsty.io/apt/percona ${distro_codename} main' ,china: 'https://repo.pigsty.cc/apt/percona ${distro_codename} main' ,origin: 'http://repo.percona.com/ppg-17.5/apt ${distro_codename} main' }}
- { name: wiltondb       ,description: 'WiltonDB'           ,module: mssql   ,releases: [         20,22,24] ,arch: [x86_64, aarch64] ,baseurl: { default: 'https://repo.pigsty.io/apt/mssql/ ${distro_codename} main', china: 'https://repo.pigsty.cc/apt/mssql/ ${distro_codename} main' , origin: 'https://ppa.launchpadcontent.net/wiltondb/wiltondb/ubuntu/ ${distro_codename} main'  }}
- { name: groonga        ,description: 'Groonga Debian'     ,module: groonga ,releases: [11,12,13         ] ,arch: [x86_64, aarch64] ,baseurl: { default: 'https://packages.groonga.org/debian/ ${distro_codename} main' }}
- { name: groonga        ,description: 'Groonga Ubuntu'     ,module: groonga ,releases: [         20,22,24] ,arch: [x86_64, aarch64] ,baseurl: { default: 'https://ppa.launchpadcontent.net/groonga/ppa/ubuntu/ ${distro_codename} main' }}
- { name: mysql          ,description: 'MySQL'              ,module: mysql   ,releases: [11,12,   20,22,24] ,arch: [x86_64, aarch64] ,baseurl: { default: 'https://repo.mysql.com/apt/${distro_name} ${distro_codename} mysql-8.0 mysql-tools', china: 'https://mirrors.tuna.tsinghua.edu.cn/mysql/apt/${distro_name} ${distro_codename} mysql-8.0 mysql-tools' }}
- { name: mongo          ,description: 'MongoDB'            ,module: mongo   ,releases: [11,12,   20,22,24] ,arch: [x86_64, aarch64] ,baseurl: { default: 'https://repo.mongodb.org/apt/${distro_name} ${distro_codename}/mongodb-org/8.0 multiverse', china: 'https://mirrors.aliyun.com/mongodb/apt/${distro_name} ${distro_codename}/mongodb-org/8.0 multiverse' }}
- { name: redis          ,description: 'Redis'              ,module: redis   ,releases: [11,12,   20,22,24] ,arch: [x86_64, aarch64] ,baseurl: { default: 'https://packages.redis.io/deb ${distro_codename} main' }}
- { name: haproxyd       ,description: 'Haproxy Debian'     ,module: haproxy ,releases: [11,12            ] ,arch: [x86_64, aarch64] ,baseurl: { default: 'http://haproxy.debian.net/ ${distro_codename}-backports-3.1 main' }}
- { name: haproxyu       ,description: 'Haproxy Ubuntu'     ,module: haproxy ,releases: [         20,22,24] ,arch: [x86_64, aarch64] ,baseurl: { default: 'https://ppa.launchpadcontent.net/vbernat/haproxy-3.1/ubuntu/ ${distro_codename} main' }}
- { name: grafana        ,description: 'Grafana'            ,module: grafana ,releases: [11,12,13,20,22,24] ,arch: [x86_64, aarch64] ,baseurl: { default: 'https://apt.grafana.com stable main' ,china: 'https://mirrors.aliyun.com/grafana/apt/ stable main' }}
- { name: kubernetes     ,description: 'Kubernetes'         ,module: kube    ,releases: [11,12,13,20,22,24] ,arch: [x86_64, aarch64] ,baseurl: { default: 'https://pkgs.k8s.io/core:/stable:/v1.33/deb/ /', china: 'https://mirrors.aliyun.com/kubernetes-new/core/stable/v1.33/deb/ /' }}
- { name: gitlab-ee      ,description: 'Gitlab EE'          ,module: gitlab  ,releases: [11,12,13,20,22,24] ,arch: [x86_64, aarch64] ,baseurl: { default: 'https://packages.gitlab.com/gitlab/gitlab-ee/${distro_name}/ ${distro_codename} main' }}
- { name: gitlab-ce      ,description: 'Gitlab CE'          ,module: gitlab  ,releases: [11,12,13,20,22,24] ,arch: [x86_64, aarch64] ,baseurl: { default: 'https://packages.gitlab.com/gitlab/gitlab-ce/${distro_name}/ ${distro_codename} main' }}
- { name: clickhouse     ,description: 'ClickHouse'         ,module: click   ,releases: [11,12,13,20,22,24] ,arch: [x86_64, aarch64] ,baseurl: { default: 'https://packages.clickhouse.com/deb/ stable main', china: 'https://mirrors.aliyun.com/clickhouse/deb/ stable main' }}
```

您可以创建 `~/.pig/repo.yml` 文件，显式修改并覆盖 pig 的仓库定义。

在编辑仓库定义文件时，您可以在 `baseurl` 处添加额外的区域镜像，例如指定中国，欧洲地区的镜像仓库 URL。当 pig 使用 `--region` 参数指定特定的区域时，pig 会优先查找对应区域的仓库 URL，如果不存在，则会 Fallback 到 `default` 的仓库 URL。




## `repo list`

`pig repo list` 将列出当前系统可用的所有仓库模块，例如，在 Debian 13 x86 系统上运行时，pig 会列出当前系统可用的仓库与模块：

```yaml
os_environment: {code: d13, arch: amd64, type: deb, major: 13}
repo_upstream:  # Available Repo: 15
  - { name: pigsty-local   ,description: 'Pigsty Local'       ,module: local    ,releases: [11,12,13,20,22,24] ,arch: [x86_64, aarch64]  ,baseurl: 'http://${admin_ip}/pigsty ./' }
  - { name: pigsty-pgsql   ,description: 'Pigsty PgSQL'       ,module: pgsql    ,releases: [11,12,13,20,22,24] ,arch: [x86_64, aarch64]  ,baseurl: 'https://repo.pigsty.io/apt/pgsql/${distro_codename} ${distro_codename} main' }
  - { name: pigsty-infra   ,description: 'Pigsty Infra'       ,module: infra    ,releases: [11,12,13,20,22,24] ,arch: [x86_64, aarch64]  ,baseurl: 'https://repo.pigsty.io/apt/infra/ generic main' }
  - { name: docker-ce      ,description: 'Docker'             ,module: infra    ,releases: [11,12,13,20,22,24] ,arch: [x86_64, aarch64]  ,baseurl: 'https://download.docker.com/linux/${distro_name} ${distro_codename} stable' }
  - { name: base           ,description: 'Debian Basic'       ,module: node     ,releases: [11,12,13]       ,arch: [x86_64, aarch64]  ,baseurl: 'http://deb.debian.org/debian/ ${distro_codename} main non-free-firmware' }
  - { name: updates        ,description: 'Debian Updates'     ,module: node     ,releases: [11,12,13]       ,arch: [x86_64, aarch64]  ,baseurl: 'http://deb.debian.org/debian/ ${distro_codename}-updates main non-free-firmware' }
  - { name: security       ,description: 'Debian Security'    ,module: node     ,releases: [11,12,13]       ,arch: [x86_64, aarch64]  ,baseurl: 'http://security.debian.org/debian-security ${distro_codename}-security main non-free-firmware' }
  - { name: pgdg           ,description: 'PGDG'               ,module: pgsql    ,releases: [11,12,13,22,24] ,arch: [x86_64, aarch64]  ,baseurl: 'http://apt.postgresql.org/pub/repos/apt/ ${distro_codename}-pgdg main' }
  - { name: pgdg-beta      ,description: 'PGDG Beta'          ,module: beta     ,releases: [11,12,13,22,24] ,arch: [x86_64, aarch64]  ,baseurl: 'http://apt.postgresql.org/pub/repos/apt/ ${distro_codename}-pgdg-testing main 19' }
  - { name: groonga        ,description: 'Groonga Debian'     ,module: groonga  ,releases: [11,12,13]       ,arch: [x86_64, aarch64]  ,baseurl: 'https://packages.groonga.org/debian/ ${distro_codename} main' }
  - { name: grafana        ,description: 'Grafana'            ,module: grafana  ,releases: [11,12,13,20,22,24] ,arch: [x86_64, aarch64]  ,baseurl: 'https://apt.grafana.com stable main' }
  - { name: kubernetes     ,description: 'Kubernetes'         ,module: kube     ,releases: [11,12,13,20,22,24] ,arch: [x86_64, aarch64]  ,baseurl: 'https://pkgs.k8s.io/core:/stable:/v1.33/deb/ /' }
  - { name: gitlab-ee      ,description: 'Gitlab EE'          ,module: gitlab   ,releases: [11,12,13,20,22,24] ,arch: [x86_64, aarch64]  ,baseurl: 'https://packages.gitlab.com/gitlab/gitlab-ee/${distro_name}/ ${distro_codename} main' }
  - { name: gitlab-ce      ,description: 'Gitlab CE'          ,module: gitlab   ,releases: [11,12,13,20,22,24] ,arch: [x86_64, aarch64]  ,baseurl: 'https://packages.gitlab.com/gitlab/gitlab-ce/${distro_name}/ ${distro_codename} main' }
  - { name: clickhouse     ,description: 'ClickHouse'         ,module: click    ,releases: [11,12,13,20,22,24] ,arch: [x86_64, aarch64]  ,baseurl: 'https://packages.clickhouse.com/deb/ stable main' }
repo_modules:   # Available Modules: 20
  - all       : pigsty-infra, pigsty-pgsql, pgdg, base, updates, extras, epel, centos-sclo, centos-sclo-rh, baseos, appstream, powertools, crb, security, backports
  - pigsty    : pigsty-infra, pigsty-pgsql
  - pgdg      : pgdg
  - node      : base, updates, extras, epel, centos-sclo, centos-sclo-rh, baseos, appstream, powertools, crb, security, backports
  - infra     : pigsty-infra, nginx, docker-ce
  - pgsql     : pigsty-pgsql, pgdg-common, pgdg-el8fix, pgdg-el9fix, pgdg13, pgdg14, pgdg15, pgdg16, pgdg17, pgdg18, pgdg
  - extra     : pgdg-extras, pgdg13-nonfree, pgdg14-nonfree, pgdg15-nonfree, pgdg16-nonfree, pgdg17-nonfree, timescaledb, citus
  - mssql     : wiltondb
...
```

如果指定了额外的 `all` 位置参数，那么输出结果不会根据您的系统进行过滤，会打印出所有的仓库定义。

```bash
pig repo list all                # 列出所有仓库（不过滤）
```






------

## `repo info`

显示特定仓库或模块的详细信息，包括 URL、元数据和区域镜像，以及 `.repo` / `.list` 仓库文件内容。

```bash
pig repo info pgdg               # 显示 pgdg 模块的信息
pig repo info pigsty pgdg        # 显示多个模块的信息
pig repo info all                # 显示所有模块的信息
```

```bash
$ pig repo info pgdg

#-------------------------------------------------
Name       : pgdg
Summary    : PGDG
Available  : Yes (debian d13 amd64)
Module     : pgsql
OS Arch    : [x86_64, aarch64]
OS Distro  : deb [11,12,13,22,24]
Meta       : trusted=yes
Base URL   : http://apt.postgresql.org/pub/repos/apt/ ${distro_codename}-pgdg main
     china : https://repo.pigsty.cc/apt/pgdg/ ${distro_codename}-pgdg main

# default repo content
# pgdg PGDG
deb [trusted=yes] http://apt.postgresql.org/pub/repos/apt/ trixie-pgdg main

# china mirror repo content
# pgdg PGDG
deb [trusted=yes] https://repo.pigsty.cc/apt/pgdg/ trixie-pgdg main
```



## `repo status`

显示系统上的当前仓库配置。

```bash
pig repo status
```




## `repo add`

添加仓库配置文件到系统。需要 root/sudo 权限。

```bash
# 基本用法
pig repo add pgdg                # 添加 PGDG 仓库
pig repo add pgdg pigsty         # 添加多个仓库
pig repo add all                 # 添加所有必要仓库 (pgdg + pigsty + node)

# 带选项
pig repo add pigsty -u           # 添加并更新缓存
pig repo add all -r              # 添加前移除现有仓库
pig repo add all -ru             # 移除、添加并更新（完全重置）

# 区域镜像
pig repo add pgdg --region=china  # 使用中国镜像
```

**选项：**

- `-r|--remove`：添加新仓库前移除现有仓库
- `-u|--update`：添加仓库后运行包缓存更新 `apt update` / `dnf makecache`
- `--region <region>`：使用区域镜像仓库（ `default` / `china` / `europe` / ... ）

|   平台   | 模块位置                                    |
|:------:|-----------------------------------------|
|   EL   | `/etc/yum.repos.d/<module>.repo`        |
| Debian | `/etc/apt/sources.list.d/<module>.list` | 



**常用工作流：**

```bash
# 全新安装 - 推荐方法
pig repo add all -ru              # 清洁方式，一步到位

# 简便写法，简单粗暴
pig repo set                      # = pig repo add all --remove --update

# 添加 PGSQL 仓库（PGDG + PIGSTY）
sudo pig repo add pgsql -u        # 假设您自己已经配置好了操作系统的仓库，只需要 PG 相关仓库

# 只要 PGDG，不要 PIGSTY 仓库
sudo pig repo add pgdg           # 仅 PGDG

# 只要 PIGSTY，不要 PGDG 仓库
sudo pig repo add pigsty         # 仅 PIGSTY
```




## `repo set`

等同于 `repo add --remove --update`。清空现有仓库并设置新仓库，然后更新缓存。

```bash
pig repo set                     # 替换为默认仓库
pig repo set pgdg pigsty         # 替换为特定仓库并更新
pig repo set all --region=china  # 使用中国镜像
```

当您想要确保干净的仓库配置，没有旧的或冲突的仓库时，这很有用。
请注意 `pig repo set` 不接受 `--remove|-r` 和 `--update|-u` 参数，这两者都是默认的行为。
您仍然可以使用 `--region=<region>` 来指定 `default` / `china` / `europe` 等区域。








## `repo rm`

移除仓库配置文件并备份它们。

```bash
pig repo rm                      # 移除所有仓库
pig repo rm pgdg                 # 移除特定仓库
pig repo rm pgdg pigsty -u       # 移除并更新缓存
```

|   平台   | 备份位置                              |
|:------:|-----------------------------------|
|   EL   | `/etc/yum.repos.d/backup/`        |
| Debian | `/etc/apt/sources.list.d/backup/` | 



## `repo update`

更新包管理器缓存以反映仓库更改。

```bash
pig repo update                  # 更新包缓存
```

|   平台   | 等效命令            |
|:------:|-----------------|
|   EL   | `dnf makecache` |
| Debian | `apt update`    | 





## `repo create`

为离线安装或缓存创建本地包仓库。

```bash
pig repo create                  # 在默认位置创建 (/www/pigsty)
pig repo create /srv/repo        # 在自定义位置创建
pig repo create /www/pigsty /www/backup  # 多个位置
```

|   平台   | 依赖软件           |
|:------:|----------------|
|   EL   | `createrepo_c` |
| Debian | `dpkg-dev`     | 


**示例工作流：**
```bash
# 创建本地仓库结构
sudo pig repo create /www/pigsty

# 复制包到仓库
cp *.rpm /www/pigsty/     # 对于 EL
cp *.deb /www/pigsty/     # 对于 Debian

# 重新创建仓库元数据
sudo pig repo create /www/pigsty

# 添加本地仓库到系统
sudo pig repo add local
```



## `repo cache`

创建仓库内容的压缩 tarball 用于离线分发。

```bash
pig repo cache                   # 默认：/www 到 /tmp/pkg.tgz
pig repo cache -f                # 强制覆盖现有文件
pig repo cache -d /srv           # 自定义源目录
pig repo cache pigsty mssql      # 仅特定仓库
```

**选项：**
- `-d, --dir`：源目录（默认：`/www/`）
- `-p, --path`：输出路径（默认：`/tmp/pkg.tgz`）
- `-f`：强制覆盖现有包

**示例工作流：**
```bash
# 在线系统上：创建离线包
pig repo add -ru                 # 设置仓库
pig install pg17 pg_duckdb       # 安装包
pig repo create                  # 创建本地仓库
pig repo cache                   # 创建离线包

# 将 pkg.tgz 传输到离线系统

# 离线系统上：从包引导
pig repo boot                    # 解压并设置
pig repo add local              # 使用本地仓库
pig install pg17 pg_duckdb      # 从本地安装
```



## `repo boot`

从离线包解压并设置本地仓库。

```bash
pig repo boot                    # 默认：/tmp/pkg.tgz 到 /www
pig repo boot -p /mnt/pkg.tgz   # 自定义包路径
pig repo boot -d /srv           # 自定义目标目录
```

**选项：**
- `-p, --path`：包路径（默认：`/tmp/pkg.tgz`）
- `-d, --dir`：目标目录（默认：`/www/`）
