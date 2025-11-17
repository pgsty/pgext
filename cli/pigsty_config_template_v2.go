package cli

// rpmTemplateV2 is the template for RPM-based distributions
const rpmTemplateV2 = `---
# License    :   AGPLv3 @ https://doc.pgsty.com/about/license
# Copyright  :   2018-2025  Ruohang Feng / Vonng (rh@vonng.com)
# {{ .OSName }}  :  {{ if eq .OSCode "el8" }}RHEL 8{{ else if eq .OSCode "el9" }}RHEL 9{{ else if eq .OSCode "el10" }}RHEL 10{{ end }} Compatible

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
  - { name: extras         ,description: 'EL 8+ Extras'       ,module: node    ,releases: [8,9,10] ,arch: [x86_64, aarch64] ,baseurl: { default: 'https://dl.rockylinux.org/pub/rocky/$releasever/extras/$basearch/os/'     ,china: 'https://mirrors.aliyun.com/rockylinux/$releasever/extras/$basearch/os/'         ,europe: 'https://mirrors.xtom.de/rocky/$releasever/extras/$basearch/os/'     }}
  - { name: powertools     ,description: 'EL 8 PowerTools'    ,module: node    ,releases: [8     ] ,arch: [x86_64, aarch64] ,baseurl: { default: 'https://dl.rockylinux.org/pub/rocky/$releasever/PowerTools/$basearch/os/' ,china: 'https://mirrors.aliyun.com/rockylinux/$releasever/PowerTools/$basearch/os/'     ,europe: 'https://mirrors.xtom.de/rocky/$releasever/PowerTools/$basearch/os/' }}
  - { name: crb            ,description: 'EL 9 CRB'           ,module: node    ,releases: [  9,10] ,arch: [x86_64, aarch64] ,baseurl: { default: 'https://dl.rockylinux.org/pub/rocky/$releasever/CRB/$basearch/os/'        ,china: 'https://mirrors.aliyun.com/rockylinux/$releasever/CRB/$basearch/os/'            ,europe: 'https://mirrors.xtom.de/rocky/$releasever/CRB/$basearch/os/'        }}
  - { name: epel           ,description: 'EL 8+ EPEL'         ,module: node    ,releases: [8,9   ] ,arch: [x86_64, aarch64] ,baseurl: { default: 'https://mirrors.edge.kernel.org/fedora-epel/$releasever/Everything/$basearch/' ,china: 'https://mirrors.aliyun.com/epel/$releasever/Everything/$basearch/'         ,europe: 'https://mirrors.xtom.de/epel/$releasever/Everything/$basearch/'     }}
  - { name: epel           ,description: 'EL 10 EPEL'         ,module: node    ,releases: [    10] ,arch: [x86_64, aarch64] ,baseurl: { default: 'https://mirrors.edge.kernel.org/fedora-epel/$releasever.0/Everything/$basearch/' ,china: 'https://mirrors.aliyun.com/epel/$releasever.0/Everything/$basearch/'         ,europe: 'https://mirrors.xtom.de/epel/$releasever.0/Everything/$basearch/' } }
  - { name: pgdg-common    ,description: 'PostgreSQL Common'  ,module: pgsql   ,releases: [8,9,10] ,arch: [x86_64, aarch64] ,baseurl: { default: 'https://download.postgresql.org/pub/repos/yum/common/redhat/rhel-$releasever-$basearch' ,china: 'https://repo.pigsty.cc/yum/pgdg/common/redhat/rhel-$releasever-$basearch' , europe: 'https://mirrors.xtom.de/postgresql/repos/yum/common/redhat/rhel-$releasever-$basearch' }}
  - { name: pgdg-el8fix    ,description: 'PostgreSQL EL8FIX'  ,module: pgsql   ,releases: [8     ] ,arch: [x86_64, aarch64] ,baseurl: { default: 'https://download.postgresql.org/pub/repos/yum/common/pgdg-centos8-sysupdates/redhat/rhel-8-$basearch/' ,china: 'https://repo.pigsty.cc/yum/pgdg/common/pgdg-centos8-sysupdates/redhat/rhel-8-x86_64/' ,europe: 'https://mirrors.xtom.de/postgresql/repos/yum/common/pgdg-centos8-sysupdates/redhat/rhel-8-x86_64/' }}
  - { name: pgdg-el9fix    ,description: 'PostgreSQL EL9FIX'  ,module: pgsql   ,releases: [  9   ] ,arch: [x86_64, aarch64] ,baseurl: { default: 'https://download.postgresql.org/pub/repos/yum/common/pgdg-rocky9-sysupdates/redhat/rhel-9-$basearch/'  ,china: 'https://repo.pigsty.cc/yum/pgdg/common/pgdg-rocky9-sysupdates/redhat/rhel-9-x86_64/'  ,europe: 'https://mirrors.xtom.de/postgresql/repos/yum/common/pgdg-rocky9-sysupdates/redhat/rhel-9-x86_64/'  }}
  - { name: pgdg14         ,description: 'PostgreSQL 14'      ,module: pgsql   ,releases: [8,9,10] ,arch: [x86_64, aarch64] ,baseurl: { default: 'https://download.postgresql.org/pub/repos/yum/14/redhat/rhel-$releasever-$basearch' ,china: 'https://repo.pigsty.cc/yum/pgdg/14/redhat/rhel-$releasever-$basearch' ,europe: 'https://mirrors.xtom.de/postgresql/repos/yum/14/redhat/rhel-$releasever-$basearch' }}
  - { name: pgdg15         ,description: 'PostgreSQL 15'      ,module: pgsql   ,releases: [8,9,10] ,arch: [x86_64, aarch64] ,baseurl: { default: 'https://download.postgresql.org/pub/repos/yum/15/redhat/rhel-$releasever-$basearch' ,china: 'https://repo.pigsty.cc/yum/pgdg/15/redhat/rhel-$releasever-$basearch' ,europe: 'https://mirrors.xtom.de/postgresql/repos/yum/15/redhat/rhel-$releasever-$basearch' }}
  - { name: pgdg16         ,description: 'PostgreSQL 16'      ,module: pgsql   ,releases: [8,9,10] ,arch: [x86_64, aarch64] ,baseurl: { default: 'https://download.postgresql.org/pub/repos/yum/16/redhat/rhel-$releasever-$basearch' ,china: 'https://repo.pigsty.cc/yum/pgdg/16/redhat/rhel-$releasever-$basearch' ,europe: 'https://mirrors.xtom.de/postgresql/repos/yum/16/redhat/rhel-$releasever-$basearch' }}
  - { name: pgdg17         ,description: 'PostgreSQL 17'      ,module: pgsql   ,releases: [8,9,10] ,arch: [x86_64, aarch64] ,baseurl: { default: 'https://download.postgresql.org/pub/repos/yum/17/redhat/rhel-$releasever-$basearch' ,china: 'https://repo.pigsty.cc/yum/pgdg/17/redhat/rhel-$releasever-$basearch' ,europe: 'https://mirrors.xtom.de/postgresql/repos/yum/17/redhat/rhel-$releasever-$basearch' }}
  - { name: pgdg18         ,description: 'PostgreSQL 18'      ,module: pgsql   ,releases: [8,9,10] ,arch: [x86_64, aarch64] ,baseurl: { default: 'https://download.postgresql.org/pub/repos/yum/18/redhat/rhel-$releasever-$basearch' ,china: 'https://repo.pigsty.cc/yum/pgdg/18/redhat/rhel-$releasever-$basearch' ,europe: 'https://mirrors.xtom.de/postgresql/repos/yum/18/redhat/rhel-$releasever-$basearch' }}
  - { name: pgdg-beta      ,description: 'PostgreSQL Testing' ,module: beta    ,releases: [8,9,10] ,arch: [x86_64, aarch64] ,baseurl: { default: 'https://download.postgresql.org/pub/repos/yum/testing/19/redhat/rhel-$releasever-$basearch' ,china: 'https://repo.pigsty.cc/yum/pgdg/testing/19/redhat/rhel-$releasever-$basearch' ,europe: 'https://mirrors.xtom.de/postgresql/repos/yum/testing/19/redhat/rhel-$releasever-$basearch' }}
  - { name: pgdg-extras    ,description: 'PostgreSQL Extra'   ,module: extra   ,releases: [8,9,10] ,arch: [x86_64, aarch64] ,baseurl: { default: 'https://download.postgresql.org/pub/repos/yum/common/pgdg-rhel$releasever-extras/redhat/rhel-$releasever-$basearch' ,china: 'https://repo.pigsty.cc/yum/pgdg/common/pgdg-rhel$releasever-extras/redhat/rhel-$releasever-$basearch' , europe: 'https://mirrors.xtom.de/postgresql/repos/yum/common/pgdg-rhel$releasever-extras/redhat/rhel-$releasever-$basearch' }}
  - { name: pgdg13-nonfree ,description: 'PostgreSQL 13+'     ,module: extra   ,releases: [8,9,10] ,arch: [x86_64         ] ,baseurl: { default: 'https://download.postgresql.org/pub/repos/yum/non-free/13/redhat/rhel-$releasever-$basearch' ,china: 'https://repo.pigsty.cc/yum/pgdg/non-free/13/redhat/rhel-$releasever-$basearch' ,europe: 'https://mirrors.xtom.de/postgresql/repos/yum/non-free/13/redhat/rhel-$releasever-$basearch' }}
  - { name: pgdg14-nonfree ,description: 'PostgreSQL 14+'     ,module: extra   ,releases: [8,9,10] ,arch: [x86_64         ] ,baseurl: { default: 'https://download.postgresql.org/pub/repos/yum/non-free/14/redhat/rhel-$releasever-$basearch' ,china: 'https://repo.pigsty.cc/yum/pgdg/non-free/14/redhat/rhel-$releasever-$basearch' ,europe: 'https://mirrors.xtom.de/postgresql/repos/yum/non-free/14/redhat/rhel-$releasever-$basearch' }}
  - { name: pgdg15-nonfree ,description: 'PostgreSQL 15+'     ,module: extra   ,releases: [8,9,10] ,arch: [x86_64         ] ,baseurl: { default: 'https://download.postgresql.org/pub/repos/yum/non-free/15/redhat/rhel-$releasever-$basearch' ,china: 'https://repo.pigsty.cc/yum/pgdg/non-free/15/redhat/rhel-$releasever-$basearch' ,europe: 'https://mirrors.xtom.de/postgresql/repos/yum/non-free/15/redhat/rhel-$releasever-$basearch' }}
  - { name: pgdg16-nonfree ,description: 'PostgreSQL 16+'     ,module: extra   ,releases: [8,9,10] ,arch: [x86_64         ] ,baseurl: { default: 'https://download.postgresql.org/pub/repos/yum/non-free/16/redhat/rhel-$releasever-$basearch' ,china: 'https://repo.pigsty.cc/yum/pgdg/non-free/16/redhat/rhel-$releasever-$basearch' ,europe: 'https://mirrors.xtom.de/postgresql/repos/yum/non-free/16/redhat/rhel-$releasever-$basearch' }}
  - { name: pgdg17-nonfree ,description: 'PostgreSQL 17+'     ,module: extra   ,releases: [8,9,10] ,arch: [x86_64         ] ,baseurl: { default: 'https://download.postgresql.org/pub/repos/yum/non-free/17/redhat/rhel-$releasever-$basearch' ,china: 'https://repo.pigsty.cc/yum/pgdg/non-free/17/redhat/rhel-$releasever-$basearch' ,europe: 'https://mirrors.xtom.de/postgresql/repos/yum/non-free/17/redhat/rhel-$releasever-$basearch' }}
  - { name: timescaledb    ,description: 'TimescaleDB'        ,module: extra   ,releases: [8,9   ] ,arch: [x86_64, aarch64] ,baseurl: { default: 'https://packagecloud.io/timescale/timescaledb/el/$releasever/$basearch'  }}
  - { name: percona        ,description: 'Percona TDE'        ,module: percona ,releases: [8,9   ] ,arch: [x86_64, aarch64] ,baseurl: { default: 'https://repo.pigsty.io/yum/percona/el$releasever.$basearch' ,china: 'https://repo.pigsty.cc/yum/percona/el$releasever.$basearch' ,origin: 'http://repo.percona.com/ppg-17.5/yum/release/$releasever/RPMS/$basearch'  }}
  - { name: wiltondb       ,description: 'WiltonDB'           ,module: mssql   ,releases: [8,9   ] ,arch: [x86_64, aarch64] ,baseurl: { default: 'https://repo.pigsty.io/yum/mssql/el$releasever.$basearch', china: 'https://repo.pigsty.cc/yum/mssql/el$releasever.$basearch' , origin: 'https://download.copr.fedorainfracloud.org/results/wiltondb/wiltondb/epel-$releasever-$basearch/' }}
  - { name: groonga        ,description: 'Groonga'            ,module: groonga ,releases: [8,9,10] ,arch: [x86_64, aarch64] ,baseurl: { default: 'https://packages.groonga.org/almalinux/$releasever/$basearch/' }}
  - { name: mysql          ,description: 'MySQL'              ,module: mysql   ,releases: [8,9   ] ,arch: [x86_64, aarch64] ,baseurl: { default: 'https://repo.mysql.com/yum/mysql-8.4-community/el/$releasever/$basearch/' }}
  - { name: mongo          ,description: 'MongoDB'            ,module: mongo   ,releases: [8,9   ] ,arch: [x86_64, aarch64] ,baseurl: { default: 'https://repo.mongodb.org/yum/redhat/$releasever/mongodb-org/8.0/$basearch/' ,china: 'https://mirrors.aliyun.com/mongodb/yum/redhat/$releasever/mongodb-org/8.0/$basearch/' }}
  - { name: redis          ,description: 'Redis'              ,module: redis   ,releases: [8,9,10] ,arch: [x86_64, aarch64] ,baseurl: { default: 'https://rpmfind.net/linux/remi/enterprise/$releasever/redis72/$basearch/' }}
  - { name: grafana        ,description: 'Grafana'            ,module: grafana ,releases: [8,9,10] ,arch: [x86_64, aarch64] ,baseurl: { default: 'https://rpm.grafana.com', china: 'https://mirrors.aliyun.com/grafana/yum/' }}
  - { name: kubernetes     ,description: 'Kubernetes'         ,module: kube    ,releases: [8,9,10] ,arch: [x86_64, aarch64] ,baseurl: { default: 'https://pkgs.k8s.io/core:/stable:/v1.33/rpm/', china: 'https://mirrors.aliyun.com/kubernetes-new/core/stable/v1.33/rpm/' }}
  - { name: gitlab-ee      ,description: 'Gitlab EE'          ,module: gitlab  ,releases: [8,9   ] ,arch: [x86_64, aarch64] ,baseurl: { default: 'https://packages.gitlab.com/gitlab/gitlab-ee/el/$releasever/$basearch' }}
  - { name: gitlab-ce      ,description: 'Gitlab CE'          ,module: gitlab  ,releases: [8,9   ] ,arch: [x86_64, aarch64] ,baseurl: { default: 'https://packages.gitlab.com/gitlab/gitlab-ce/el/$releasever/$basearch' }}
  - { name: clickhouse     ,description: 'ClickHouse'         ,module: click   ,releases: [8,9,10] ,arch: [x86_64, aarch64] ,baseurl: { default: 'https://packages.clickhouse.com/rpm/stable/', china: 'https://mirrors.aliyun.com/clickhouse/rpm/stable/' }}

# package alias map to simplify package download & installation
package_map:
  #--------------------------------#
  # REPO: download packages
  #--------------------------------#
  node-bootstrap:          "{{ getBootstrap }}"
  infra-package:           "{{ index .Constants.RPMCommonPkg 0 }}"
  infra-addons:            "{{ index .Constants.RPMCommonPkg 1 }}"
  extra-modules:           "{{ index .Constants.RPMCommonPkg 2 }}"
  node-package1:           "{{ index .Constants.RPMCommonPkg 3 }}"
  node-package2:           "{{ index .Constants.RPMCommonPkg 4 }}"
  pgsql-utility:           "{{ index .Constants.RPMCommonPkg 5 }}"

  #--------------------------------#
  # PGSQL: kernel packages
  #--------------------------------#
  postgresql:              "postgresql$v*"{{ range .Constants.PGSQLKernelMap }}
  {{ printf "%-24s" (printf "%s:" .Key) }} "{{ .RPM }}"{{ end }}

  #--------------------------------#
  # MISC: pro modules
  #--------------------------------#{{ range .Constants.PGSQLExoticMap }}{{ if and (or (eq .Key "greenplum") (eq .Key "gpsql")) (eq $.Arch "aarch64") }}{{ else if .RPM }}
  {{ printf "%-24s" (printf "%s:" .Key) }} "{{ .RPM }}"{{ end }}{{ end }}
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
  # PGSQL: tools & utils
  #--------------------------------#{{ range .Constants.PGSQLUtilMap }}
  {{ printf "%-24s" (printf "%s:" .Key) }} "{{ .RPM }}"{{ end }}

  #--------------------------------#
  # PGSQL: extensions (default pg18)
  #--------------------------------#
{{ range $cat := .Categories }}{{ $pkgs := index $.CategoryExts 18 $cat }}{{ if $pkgs }}  {{ printf "%-24s" (printf "pgsql-%s:" $cat) }} "{{ replaceAll "_18" "_$v" $pkgs.Visible }}"{{ if $pkgs.Hidden }} # {{ $pkgs.Hidden }}{{ end }}
{{ end }}{{ end }}

{{ range $pgVer := .PGVersions }}
  #--------------------------------#
  # PG{{ $pgVer }}: packages
  #--------------------------------#{{ range $.Constants.PGSQLKernelMap }}{{ if eq .Key "pgsql" }}
  {{ printf "%-24s" (printf "pg%d:" $pgVer) }} "{{ replaceAll "$v" (printf "%d" $pgVer) .RPM }}"{{ else }}
  {{ printf "%-24s" (printf "pg%d-%s:" $pgVer (trimPrefix "pgsql-" .Key)) }} "{{ replaceAll "$v" (printf "%d" $pgVer) .RPM }}"{{ end }}{{ end }}
{{ range $cat := $.Categories }}{{ $pkgs := index $.CategoryExts $pgVer $cat }}{{ if $pkgs }}  {{ printf "%-24s" (printf "pg%d-%s:" $pgVer $cat) }} "{{ $pkgs.Visible }}"{{ if $pkgs.Hidden }} # {{ $pkgs.Hidden }}{{ end }}
{{ end }}{{ end }}
{{ end }}
{{ $lastCat := "" }}{{ range .ExtMappings }}{{ if ne .Category $lastCat }}{{ $lastCat = .Category }}

  #--------------------------------#
  # {{ toUpper .Category }}: extension
  #--------------------------------#
{{ end }}  {{ printf "%-24s %-42s" (printf "%s:" .Name) (printf "\"%s\"" .Package) }} # {{ formatPGVersions .PGVers }}
{{ end }}

...`

// debTemplateV2 is the template for DEB-based distributions
const debTemplateV2 = `---
# License    :   AGPLv3 @ https://doc.pgsty.com/about/license
# Copyright  :   2018-2025  Ruohang Feng / Vonng (rh@vonng.com)
# {{ .OSName }} :   {{ if eq .OSCode "d11" }}Debian 11{{ else if eq .OSCode "d12" }}Debian 12{{ else if eq .OSCode "d13" }}Debian 13{{ else if eq .OSCode "u20" }}Ubuntu 20.04{{ else if eq .OSCode "u22" }}Ubuntu 22.04{{ else if eq .OSCode "u24" }}Ubuntu 24.04{{ end }} Compatible

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
  - { name: base           ,description: 'Ubuntu Basic'       ,module: node    ,releases: [      20,22,24] ,arch: [x86_64, aarch64] ,baseurl: { default: 'https://mirrors.edge.kernel.org/ubuntu/ ${distro_codename} main universe multiverse restricted'           ,china: 'https://mirrors.aliyun.com/ubuntu/ ${distro_codename} main restricted universe multiverse' }}
  - { name: updates        ,description: 'Ubuntu Updates'     ,module: node    ,releases: [      20,22,24] ,arch: [x86_64, aarch64] ,baseurl: { default: 'https://mirrors.edge.kernel.org/ubuntu/ ${distro_codename}-updates main universe multiverse restricted'   ,china: 'https://mirrors.aliyun.com/ubuntu/ ${distro_codename}-updates main restricted universe multiverse' }}
  - { name: backports      ,description: 'Ubuntu Backports'   ,module: node    ,releases: [      20,22,24] ,arch: [x86_64, aarch64] ,baseurl: { default: 'https://mirrors.edge.kernel.org/ubuntu/ ${distro_codename}-backports main universe multiverse restricted' ,china: 'https://mirrors.aliyun.com/ubuntu/ ${distro_codename}-backports main restricted universe multiverse' }}
  - { name: security       ,description: 'Ubuntu Security'    ,module: node    ,releases: [      20,22,24] ,arch: [x86_64, aarch64] ,baseurl: { default: 'https://mirrors.edge.kernel.org/ubuntu/ ${distro_codename}-security main universe multiverse restricted'  ,china: 'https://mirrors.aliyun.com/ubuntu/ ${distro_codename}-security main restricted universe multiverse' }}
  - { name: pgdg           ,description: 'PostgreSQL'         ,module: pgsql   ,releases: [11,12,13,20,22,24] ,arch: [x86_64, aarch64] ,baseurl: { default: 'https://apt.postgresql.org/pub/repos/apt/ ${distro_codename}-pgdg main', china: 'https://mirrors.aliyun.com/postgresql/repos/apt/ ${distro_codename}-pgdg main' }}
  - { name: grafana        ,description: 'Grafana'            ,module: grafana ,releases: [11,12,13,20,22,24] ,arch: [x86_64, aarch64] ,baseurl: { default: 'https://apt.grafana.com stable main' ,china: 'https://mirrors.aliyun.com/grafana-apt/ stable main' }}
  - { name: victoriametrics,description: 'VictoriaMetrics'    ,module: victoria,releases: [11,12,13,20,22,24] ,arch: [x86_64, aarch64] ,baseurl: { default: 'https://repo.victoriametrics.com/deb/ victoria-metrics main' }}
  - { name: kubernetes     ,description: 'Kubernetes'         ,module: kube    ,releases: [11,12,13,20,22,24] ,arch: [x86_64, aarch64] ,baseurl: { default: 'https://pkgs.k8s.io/core:/stable:/v1.33/deb/ /', china: 'https://mirrors.aliyun.com/kubernetes-new/core/stable/v1.33/deb/' }}
  - { name: gitlab-ce      ,description: 'GitLab CE'          ,module: gitlab  ,releases: [11,12,13,20,22,24] ,arch: [x86_64, aarch64] ,baseurl: { default: 'https://packages.gitlab.com/gitlab/gitlab-ce/${distro_name}/ ${distro_codename} main' }}
  - { name: gitlab-ee      ,description: 'GitLab EE'          ,module: gitlab  ,releases: [11,12,13,20,22,24] ,arch: [x86_64, aarch64] ,baseurl: { default: 'https://packages.gitlab.com/gitlab/gitlab-ee/${distro_name}/ ${distro_codename} main' }}
  - { name: clickhouse     ,description: 'ClickHouse'         ,module: click   ,releases: [11,12,13,20,22,24] ,arch: [x86_64, aarch64] ,baseurl: { default: 'https://packages.clickhouse.com/deb stable main' ,china: 'https://mirrors.aliyun.com/clickhouse/deb stable main' }}

# package alias map to simplify package download & installation
package_map:
  #--------------------------------#
  # REPO: download packages
  #--------------------------------#
  node-bootstrap:          "{{ getBootstrap }}"
  infra-package:           "{{ index .Constants.DEBCommonPkg 0 }}"
  infra-addons:            "{{ index .Constants.DEBCommonPkg 1 }}"
  extra-modules:           "{{ index .Constants.DEBCommonPkg 2 }}"
  node-package1:           "{{ index .Constants.DEBCommonPkg 3 }}"
  node-package2:           "{{ index .Constants.DEBCommonPkg 4 }}"
  pgsql-utility:           "{{ index .Constants.DEBCommonPkg 5 }}"

  #--------------------------------#
  # PGSQL: kernel packages
  #--------------------------------#
  postgresql:              "postgresql-$v postgresql-client-$v postgresql-plpython3-$v postgresql-plperl-$v postgresql-pltcl-$v postgresql-server-dev-$v"{{ range .Constants.PGSQLKernelMap }}
  {{ printf "%-24s" (printf "%s:" .Key) }} "{{ .DEB }}"{{ end }}

  #--------------------------------#
  # MISC: pro modules
  #--------------------------------#{{ range .Constants.PGSQLExoticMap }}{{ if and (or (eq .Key "greenplum") (eq .Key "cloudberry")) }}{{ else if .DEB }}
  {{ printf "%-24s" (printf "%s:" .Key) }} "{{ .DEB }}"{{ end }}{{ end }}
  java-runtime:            "openjdk-17-jdk"
  kafka:                   "kafka kafka-exporter"
  {{ if or (eq .OSCode "u22") (eq .OSCode "d12") }}kube-docker:             "containerd.io cri-dockerd"{{ else if or (eq .OSCode "d11") (eq .OSCode "u20") }}kube-docker:             "containerd.io"{{ end }}
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
  # PGSQL: tools & utils
  #--------------------------------#{{ range .Constants.PGSQLUtilMap }}
  {{ printf "%-24s" (printf "%s:" .Key) }} "{{ .DEB }}"{{ end }}

  #--------------------------------#
  # PGSQL: extensions (default pg17)
  #--------------------------------#
{{ range $cat := .Categories }}{{ $pkgs := index $.CategoryExts 17 $cat }}{{ if $pkgs }}  {{ printf "%-24s" (printf "pgsql-%s:" $cat) }} "{{ $pkgs.Visible | replaceAll "postgresql-17-" "postgresql-$v-" }}"{{ if $pkgs.Hidden }} # {{ $pkgs.Hidden }}{{ end }}
{{ end }}{{ end }}

{{ range $pgVer := .PGVersions }}
  #--------------------------------#
  # PG{{ $pgVer }}: packages
  #--------------------------------#{{ range $.Constants.PGSQLKernelMap }}{{ if eq .Key "pgsql" }}{{ if eq $pgVer 18 }}
  {{ printf "%-24s" (printf "pg%d:" $pgVer) }} "{{ replaceAll "$v" (printf "%d" $pgVer) .DEB | replaceAll "postgresql-client-18" "postgresql-client-18 libpq5=18*" }}"{{ else }}
  {{ printf "%-24s" (printf "pg%d:" $pgVer) }} "{{ replaceAll "$v" (printf "%d" $pgVer) .DEB }}"{{ end }}{{ else }}{{ if and (eq $pgVer 18) (hasPrefix .Key "pgsql") }}
  {{ printf "%-24s" (printf "pg%d-%s:" $pgVer (trimPrefix "pgsql-" .Key)) }} "{{ replaceAll "$v" (printf "%d" $pgVer) .DEB | replaceAll "postgresql-client-18" "postgresql-client-18 libpq5=18*" }}"{{ else }}
  {{ printf "%-24s" (printf "pg%d-%s:" $pgVer (trimPrefix "pgsql-" .Key)) }} "{{ replaceAll "$v" (printf "%d" $pgVer) .DEB }}"{{ end }}{{ end }}{{ end }}
{{ range $cat := $.Categories }}{{ $pkgs := index $.CategoryExts $pgVer $cat }}{{ if $pkgs }}  {{ printf "%-24s" (printf "pg%d-%s:" $pgVer $cat) }} "{{ $pkgs.Visible }}"{{ if $pkgs.Hidden }} # {{ $pkgs.Hidden }}{{ end }}
{{ end }}{{ end }}
{{ end }}
{{ $lastCat := "" }}{{ range .ExtMappings }}{{ if ne .Category $lastCat }}{{ $lastCat = .Category }}

  #--------------------------------#
  # {{ toUpper .Category }}: extension
  #--------------------------------#
{{ end }}  {{ printf "%-24s %-42s" (printf "%s:" .Name) (printf "\"%s\"" .Package) }} # {{ formatPGVersions .PGVers }}
{{ end }}

...`