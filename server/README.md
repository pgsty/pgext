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
2. **ttlCache（`cache.go`）**——逐扩展的 `pkg` 矩阵、`bin` 文件表、`doc` 正文按需查询后缓存；
   键内嵌快照版本号，快照刷新即自然失效。

中间件（`middleware.go`）：访问日志（debug 级）、panic 恢复、只读 `/api/*` CORS、安全响应头、
gzip（204/304 除外）。`POST /api/v1/reload` 默认关闭；仅在设置 `--reload-token` 或
`PGEXT_RELOAD_TOKEN` 后启用，并要求 `Authorization: Bearer …` 或 `X-PGEXT-Reload-Token`。
优雅退出：SIGINT/SIGTERM → `http.Server.Shutdown`（5s 期限）。

## 查询 API（/api/v1，全 JSON，只读）

| 端点 | 说明 |
|---|---|
| `GET /api/v1/meta` | 服务信息、计数、分类（双语描述）、活跃 PG/OS |
| `GET /api/v1/ext` | 列表/搜索：除分类、仓库、PG/OS、形态、生命周期与来源外，还支持 `tag pkg capability build docs relation pgrx active`；`q` 可组合 `tag:vector build:pgrx doc:bilingual is:packaged` 等操作符 |
| `GET /api/v1/ext/{name}` | 单个扩展完整记录（新版 universe 字段 + `state/repo/ext_*` v1 兼容别名） |
| `GET /api/v1/ext/{name}/matrix` | 可用性矩阵：`pgext.pkg` 的 pg × os 单元格（state/org/version/count） |
| `GET /api/v1/ext/{name}/files` | 二进制包文件（`pgext.bin` ⋈ `repository`，含下载 URL 与 SHA256），可选 `?pg= &os=` |
| `GET /api/v1/ext/{name}/doc` | 用法文档 markdown（`pgext.doc`），`?lang=en\|zh` |
| `GET /api/v1/dim/{key}` | 19 个维度聚合：`category tag package kind lifecycle license lang distribution repo pg os build pgrx capability docs relation vendor kernel activity`（旧 `type` 仍是 `kind` 的兼容别名） |
| `GET /api/v1/bootstrap` | SPA 引导载荷（紧凑位置数组，ETag/304，gzip 后约 100 KiB） |
| `POST /api/v1/reload` | 带令牌刷新内存快照；未配置令牌时返回 404 |
| `GET /healthz` | 存活探针：db ping + 快照年龄 |

组合示例：

```bash
curl 'localhost:8432/api/v1/ext?q=vector&cat=RAG&limit=5'      # RAG 分类下搜 vector
curl 'localhost:8432/api/v1/ext?license=MIT&pg=18&sort=stars'  # MIT 且支持 PG18，按星标
curl 'localhost:8432/api/v1/ext?pg=18,17'                      # 同时支持 PG18 与 PG17（pg_ver @>）
curl 'localhost:8432/api/v1/ext?q=vector&pg=18&os=el9.x86_64' # 该精确二进制目标可安装
curl 'localhost:8432/api/v1/ext?tag=analytics&build=pgrx&docs=bilingual'
curl 'localhost:8432/api/v1/dim/capability'                    # 运行时能力聚合
curl 'localhost:8432/api/v1/ext/postgis/files?pg=18&os=el9.x86_64'
curl 'localhost:8432/api/v1/ext/timescaledb/doc?lang=zh'
curl -X POST -H 'Authorization: Bearer secret' localhost:8432/api/v1/reload
```

## 前端（web/）

`design/prototype/` 静态原型的 API 化版本：同一套设计系统，使用 History API 的干净路由
（`/`、`/p/{pkg}`、`/e/{name}`、`/c/{CODE}`、`/browse`、`/dim/{key}`、`/about`）。旧的
`#/pkg/{pkg}`、`#/ext/{name}` 与 `#/cat/{CODE}` 链接会在客户端自动迁移到对应新地址。差异：

- 启动时拉取 `/api/v1/bootstrap`（位置数组，列序与 `handleBootstrap` 注释同步维护）；
- 首页默认展示扩展卡片（导航中的“扩展”直接进入 EXT 目录），可显式切换项目/共享包族或表格；
  `packaged/source/kernel/vendor/contrib` 独立编码，并支持 `PG + OS + 架构` 精确包可用性筛选；
- 多维索引按身份分类、构建交付、运行时文档、生态活跃度组织 19 个维度。标签、包族、二进制目标、
  构建链、pgrx 版本、能力位、文档覆盖与依赖信号均直接利用 `universe/doc/pkg` 快照；点击任意取值
  会生成可组合、可分享的首页查询参数，长维度页可即时搜索取值；
- 详情页沿用 `pgext gen cc/io` 的内容逻辑组织成完整扩展手册：概览、分组元数据、正反向依赖与
  同包扩展、版本/包名定义、可用性矩阵与 SHA256 下载、源码构建、仓库/安装/预加载/启用/验证；
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
