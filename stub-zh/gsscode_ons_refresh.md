## 用法

来源：

- [PGXN 1.1.1 README](https://pgxn.org/dist/gsscode/1.1.1/README.html)
- [gsscode_ons_refresh 控制文件](https://api.pgxn.org/src/gsscode/gsscode-1.1.1/gsscode_ons_refresh.control)
- [gsscode_ons_refresh 1.0.0 SQL 定义](https://api.pgxn.org/src/gsscode/gsscode-1.1.1/gsscode_ons_refresh--1.0.0.sql)

`gsscode_ons_refresh` 是可选配套扩展，用于刷新 `gsscode` 使用的 `gsscode_types` 登记表。只有希望在数据库内完成刷新时才需要它；核心压缩类型与操作符并不依赖它。

### 核心流程

```sql
CREATE EXTENSION gsscode_ons_refresh CASCADE;

SELECT update_gsscode_types();
```

控制文件要求 `gsscode` 与 `http`；当依赖文件可用时，`CASCADE` 可以一并创建它们。`update_gsscode_types()` 下载项目生成的 JSON 登记表，执行行级 upsert，并返回处理数量。

### 信任与运维

1.0.0 的实际 SQL 实现使用 PL/pgSQL 与 `http`，尽管过时的 PGXN 元数据仍描述早期的 `plpython3u` 方案。启用 `http` 通常需要超级用户。数据库服务器必须能出站访问 `raw.githubusercontent.com` 的 HTTPS，因此网络策略、DNS、TLS 信任、超时和代理行为都会进入刷新链路。

JSON 文件由上游仓库的 GitHub Actions 工作流根据 ONS 发布生成。因此刷新同时信任 ONS 数据、项目自动化与仓库写权限；将其用于权威数据前应审查这条供应链。该函数执行 upsert，不会删除下载载荷中缺失的本地行；应在受控维护会话中运行，并核对返回数量与代表性记录。

上游 README 建议大多数安装使用外部 `update_gsscode_types.py` 路径，因为它无需数据库服务器出站访问，也不依赖 `http`。1.0.0 没有发布 PostgreSQL 主版本支持矩阵，因此应在目标服务器上验证两个依赖。
