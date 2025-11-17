package cli

// RPM distribution configuration template
const rpmConfigTemplate = `---
# License    :   AGPLv3 @ https://doc.pgsty.com/about/license
# Copyright  :   2018-2025  Ruohang Feng / Vonng (rh@vonng.com)
# {{ .OSName }}  :  {{ .OSCode }} {{ .Arch }} OS specific vars

# where to register systemd files
systemd_dir: /usr/lib/systemd/system
syslog_path: /var/log/messages

# default packages to be downloaded (if ` + "`repo_packages`" + ` is not explicitly set)
repo_packages_default: [ node-bootstrap, infra-package, infra-addons, node-package1, node-package2, pgsql-utility, extra-modules ]

# default postgres packages to be downloaded
repo_extra_packages_default: [ pgsql-main ]

# default node packages to be installed (if ` + "`node_default_packages`" + ` is not explicitly set)
node_packages_default:
  - lz4,unzip,bzip2,pv,jq,git,ncdu,make,patch,bash,lsof,wget,uuid,tuned,nvme-cli,numactl,sysstat,iotop,htop,rsync,tcpdump
  - python3,python3-pip,socat,lrzsz,net-tools,ipvsadm,telnet,ca-certificates,openssl,keepalived,etcd,haproxy,chrony,pig
  - zlib,yum,audit,bind-utils,readline,vim-minimal,node_exporter,grubby,openssh-server,openssh-clients,chkconfig

# default infra packages to be installed (if ` + "`infra_packages`" + ` is not explicitly set)
infra_packages_default:
  - grafana,loki,logcli,prometheus,alertmanager,pushgateway,grafana-plugins,restic,certbot,python3-certbot-nginx
  - node_exporter,blackbox_exporter,nginx_exporter,pg_exporter,pev2,nginx,dnsmasq,ansible,etcd,python3-requests,redis,mcli

# postgres home dir in various mode
pg_home_map:
  pgsql:  '/usr/pgsql-$v'
  citus:  '/usr/pgsql-$v'
  mssql:  '/usr/'
  ivory:  '/usr/ivory-4'
  mysql:  '/usr/halo-14'
  gpsql:  '/usr/gpsql'
  polar:  '/u01/polardb_pg'
  oracle: '/u01/polardb_pg'
  oriole: '/usr/oriole-$v'

# default upstream repo (if ` + "`repo_upstream`" + ` is not explicitly set)
repo_upstream_default:
  - { name: pigsty-local   ,description: 'Pigsty Local'       ,module: local   ,releases: [8,9,10] ,arch: [x86_64, aarch64] ,baseurl: { default: 'http://${admin_ip}/pigsty'  }} # used by intranet nodes
  - { name: pigsty-infra   ,description: 'Pigsty INFRA'       ,module: infra   ,releases: [8,9,10] ,arch: [x86_64, aarch64] ,baseurl: { default: 'https://repo.pigsty.io/yum/infra/$basearch' ,china: 'https://repo.pigsty.cc/yum/infra/$basearch' }}
  - { name: pigsty-pgsql   ,description: 'Pigsty PGSQL'       ,module: pgsql   ,releases: [8,9,10] ,arch: [x86_64, aarch64] ,baseurl: { default: 'https://repo.pigsty.io/yum/pgsql/el$releasever.$basearch' ,china: 'https://repo.pigsty.cc/yum/pgsql/el$releasever.$basearch' }}
  - { name: nginx          ,description: 'Nginx Repo'         ,module: infra   ,releases: [8,9,10] ,arch: [x86_64, aarch64] ,baseurl: { default: 'https://nginx.org/packages/rhel/$releasever/$basearch/' }}
  - { name: docker-ce      ,description: 'Docker CE'          ,module: infra   ,releases: [8,9,10] ,arch: [x86_64, aarch64] ,baseurl: { default: 'https://download.docker.com/linux/centos/$releasever/$basearch/stable'    ,china: 'https://mirrors.aliyun.com/docker-ce/linux/centos/$releasever/$basearch/stable' ,europe: 'https://mirrors.xtom.de/docker-ce/linux/centos/$releasever/$basearch/stable' }}
  - { name: baseos         ,description: 'EL 8+ BaseOS'       ,module: node    ,releases: [8,9,10] ,arch: [x86_64, aarch64] ,baseurl: { default: 'https://dl.rockylinux.org/pub/rocky/$releasever/BaseOS/$basearch/os/'     ,china: 'https://mirrors.aliyun.com/rockylinux/$releasever/BaseOS/$basearch/os/'         ,europe: 'https://mirrors.xtom.de/rocky/$releasever/BaseOS/$basearch/os/'     }}
  - { name: appstream      ,description: 'EL 8+ AppStream'    ,module: node    ,releases: [8,9,10] ,arch: [x86_64, aarch64] ,baseurl: { default: 'https://dl.rockylinux.org/pub/rocky/$releasever/AppStream/$basearch/os/'  ,china: 'https://mirrors.aliyun.com/rockylinux/$releasever/AppStream/$basearch/os/'      ,europe: 'https://mirrors.xtom.de/rocky/$releasever/AppStream/$basearch/os/'  }}
  - { name: extras         ,description: 'EL 8+ Extras'       ,module: node    ,releases: [8,9,10] ,arch: [x86_64, aarch64] ,baseurl: { default: 'https://dl.rockylinux.org/pub/rocky/$releasever/extras/$basearch/os/'     ,china: 'https://mirrors.aliyun.com/rockylinux/$releasever/extras/$basearch/os/'         ,europe: 'https://mirrors.xtom.de/rocky/$releasever/extras/$basearch/os/'     }}{{ if eq .OSCode "el8" }}
  - { name: powertools     ,description: 'EL 8 PowerTools'    ,module: node    ,releases: [8     ] ,arch: [x86_64, aarch64] ,baseurl: { default: 'https://dl.rockylinux.org/pub/rocky/$releasever/PowerTools/$basearch/os/' ,china: 'https://mirrors.aliyun.com/rockylinux/$releasever/PowerTools/$basearch/os/'     ,europe: 'https://mirrors.xtom.de/rocky/$releasever/PowerTools/$basearch/os/' }}{{ else }}
  - { name: crb            ,description: 'EL 9 CRB'           ,module: node    ,releases: [  9,10] ,arch: [x86_64, aarch64] ,baseurl: { default: 'https://dl.rockylinux.org/pub/rocky/$releasever/CRB/$basearch/os/'        ,china: 'https://mirrors.aliyun.com/rockylinux/$releasever/CRB/$basearch/os/'            ,europe: 'https://mirrors.xtom.de/rocky/$releasever/CRB/$basearch/os/'        }}{{ end }}
  - { name: epel           ,description: 'EL 8+ EPEL'         ,module: node    ,releases: [8,9   ] ,arch: [x86_64, aarch64] ,baseurl: { default: 'https://mirrors.edge.kernel.org/fedora-epel/$releasever/Everything/$basearch/' ,china: 'https://mirrors.aliyun.com/epel/$releasever/Everything/$basearch/'         ,europe: 'https://mirrors.xtom.de/epel/$releasever/Everything/$basearch/'     }}
  - { name: epel           ,description: 'EL 10 EPEL'         ,module: node    ,releases: [    10] ,arch: [x86_64, aarch64] ,baseurl: { default: 'https://mirrors.edge.kernel.org/fedora-epel/$releasever.0/Everything/$basearch/' ,china: 'https://mirrors.aliyun.com/epel/$releasever.0/Everything/$basearch/'         ,europe: 'https://mirrors.xtom.de/epel/$releasever.0/Everything/$basearch/' } }
  - { name: pgdg-common    ,description: 'PostgreSQL Common'  ,module: pgsql   ,releases: [8,9,10] ,arch: [x86_64, aarch64] ,baseurl: { default: 'https://download.postgresql.org/pub/repos/yum/common/redhat/rhel-$releasever-$basearch' ,china: 'https://repo.pigsty.cc/yum/pgdg/common/redhat/rhel-$releasever-$basearch' , europe: 'https://mirrors.xtom.de/postgresql/repos/yum/common/redhat/rhel-$releasever-$basearch' }}{{ if eq .OSCode "el8" }}
  - { name: pgdg-el8fix    ,description: 'PostgreSQL EL8FIX'  ,module: pgsql   ,releases: [8     ] ,arch: [x86_64, aarch64] ,baseurl: { default: 'https://download.postgresql.org/pub/repos/yum/common/pgdg-centos8-sysupdates/redhat/rhel-8-$basearch/' ,china: 'https://repo.pigsty.cc/yum/pgdg/common/pgdg-centos8-sysupdates/redhat/rhel-8-x86_64/' ,europe: 'https://mirrors.xtom.de/postgresql/repos/yum/common/pgdg-centos8-sysupdates/redhat/rhel-8-x86_64/' }}{{ else if eq .OSCode "el9" }}
  - { name: pgdg-el9fix    ,description: 'PostgreSQL EL9FIX'  ,module: pgsql   ,releases: [  9   ] ,arch: [x86_64, aarch64] ,baseurl: { default: 'https://download.postgresql.org/pub/repos/yum/common/pgdg-rocky9-sysupdates/redhat/rhel-9-$basearch/'  ,china: 'https://repo.pigsty.cc/yum/pgdg/common/pgdg-rocky9-sysupdates/redhat/rhel-9-x86_64/'  ,europe: 'https://mirrors.xtom.de/postgresql/repos/yum/common/pgdg-rocky9-sysupdates/redhat/rhel-9-x86_64/'  }}{{ end }}
  - { name: pgdg14         ,description: 'PostgreSQL 14'      ,module: pgsql   ,releases: [8,9,10] ,arch: [x86_64, aarch64] ,baseurl: { default: 'https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-$releasever-$basearch' ,china: 'https://repo.pigsty.cc/yum/pgdg/14/redhat/rhel-$releasever-$basearch' ,europe: 'https://mirrors.xtom.de/postgresql/repos/yum/14/redhat/rhel-$releasever-$basearch' }}
  - { name: pgdg15         ,description: 'PostgreSQL 15'      ,module: pgsql   ,releases: [8,9,10] ,arch: [x86_64, aarch64] ,baseurl: { default: 'https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-$releasever-$basearch' ,china: 'https://repo.pigsty.cc/yum/pgdg/15/redhat/rhel-$releasever-$basearch' ,europe: 'https://mirrors.xtom.de/postgresql/repos/yum/15/redhat/rhel-$releasever-$basearch' }}
  - { name: pgdg16         ,description: 'PostgreSQL 16'      ,module: pgsql   ,releases: [8,9,10] ,arch: [x86_64, aarch64] ,baseurl: { default: 'https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-$releasever-$basearch' ,china: 'https://repo.pigsty.cc/yum/pgdg/16/redhat/rhel-$releasever-$basearch' ,europe: 'https://mirrors.xtom.de/postgresql/repos/yum/16/redhat/rhel-$releasever-$basearch' }}
  - { name: pgdg17         ,description: 'PostgreSQL 17'      ,module: pgsql   ,releases: [8,9,10] ,arch: [x86_64, aarch64] ,baseurl: { default: 'https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-$releasever-$basearch' ,china: 'https://repo.pigsty.cc/yum/pgdg/17/redhat/rhel-$releasever-$basearch' ,europe: 'https://mirrors.xtom.de/postgresql/repos/yum/17/redhat/rhel-$releasever-$basearch' }}
  - { name: pgdg18         ,description: 'PostgreSQL 18'      ,module: pgsql   ,releases: [8,9,10] ,arch: [x86_64, aarch64] ,baseurl: { default: 'https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-$releasever-$basearch' ,china: 'https://repo.pigsty.cc/yum/pgdg/18/redhat/rhel-$releasever-$basearch' ,europe: 'https://mirrors.xtom.de/postgresql/repos/yum/18/redhat/rhel-$releasever-$basearch' }}

# package alias map to simplify package download & installation
package_map:
  #--------------------------------#
  # REPO: download packages
  #--------------------------------#
  node-bootstrap:          "ansible python3 python3-pip python3-virtualenv python3-requests python3-jmespath python3-cryptography dnf-utils modulemd-tools createrepo_c sshpass"
  infra-package:           "nginx dnsmasq etcd haproxy vip-manager node_exporter keepalived_exporter pg_exporter pgbackrest_exporter redis_exporter redis minio mcli pig"
  infra-addons:            "grafana grafana-plugins loki logcli vector promtail prometheus alertmanager pushgateway blackbox_exporter nginx_exporter pev2 certbot python3-certbot-nginx"
  extra-modules:           "docker-ce docker-compose-plugin ferretdb2 duckdb restic juicefs vray grafana-infinity-ds"
  node-package1:           "lz4 unzip bzip2 zlib yum pv jq git ncdu make patch bash lsof wget uuid tuned nvme-cli numactl grubby sysstat iotop htop rsync tcpdump perf flamegraph chkconfig"
  node-package2:           "netcat socat ftp lrzsz net-tools ipvsadm bind-utils telnet audit ca-certificates readline vim-minimal keepalived chrony openssl openssh-server openssh-clients"
  pgsql-utility:           "patroni patroni-etcd pgbouncer pgbackrest pgbadger pg_timetable pgFormatter pg_filedump pgxnclient timescaledb-tools timescaledb-event-streamer"

  #--------------------------------#
  # PGSQL: kernel packages
  #--------------------------------#
  postgresql:              "postgresql$v*"
  pgsql:                   "postgresql$v postgresql$v-server postgresql$v-libs postgresql$v-contrib postgresql$v-plperl postgresql$v-plpython3 postgresql$v-pltcl postgresql$v-llvmjit"
  pgsql-mini:              "postgresql$v postgresql$v-server postgresql$v-libs postgresql$v-contrib"
  pgsql-core:              "postgresql$v postgresql$v-server postgresql$v-libs postgresql$v-contrib postgresql$v-plperl postgresql$v-plpython3 postgresql$v-pltcl postgresql$v-llvmjit"
  pgsql-full:              "postgresql$v postgresql$v-server postgresql$v-libs postgresql$v-contrib postgresql$v-plperl postgresql$v-plpython3 postgresql$v-pltcl postgresql$v-llvmjit postgresql$v-test postgresql$v-devel"
  pgsql-main:              "postgresql$v postgresql$v-server postgresql$v-libs postgresql$v-contrib postgresql$v-plperl postgresql$v-plpython3 postgresql$v-pltcl postgresql$v-llvmjit pg_repack_$v* wal2json_$v* pgvector_$v*"
  pgsql-client:            "postgresql$v"
  pgsql-server:            "postgresql$v-server postgresql$v-libs postgresql$v-contrib"
  pgsql-devel:             "postgresql$v-devel"
  pgsql-basic:             "pg_repack_$v* wal2json_$v* pgvector_$v*"

  #--------------------------------#
  # MISC: pro modules
  #--------------------------------#
  mssql:                   "wiltondb* libicu-devel"
  orioledb:                "orioledb-postgresql-$v orioledb-postgresql-$v-contrib orioledb-postgresql-$v-server orioledb-postgresql-$v-libs"
  percona:                 "percona-postgresql$v-server percona-postgresql$v-libs percona-postgresql$v percona-postgresql$v-contrib"
  polardb:                 "PolarDB-$v"
  ivorysql:                "ivorysql4 ivorysql4-server ivorysql4-contrib ivorysql4-libs"
  gpsql:                   "PolarDB-PG_$v"
  java-runtime:            "java-17-openjdk-src java-17-openjdk-headless"
  kafka:                   "kafka kafka_exporter"
  kube-runtime:            "containerd.io"
  sealos:                  "sealos"
  kubernetes:              "kubeadm kubelet kubectl"
  docker:                  "docker-ce docker-compose-plugin"
  infra-extra:             "victoria-metrics victoria-metrics-cluster vmutils grafana-victoriametrics-ds victoria-logs vlogscli vlagent grafana-victorialogs-ds rclone mysqld_exporter mongodb_exporter kafka_exporter"
  victoria:                "victoria-metrics victoria-metrics-cluster vmutils grafana-victoriametrics-ds victoria-logs vlogscli vlagent grafana-victorialogs-ds"
  vmetrics:                "victoria-metrics victoria-metrics-cluster vmutils grafana-victoriametrics-ds"
  vlogs:                   "victoria-logs vlogscli vlagent grafana-victorialogs-ds"
  tigerbeetle:             "tigerbeetle"
  clickhouse:              "clickhouse-server clickhouse-client clickhouse-common-static"

  #--------------------------------#
  # PGSQL: extensions (default pg18)
  #--------------------------------#{{ range $pgVer, $mappings := .PGMappings }}{{ if eq $pgVer 18 }}{{ range $mappings }}
  {{ formatPkgLine . }}{{ end }}{{ end }}{{ end }}

{{ range .PGVersions }}
  #--------------------------------#
  # PG{{ . }}: packages
  #--------------------------------#
  pg{{ . }}:                    "postgresql{{ . }} postgresql{{ . }}-server postgresql{{ . }}-libs postgresql{{ . }}-contrib postgresql{{ . }}-plperl postgresql{{ . }}-plpython3 postgresql{{ . }}-pltcl postgresql{{ . }}-llvmjit"
  pg{{ . }}-mini:               "postgresql{{ . }} postgresql{{ . }}-server postgresql{{ . }}-libs postgresql{{ . }}-contrib"
  pg{{ . }}-core:               "postgresql{{ . }} postgresql{{ . }}-server postgresql{{ . }}-libs postgresql{{ . }}-contrib postgresql{{ . }}-plperl postgresql{{ . }}-plpython3 postgresql{{ . }}-pltcl postgresql{{ . }}-llvmjit"
  pg{{ . }}-full:               "postgresql{{ . }} postgresql{{ . }}-server postgresql{{ . }}-libs postgresql{{ . }}-contrib postgresql{{ . }}-plperl postgresql{{ . }}-plpython3 postgresql{{ . }}-pltcl postgresql{{ . }}-llvmjit postgresql{{ . }}-test postgresql{{ . }}-devel"
  pg{{ . }}-main:               "postgresql{{ . }} postgresql{{ . }}-server postgresql{{ . }}-libs postgresql{{ . }}-contrib postgresql{{ . }}-plperl postgresql{{ . }}-plpython3 postgresql{{ . }}-pltcl postgresql{{ . }}-llvmjit pg_repack_{{ . }}* wal2json_{{ . }}* pgvector_{{ . }}*"
  pg{{ . }}-client:             "postgresql{{ . }}"
  pg{{ . }}-server:             "postgresql{{ . }}-server postgresql{{ . }}-libs postgresql{{ . }}-contrib"
  pg{{ . }}-devel:              "postgresql{{ . }}-devel"
  pg{{ . }}-basic:              "pg_repack_{{ . }}* wal2json_{{ . }}* pgvector_{{ . }}*"{{ range index $.PGMappings . }}
  {{ formatPkgLine . }}{{ end }}
{{ end }}

...`

// DEB distribution configuration template
const debConfigTemplate = `---
# License    :   AGPLv3 @ https://doc.pgsty.com/about/license
# Copyright  :   2018-2025  Ruohang Feng / Vonng (rh@vonng.com)
# {{ .OSName }} :   {{ .OSCode }} {{ .Arch }} OS specific vars

# where to register systemd files
systemd_dir: /lib/systemd/system
syslog_path: /var/log/syslog

# default packages to be downloaded (if ` + "`repo_packages`" + ` is not explicitly set)
repo_packages_default: [ node-bootstrap, infra-package, infra-addons, node-package1, node-package2, pgsql-utility, extra-modules ]

# default postgres packages to be downloaded
repo_extra_packages_default: [ pgsql-main ]

# default node packages to be installed (if ` + "`node_default_packages`" + ` is not explicitly set)
node_packages_default:
  - lz4,unzip,bzip2,pv,jq,git,ncdu,make,patch,bash,lsof,wget,uuid,tuned,nvme-cli,numactl,sysstat,iotop,htop,rsync,tcpdump
  - python3,python3-pip,socat,lrzsz,net-tools,ipvsadm,telnet,ca-certificates,openssl,keepalived,etcd,haproxy,chrony,pig
  - zlib1g,acl,dnsutils,libreadline-dev,vim-tiny,node-exporter,openssh-server,openssh-client

# default infra packages to be installed (if ` + "`infra_packages`" + ` is not explicitly set)
infra_packages_default:
  - grafana,grafana-plugins,loki,logcli,prometheus,alertmanager,pushgateway,restic,certbot,python3-certbot-nginx
  - node-exporter,blackbox-exporter,nginx-exporter,pg-exporter,pev2,nginx,dnsmasq,ansible,etcd,python3-requests,redis,mcli

# postgres home dir in various mode
pg_home_map:
  pgsql:  '/usr/lib/postgresql/$v'
  citus:  '/usr/lib/postgresql/$v'
  mssql:  '/usr/lib/postgresql/$v'
  ivory:  '/usr/ivory-4'
  mysql:  '/usr/halo-14'
  polar:  '/u01/polardb_pg'
  oracle: '/u01/polardb_pg'
  oriole: '/usr/oriole-$v'

# default upstream repo (if ` + "`repo_upstream`" + ` is not explicitly set)
repo_upstream_default:
  - { name: pigsty-local   ,description: 'Pigsty Local'       ,module: local   ,releases: [11,12,13,20,22,24] ,arch: [x86_64, aarch64] ,baseurl: { default: 'http://${admin_ip}/pigsty ./' }}
  - { name: pigsty-pgsql   ,description: 'Pigsty PgSQL'       ,module: pgsql   ,releases: [11,12,13,20,22,24] ,arch: [x86_64, aarch64] ,baseurl: { default: 'https://repo.pigsty.io/apt/pgsql/${distro_codename} ${distro_codename} main', china: 'https://repo.pigsty.cc/apt/pgsql/${distro_codename} ${distro_codename} main' }}
  - { name: pigsty-infra   ,description: 'Pigsty Infra'       ,module: infra   ,releases: [11,12,13,20,22,24] ,arch: [x86_64, aarch64] ,baseurl: { default: 'https://repo.pigsty.io/apt/infra/ generic main' ,china: 'https://repo.pigsty.cc/apt/infra/ generic main' }}
  - { name: nginx          ,description: 'Nginx'              ,module: infra   ,releases: [11,12,   20,22,24] ,arch: [x86_64, aarch64] ,baseurl: { default: 'http://nginx.org/packages/${distro_name} ${distro_codename} nginx' }}
  - { name: docker-ce      ,description: 'Docker'             ,module: infra   ,releases: [11,12,13,20,22,24] ,arch: [x86_64, aarch64] ,baseurl: { default: 'https://download.docker.com/linux/${distro_name} ${distro_codename} stable'                               ,china: 'https://mirrors.aliyun.com/docker-ce/linux/${distro_name} ${distro_codename} stable' }}
  - { name: base           ,description: 'Debian Basic'       ,module: node    ,releases: [11,12,13         ] ,arch: [x86_64, aarch64] ,baseurl: { default: 'http://deb.debian.org/debian/ ${distro_codename} main non-free-firmware'                                  ,china: 'https://mirrors.aliyun.com/debian/ ${distro_codename} main restricted universe multiverse' }}
  - { name: updates        ,description: 'Debian Updates'     ,module: node    ,releases: [11,12,13         ] ,arch: [x86_64, aarch64] ,baseurl: { default: 'http://deb.debian.org/debian/ ${distro_codename}-updates main non-free-firmware'                          ,china: 'https://mirrors.aliyun.com/debian/ ${distro_codename}-updates main restricted universe multiverse' }}
  - { name: security       ,description: 'Debian Security'    ,module: node    ,releases: [11,12,13         ] ,arch: [x86_64, aarch64] ,baseurl: { default: 'http://security.debian.org/debian-security ${distro_codename}-security main non-free-firmware'            ,china: 'https://mirrors.aliyun.com/debian-security/ ${distro_codename}-security main non-free-firmware' }}
  - { name: pgdg           ,description: 'PostgreSQL'         ,module: pgsql   ,releases: [11,12,13,20,22,24] ,arch: [x86_64, aarch64] ,baseurl: { default: 'https://apt.postgresql.org/pub/repos/apt/ ${distro_codename}-pgdg main', china: 'https://mirrors.aliyun.com/postgresql/repos/apt/ ${distro_codename}-pgdg main' }}

# package alias map to simplify package download & installation
package_map:
  #--------------------------------#
  # REPO: download packages
  #--------------------------------#
  node-bootstrap:          "ansible python3 python3-pip python3-venv python3-jmespath dpkg-dev sshpass"
  infra-package:           "nginx dnsmasq etcd haproxy vip-manager node-exporter keepalived-exporter pg-exporter pgbackrest-exporter redis-exporter redis minio mcli pig"
  infra-addons:            "grafana grafana-plugins loki logcli vector promtail prometheus alertmanager pushgateway blackbox-exporter nginx-exporter pev2 certbot python3-certbot-nginx"
  extra-modules:           "docker-ce docker-compose-plugin ferretdb2 duckdb restic juicefs vray grafana-infinity-ds"
  node-package1:           "lz4 unzip bzip2 zlib1g pv jq git ncdu make patch bash lsof wget uuid tuned nvme-cli numactl sysstat iotop htop rsync tcpdump acl chrony"
  node-package2:           "netcat-openbsd socat lrzsz net-tools ipvsadm dnsutils telnet ca-certificates libreadline-dev vim-tiny keepalived openssl openssh-server openssh-client"
  pgsql-utility:           "patroni pgbouncer pgbackrest pgbadger pg-timetable pgformatter postgresql-filedump pgxnclient timescaledb-tools timescaledb-event-streamer"

  #--------------------------------#
  # PGSQL: kernel packages
  #--------------------------------#
  postgresql:              "postgresql-$v"
  pgsql:                   "postgresql-$v postgresql-client-$v postgresql-$v-plpython3 postgresql-$v-plperl postgresql-$v-pltcl"
  pgsql-mini:              "postgresql-$v postgresql-client-$v"
  pgsql-core:              "postgresql-$v postgresql-client-$v postgresql-$v-plpython3 postgresql-$v-plperl postgresql-$v-pltcl"
  pgsql-full:              "postgresql-$v postgresql-client-$v postgresql-$v-plpython3 postgresql-$v-plperl postgresql-$v-pltcl postgresql-server-dev-$v"
  pgsql-main:              "postgresql-$v postgresql-client-$v postgresql-$v-plpython3 postgresql-$v-plperl postgresql-$v-pltcl postgresql-$v-repack postgresql-$v-wal2json postgresql-$v-pgvector"
  pgsql-client:            "postgresql-client-$v"
  pgsql-server:            "postgresql-$v"
  pgsql-devel:             "postgresql-server-dev-$v"
  pgsql-basic:             "postgresql-$v-repack postgresql-$v-wal2json postgresql-$v-pgvector"

  #--------------------------------#
  # MISC: pro modules
  #--------------------------------#
  mssql:                   "wiltondb* libicu-dev"
  gpsql:                   "postgres-polar-$v"
  java-runtime:            "openjdk-17-jdk-headless openjdk-17-jdk"
  kafka:                   "kafka kafka-exporter"
  kube-runtime:            "containerd.io"
  sealos:                  "sealos"
  kubernetes:              "kubeadm kubelet kubectl"
  docker:                  "docker-ce docker-compose-plugin"
  infra-extra:             "victoria-metrics victoria-metrics-cluster vmutils grafana-victoriametrics-ds victoria-logs vlogscli vlagent grafana-victorialogs-ds rclone mysqld-exporter mongodb-exporter kafka-exporter"
  victoria:                "victoria-metrics victoria-metrics-cluster vmutils grafana-victoriametrics-ds victoria-logs vlogscli vlagent grafana-victorialogs-ds"
  vmetrics:                "victoria-metrics victoria-metrics-cluster vmutils grafana-victoriametrics-ds"
  vlogs:                   "victoria-logs vlogscli vlagent grafana-victorialogs-ds"
  tigerbeetle:             "tigerbeetle"
  clickhouse:              "clickhouse-server clickhouse-client clickhouse-common-static"

  #--------------------------------#
  # PGSQL: extensions (default pg18)
  #--------------------------------#{{ range $pgVer, $mappings := .PGMappings }}{{ if eq $pgVer 18 }}{{ range $mappings }}
  {{ formatPkgLine . }}{{ end }}{{ end }}{{ end }}

{{ range .PGVersions }}
  #--------------------------------#
  # PG{{ . }}: packages
  #--------------------------------#
  pg{{ . }}:                    "postgresql-{{ . }} postgresql-client-{{ . }} postgresql-{{ . }}-plpython3 postgresql-{{ . }}-plperl postgresql-{{ . }}-pltcl"
  pg{{ . }}-mini:               "postgresql-{{ . }} postgresql-client-{{ . }}"
  pg{{ . }}-core:               "postgresql-{{ . }} postgresql-client-{{ . }} postgresql-{{ . }}-plpython3 postgresql-{{ . }}-plperl postgresql-{{ . }}-pltcl"
  pg{{ . }}-full:               "postgresql-{{ . }} postgresql-client-{{ . }} postgresql-{{ . }}-plpython3 postgresql-{{ . }}-plperl postgresql-{{ . }}-pltcl postgresql-server-dev-{{ . }}"
  pg{{ . }}-main:               "postgresql-{{ . }} postgresql-client-{{ . }} postgresql-{{ . }}-plpython3 postgresql-{{ . }}-plperl postgresql-{{ . }}-pltcl postgresql-{{ . }}-repack postgresql-{{ . }}-wal2json postgresql-{{ . }}-pgvector"
  pg{{ . }}-client:             "postgresql-client-{{ . }}"
  pg{{ . }}-server:             "postgresql-{{ . }}"
  pg{{ . }}-devel:              "postgresql-server-dev-{{ . }}"
  pg{{ . }}-basic:              "postgresql-{{ . }}-repack postgresql-{{ . }}-wal2json postgresql-{{ . }}-pgvector"{{ range $.PGMappings {{ . }} }}{{ range . }}
  {{ formatPkgLine . }}{{ end }}{{ end }}
{{ end }}

...`