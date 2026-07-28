# 06 · 路线图、风险与开放问题

## 分期（节奏按"单人 + Claude 结对"估）

### P0 · 骨架（~3 天）
`server` 子命令、路由表、模板引擎 + i18n、pgxpool、TTL 缓存、hit 记录器、
healthz/metrics、Tailwind 构建链、部署脚本。
**验收**：`pgext server` 能用两种语言渲染一个真实详情页。

### P1 · 平价迁移 + 解锁（~1–2 周）→ 切换 DNS
- 全部现有 URL 重实现（详情 / 索引 / list / os 矩阵 / repo / pig / release）；
- 详情页扩到 **universe 全量 1,633**；`/zh` 全镜像（解锁 531 篇中文文档）；
- `/api/search` + ⌘K + `/search` 页；`/categories/<slug>` 分类页重实现；
- `static/` 下现有 `/galaxy`、`/matrix` SPA 原样托管（P2 再重建）；
- 首页 P1 版（静态 SVG 星图 hero + 搜索 + 统计带 + 分类格 + 榜单）；
- SEO 全套（canonical/hreflang/JSON-LD/sitemap/RSS 路径保留）；
- staging 对拍报告 → DNS 切换。
**验收**：旧站 sitemap 100% URL 在新站 200；中文详情页可用；Lighthouse ≥ 95。

### P2 · 宇宙（~1 周）
活的星图 hero + `/galaxy` 全屏、`/stats` 仪表盘、Trending（hit 积累两周后开榜）、
`/vendor/*` 云厂商页、README 徽章服务、动态 OG 图、`llms.txt` + `.md` 端点。

### P3 · 生态（按需排期）
pgvector 语义搜索、对比页 `/compare/a-vs-b`、MCP server、
admin 策展台（或继续用 psql，动态渲染本来就即时生效）、
云厂商支持矩阵（需新数据抓取）。

## 风险表

| 风险 | 等级 | 对策 |
|---|---|---|
| SEO 迁移抖动 | 中 | URL 100% 保留、SSR、sitemap lastmod、并行期对拍、Search Console 盯两周 |
| 运维面扩大（静态站 → 服务） | 中 | 单二进制 + systemd；CDN stale-while-revalidate 兜底源站故障；PG 日备份（数据可由管线重建） |
| 中文搜索质量 | 低 | MVP trgm 够用；不满意再上 zhparser/pg_jieba（都在自家仓库里，一条 pig 命令的事） |
| star 数失真（osgeo 等非 GitHub 仓库） | 低 | 榜单白名单修正；详情页标注数据口径 |
| 星图沦为花瓶 | 低 | 星图=真实数据投影 + 搜索联动；任何信息都有 HTML 等价物 |
| 双语界面文案维护 | 低 | catalog 就两个文件；数据双语已在 DB 里 |

## 成本估算

- **托管**：1 台小 VPS（同机 PG，数据 < 300MB）+ Cloudflare 免费档；
- **不新增**任何 SaaS 依赖；GA / giscus 去留是独立决策；
- 数据管线成本不变（本来就每天在跑）。

## 开放问题（请逐条表态，这是头脑风暴的议程）

1. **技术路线**：方案 A（Go SSR + 轻岛屿，我的推荐）还是方案 B（Next.js + Go API）？
   ——见 02 的对照表。
2. **详情页范围**：universe 全量 1,633 都出页（我的推荐，3× SEO 面），
   还是先只做打包的 531？
3. **`/zh` 策略**：全站镜像，还是内容页优先、工具页（stats/galaxy）后补？
4. **pig / release 文档的归宿**：goldmark 渲染 repo 内 markdown（embed），
   还是入库 `pgext.doc` 走同一套渲染？（我倾向 embed：这些是"关于软件的文档"，
   跟着代码版本走更自然。）
5. **GA 与 giscus**：自家 hit 统计上线后 GA 还留吗？评论区（giscus）要不要保留？
6. **首页主题**：跟随系统（现状，我的推荐——昼夜两种星图都已设计）还是强制深色开场？
7. **Hugo 退役时点**：DNS 切换后立即归档，还是保留一个发版周期？
8. **命名主线**："Extension Universe"（宇宙叙事，我的推荐）还是沿用
   "Extension Catalog"（目录叙事）？影响首页文案与 OG 呈现。
9. **P2 里你最想先要哪个**：活星图 / stats / vendor 页 / 徽章 / MCP？
   （决定 P2 内部顺序。）
