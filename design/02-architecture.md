# 02 · 技术架构

## 总览

```
                 ┌─ 匿名 GET 全部可缓存 ─┐
Browser ───► CDN (Cloudflare) ───► pgext server (Go 单二进制)
                                        │  SSR (html/template)
                                        │  /api/* (JSON)
                                        │  内存 TTL 缓存
                                        ▼
                                  PostgreSQL (data · pgext schema)
                                        ▲
              现有数据管线不动：fetch → parse → recap → reload（每日）
```

一句话：**`pgext server` 是一个连接 PostgreSQL、服务端渲染 HTML、
单二进制部署的 Go Web 服务器**。数据日更、读多写少，动态渲染 + 三层缓存
让它同时拥有静态站的性能和数据库站的灵活性。

## `server` 子命令

```
pgext server [--listen :3000] [--pgurl $PGURL] [--dev] [--cache-ttl 60s]
```

- 复用现有 cobra 骨架（与 `fetch` / `parse` / `recap` / `reload` 平级）；
- `go:embed` 打包全部模板、CSS/JS、字体、图标 → **单二进制部署**，延续 pig/pigsty 的分发哲学；
- `--dev` 模式从磁盘热读模板，改模板不用重编译；
- `/healthz`、`/metrics`（Prometheus）、结构化日志（logrus，与现状一致）、graceful shutdown。

## 技术选型：方案 A（推荐）——Go SSR + 轻量 JS 岛屿

| 层 | 选择 | 理由 |
|---|---|---|
| 路由 | `net/http`（Go 1.22+ pattern routing） | 标准库足够，路由总数 < 30 条 |
| 模板 | `html/template` + 自定义 FuncMap | 标准库、零依赖；若嫌弃可后续迁 `templ` |
| 数据访问 | `pgx/v5` + `pgxpool`（**已是依赖**） | 手写 SQL——本项目的灵魂就是 SQL，视图承担聚合逻辑 |
| Markdown | `goldmark` + `chroma` | 服务端渲染 `doc.en_doc/zh_doc`，配置对齐现有 hugo.yaml（footnote/table/typographer） |
| 样式 | Tailwind CSS v4 **standalone CLI** | Makefile 一条命令构建，不引入 Node 运行时依赖 |
| 交互 | Alpine.js（~15KB）+ 少量 vanilla JS | 微交互（tab、复制、facet）；重交互只有星图一处 |
| 星图 | 自研 Canvas 2D 模块（~500 行，零依赖） | 1,633 个点对 canvas 是小菜，不需要 three.js/pixi |
| i18n | URL 前缀（en 在根、`/zh` 前缀）+ 嵌入式 message catalog | 与 Hugo 现状完全一致 → URL 天然兼容 |

**前端"现代感"来自设计质量，不来自框架选择。** 内容形态是文档 + 数据表格，
SSR HTML 是这类站点的最优解（SEO、首屏、可访问性、维护成本全部占优）。

### 方案 B（对照）：Next.js / Astro SSR + Go JSON API

| 维度 | A：Go SSR | B：Next.js SSR |
|---|---|---|
| 运行时 | 1 个二进制 | Go API + Node SSR，两套部署 |
| SEO / 首屏 | 原生最优 | 可以做好，但要管好水合与缓存 |
| 组件生态 | 手写（量可控：~10 个组件） | shadcn/ui 等现成生态 |
| i18n | 一套 catalog | 前后端两套 |
| 维护者技能匹配 | Go + SQL（完全命中） | 需持续维护 React 工具链 |
| 构建链 | tailwind + esbuild，两个静态二进制 | node_modules 全家桶 |

结论：B 的收益（组件生态）对本站页面形态收益很低，成本（双运行时）持续存在。
**且 A 不封死后路**：`/api/*` 是一等公民，未来任何重交互页面（如矩阵浏览器）
都可以作为独立前端岛屿挂在子路径上消费同一套 API。

## 性能与缓存：三层

数据本质是**日更**的，这是整个架构最大的红利——动态渲染但几乎一切可缓存。

1. **PostgreSQL**：所有页面查询走索引；搜索走物化视图 `search_idx`（见 03）；目标 p99 < 10ms；
2. **进程内**：view-model 级 TTL 缓存（默认 60s，维度表/汇总 5min）；数据日更 → 命中率 > 99%；
   `reload` 完成后 `NOTIFY pgext_reload` 主动失效（P2，可选）；
3. **HTTP/CDN**：`Cache-Control: s-maxage=300, stale-while-revalidate=86400` + ETag
   （从 `pgext.status` 的 recap 时间戳派生）。匿名 GET 全部可被 Cloudflare 缓存，
   即使源站挂了 CDN 也能继续吐陈旧页面。

**预算**：服务端渲染 p50 < 5ms，搜索 API < 30ms；单台 1C1G VPS 足以支撑全站。

## 访问统计（一石二鸟的落地）

```
middleware ──► buffered channel ──► 批量 INSERT（每 5s / 500 条）
                                        ▼
                    pgext.hit (UNLOGGED)  ──日汇总──►  pgext.hit_daily
```

- 无 cookie、不存原始 IP（只记 path / ext / lang / referrer 域名 / UA 大类 / 日期）；
- `UNLOGGED` 表：写入零 WAL 开销，崩了就丢几分钟数据，无所谓；
- 产出三个别处买不到的信号：**站内热度排序、Trending 榜（7 日 z-score）、中英文流量比**；
- GA 保留与否是独立决策（见 06 开放问题）；站内 popularity 一律用自家数据。

## SQLite 的去留

检视结论：SQLite 只在 `cli/parse_dnf.go` 中用于**读取上游 RPM 仓库的
`primary.sqlite` 元数据**，不是应用存储。与"网站用 PG"零冲突，原样保留。
唯一影响：`go-sqlite3` 是 CGO 依赖（goreleaser 现有 arm/amd 目标已在带 CGO
发布，不构成新问题）；若日后想要纯静态构建，可换 `modernc.org/sqlite`，与本设计无关。

## 部署与迁移（零风险切换）

1. **并行期**：新站跑在 `staging.pgext.cloud`，Hugo 站照常；
2. **对拍**：脚本遍历旧站 sitemap，逐 URL 请求新站，比对 status code 与关键内容（标题、安装命令），出 diff 报告；
3. **切换**：DNS 指向新站（Cloudflare 上改 origin，秒级）；旧 URL 100% 保留（见 04），无需 301 大迁移；
4. **回滚**：DNS 切回即可；
5. **退役**：稳定运行 2–4 周后归档 Hugo 相关目录（`layouts/`、`content/` 中已被 DB 取代的部分）。
   注意退役范围只限本站：`gen io` / `gen cc` / `gen conf` 是给 pigsty.io、pigsty.cc、
   Pigsty 配置文件用的生成器，**继续保留**；退役的只是 `gen page/list/os/matrix`（本站 Markdown 生成）。

运行拓扑：1 台 VPS（systemd 或 docker），PG 同机即可（数据 < 300MB）；
或直接挂在现有 Pigsty 基建上——这本身又是一次 dogfooding。

## 代码组织（新增部分）

```
cmd/server.go          # cobra 子命令入口
srv/
  server.go            # 路由表、中间件（gzip / recover / cache headers / hits）
  render.go            # 模板引擎、i18n FuncMap、goldmark 渲染器
  cache.go             # 进程内 TTL 缓存
  hits.go              # 异步访问统计写入器
  query/               # 每个页面一个查询文件，返回 view-model 结构体
  assets/              # go:embed — templates/ static/ dist/
i18n/
  en.toml  zh.toml     # 界面文案 catalog（数据内容的双语来自 DB）
```

原则：**`cli/`（数据管线）与 `srv/`（网站）互不 import**，共享的只有数据库 schema。
现有 `cli/cache.go` 的 `ExtensionCache` / `Extension`（41 字段，1:1 映射 `extension` 表）
是给 Markdown 生成器用的读模型；新站以 `universe` 为主表、按页面裁剪字段，
建议 `srv/query` 定义自己的 view-model 而不迁就旧结构（若确有共用价值再上提为 `model/` 包）。
