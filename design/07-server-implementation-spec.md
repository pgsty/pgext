# 07 · pgext server 实施规格

> 状态：已确认，可执行。目标是先做一条真实的 Go SSR 垂直切片，再扩展到完整目录站。

## 目标

新增 `pgext server`，把 pgext 从 Hugo 静态生成站升级为 PostgreSQL 驱动的 Go SSR 单二进制 Web 服务。新站的核心体验是 PostgreSQL Extension Universe：搜索、浏览、评估、安装 PostgreSQL 扩展。

## 已确认决策

- 技术路线锁定为 Go SSR 单二进制，不引入 Next.js、Astro 或 Node SSR。
- P1 详情页覆盖 `pgext.universe` 全量扩展，而不是只迁移已打包的 `pgext.extension`。
- `/zh` 在 P1 做全站镜像：模板同构，界面文案走 catalog，数据按 locale 选择 `en_*` 或 `zh_*` 字段。
- P1 不迁移旧 Hugo 里的 `pig/` 与 `release/` 页面；它们暂时不是新 pgext 主题站的一部分。
- 搜索以数据库函数作为稳定接口，MVP 用 PostgreSQL 内置全文检索、前缀/模糊匹配和 `pg_trgm` 排序，后续可以替换为 `zhparser`、`pg_jieba` 或 `pgvector`，不改变 Go handler 和 JSON API。
- P0/P1 先做可维护的 SSR 和数据查询闭环；live Canvas 星图、stats、vendor 页、badge、MCP、语义搜索放到 P2/P3。

## 范围

P0 垂直切片必须包含：

- `pgext server --listen :3000 --dev --cache-ttl 60s` 子命令。
- `srv/` 包，包含路由、中间件、模板渲染、查询层、内存 TTL 缓存和 embedded assets。
- `/healthz`。
- `/` 首页的 SSR 骨架，展示真实 universe/packaged/doc/repo/package summary。
- `/e/` 扩展索引页的 SSR 骨架。
- `/e/<name>/` 和 `/zh/e/<name>/` 详情页，能够渲染一个 packaged 扩展和一个未打包扩展。
- `/search?q=...` HTML 搜索页。
- `/api/search?q=...` JSON 搜索接口。
- `/api/summary`、`/api/ext/<name>` JSON 接口。
- `/robots.txt`、`/sitemap.xml`、`/metrics` 运维与 SEO 基础端点。
- 搜索数据库函数和 `pg_trgm` 安装。

P1 扩展范围：

- 全量 `/e/<name>/` 与 `/zh/e/<name>/`。
- 分类、list、OS/repo 可用性页面。
- 核心旧 URL 兼容，不包含 `pig/` 和 `release/`。
- SEO 基础：canonical、hreflang、sitemap、robots、JSON-LD。其中 P0/P1 垂直切片先落地 sitemap/robots，canonical/hreflang/JSON-LD 随模板深化补齐。
- staging 对拍核心 URL。

## 架构

新增命令：

```bash
pgext server --listen :3000 --database "$PGURL" --dev --cache-ttl 60s
```

文件组织：

```text
cmd/server.go
srv/
  server.go
  config.go
  render.go
  cache.go
  middleware.go
  query/
    db.go
    summary.go
    extension.go
    search.go
  assets/
    assets.go
    templates/
    static/
```

原则：

- `srv/` 不 import 现有 Hugo Markdown 生成器。
- 现有 `cli/` 数据管线继续保留。
- 页面查询使用手写 SQL 和小型 view-model，不复用 `cli.ExtensionCache` 这种生成器读模型。
- `html/template` 负责 SSR；`goldmark` 负责渲染 `doc.en_doc` / `doc.zh_doc`。
- Tailwind standalone 作为 CSS 构建工具；P0 可先提供最小静态 CSS，P1 再引入完整构建链。

## 数据

live DB 的当前关键事实：

- `pgext.universe`: 1,633 rows。
- `pgext.extension`: 531 rows。
- `pgext.doc.en_doc`: 531 rows。
- `pgext.doc.zh_doc`: 531 rows。
- `pgext.gh_repo`: 1,342 rows。
- `pgext.pkg`: 30,720 rows，其中可用 rows 当前约 27k。

实施要求：

- `db/schema.sql` 需要补齐 `pgext.universe` 和 `pgext.gh_repo` 的 DDL，因为 live DB 已存在但仓库 schema 不能完整重建。
- `pgext.summary` 当前只覆盖 `pgext.extension`，不能代表 universe 首页统计。server 需要自己的 summary 查询或新增 view。
- `CREATE EXTENSION IF NOT EXISTS pg_trgm` 必须进入 schema/migration。
- 搜索函数作为 Go 层唯一搜索依赖，例如：

```sql
CREATE OR REPLACE FUNCTION pgext.search_ext(
    q text,
    p_lang text DEFAULT 'en',
    limit_n int DEFAULT 20,
    offset_n int DEFAULT 0
)
RETURNS TABLE (
    name text,
    pkg text,
    category text,
    state text,
    repo text,
    version text,
    description text,
    star_cnt int,
    score real
)
LANGUAGE sql
STABLE
AS $$
    WITH input AS (
        SELECT lower(trim(q)) AS term
    )
    SELECT
        u.name,
        u.pkg,
        u.category,
        u.state,
        u.repo,
        u.version,
        CASE WHEN p_lang = 'zh' THEN coalesce(u.zh_desc, u.en_desc, '')
             ELSE coalesce(u.en_desc, u.zh_desc, '')
        END AS description,
        u.star_cnt,
        (
            CASE WHEN lower(u.name) = (SELECT term FROM input) THEN 1000 ELSE 0 END
            + CASE WHEN lower(u.name) LIKE (SELECT term FROM input) || '%' THEN 100 ELSE 0 END
            + similarity(lower(u.name), (SELECT term FROM input)) * 50
            + log(coalesce(u.star_cnt, 0) + 1)
            + CASE WHEN u.state = 'available' THEN 20 ELSE 0 END
        )::real AS score
    FROM pgext.universe u
    WHERE (SELECT term FROM input) <> ''
      AND (
          lower(u.name) LIKE '%' || (SELECT term FROM input) || '%'
          OR lower(coalesce(u.pkg, '')) LIKE '%' || (SELECT term FROM input) || '%'
          OR lower(coalesce(u.en_desc, '')) LIKE '%' || (SELECT term FROM input) || '%'
          OR lower(coalesce(u.zh_desc, '')) LIKE '%' || (SELECT term FROM input) || '%'
          OR similarity(lower(u.name), (SELECT term FROM input)) > 0.2
      )
    ORDER BY score DESC, u.name
    LIMIT limit_n OFFSET offset_n
$$;
```

后续可以把函数内部替换成 materialized view、全文 rank、中文分词或 vector search。

## 页面合同

已打包扩展详情页：

- header：名称、分类、版本、license、语言、repo/source、stars。
- install block：pig/apt/dnf/sql 命令，PG/OS selector 在 P1 联动。
- availability grid：来自 `pgext.pkg`。
- docs：来自 `pgext.doc`，按 locale 渲染 Markdown。
- relationships：`requires`、`require_by`、`see_also`。
- activity：来自 `universe` 和 `gh_repo`。

未打包扩展详情页：

- 隐藏 install block 和 availability grid。
- 展示 source-only、vendor-only、kernel-specific 或 discovered 状态。
- 展示 URL、vendor/kernel/type、描述、活跃度和相关/替代扩展。
- 页面结构与 packaged 扩展保持一致，未来打包后不需要换 URL。

索引与搜索：

- `/e/` 支持分页、排序和基础 facet。
- 筛选状态放在 URL query。
- `/search` 与 `/api/search` 使用同一个 DB 函数。
- 无 JS 时仍能提交表单并得到 HTML 结果。

## 验收

P0 完成时必须满足：

- `pgext server --listen :3000` 能启动。
- `/healthz` 返回 200。
- `/` 返回真实 DB summary。
- `/e/postgis/` 和 `/zh/e/postgis/` 返回 200。
- 任意一个 `state='n/a'` 的扩展详情页返回 200，并且不展示安装矩阵。
- `/api/search?q=vector` 返回 JSON。
- `/search?q=vector` 返回 HTML。
- 新增 server 单元测试和 handler/query 测试通过。
- 全量 `go test ./...` 的结果要报告；若旧测试仍失败，明确列为既有失败，不把它计入 server 验收。
