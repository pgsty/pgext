# pgext serve — PGEXT.CLOUD 自包含 Web 服务

单二进制交付：网页资产经 `go:embed` 内嵌，数据全部来自 PostgreSQL 的 `pgext` schema。

```bash
pgext serve                                    # 本地开发：postgres:///data（或 $PGURL）
pgext serve --db postgres://user:pw@host/data  # 显式连接串
pgext serve --db mydb --listen :8080           # 库名简写 + 自定义端口
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

1. **Snapshot（`snapshot.go`）**——启动时把 `universe`（1,633 行）、`category`、`active_pg/os`、
   `doc` 元数据整体载入内存（`atomic.Pointer` 原子换入），列表/详情/维度聚合全部零 SQL 应答。
   按 `--cache-ttl`（默认 5m）后台自动刷新；`POST /api/v1/reload` 立即刷新（5s 内合并去重）。
2. **ttlCache（`cache.go`）**——逐扩展的 `pkg` 矩阵、`bin` 文件表、`doc` 正文按需查询后缓存；
   键内嵌快照版本号，快照刷新即自然失效。

中间件（`middleware.go`）：访问日志（debug 级）、panic 恢复、`/api/*` CORS、gzip（204/304 除外）。
优雅退出：SIGINT/SIGTERM → `http.Server.Shutdown`（5s 期限）。

## 查询 API（/api/v1，全 JSON，只读）

| 端点 | 说明 |
|---|---|
| `GET /api/v1/meta` | 服务信息、计数、分类（双语描述）、活跃 PG/OS |
| `GET /api/v1/ext` | 列表/搜索：`q cat repo license lang pg type vendor scope sort limit offset`；`q` 支持 `cat:GIS lang:rust pg:18 is:packaged` 操作符，相关度排序与前端一致 |
| `GET /api/v1/ext/{name}` | 单个扩展完整记录（含 requires/require_by/see_also/siblings/doc 链接） |
| `GET /api/v1/ext/{name}/matrix` | 可用性矩阵：`pgext.pkg` 的 pg × os 单元格（state/org/version/count） |
| `GET /api/v1/ext/{name}/files` | 二进制包文件（`pgext.bin` ⋈ `repository`，含完整下载 URL），可选 `?pg= &os=` |
| `GET /api/v1/ext/{name}/doc` | 用法文档 markdown（`pgext.doc`），`?lang=en\|zh` |
| `GET /api/v1/dim/{key}` | 维度聚合：`category license lang repo vendor kernel type pg` |
| `GET /api/v1/bootstrap` | SPA 引导载荷（紧凑位置数组，ETag/304，gzip 后约 100 KiB） |
| `POST /api/v1/reload` | 立即刷新内存快照 |
| `GET /healthz` | 存活探针：db ping + 快照年龄 |

组合示例：

```bash
curl 'localhost:8432/api/v1/ext?q=vector&cat=RAG&limit=5'      # RAG 分类下搜 vector
curl 'localhost:8432/api/v1/ext?license=MIT&pg=18&sort=stars'  # MIT 且支持 PG18，按星标
curl 'localhost:8432/api/v1/ext/postgis/files?pg=18&os=el9.x86_64'
curl 'localhost:8432/api/v1/ext/timescaledb/doc?lang=zh'
```

## 前端（web/）

`design/prototype/` 静态原型的 API 化版本：同一套设计系统与路由
（`#/`、`#/ext/{name}`、`#/cat/{CODE}`、`#/browse`、`#/dim/{key}`、`#/about`）。差异：

- 启动时拉取 `/api/v1/bootstrap`（位置数组，列序与 `handleBootstrap` 注释同步维护）；
- 详情页骨架屏 + 四路并行水合：完整记录、16 OS × 5 PG 全量矩阵、按 PG 分页的包文件
  下载表（默认每 OS 只展示最新构建，可展开全部）、数据库直出的双语用法文档（531 篇全量可用）；
- 会话级缓存（FULLC/MXC/FILEC/DOCC），语言切换不重复拉取。

## 测试

```bash
go test ./server        # 过滤引擎、OS 排序、TTL 缓存（无需数据库）
go build -o pgext . && ./pgext serve   # 端到端
```
