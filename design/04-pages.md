# 04 · 信息架构与路由

## 原则

1. **旧 URL 一个不死**：Hugo 发布过的路径在新站原样可达（不是 301，是同路径直接渲染）；
2. **每个 HTML 页面都有对应的 JSON / Markdown 形态**（`/api/...`、`.md` 后缀）——人、程序、AI 三种读者；
3. `/zh` 前缀镜像全站（Hugo 现状即如此，`defaultContentLanguage: en` 不带子目录）。

## 路由总表

| 路径 | 今天（Hugo） | 新站 | 阶段 |
|---|---|---|---|
| `/` | hextra 首页 | **扩展宇宙首页**（星图 + 搜索） | P1 简版 / P2 完整 |
| `/e/` | 扩展索引 | SSR 数据表 + facet 筛选 | P1 |
| `/e/<name>/` | 531 个详情页 | **1,633 个详情页**（universe 全量） | P1 |
| `/list/{cate,lang,license,pkg,ext}/` | 5 张分组列表 | 同 URL 重实现 | P1 |
| `/os/matrix/`、`/os/<os>/` | 17 张矩阵页 | 同 URL 重实现（交互升级 P2） | P1 |
| `/categories/<slug>/` | Hugo taxonomy 自动生成 ×16 | 同 URL 重实现为正式分类页（16 个 SEO 入口） | P1 |
| `/repo/<name>/` | 仓库页 ×10 | 同 URL 重实现 | P1 |
| `/pig/`、`/pig/cmd/*` | 手写 pig 文档 | goldmark 渲染（embed 或入库，见 06 问题 6） | P1 |
| `/release/*` | 发版记录 | 同上 | P1 |
| `/zh/**` | 仅部分页面有中文 | **全站镜像**，解锁 531 篇中文文档 | P1 |
| `/galaxy/` | `static/` 手工 SPA（一次性拉 ~3.2MB JSON） | P1 原样托管；P2 用 `layout` 表重建，首屏载荷降到 ~100KB | P1/P2 |
| `/matrix/` | `static/` 自包含 HTML（~7.8MB） | P1 原样托管；P2 动态矩阵浏览器 | P1/P2 |
| `/search?q=` | flexsearch 客户端 | SSR 结果页（同一 API 喂 ⌘K） | P1 |
| `/stats/` | 无 | 生态统计仪表盘 | P2 |
| `/vendor/<name>/` | 无 | 云厂商专页（aws / alibaba / tencent …） | P2 |
| `/api/search`、`/api/ext/<name>`、`/api/summary` | 无 | 公共 JSON API | P1 |
| `/badge/<name>.svg` | 无 | README 徽章 | P2 |
| `/llms.txt`、`/e/<name>.md` | 无 | AI 可读形态 | P2 |
| `sitemap.xml`、`robots.txt`、`/feed/*.xml` | Hugo 生成 | 动态生成 | P1 |

## 详情页（本站的原子单位，做到"高质量"的主战场）

现状是巨型 Markdown 表格墙；新版按"回答问题的顺序"重排：

```
┌──────────────────────────────────────────────────────────┬───────────────┐
│ ⌂ / GIS / postgis                                        │  METADATA     │
│                                                          │  版本   3.6.4 │
│ postgis   [GIS] [GPL-2.0] [C] [PGDG]  ★ 1.2k            │  语言   C     │
│ PostGIS 几何与地理空间类型与函数                            │  License ...  │
│                                                          │  DDL / 加载 / │
│ ┌ INSTALL ─────────────────────────────────────────────┐ │  信任 / 可迁移 │
│ │ [pig] [apt] [dnf] [SQL]      PG: 18▾   OS: u24.arm▾ │ │  ──────────── │
│ │ $ pig ext install postgis                        ⧉  │ │  ACTIVITY     │
│ │ ▸ apt install postgresql-18-postgis-3               │ │  commit  5月   │
│ └──────────────────────────────────────────────────────┘ │  release 4月   │
│                                                          │  ──────────── │
│ AVAILABILITY  （pkg 矩阵 → 紧凑热力格）                    │  LINKS        │
│        PG18 PG17 PG16 PG15 PG14                          │  repo/docs/…  │
│  el8    ██   ██   ██   ██   ██     ← hover 显示包名+版本  │  ──────────── │
│  el9    ██   ██   ██   ██   ██                           │  RELATED      │
│  d12    ██   ██   ██   ██   ██                           │  h3, pgrouting│
│  u24    ██   ██   ██   ██   ██                           │  pointcloud … │
│                                                          │               │
│ USAGE  （doc.en_doc / zh_doc → goldmark 渲染，带 TOC）     │               │
│ DEPENDENCIES （requires ↔ require_by 局部关系图/列表）      │               │
│ PACKAGES （RPM/DEB 包名、版本、依赖明细表）                 │               │
└──────────────────────────────────────────────────────────┴───────────────┘
```

关键交互：**Install 块的 PG / OS 选择器是全页状态**——切换后命令、矩阵高亮、
包名全部联动；选择记入 localStorage（DBA 通常只关心自己的一种环境）。

未打包扩展（1,102 个）：隐藏 Install/矩阵区块，以元数据 + GitHub 活跃度 +
"这是 XX 云专有扩展" / "尚未打包" 状态说明为主体，并给出 see_also 的已打包替代品。
**这是 3 倍于现状的 SEO 表面积。**

## 索引页 `/e/`

```
┌ FACETS ──────┬───────────────────────────────────────────────┐
│ Category 16  │  🔍 filter…        1,633 exts   [表格|卡片]    │
│ PG  18…14    │  NAME      CATEGORY  VER   PG    OS    ★  DESC │
│ OS  el/d/u   │  postgis   GIS      3.6.4 14-18 ████  1k  ...  │
│ Repo         │  timescaledb TIME   2.x   14-18 ███░  18k ...  │
│ License      │  citus     OLAP     13.x  14-18 ████  11k ...  │
│ Lang         │  …（服务端分页，URL query 承载全部筛选状态）      │
└──────────────┴───────────────────────────────────────────────┘
```

- 筛选状态全部在 URL query（可分享、可收藏、可被搜索引擎收录）；
- htmx 局部刷新表格；**无 JS 时表单提交整页刷新，功能不减**；
- 列可排序：name / stars / 最近更新 / 站内热度。

## 首页（概述；视觉细节见 05）

```
┌────────────────────────────────────────────────┐
│            ✦ 星图 hero（真实数据全屏 canvas）      │
│     The PostgreSQL Extension Universe          │
│   ┌──────────────────────────────────────┐     │
│   │ 🔍  Search 1,633 extensions…    ⌘K  │     │
│   └──────────────────────────────────────┘     │
│   [TIME] [GIS] [RAG] [OLAP] … ← 16 星座图例     │
├────────────────────────────────────────────────┤
│  1,633 exts · 531 packaged · 5 PG · 16 OS      │  ← summary 视图，数字滚动
├────────────────────────────────────────────────┤
│  16 分类卡片格（icon + 计数 + top3 扩展）          │
├────────────────────────────────────────────────┤
│  Top Starred │ Recently Updated │ Trending 7d   │
├────────────────────────────────────────────────┤
│  最近收录时间线 · Powered-by 徽章 · 页脚           │
└────────────────────────────────────────────────┘
```

## SEO 与元数据

- 每页 `canonical` + en/zh 成对 `hreflang`（**不做**基于 Accept-Language 的自动跳转，只做顶部提示条——避免搜索引擎收录混乱）；
- 详情页注入 JSON-LD `SoftwareApplication`（name/description/license/softwareVersion/operatingSystem）；
- `sitemap.xml` 从 DB 动态生成（含 lastmod = `mtime`）；
- OG 卡片：P1 用统一模板图，P2 每扩展动态生成 PNG（名称 + 分类色 + 矩阵缩略）；
- 保留 Hugo 时代的 section RSS 路径，避免订阅者断链。
