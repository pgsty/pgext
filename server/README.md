# pgext serve — PGEXT.CLOUD 自包含 Web 服务

单二进制交付：网页资产经 `go:embed` 内嵌，数据全部来自 PostgreSQL 的 `pgext` schema。

```bash
pgext serve                                    # 本地开发：postgres:///data（或 $PGURL）
pgext serve --db postgres://user:pw@host/data  # 显式连接串
pgext serve --db mydb --listen :8080           # 库名简写 + 自定义端口
PGEXT_RELOAD_TOKEN=secret pgext serve           # 可选：启用带认证的手工刷新
```

连接串优先级：`--db` > 全局 `-d/--database` > 环境变量 `PGURL` > `postgres:///data`。
前提：目标库已有完整 `pgext` schema 与目录数据（`pgext init`）。

## 架构

```
┌──────────────┐   go:embed    ┌────────────────────────────┐
│ web/          │──────────────▶│ pgext serve (:8432)        │
│  index.html   │               │  ┌──────────┐  ┌─────────┐ │   pgxpool   ┌────────────┐
│  app.js      │   /assets/*   │  │ Snapshot │  │ ttlCache│ │◀───────────▶│ PostgreSQL │
│  style.css   │               │  │ (内存快照)│  │ (响应缓存)│ │             │ pgext.*    │
└──────────────┘               │  └──────────┘  └─────────┘ │             └────────────┘
                               └────────────────────────────┘
```

两级缓存策略：

1. **Snapshot（`snapshot.go`）**——启动时把新版 `universe`、`category`、`active_pg/os`、
   `doc` 元数据与 `pkg` 的 AVAIL 目标整体载入内存（`atomic.Pointer` 原子换入）。列表、详情、
   维度聚合及首页的 PG × OS 精确筛选均零 SQL 应答。内容 ETag 不包含加载时间，相同数据刷新后
   浏览器缓存仍然有效。按 `--cache-ttl`（默认 5m）后台自动刷新。
2. **矩阵缓存（`matrix.go` / `pgext.matrix_cache`）**——`/api/v1/matrix` 的完整响应在写入侧预计算：
   一条 `jsonb_object_agg` 聚合把 `pgext.pkg` 折叠成每包一行（键 `el8i.17`，值 `版本 + #/@ + G/P/M/N`），
   Go 侧铺到规范 OS × PG 网格后连同统计一起序列化，落入单行 mtime 缓存表。读路径是一次主键查询；
   缓存在早于 `pgext.status` 装载时间或超过 24h 时惰性重建，`POST /api/v1/reload` 与启动预热会主动重建。
3. **ttlCache（`cache.go`）**——全局物化矩阵、逐扩展 `pkg` 矩阵、`bin` 文件表、`doc` 正文按需查询后缓存；
   键内嵌快照版本号，快照刷新即自然失效。

中间件（`middleware.go`）：访问日志（debug 级）、panic 恢复、只读 `/api/*` CORS、安全响应头、
gzip（204/304 除外）。`POST /api/v1/reload` 默认关闭；仅在设置 `--reload-token` 或
`PGEXT_RELOAD_TOKEN` 后启用，并要求 `Authorization: Bearer …` 或 `X-PGEXT-Reload-Token`。
优雅退出：SIGINT/SIGTERM → `http.Server.Shutdown`（5s 期限）。

## 查询 API（/api/v1，全 JSON，只读）

| 端点 | 说明 |
|---|---|
| `GET /api/v1/meta` | 服务信息、计数（含 `packages` 包族数与 `build_slots` 构建槽位）、分类（双语描述）、活跃 PG/OS |
| `GET /api/v1/ext` | 列表/搜索：除分类、仓库、PG/OS、形态、生命周期与来源外，还支持 `tag pkg capability build docs relation pgrx active`；`q` 可组合 `tag:vector build:pgrx doc:bilingual is:packaged` 等操作符 |
| `GET /api/v1/ext/{name}` | 单个扩展完整记录（新版 universe 字段 + `state/repo/ext_*` v1 兼容别名） |
| `GET /api/v1/ext/{name}/matrix` | 可用性矩阵：`pgext.pkg` 的 pg × os 单元格（state/org/version/count） |
| `GET /api/v1/matrix` | 全局构建矩阵（`matrix-row.v2`）：`p/e/c` 每格一个状态字节（`G/P/M/N`），`v/i` 为行级版本字典与 base36 索引（gzip 后约 7 KiB，读路径为 `pgext.matrix_cache` 单行查询；前端以单张 canvas 一次绘制，/pg/{n} 与 /os/{target} 复用同一载荷切片渲染） |
| `GET /api/v1/ext/{name}/files` | 二进制包文件（`pgext.bin` ⋈ `repository`，含下载 URL 与 SHA256），可选 `?pg= &os=` |
| `GET /api/v1/ext/{name}/doc` | 用法文档 markdown（`pgext.doc`），`?lang=en\|zh` |
| `GET/POST /api/v1/ext/{name}/visit` | 页面访问计数：POST 自增并返回 `{visits, downloads}`，GET 只读；下载计数暂为占位 0。由 `--visits-file`（默认 `~/.pgext/visits.json`）持久化，30s 周期 + 退出时落盘 |
| `GET /api/v1/dim/{key}` | 19 个维度聚合：`category tag package kind lifecycle license lang distribution repo pg os build pgrx capability docs relation vendor kernel activity`（旧 `type` 仍是 `kind` 的兼容别名） |
| `GET /api/v1/bootstrap` | SPA 引导载荷（紧凑位置数组 32 列，含 `repo_url/url/license_url`；ETag/304，gzip 后约 200 KiB）。载荷格式版本 `bootFormat` 同时参与 ETag 与前端请求 URL（`?fmt=b32`）：改动列布局时必须同步递增，否则浏览器缓存可能把旧列布局供给新前端 |
| `POST /api/v1/reload` | 带令牌刷新内存快照；未配置令牌时返回 404 |
| `GET /healthz` | 存活探针：db ping + 快照年龄 |

路由层面：规范深链为 `/ext/{name}`、`/pkg/{name}`、`/cate/{CODE}` 与 `/list[/{dim}]`；旧的
`/e/ /p/ /c/` 由服务端 301 永久重定向（保留查询串，含尾斜杠变体），`/browse`、`/dim/{key}`、`/cat/` 由客户端就地迁移。
Hugo 时代的旧页面由 `server.go` 中集中的 `legacyRedirects` 表做 302 兼容跳转（保留查询串、覆盖有无尾斜杠）：
`/zh(/...)`与 `/en(/...)` 去语言前缀、`/e` `/list/ext` → `/`、`/list/pkg|cate` → `/list/package|category`、
`/os` → `/list/os`、`/os/matrix` → `/matrix`、`/categories[/{code}]` `/tags[/{tag}]` 分别映射到
`/list/category`、`/cate/{code}`、`/list/tag`、`/?tag={tag}`、`/repo` → `/list/repo`、
`/repo/pgsql` → `/repo/PIGSTY`、`/pig/*` → Pigsty 文档；新动态 `/repo/{value}`、`/os/{target}` 不受影响。

组合示例：

```bash
curl 'localhost:8432/api/v1/ext?q=vector&cat=RAG&limit=5'      # RAG 分类下搜 vector
curl 'localhost:8432/api/v1/ext?license=MIT&pg=18&sort=stars'  # MIT 且支持 PG18，按星标
curl 'localhost:8432/api/v1/ext?pg=18,17'                      # 同时支持 PG18 与 PG17（pg_ver @>）
curl 'localhost:8432/api/v1/ext?q=vector&pg=18&os=el9.x86_64' # 该精确二进制目标可安装
curl 'localhost:8432/api/v1/ext?tag=analytics&build=pgrx&docs=bilingual'
curl 'localhost:8432/api/v1/dim/capability'                    # 运行时能力聚合
curl 'localhost:8432/api/v1/matrix'                           # 全局扩展包构建矩阵
curl 'localhost:8432/api/v1/ext/postgis/files?pg=18&os=el9.x86_64'
curl 'localhost:8432/api/v1/ext/timescaledb/doc?lang=zh'
curl -X POST -H 'Authorization: Bearer secret' localhost:8432/api/v1/reload
```

全局矩阵缓存表 `pgext.matrix_cache` 由服务端按需自建（`CREATE TABLE IF NOT EXISTS`，只读角色下
自动降级为纯内存计算）。任何路径更新 `pgext.pkg` 后，只要 `pgext.status` 的装载时间随之推进，
服务端就会在下一次读取时自动重建缓存；`POST /api/v1/reload` 会立即重建。旧的
`pgext.matrix` 物化视图不再被服务端读取，仍由 `pgext reload` 的发布事务顺带刷新。

## 前端（web/）

`design/prototype/` 静态原型的 API 化版本：同一套设计系统，使用 History API 的干净路由
（`/`、`/pkg/{pkg}`、`/ext/{name}`、`/c/{CODE}`、`/matrix`、`/list`、`/list/{dim}`、`/about`）。旧的
`/e/ /p/ /browse /dim/` 与 `#/…` 哈希链接会经服务端 301 或客户端迁移到对应新地址。差异：

- 启动时拉取 `/api/v1/bootstrap`（位置数组，列序与 `handleBootstrap` 注释同步维护）；上一次
  载荷缓存在 localStorage，回访先用缓存目录即时渲染，再以 `If-None-Match` 条件请求后台校验，
  全部筛选、排序与搜索都在前端对这份数据完成；
- 首页 Hero：单行主标题与说明、六项关键指标（收录/已打包/包族/Linux 目标/PG 大版本/构建槽位）、
  放大后的宇宙点阵导航（分类图例并入筛选面板）；检索区为 搜索框 → SQL 读出 → 筛选面板 → 结果；
- 筛选面板收敛为 SCOPE（全部/可交付/未交付/分支限定/特定厂商/PG 自带）+ 分类 + PostgreSQL +
  语言 + 许可证（≤2 个成员的取值折叠进"全部 N 种 →"索引链接）。SCOPE 级联：其余各行计数都在
  所选范围子集上重新聚合；分类固定 8 × 2 对齐、Any 独立成列；每行标题可点击直达对应
  `/list/{dim}` 索引页；无独立行的维度（形态、仓库来源、生命周期、二进制目标等）在激活时显示于
  "当前筛选"并可一键清除。筛选参数统一为 `lang/license/repo/...`（`lng` 兼容保留，
  `?lang=zh|en` 仍是界面语言开关）；
- 扩展卡片固定四列：标题 + 角标（GitHub 仓库显示 GitHub 图标与星数，其余上游显示 ↗ 箭头，点击均
  跳首要 URL）、灰色 `pkg · version` 副标题、双语描述与维度标签（点击即筛选）；`pgext.extension`
  （packaged）条目带强调边框以示开箱即用。软件包卡片同构：标题携主扩展版本、正文列出全部交付
  扩展按钮（直达 `/ext`）、标签取自主扩展。卡片与表格行悬停片刻会弹出可交互的详情浮窗（完整描述、
  元数据与上游链接）。扩展表格列序为 name/version/description/category/license/lang/pg/stars，
  名称与许可证列设宽度上限，描述占满余宽；软件包表格为 扩展包/主扩展/描述/分类/许可证/PG/星标，
  多扩展的包在首列于包名下方逐行列出全部扩展并可直达；
- 代码块内置轻量高亮（SQL / shell / yaml·conf 键值），Usage 文档、安装与构建命令统一生效；
- `/matrix` 重写为整页滚动的全宽格子表：行虚拟化改挂窗口滚动，双行表头（OS 释出 + 架构配色区分
  x86_64/aarch64 + PG 大版本）吸顶，每格是首页点阵同款的纯色方块；除状态筛选外新增
  全部 / PGDG / Pigsty 三种视角，按所选仓库的覆盖情况重新着色所有有效槽位，直观呈现
  “PGDG 提供了哪些、缺了哪些，Pigsty 补齐了哪些”。性能：物化载荷缓存于 localStorage
  （回访即时渲染，按 `generated` 时间戳后台校验一次），每行单元格 HTML 按视角惰性预计算，
  滚动帧只拼接缓存字符串；
- 首页默认展示扩展卡片（导航中的“扩展”直接进入 EXT 目录），可显式切换项目/共享包族或表格；
  `packaged/source/kernel/vendor/contrib` 独立编码，并支持 `PG + OS + 架构` 精确包可用性筛选；
- 多维索引按身份分类、构建交付、运行时文档、生态活跃度组织 19 个维度。标签、包族、二进制目标、
  构建链、pgrx 版本、能力位、文档覆盖与依赖信号均直接利用 `universe/doc/pkg` 快照；点击任意取值
  会生成可组合、可分享的首页查询参数，长维度页可即时搜索取值；
- 扩展详情页重构：Hero 左侧为身份区（名称/状态/标签/GitHub 活跃度 + 访问计数，下载计数占位），
  右侧是醒目的软件包导航卡（直达 `/pkg/{pkg}` 的定义、矩阵与下载）；正文为 概览 → 安装 →
  版本与可用性 → 构建 → 关联 → 元数据 → 用法。软件下载表只保留在 PKG 页；用法文档按
  GitHub 风味 Markdown 渲染，支持 `> [!NOTE]` 等 Alert 标注（含标注行内正文的变体）；
- 预加载配置会合并依赖扩展所需的动态库；安装命令会按选定的 PG + OS + 架构使用矩阵里的
  精确包名，并区分预编译包、PostgreSQL contrib、源码目录与供应商专有扩展；
- `pgext.doc` 的中英文 Markdown 正文会去掉重复的顶层 Usage 标题，并渲染目录、锚点、表格、
  嵌套列表、引用、图片及带语言/复制按钮的代码块；
- 包族页只呈现包层信息：源码/RPM/DEB 定义、精确可用性、校验和下载、构建手册与按目标安装命令；
  页面末尾再链接该软件包交付的全部扩展，预加载和 `CREATE EXTENSION` 等扩展级步骤留在 EXT 页；
- HTML 外壳始终重验证，内嵌 CSS/JS URL 带内容指纹，二进制升级后不会混用旧前端缓存；
- 会话级缓存（FULLC/MXC/FILEC/DOCC），语言切换不重复拉取。

这里不再使用 Hash Router。`/#/...` 只适合无法把深层路径交回应用的纯静态托管；`pgext serve`
本身可以为 `/e/...`、`/p/...` 等任意 History API 深链返回 SPA 外壳，所以公开链接始终是干净路径。
客户端仅保留一次性的旧链接迁移：访问旧 `/#/ext/vector` 会立即 `replaceState` 为 `/e/vector`。

如果页面仍生成 `/#/`，说明当前端口运行的是升级前编译、内嵌了旧 `web/app.js` 的进程；磁盘上的
前端文件不会热替换已经运行的 Go 二进制。请重启开发进程（例如 `go run . serve --listen :8432`），
再用页面 HTML 中带内容指纹的 `/assets/app.js?v=...` 确认新资产已经生效。

## 测试

```bash
go test ./server        # 过滤引擎、OS 排序、TTL 缓存（无需数据库）
go build -o pgext . && ./pgext serve   # 端到端
```
