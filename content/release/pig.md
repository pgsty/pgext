---
title: PIG Releases
description: The pig releasenote and changelog
weight: 100
breadcrumbs: false
---

## v0.7.2

- Extension list update, + 6 new extensions, 437 total  
- Add PGDG EL10 Sysupdate repo
- Add LLVM APT repo
- use local extension.csv catalog in pig build sub command
- update: vchord pg_later pgvectorscale pglite_fusion pgx_ulid pg_search citus timescaledb pg_profile pg_stat_monitor documentdb
- new: pglinter pg_typeid pg_enigma pg_retry pg_biscuit pg_weighted_statistics

**Checksums**

```bash
f303c391fc28bc74832712e0aa58319abe0ebcae4f6c07fdf9a9e542b735d2ec  pig-0.7.2-1.aarch64.rpm
c096a61a4e3a49b1238659664bbe2cd7f29954c43fb6bb8e8e9fb271f95a612e  pig-0.7.2-1.x86_64.rpm
5e037c891dff23b46856485108d6f64bede5216dfbd4f38a481f0d0672ee910b  pig-v0.7.2.darwin-amd64.tar.gz
736b4b47999c543c3c886781f4d8dddbf4276f363c35c7bf50094b6f18d14600  pig-v0.7.2.darwin-arm64.tar.gz
20b13f059efed29dd76f6927b3e8d7b597c0c8d734f9e22ba3d0a2af6dbcd3bf  pig-v0.7.2.linux-amd64.tar.gz
9548b530c05f2ffdc8d73b8f890718d47b74a51eb62852a99c08b1b52e47f014  pig-v0.7.2.linux-arm64.tar.gz
b6faad9f92b926546a10f590274f2cb2afff21b9cea878094cfc5caf09e67d2c  pig_0.7.2-1_amd64.deb
452f73f1fa035e5417ab49fc51d797925550179ffcc023e8f03d80144309212a  pig_0.7.2-1_arm64.deb
```


## v0.7.1

- The brand-new website: https://pgext.cloud
- remove unnecessary sudo usage, now can be used inside docker
- allow using `pg18`, `pg17` arg format in pig ext link command
- add environment var `PIG_NO_SUDO` to force not using sudo
- [RPM Changelog](/release/rpm#2025-11-10): Add PG 18 support to almost all extensions
- [DEB Changelog](/release/deb#2025-11-10): Add PG 18 support to almost all extensions
- [Infra Changelog](/release/infra#2025-11-08): Routine update to the latest version

**Checksums**

```bash
a696c9ec784e2fc248e5f3d87cc8aae4116e890f78c5997957d30593f2c85ca6  pig-0.7.1-1.aarch64.rpm
f669538a99cd1dc592d3005b949628fcceb9e78114fc78862d7726b340ee194d  pig-0.7.1-1.x86_64.rpm
e42bdaaf93b720c5b76b32b57362320e4b447109740c76089aefe030b7c8b836  pig-v0.7.1.darwin-amd64.tar.gz
b4c240aadad34e785666ee0a755d9b7455724f790c2d088a1dd7c37ad3b2a457  pig-v0.7.1.darwin-arm64.tar.gz
ffc687add0ca71ac90cba5749c8a7a6075cf7618cba85584072831cf3eb182f7  pig-v0.7.1.linux-amd64.tar.gz
7b0d1f158150d0a40c525692f02b6bce9f5b4ac523a4e59278d702c334e222e1  pig-v0.7.1.linux-arm64.tar.gz
43e91a3bea273d7cacb2d7a58c0a5745501dbd06348b5cb3af971171fae70268  pig_0.7.1-1_amd64.deb
fc2a34aeb46e07cb0ae93611de47d6622c3bd46fe4c415ce4c9091840e0e08a2  pig_0.7.1-1_arm64.deb
```


## v0.7.0

- Bump tons of extension to the latest version, with pg 18 support.
- Most rust extension are bump to the latest pgrx 0.16.1
- `pig build` overhaul
    - `pig build pkg <pkg>` will download source, prepare deps, and build rpm/deb
    - `pig build pgrx` is separated from `pig build rust` with its own subcommand
    - `pig build pgrx [-v pgrx_version]` now can scan existing pg and init with them
    - `pig build dep` will now process EL rpm spec dependencies directly
    - `pig build ext` will now build EL rpm spec directly without the `build` script
    - `pig build spec` now support downloading spec files directly from pigsty repo
    - `pig build repo` / `pig repo add` / `pig repo set` will now using `node,pgsql,infra` as default repo modules instead of previous `node,pgdg,pigsty`
- Optimized error logging
- Reworked catalog website based on hugo and hextra

**Checksums**

```bash
ad60f9abcde954769e46eb23de61965e  pig_0.7.0-1_amd64.deb
aa15d7088d561528e38b2778fe8f7cf9  pig_0.7.0-1_arm64.deb
05549fe01008e04f8d5a59d4f2a5f0b8  pig-0.7.0-1.aarch64.rpm
0cc9e46c7c72d43c127a6ad115873b67  pig-0.7.0-1.x86_64.rpm
ddacfb052f3f3e5567a02e92fdb31cdd  pig-v0.7.0.darwin-amd64.tar.gz
17d25b565308d3d35513e4b0d824946b  pig-v0.7.0.darwin-arm64.tar.gz
ee7e055ceff638039956765fb747f80b  pig-v0.7.0.linux-amd64.tar.gz
284e674807b87447d4b33691fd7a420d  pig-v0.7.0.linux-arm64.tar.gz
```

## v0.6.2

- Use official PG 18 repo instead of testing repo
- Add `v` prefix when specify pigsty version string
- Improved network connectivity check

**Checksums**

```bash
01f5b7dc20644226c762dbb229768347  pig_0.6.2-1_amd64.deb
ce4f00256adc12cbea91467b7f2241cd  pig_0.6.2-1_arm64.deb
cefc36ae8f348aede533b30836fba720  pig-0.6.2-1.aarch64.rpm
d04a287c6eb92b11ecbf99542c2db602  pig-0.6.2-1.x86_64.rpm
e637ca86a7f38866c67686b060223d9a  pig-v0.6.2.darwin-amd64.tar.gz
79749bc69c683586bd8d761bdf6af98e  pig-v0.6.2.darwin-arm64.tar.gz
ad4f02993c7d7d8eec142f0224551bb4  pig-v0.6.2.linux-amd64.tar.gz
9793affa4a0cb60e9753e65b7cba3dca  pig-v0.6.2.linux-arm64.tar.gz
```

## v0.6.1

- Add el10 and debian 13 trixie support stub
- Dedicate website: https://pig.pgsty.com
- rebuild with go 1.25 and CI/CD pipeline
- Use the PIGSTY mirror in mainland china due to pgdg ftp rsync mirror break
- Remove unused `pgdg-el10fix` repo
- Use pigsty mirror for WiltonDB
- Add EL 10 dedicate epel repo (with minor suffix)
- pig version output with go build environment

**Checksums**

```bash
871d2f3abb90afd77e943a126e917997  pig_0.6.1-1_amd64.deb
0ffbe364c9a64e997e87a7dd1937d37f  pig_0.6.1-1_arm64.deb
8410713cb946be5a0cfde078d375c5b5  pig-0.6.1-1.aarch64.rpm
1a28b44dc53c3e5052f2227a7d76b860  pig-0.6.1-1.x86_64.rpm
d02239f82c1bcf1674ec16a25d62d3fe  pig-v0.6.1.darwin-amd64.tar.gz
72e850ce163476ddbf9c6404a624086f  pig-v0.6.1.darwin-arm64.tar.gz
2b2fdbf00b610dea648d2212452316da  pig-v0.6.1.linux-amd64.tar.gz
060ab4314b873832cce8c3602a95579b  pig-v0.6.1.linux-arm64.tar.gz
```


## v0.6.0

- New extension catalog: [https://ext.pgsty.com](https://ext.pgsty.com)
- New subcommand: `pig install` to simplify `pig ext install`
- Add new kernel support: percona with pg_tde
- Add new package: Google GenAI MCP toolbox for databases
- Add new repo: percona repo and clickhouse repo
- Change extension summary info links to https://ext.pgsty.com
- fix orioledb broken on the Debian/Ubuntu system
- fix epel repo on EL distributions
- Bump golang to 1.24.5
- Bump pigsty to v3.6.0

**Checksums**

```bash
1804766d235b9267701a08f95903bc3b  pig_0.6.0-1_amd64.deb
35f4efa35c1eaecdd12aa680d29eadcb  pig_0.6.0-1_arm64.deb
b523b54d9f2d7dcc5999bcc6bd046b1d  pig-0.6.1-1.aarch64.rpm
9434d9dca7fd9725ea574c5fae1a7f52  pig-0.6.1-1.x86_64.rpm
f635c12d9ad46a779aa7174552977d11  pig-v0.6.0.linux-amd64.tar.gz
165af4e63ec0031d303fe8b6c35c5732  pig-v0.6.0.linux-arm64.tar.gz
```


## v0.5.0

- Update the extension list to 422
- New extension: [pgactive](https://github.com/aws/pgactive) from AWS
- bump timescaledb to 2.20.3
- bump citus to 13.1.0
- bump vchord to 0.4.3
- bug fix pgvectorscale debian/ubuntu pg17 failure
- bump kubernetes repo to 1.33
- bump default pigsty version to 3.5.0

**Checksums**

```
9ec6f3caf3edbe867caab5de0e0ccb33  pig_0.5.0-1_amd64.deb
4fbb0a42cd8a88bce50b3c9d85745d77  pig_0.5.0-1_arm64.deb
9cf8208396b068cab438f72c90d39efe  pig-0.5.0-1.aarch64.rpm
d9a8d78c30f45e098b29c3d16471aa8d  pig-0.5.0-1.x86_64.rpm
761df804ff7b83965c41492700717674  pig-v0.5.0.linux-amd64.tar.gz
5d1830069d98030728f08835f883ea39  pig-v0.5.0.linux-arm64.tar.gz
```

Release: https://github.com/pgsty/pig/releases/tag/v0.6.0


## v0.4.2

- Update the extension list to 421
- Add openhalo/orioledb support for Debian / Ubuntu
- pgdd [0.6.0](https://github.com/rustprooflabs/pgdd) (pgrx 0.14.1)
- convert [0.0.4](https://github.com/rustprooflabs/convert) (pgrx 0.14.1)
- pg_idkit [0.3.0](https://github.com/VADOSWARE/pg_idkit) (pgrx 0.14.1)
- pg_tokenizer.rs [0.1.0](https://github.com/tensorchord/pg_tokenizer.rs) (pgrx 0.13.1)
- pg_render [0.1.2](https://github.com/mkaski/pg_render) (pgrx 0.12.8)
- pgx_ulid [0.2.0](https://github.com/pksunkara/pgx_ulid) (pgrx 0.12.7)
- pg_ivm [1.11.0](https://github.com/sraoss/pg_ivm) for debian/ubuntu
- orioledb [1.4.0 beta11](https://github.com/orioledb/orioledb)
- Add el7 repo back

**Checksums**

```bash
bbf83fa3e3ec9a4dca82eeed921ae90a  pig_0.4.2-1_amd64.deb
e45753335faf80a70d4f2ef1d3100d72  pig_0.4.2-1_arm64.deb
966d60bbc2025ba9cc53393011605f9f  pig-0.4.2-1.aarch64.rpm
1f31f54da144f10039fa026b7b6e75ad  pig-0.4.2-1.x86_64.rpm
1eec26c4e69b40921e209bcaa4fe257a  pig-v0.4.2.linux-amd64.tar.gz
768d43441917a3625c462ce9f2b9d4ef  pig-v0.4.2.linux-arm64.tar.gz
```

Release: https://github.com/pgsty/pig/releases/tag/v0.4.2


## v0.4.1

- Update the extension list to 414
- add `citus_wal2json` and `citus_pgoutput` to `pig ext scan` mapping
- Add PG 18 beta repo
- Add PG 18 package alias

**Extension Package Updates**

- omnigres 20250507
- citus [12.0.3](https://github.com/citusdata/citus/releases/tag/v13.0.3)
- timescaledb [2.19.3](https://github.com/timescale/timescaledb/releases/tag/2.19.3)
- supautils [2.9.1](https://github.com/supabase/supautils/releases/tag/v2.9.1)
- pg_envvar [1.0.1](https://github.com/theory/pg-envvar/releases/tag/v1.0.1)
- pgcollection [1.0.0](https://github.com/aws/pgcollection/releases/tag/v1.0.0)
- aggs_for_vecs [1.4.0](https://github.com/pjungwir/aggs_for_vecs/releases/tag/1.4.0)
- pg_tracing [0.1.3](https://github.com/DataDog/pg_tracing/releases/tag/v0.1.3)
- pgmq [1.5.1](https://github.com/pgmq/pgmq/releases/tag/v1.5.1)
- tzf-pg [0.2.0](https://github.com/ringsaturn/tzf-pg/releases/tag/v0.2.0) (pgrx 0.14.1)
- pg_search [0.15.18](https://github.com/paradedb/paradedb/releases/tag/v0.15.18) (pgrx 0.14.1)
- anon [2.1.1](https://gitlab.com/dalibo/postgresql_anonymizer/-/tree/latest/debian?ref_type=heads) (pgrx 0.14.1)
- pg_parquet [0.4.0](https://github.com/CrunchyData/pg_parquet/releases/tag/v0.3.2) (0.14.1)
- pg_cardano [1.0.5](https://github.com/Fell-x27/pg_cardano/commits/master/) (pgrx 0.12) -> 0.14.1
- pglite_fusion [0.0.5](https://github.com/frectonz/pglite-fusion/releases/tag/v0.0.5) (pgrx 0.12.8) -> 14.1
- vchord_bm25 [0.2.1](https://github.com/tensorchord/VectorChord-bm25/releases/tag/0.2.1) (pgrx 0.13.1)
- vchord [0.3.0](https://github.com/tensorchord/VectorChord/releases/tag/0.3.0) (pgrx 0.13.1)
- pg_vectorize [0.22.1](https://github.com/ChuckHend/pg_vectorize/releases/tag/v0.22.1) (pgrx 0.13.1)
- wrappers [0.4.6](https://github.com/supabase/wrappers/releases/tag/v0.4.6) (pgrx 0.12.9)
- timescaledb-toolkit [1.21.0](https://github.com/timescale/timescaledb-toolkit/releases/tag/1.21.0) (0.12.9)
- pgvectorscale [0.7.1](https://github.com/timescale/pgvectorscale/releases/tag/0.7.1) (pgrx 0.12.9)
- pg_session_jwt [0.3.1](https://github.com/neondatabase/pg_session_jwt/releases/tag/v0.3.1) (pgrx 0.12.6) -> 0.12.9

**Checksums**

```
e2c1037c20f97c6f5930876ee82b6392  pig_0.4.1-1_amd64.deb
8197b6b5b95d1d1ae95e0a0e50355ecb  pig_0.4.1-1_arm64.deb
9d3a261d31c92fc73fe5bbfcd5b8e8ba  pig-0.4.1-1.aarch64.rpm
ffcec2a2ae965d14b9d3d80278fd340c  pig-0.4.1-1.x86_64.rpm
01d3128e782f35a20f0c81480cbe9025  pig-v0.4.1.linux-amd64.tar.gz
b2655628df326a1d0ed13f3dd8762c65  pig-v0.4.1.linux-arm64.tar.gz
```

Release: https://github.com/pgsty/pig/releases/tag/v0.4.1


## v0.4.0

- Updated extension list, available extensions reached **407**
- Added `pig do` subcommand for executing Pigsty playbook tasks
- Added `pig pt` subcommand for wrapping Patroni command-line tools
- Added extension aliases: `openhalo` and `orioledb`
- Added `gitlab-ce` / `gitlab-ee` repository distinction
- Built with the latest Go 1.24.2 and upgraded dependency versions
- Fixed `pig ext status` panic issue under specific conditions
- Fixed `pig ext scan` unable to match several extensions

**Extension Package Updates**

- Updated pg_search to 0.15.13
- Updated citus to 13.0.3
- Updated timescaledb to 2.19.1
- Updated pgcollection RPM to 1.0.0
- Updated pg_vectorize RPM to 0.22.1
- Updated pglite_fusion RPM to 0.0.4
- Updated aggs_for_vecs RPM to 1.4.0
- Updated pg_tracing RPM to 0.1.3
- Updated pgmq RPM to 1.5.1

**Checksums**

```
bbc0adf94b342ac450c7999ea1c5ab76  pig_0.4.0-1_amd64.deb
7445b819624e7498b496edb12a36f426  pig_0.4.0-1_arm64.deb
835ce929afac0fb1f249f55571fbed97  pig-0.4.0-1.aarch64.rpm
25ba5a846095e17d2bfa2f15fe4e4b44  pig-0.4.0-1.x86_64.rpm
1568b163ffa23cb921ee439452ca4de9  pig-v0.4.0.linux-amd64.tar.gz
9f2ab3f5d1e29807a9642dfbe1dc9b0e  pig-v0.4.0.linux-arm64.tar.gz
```

Release: https://github.com/pgsty/pig/releases/tag/v0.4.0


## v0.3.4

```bash
curl https://repo.pigsty.io/pig | bash -s 0.3.4
```

- routine extension metadata update
- use aliyun epel mirror instead of broken tsinghua tuna mirror
- bump pigsty version string
- add `gitlab` repo to the repo list

**Checksums**

```
5c0bba04d955bbe6a29d24d31aa17c6b  pig-0.3.4-1.aarch64.rpm
42636b9fc64d7882391d856d36d715e7  pig-0.3.4-1.x86_64.rpm
1a6296421d642000ad75a5a41bc9ab96  pig-v0.3.4.linux-amd64.tar.gz
f7ea5ba8abaa89e866811e5b2508e82f  pig-v0.3.4.linux-arm64.tar.gz
2dd63cdb5965f78a48da462a0453001d  pig_0.3.4-1_amd64.deb
094b9e028e81c46d71ee315d8a223ada  pig_0.3.4-1_arm64.deb
```

Release: https://github.com/pgsty/pig/releases/tag/v0.3.4


## v0.3.3

- add `pig build dep` command to install extension build dependencies
- update default repo list
- use pigsty.io mirror for `mssql` module (wiltondb/babelfish)
- merge docker module into `infra`
- remove pg16/17 from el7 target
- allow installing extensions in el7
- update package alias
- `pgsql`, `pgsql-main`, `pgsql-core`, `pgsql-mini`, `pgsql-full`
- `ivorysql` now maps to `ivorysql4`
- `timescaledb-utils`
- `pgbackrest_exporter`
- remove `pgsql-simple`
- Pulls [#13 Bump github.com/golang-jwt/jwt/v5 from 5.2.1 to 5.2.2](https://github.com/pgsty/pig/pull/13)
- Bump polardb to 15.12.3.0-e1e6d85b
- `pig repo set` now will auto update the meta-cache
- clean up embedded pigsty tarball

**What’s Changed**

- Bump github.com/golang-jwt/jwt/v5 from 5.2.1 to 5.2.2 by @dependabot in https://github.com/pgsty/pig/pull/13

**New Contributors**

- @dependabot made their first contribution in https://github.com/pgsty/pig/pull/13

**Full Changelog**: https://github.com/pgsty/pig/compare/v0.3.2...v0.3.3

Release: https://github.com/pgsty/pig/releases/tag/v0.3.3

**Checksums**

```
4e10567077e5d8cefd94d1c7aeb9478b  pig-0.3.3-1.aarch64.rpm
cc8a423abeb0f5316b427097993b9c6e  pig-0.3.3-1.x86_64.rpm
835d4f63b4ee0b36e2322a4ffef6527a  pig-v0.3.3.linux-amd64.tar.gz
c43e082c661e75d91f1c726e60911ea3  pig-v0.3.3.linux-arm64.tar.gz
938db83c5ca065419b8185adb285ed5a  pig_0.3.3-1_amd64.deb
75af6731adc4d31aa3458d70fc7f4e42  pig_0.3.3-1_arm64.deb
```

Release: https://github.com/pgsty/pig/releases/tag/v0.3.3


## v0.3.2

**Enhancement**

- new extensions
- use `upx` to reduce binary size
- remove embeded pigsty to reudce binary size
- allow specifiy `-y` to force re-install rust in `pig build rust`
- allow specifiy `-v` to specifiy `pgrx` version to be installed

**New Extensions**

List of **405** PG extensions:

- apache age 13 - 17 el rpm (1.5.0)
- pgspider_ext 1.3.0 (new extension)
- timescaledb 2.18.2 -> 2.19.0
- citus 13.0.1 -> 13.0.2
- documentdb 1.101-0 -> 1.102-0
- pg_analytics: 0.3.4 -> 0.3.7
- pg_search: 0.15.2 -> 0.15.8
- pg_ivm 1.9 -> 1.10
- emaj 4.4.0 -> 4.6.0
- pgsql_tweaks 0.10.0 -> 0.11.0
- pgvectorscale 0.4.0 -> 0.6.0 (pgrx 0.12.5)
- pg_session_jwt 0.1.2 -> 0.2.0 (pgrx 0.12.6)
- wrappers 0.4.4 -> 0.4.5 (pgrx 0.12.9)
- pg_parquet 0.2.0 -> 0.3.1 (pgrx 0.13.1)
- vchord 0.2.1 -> 0.2.2 (pgrx 0.13.1)
- pg_tle 1.2.0 -> 1.5.0
- supautils 2.5.0 -> 2.6.0
- sslutils 1.3 -> 1.4
- pg_profile 4.7 -> 4.8
- pg_snakeoil 1.3 -> 1.4
- pg_jsonschema 0.3.2 -> 0.3.3
- pg_incremental: 1.1.1 -> 1.2.0
- pg_stat_monitor 2.1.0 -> 2.1.1
- fix ddl_historization ver 0.7 -> 0.0.7
- fix pg_sqlog 3.1.7 -> 1.6
- fix pg_random remove dev suffix
- asn1oid 1.5 -> 1.6
- table_log 0.6.1 -> 0.6.4

**Checksums**

```
f773aedf4a76d031f411cb38bc623134  pig-0.3.2-1.aarch64.rpm
fa9084877deb57d4882b7d9531ea0369  pig-0.3.2-1.x86_64.rpm
7f9a03c9dd23cba094191a8044fa0263  pig-v0.3.2.linux-amd64.tar.gz
adda8986efc048565834cda1ef206a20  pig-v0.3.2.linux-arm64.tar.gz
5b27cefdc716629db8f1fbc534f58691  pig_0.3.2-1_amd64.deb
936e85bda5818da4c20b758ebd65e618  pig_0.3.2-1_arm64.deb
```

Release: https://github.com/pgsty/pig/releases/tag/v0.3.2


## v0.3.1

Routine bugfix

- fix repo format string
- fix ext info links
- update pg_mooncake metadata

**Checksums**

```
9251aa18e663f1ecf239adcba3a798b9  pig-0.3.1-1.aarch64.rpm
3b91e7faa78c5f0283d27ffe632dda46  pig-0.3.1-1.x86_64.rpm
87c75dfd114252230c53ee8c5d60dac4  pig-v0.3.1.linux-amd64.tar.gz
82832ae767e226627087b97a87982daf  pig-v0.3.1.linux-arm64.tar.gz
4d99f9c03915accf413b6374b75f1bdb  pig_0.3.1-1_amd64.deb
e38e8a21ed73a37d4588053f8c900f7c  pig_0.3.1-1_arm64.deb
```

Release: https://github.com/pgsty/pig/releases/tag/v0.3.1


## v0.3.0

The [`pig`](/pig/) project now has a new [homepage](https://pig.pgsty.com), alone with the PostgreSQL Extension [Catalog](https://ext.pgsty.com/list).

```bash
curl https://repo.pigsty.io/pig | bash    # cloudflare
curl https://repo.pigsty.cc/pig | bash    # china cdn
```

You can install PostgreSQL Kernels along with & [**404 extensions**](https://ext.pgsty.com/list) with a simple command. Besides,
pig v0.3 is also embedded & shipped with the latest Pigsty [v3.3.0](https://doc.pgsty.com/release).

**New Features**

`pig build` subcommand with the [ability](https://pig.pgsty.com/cmd/build/) to set up extension building environment

```bash
pig build repo     # init build repo (=repo set -ru)
pig build tool     # init build toolset
pig build rust     # init rustc & pgrx (0.12.9)
pig build spec     # init rpm/deb spec repo
pig build get      # get extension src tarball
pig build ext      # build extension
## download big tarball
pig build get std          # download std small tarball
pig build get all          # download all source tarball
pig build get pg_mooncake
pig build get pg_duckdb
pig build get omnigres
pig build get plv8
pig build get citus

pig build ext citus
pig build ext timescaledb
```

And other utils such as building proxy:

```bash
pig build proxy                  # install v2ray proxy
pig build proxy [user@host:port] # init & setup proxy
```

And `pig` 0.3.0 is shipped with Pigsty [3.3.0](https://github.com/Vonng/pigsty/releases/tag/v3.3.0)

**New Extensions**

The [`ext.pgsty.com`](https://ext.pgsty.com/) catalog is moving to [`https://ext.pgsty.com/list`](https://ext.pgsty.com/ext) with more information!

[![ecosystem](/img/pigsty/ecosystem.gif)](https://pigsty.io/about/ext)

**Checksums**

```bash
9cc3848ab13c41a0415f1fea6294ad2d  pig-0.3.0-1.aarch64.rpm
ee99a6c1ff17975ed184f009a4b1aac5  pig-0.3.0-1.x86_64.rpm
b06f6b5aeaa83a9d76c9b563b2516e1c  pig-v0.3.0.linux-amd64.tar.gz
d783732413e4f32074adeab2d5d092c3  pig-v0.3.0.linux-arm64.tar.gz
7c942b8dbd78458d5371c1abca2571c6  pig_0.3.0-1_amd64.deb
c0a411cf53cb58706ca81b49b4fc840e  pig_0.3.0-1_arm64.deb
```

Release: https://github.com/pgsty/pig/releases/tag/v0.3.0


## v0.2.2

[**404**](/list) Extensions Available in Pig v0.2.2

```bash
curl https://repo.pigsty.io/pig | bash -s v0.2.2
```

- **documentdb** 0.101-0
- **pgcollection** (new) 0.9.1
- **pg_bzip** (new) 1.0.0
- pg_net 0.14.0 (some distro)
- pg_curl 2.4.2
- vault 0.3.1 (SQL -> C)
- table_version 1.10.3 -> 1.11.0
- pg_duration 1.0.2
- timescaledb 2.18.2
- pg_analytics 0.3.4
- pg_search 0.15.2
- pg_graphql 1.5.11
- vchord 0.1.1 -> 0.2.1 ((+13))
- vchord_bm25 0.1.0 -> 0.1.1
- **pg_mooncake** 0.1.1 -> 0.1.2
- **pg_duckdb** 0.2.0 -> 0.3.1
- pgddl 0.29
- pgsql_tweaks 0.11.0

Release: https://github.com/pgsty/pig/releases/tag/v0.2.2


## v0.2.0

Install the latest `pig` version with:

```bash
curl -fsSL https://repo.pigsty.io/pig | bash
```

**New Extensions**

- [pg_documentdb_core](https://github.com/microsoft/documentdb/tree/main/pg_documentdb_core), and ferretdb
- [VectorChord-bm25](https://github.com/tensorchord/VectorChord-bm25) (vchord_bm25) 0.1.0
- [pg_tracing](https://github.com/DataDog/pg_tracing) 0.1.2
- [pg_curl](https://github.com/RekGRpth/pg_curl) 2.4
- [pgxicor](https://github.com/Florents-Tselai/pgxicor) 0.1.0
- [pgsparql](https://github.com/lacanoid/pgsparql) 1.0
- [pgjq](https://github.com/Florents-Tselai/pgJQ) 0.1.0
- [hashtypes](https://github.com/adjust/hashtypes/) 0.1.5
- [db_migrator](https://github.com/cybertec-postgresql/db_migrator) 1.0.0
- [pg_cooldown](https://github.com/rbergm/pg_cooldown) 0.1

**Update Extension Version**

- citus 13.0.0 -> 13.0.1
- pg_mooncake 0.1.0 -> 0.1.1
- timescaledb 2.17.2 -> 2.18.1
- supautils 2.5.0 -> 2.6.0
- VectorChord 0.1.0 -> 0.2.0
- pg_bulkload 3.1.22 (+pg17)
- pg_store_plan 1.8 (+pg17)
- pg_search 0.14 -> 0.15.1
- pg_analytics 0.3.0 -> 0.3.2
- pgroonga 3.2.5 -> 4.0.0
- zhparser 2.2 -> 2.3
- pg_vectorize 0.20.0 -> 0.21.1

Release: https://github.com/pgsty/pig/releases/tag/v0.2.0


## v0.1.4

Install the latest `pig` version with:

```bash
curl -fsSL https://repo.pigsty.io/pig | bash
```

**New Extensions**

- [pg_documentdb_core](https://github.com/microsoft/documentdb/tree/main/pg_documentdb_core), and ferretdb
- [VectorChord-bm25](https://github.com/tensorchord/VectorChord-bm25) (vchord_bm25) 0.1.0
- [pg_tracing](https://github.com/DataDog/pg_tracing) 0.1.2
- [pg_curl](https://github.com/RekGRpth/pg_curl) 2.4
- [pgxicor](https://github.com/Florents-Tselai/pgxicor) 0.1.0
- [pgsparql](https://github.com/lacanoid/pgsparql) 1.0
- [pgjq](https://github.com/Florents-Tselai/pgJQ) 0.1.0
- [hashtypes](https://github.com/adjust/hashtypes/) 0.1.5
- [db_migrator](https://github.com/cybertec-postgresql/db_migrator) 1.0.0
- [pg_cooldown](https://github.com/rbergm/pg_cooldown) 0.1

**Update Extension Version**

- citus 13.0.0 -> 13.0.1
- pg_mooncake 0.1.0 -> 0.1.1
- timescaledb 2.17.2 -> 2.18.1
- supautils 2.5.0 -> 2.6.0
- VectorChord 0.1.0 -> 0.2.0
- pg_bulkload 3.1.22 (+pg17)
- pg_store_plan 1.8 (+pg17)
- pg_search 0.14 -> 0.15.1
- pg_analytics 0.3.0 -> 0.3.2
- pgroonga 3.2.5 -> 4.0.0
- zhparser 2.2 -> 2.3
- pg_vectorize 0.20.0 -> 0.21.1

**Checksums**

```
6da06705be1c179941327c836d455d35  pig-0.1.4-1.aarch64.rpm
9fa5712e3cfe56e0dcf22a11320b01b1  pig-0.1.4-1.x86_64.rpm
af506dc37f955a7a2e31ff11e227450c  pig-v0.1.4.linux-amd64.tar.gz
1e6eb3dc1ad26f49b07afabdd9142d4e  pig-v0.1.4.linux-arm64.tar.gz
83ae89b58bff003da5c3022eeac1786e  pig_0.1.4_amd64.deb
d6778e628d82bddf3fae1e058e1e05e4  pig_0.1.4_arm64.deb
```

Release: https://github.com/pgsty/pig/releases/tag/v0.1.4


## v0.1.3

v0.1.3, routine update, with 390 extensions available now!

- New Extension(s): [`Omnigres`](https://ext.pgsty.com/e/omni) 33 extensions, postgres as platform
- New Extension: [`pg_mooncake`](https://ext.pgsty.com/e/pg_mooncake): duckdb in postgres
- New Extensions: [`pg_xxhash`](https://ext.pgsty.com/e/xxhash)
- New Extension: [`timescaledb_toolkit`](https://ext.pgsty.com/e/timescaledb_toolkit)
- New Extension: [`pg_xenophile`](https://ext.pgsty.com/e/pg_xenophile)
- New Extension: [`pg_drop_events`](https://ext.pgsty.com/e/pg_drop_events)
- New Extension: [`pg_incremental`](https://ext.pgsty.com/e/pg_incremental)
- Bump [`citus`](https://github.com/citusdata/citus/tree/v13.0.0) to 13.0.0 with PostgreSQL 17 support.
- Bump [`pgml`](https://github.com/postgresml/postgresml/releases/tag/v2.10.0) to 2.10.0
- Bump [`pg_extra_time`](https://ext.pgsty.com/e/pg_extra_time) to 2.0.0
- Bump [`pg_vectorize`](https://ext.pgsty.com/e/pg_vectorize) to 0.20.0

```bash
curl https://repo.pigsty.io/pig | bash
curl https://repo.pigsty.cc/pig | bash
```

**Checksums**

```
c79b74f676b03482859f5519b279b657  pig-0.1.3-1.aarch64.rpm
1d00a7cd5855a65e4db964075a5e49f6  pig-0.1.3-1.x86_64.rpm
6cd8507b130fca093247278e36d9478b  pig-v0.1.3.linux-amd64.tar.gz
5eee92908701b0d456ec3c15bc817c0b  pig-v0.1.3.linux-arm64.tar.gz
cb376ef2c3512ad35ff43132942c0052  pig_0.1.3_amd64.deb
2b545abc617670a96c2edd13878e0227  pig_0.1.3_arm64.deb
```

Release: https://github.com/pgsty/pig/releases/tag/v0.1.3


## v0.1.2

[**351**](https://ext.pgsty.com/list) PostgreSQL Extensions, including the powerful [postgresql-anonymizer 2.0](https://postgresql-anonymizer.readthedocs.io/en/stable/)

Now you can install pig with:

```bash
curl -fsSL https://repo.pigsty.io/pig | bash
curl -fsSL https://repo.pigsty.cc/pig | bash
```

**Add New Extension**

- add pg_anon 2.0.0
- add omnisketch 1.0.2
- add ddsketch 1.0.1
- add pg_duration 1.0.1
- add ddl_historization 0.0.7
- add data_historization 1.1.0
- add schedoc 0.0.1
- add floatfile 1.3.1
- add pg_upless 0.0.3
- add pg_task 1.0.0
- add pg_readme 0.7.0
- add vasco 0.1.0
- add pg_xxhash 0.0.1

**Update Extension**

- lower_quantile 1.0.3
- quantile 1.1.8
- sequential_uuids 1.0.3
- pgmq 1.5.0 (subdir)
- floatvec 1.1.1
- pg_parquet 0.2.0
- wrappers 0.4.4
- pg_later 0.3.0
- topn fix for deb.arm64
- add age 17 on debian
- powa + pg17, 5.0.1
- h3 + pg17
- ogr_fdw + pg17
- age + pg17 1.5 on debian
- pgtap + pg17 1.3.3
- repmgr
- topn + pg17
- pg_partman 5.2.4
- credcheck 3.0
- ogr_fdw 1.1.5
- ddlx 0.29
- postgis 3.5.1
- tdigest 1.4.3
- pg_repack 1.5.2

Release: https://github.com/pgsty/pig/releases/tag/v0.1.2


## v0.1.0

The pig CLI v0.1 is out, with the following new features:

**Install Script**

```bash
curl -fsSL https://repo.pigsty.io/pig | bash     # cloudflare, default
curl -fsSL https://repo.pigsty.cc/pig | bash     # mainland china mirror
```

**Extension Management**

You can download the extension and its dependencies with `import` subcommand, activate different postgres major versions with `link`, and prepare building env with `build` subcommand

```bash
pig ext list    [query]      # list & search extension
pig ext info    [ext...]     # get information of a specific extension
pig ext status  [-v]         # show installed extension and pg status
pig ext add     [ext...]     # install extension for current pg version
pig ext rm      [ext...]     # remove extension for current pg version
pig ext update  [ext...]     # update extension to the latest version
pig ext import  [ext...]     # download extension to local repo
pig ext link    [ext...]     # link postgres installation to path
pig ext build   [ext...]     # setup building env for extension
```

**Repo Management**

You can now create a local repo and create a tarball (offline package) from it, copy it to somewhere (e.g., without internet access), and create a repo from that offline package:

```bash
pig repo list                    # available repo list             (info)
pig repo info   [repo|module...] # show repo info                  (info)
pig repo status                  # show current repo status        (info)
pig repo add    [repo|module...] # add repo and modules            (root)
pig repo rm     [repo|module...] # remove repo & modules           (root)
pig repo update                  # update repo pkg cache           (root)
pig repo create                  # create repo on current system   (root)
pig repo boot                    # boot repo from offline package  (root)
pig repo cache                   # cache repo as offline package   (root)
```

**Pigsty Management**

The **pig** can also be used as a CLI tool for [Pigsty](https://doc.pgsty.com/) — the battery-include free PostgreSQL RDS

```bash
pig sty init     # install embed pigsty to ~/pigsty
pig sty boot     # install ansible and other pre-deps
pig sty conf     # auto-generate pigsty.yml config file
pig sty install  # run the install.yml playbook
```

**Self-Updating**

To update pig itself to the latest version, you can use the following command:

```bash
pig update
```

**Info**

Now pig info provides more details about your OS & PG environment:

```
$ pig info

# [Configuration] ================================
Pig Version      : 0.1.0
Pig Config       : /home/vagrant/.pig/config.yml
Log Level        : info
Log Path         : stderr

# [OS Environment] ===============================
OS Distro Code   : el9
OS Architecture  : amd64
OS Package Type  : rpm
OS Vendor ID     : rocky
OS Version       : 9
OS Version Full  : 9.3
OS Version Code  : el9

# [PG Environment] ===============================
Installed:
* PostgreSQL 17.2  74  Extensions

Active:
PG Version      :  PostgreSQL 17.2
Config Path     :  /usr/pgsql-17/bin/pg_config
Binary Path     :  /usr/pgsql-17/bin
Library Path    :  /usr/pgsql-17/lib
Extension Path  :  /usr/pgsql-17/share/extension

# [Pigsty Environment] ===========================
Inventory Path   : /home/vagrant/pigsty/pigsty.yml
Pigsty Home      : /home/vagrant/pigsty
Embedded Version : 3.2.0

# [Network Conditions] ===========================
pigsty.cc  ping ok: 141 ms
pigsty.io  ping ok: 930 ms
google.com request error
Internet Access   :  true
Pigsty Repo       :  pigsty.io
Inferred Region   :  china
Latest Pigsty Ver :  v3.2.0
```

Enjoy PostgreSQL!

**What’s Changed**

- fix typos by @kianmeng in https://github.com/pgsty/pig/pull/4

**New Contributors**

- @kianmeng made their first contribution in https://github.com/pgsty/pig/pull/4

**Full Changelog**: https://github.com/pgsty/pig/compare/v0.0.1...v0.1.0

**Checksums**

```
46165beec97ab9ff1314f80af953bd59  pig-0.1.0-1.aarch64.rpm
1320a6f9bfbd79948515657d6becbf37  pig-0.1.0-1.x86_64.rpm
bd078a5dc0c41454fcbbe0d8693d5fa0  pig-v0.1.0.linux-amd64.tar.gz
8a15e52f96735b78afa7da42843f1504  pig-v0.1.0.linux-arm64.tar.gz
4d25597cff8425c7e52a2b411344aa4a  pig_0.1.0_amd64.deb
d5f0874601bc1bbd0dd40b5c9982ea9f  pig_0.1.0_arm64.deb
```

Release: https://github.com/pgsty/pig/releases/tag/v0.1.0

------

## v0.0.1

**Get Started**

Installthe `pig` package first, you can also install via command:

```bash
curl -fsSL https://repo.pigsty.io/pig | bash     # cloudflare, default
curl -fsSL https://repo.pigsty.cc/pig | bash     # mainland china mirror
```

Then it’s ready to use, assume you want to install the [`pg_duckdb`](https://ext.pgsty.com/e/pg_duckdb) extension:

```bash
$ pig repo add pigsty pgdg -u  # add pgdg & pigsty repo, update cache
$ pig ext install pg17         # install PostgreSQL 17 kernels with PGDG native packages
$ pig ext install pg_duckdb    # install the pg_duckdb extension (for current pg17)
```

That’s it! All sets! you can check with the `pig ext status` sub command:

```bash
$ pig ext status               # show installed extension and pg status
                               # to print built-in contrib extension, use -c|--contrib flag
Installed PG Vers :  17 (active)
Active PostgreSQL :  PostgreSQL 17.2
PostgreSQL        :  PostgreSQL 17.2
Binary Path       :  /usr/pgsql-17/bin
Library Path      :  /usr/pgsql-17/lib
Extension Path    :  /usr/pgsql-17/share/extension
Extension Stat    :  1 Installed (PIGSTY 1, PGDG 0) + 67 CONTRIB = 68 Total

Name       Version  Cate  Flags   License  Repo    Package        Description
----       -------  ----  ------  -------  ------  ------------   ---------------------
pg_duckdb  0.2.0    OLAP  -dsl--  MIT      PIGSTY  pg_duckdb_17*  DuckDB Embedded in Postgres

(1 Rows) (Flags: b = HasBin, d = HasDDL, s = HasSolib, l = NeedLoad, t = Trusted, r = Relocatable, x = Unknown)
```

Check the advanced usage for details and [list 340 available extensions](https://ext.pgsty.com/list).


**Installation**

The `pig` util is a standalone go binary with no dependencies. you can just download the binary or use the following commands to add the repo and install it via package manager (recommended).

For Ubuntu 22.04 / 24.04 & Debian 12 or any compatible platforms:

```bash
sudo tee /etc/apt/sources.list.d/pigsty.list > /dev/null <<EOF
deb [trusted=yes] https://repo.pigsty.io/apt/infra generic main
EOF
sudo apt update; sudo apt install -y pig
```

For EL 8/9 and compatible platforms:

```bash
sudo tee /etc/yum.repos.d/pigsty.repo > /dev/null <<-'EOF'
[pigsty-infra]
name=Pigsty Infra for $basearch
baseurl=https://repo.pigsty.io/yum/infra/$basearch
enabled = 1
gpgcheck = 0
module_hotfixes=1
EOF
sudo yum makecache; sudo yum install -y pig
```

> For mainland china user: consider replace the `repo.pigsty.io` with `repo.pigsty.cc`

**Compatibility**

`pig` runs on: RHEL 8/9, Ubuntu 22.04/24.04, and Debian 12, on both `amd64/arm64` arch

|  Code   | Distribution                |  `x86_64`  | `aarch64`  |
|:-------:|-----------------------------|:----------:|:----------:|
| **el9** | RHEL 9 / Rocky9 / Alma9 / … | PG 17 - 13 | PG 17 - 13 |
| **el8** | RHEL 8 / Rocky8 / Alma8 / … | PG 17 - 13 | PG 17 - 13 |
| **u24** | Ubuntu 24.04 (`noble`)      | PG 17 - 13 | PG 17 - 13 |
| **u22** | Ubuntu 22.04 (`jammy`)      | PG 17 - 13 | PG 17 - 13 |
| **d12** | Debian 12 (`bookworm`)      | PG 17 - 13 | PG 17 - 13 |

Here are some bad cases and limitations for the above distros:

- [`citus`](https://ext.pgsty.com/e/citus) is not available on `aarch64` and ubuntu 24.04
- [`pljava`](https://ext.pgsty.com/e/pljava) is missing on `el8`
- [`jdbc_fdw`](https://ext.pgsty.com/e/jdbc_fdw) is missing on `el8.aarch64` and `el9.aarch64`
- [`pllua`](https://ext.pgsty.com/e/pllua) is missing on `el8.aarch64` for pg 13,14,15
- [`topn`](https://ext.pgsty.com/e/topn) is missing on `el8.aarch64` and `el9.aarch64` for pg13, and all `deb.aarch64`
- [`pg_partman`](https://ext.pgsty.com/e/pg_partman) and [`timeseries`](https://ext.pgsty.com/e/timeseries) is missing on `u24` for pg13
- [`wiltondb`](https://ext.pgsty.com/e/wiltondb) is missing on `d12`

Release: https://github.com/pgsty/pig/releases/tag/v0.0.1