# PG19 非 Rust 扩展元数据回填 · 2026-11-21

本次将 RPM / DEB 完整核验的最终结果回填到本地 `/tmp:5432/data` 的 `pgext.extension`、`pgext.universe`，并按仓库 `make dump` 的 SELECT / COPY 格式导出这两张表。构建核验与元数据回填是两个独立批次；原 HTML 报告中的「PGEXT 未写」描述的是回填前的审计时点。

输入矩阵：`../rpm/doc/nonrust-pg19-20261121.json`；SHA256：`170e46ec0237ebc4fa3180b540739f0f0a551f2351e365242dd530c71d353fa6`。核验平台为 Rocky Linux 9.8 / aarch64 的 `pgsty/el9a:build`，以及 m3 上 Ubuntu 24.04.5 / amd64 的 `pgsty/u24:build`。

## 结果

| 项目 | extension | universe |
| --- | ---: | ---: |
| 实际变更记录数 | 338 | 340 |
| 更新 version | 30 | 3 |
| 更新 rpm_ver | 37 | 37 |
| 更新 deb_ver | 34 | 36 |
| pg_ver 增加 19 | 278 | 279 |
| rpm_pg 增加 19 | 321 | 321 |
| deb_pg 增加 19 | 275 | 277 |
| 源码文件名同步（source / tarball） | 30 | 28 |

`extension` 中有 40 条记录的包版本得到更新；Universe 额外为已有的 `timeit`、`mysqlcompat` 记录补齐了本轮验证的 DEB 版本、包名模板、供应来源、构建标记和 PG19 范围。这两条仍保留 `packaged=false`，本次没有新增 packaged baseline 成员。

回填后两表总记录数分别为 627、2814。PG19 RPM 范围均覆盖 321 条记录；PG19 DEB 范围分别覆盖 429、431 条记录。后两个总数包含既有支持，不能当作本次新增数或独立源码构建单元数。

## 回填规则

- `version`、`rpm_ver`、`deb_ver` 为规范化扩展／软件包版本；`pg_ver`、`rpm_pg`、`deb_pg` 为 PG 主版本范围。分别更新，不能把二者混用。
- 本地 PG19 证据必须同时具有最终 `local_recipe_integrated=true`、验包结果、相应主包名称和 SHA256；需要 DDL 的扩展还须在主包中找到其 control 文件。一个源码包有多个扩展时逐一确认。
- 使用最终成功产物的版本，去除 epoch、PGSTY/PGDG 修订、发行版后缀和 Git 快照后缀。不会因源码包的版本而机械改写具有独立 control 版本的子扩展。
- 既有 PGDG / PIGSTY 供应商策略保持原值。本轮 ARM RPM 或 amd64 DEB 构建成功，表示该包族新增了一个已验证 PG19 构建坐标；不表示全部系统／架构已经构建或发布。
- 保留目录中的历史支持范围，只增加有证据的 PG19。默认构建目标由 `pgext.active_pg` 决定，已经是 15–19；PG14 历史记录不表示重新启用 PG14。本轮回填时 EL8 尚未恢复活跃标记（后续统一策略恢复 EL8 服务 PG15–18，PG19 在 EL8 上标记为 `N/A`）。
- 同步升级涉及的 source / tarball，避免版本更新后仍下载旧源码。真实源码文件名保留日期版和上游修订等原始信息。

## 必须保留的边界

| 项目 | 本次结论 |
| --- | --- |
| Rust/pgrx、大型后置、内核配套后置 | 本次回填不改变这些分组。其他任务对 Universe 的独立修改以新基线保留。 |
| pg_variables、pg_datasentinel、pg_failover_slots | 运行验证失败／超时，PG19 范围保持禁用。 |
| Citus、AGE、TimescaleDB | PG19 保持禁用；本次 PG19 回填不顺带提升它们的包版本。 |
| pg_net | 仅 DEB 增加 PG19；RPM/EL9 的 0.9.2 分支失败。rpm_ver 继续保留原有跨平台 0.20.5 基线，不改成 EL9 单个平台的旧值。 |
| pg_oidc_validator | 仅 DEB 增加 PG19；RPM 的 EL10/GCC ≥13 平台限制保持，外部 IdP 认证未测试。 |
| PostGIS | 依据本轮两平台 PGDG 正式 3.7.0 包证据补齐 RPM PG19；不扩大本地备用 recipe 范围。 |
| pgaudit | 两侧包版本为 19.0，RPM PG19 补齐；PG15 本地兼容配方 1.7.2 不覆盖目录的最新主版本分支。 |
| OneSparse | version=0.1.0；rpm_ver=deb_ver=1.0.0。已读取实际包内 control 确认。 |
| pg_bigm | 三个版本字段保留 1.2；源码改为 pg_bigm-1.2-20261118.tar.gz，原始日期版本留在证据中。 |
| pg_crash | rpm_ver / deb_ver 从历史 1.0 修正为 0.3；新包为带 epoch 的 1:0.3。epoch 不进入规范化目录字段。 |
| pg_ash | 本地包提供 ash.control，但两表尚无对应 canonical 行，留为单独收录事项；未凭包名虚构目录记录。 |
| pgcopydb | 独立工具，无现有扩展目录行，不新增 CREATE EXTENSION 类型记录。 |
| 源码内未随主包交付的子扩展、旧目录别名 | documentdb_distributed、PL/PHP transforms、uuid_v7 等未获相应 payload 证据，保持原值；包名别名不冒充 canonical control 名。 |

## 包版本更新明细

下表列出 `extension` 中包版本发生变化的 40 条记录；对应 Universe 包版本已一致。`保持`表示该侧沿用此前记录。

| 扩展 | RPM | DEB |
| --- | --- |
| pg_background | 2.0.4 → 2.0.5 | 2.0.4 → 2.0.5 |
| mobilitydb | 1.3.1 → 1.3.2 | 1.3.1 → 1.3.2 |
| mobilitydb_datagen | 1.3.1 → 1.3.2 | 1.3.1 → 1.3.2 |
| biscuit | 3.1.0 → 3.1.1 | 3.1.0 → 3.1.1 |
| pg_fts | 1.16.0 → 1.18.0 | 1.16.0 → 1.18.0 |
| pg_tre | 4.6.2 → 4.6.4 | 4.6.2 → 4.6.4 |
| pg_partman | 5.5.0 → 5.5.1 | 5.5.0 → 5.5.1 |
| pgq | 3.5.1 → 3.5.2 | 3.5.1 → 3.5.2 |
| pgmq | 1.13.1 → 1.13.2 | 1.13.1 → 1.13.2 |
| ulak | 0.2.0 → 0.3.0 | 0.2.0 → 0.3.0 |
| fsm_core | 1.1.0 → 1.2.0 | 1.1.0 → 1.2.0 |
| provsql | 1.13.1 → 1.14.0 | 1.13.1 → 1.14.0 |
| plpgsql_check | 2.10.16 → 2.11.0 | 2.10.16 → 2.11.0 |
| plr | 8.4.8.10 → 8.4.8.11 | 8.4.8.10 → 8.4.8.11 |
| semver | 0.42.0 → 0.43.0 | 0.42.0 → 0.43.0 |
| roaringbitmap | 1.3.0（保持） | 1.2.0 → 1.3.0 |
| uint128 | 1.2.0 → 1.3.0 | 1.2.0 → 1.3.0 |
| ddlx | 0.31 → 0.32 | 0.31 → 0.32 |
| pg_checksums | 1.3 → 1.4 | 1.4（保持） |
| pg_integrity_checksums | 2.2.0 → 2.9.1 | 2.2.0 → 2.9.1 |
| pg_permissions | 1.4.1 → 1.6 | 1.4.1 → 1.6 |
| pg_crash | 1.0 → 0.3 | 1.0 → 0.3 |
| safeupdate | 1.5 → 1.7 | 1.7（保持） |
| pg_stat_plans | 2.1.0 → 2.1.1 | 2.1.1（保持） |
| pg_stat_monitor | 2.4.0 → 2.4.1 | 2.4.0 → 2.4.1 |
| pgexporter_ext | 0.2.4 → 0.2.5 | 0.2.5（保持） |
| powa | 5.1.0 → 5.3.0 | 5.3.0（保持） |
| supautils | 4.0.0 → 4.0.1 | 4.0.0 → 4.0.1 |
| pgaudit | 18.0 → 19.0 | 18.0 → 19.0 |
| logerrors | 2.1.5（保持） | 2.1.5 → 2.1.6 |
| pg_auth_mon | 3.0 → 5.0 | 5.0（保持） |
| sslutils | 1.4.1 → 1.4.2 | 1.4.1 → 1.4.2 |
| documentdb | 1.0 → 1.2 | 1.0 → 1.2 |
| documentdb_core | 1.0 → 1.2 | 1.0 → 1.2 |
| documentdb_extended_rum | 1.0 → 1.2 | 1.0 → 1.2 |
| pg_dbms_metadata | 1.0.0 → 1.1.0 | 1.0.0 → 1.1.0 |
| pgactive | 2.1.9 → 2.1.10 | 2.1.9 → 2.1.10 |
| pgclone | 4.4.2 → 4.9.0 | 4.4.2 → 4.9.0 |
| decoderbufs | 3.5.0 → 3.7.1 | 3.7.0 → 3.7.1 |
| repmgr | 5.5.1（保持） | 5.5.0 → 5.5.1 |

## 验证与并发保护

- 首次预检与原 CSV 一致；写入前出现并发 Universe 收录，旧基线守卫拒绝了写入。重新读取 2814 行的新基线，保留新增的 7 条记录及 34 条既有记录的独立修改。
- 在 SHARE ROW EXCLUSIVE 表锁内逐行核对 ID、名称和完整前值，事务回滚演练通过后提交 678 条精确更新。未插入或删除目录行。
- 逐字段比对数据库、预期变更和 CSV；确认所有未规划字段不变，两表的 rpm_ver / deb_ver / rpm_pg / deb_pg 一致；核对 1244 个本地构建证据引用。
- 只更新这两张目录表及 `pgext.status` 的目录时间／派生缓存失效标记。没有将本地验收包写入 `pgext.bin` / `pgext.pkg` 来冒充已发布包；仓库下载可用性仍由真实仓库数据决定。
- `go test ./...` 通过；定向 `git diff --check` 通过。

## 可复查材料

完整逐字段清单、证据关联、事务脚本、执行日志、前后数据库快照及 CSV 快照保存在仓库外：

```text
~/pgsty/artifacts/nonrust-pg19-20261121/metadata-update/
  changes.csv
  plan.json
  verification.json
  committed.json
  control-version-review.json
  early-build-sources.json
  apply.psql
  dry-run-rebased.log
  apply-rebased.log
  go-test.log
  rebase01/concurrent-changes.json
```

原完整 HTML 核验报表保留在 `~/pgsty/artifacts/nonrust-pg19-20261121/nonrust-pg19-final-report.html`。本文件记录随后完成的元数据回填，不改写原构建审计历史。
