# 03 · 数据与查询设计

## 现有 schema → 页面的映射

现有数据管线（fetch → parse → recap → reload）**一行不改**。网站是纯读者：

| 页面 / 功能 | 数据来源 |
|---|---|
| 扩展详情页 | `universe`（主档）+ `doc`（双语文档与链接）+ `pkg`（矩阵）+ `gh_repo`（活跃度） |
| 扩展索引 / facet 筛选 | `universe` + `category` + 维度表 `pg` / `os` |
| 分类页 | `category` + `universe` |
| OS 矩阵页 | `pkg` × `active_pg` × `active_os` |
| 仓库页 | `repository` + `pkg`（org 维度） |
| 首页统计带 | `summary` 视图（现成的） |
| 星图 / 宇宙 | `universe`（category / star_cnt / requires）+ 新增 `layout` 表 |
| 榜单 | `gh_repo`（stars / last_commit / release）+ 新增 `hit_daily`（站内热度） |

## 网站专用新对象（全部加性，不碰现有表）

### 1. 搜索物化视图

```sql
CREATE MATERIALIZED VIEW pgext.search_idx AS
SELECT u.id, u.name, u.pkg, u.category, u.repo, u.state, u.star_cnt,
       u.en_desc, u.zh_desc,
       (u.state = 'available')                          AS packaged,
       setweight(to_tsvector('english', u.name), 'A')
       || setweight(to_tsvector('english', coalesce(u.en_desc,'')), 'B')
       || setweight(to_tsvector('simple',  coalesce(array_to_string(u.tags,' '),'')), 'B')
                                                        AS tsv_en,
       to_tsvector('simple', coalesce(u.zh_desc,''))    AS tsv_zh
FROM pgext.universe u;

CREATE INDEX ON pgext.search_idx USING gin (name gin_trgm_ops);
CREATE INDEX ON pgext.search_idx USING gin (tsv_en);
CREATE INDEX ON pgext.search_idx USING gin (zh_desc gin_trgm_ops);  -- CJK 走 trigram
```

`pgext reload` 末尾追加 `REFRESH MATERIALIZED VIEW`。

### 2. 访问统计

```sql
CREATE UNLOGGED TABLE pgext.hit (
    ts       timestamptz NOT NULL DEFAULT now(),
    path     text NOT NULL,          -- 归一化路径，如 /e/postgis
    ext      text,                   -- 命中的扩展名（若有）
    lang     text NOT NULL,          -- en / zh
    referrer text,                   -- 只保留域名
    ua_class text                    -- browser / bot / cli / ai-agent
);
CREATE TABLE pgext.hit_daily (       -- 每日 rollup（cron 或 reload 时顺手做）
    day date, path text, ext text, lang text, hits bigint,
    PRIMARY KEY (day, path, lang)
);
```

### 3. 星图布局（让宇宙稳定）

```sql
CREATE TABLE pgext.layout (
    name  text PRIMARY KEY REFERENCES pgext.universe(name),
    x     real NOT NULL, y real NOT NULL,
    r     real NOT NULL              -- 视觉半径 = f(log(star_cnt))
);
```

由 `pgext reload` 用**固定随机种子**跑一次力导向布局（Go 实现，按 category 聚簇 +
requires 边吸引）写入。前端只读坐标 → **星星的位置每天稳定**，新扩展"诞生"在
所属星座边缘。首页 canvas 与降级 SVG 共享同一份坐标。

## 搜索设计（MVP 一条 SQL 打天下）

```
score = exact_name(1000) + name_prefix(100) + trgm_sim(name)×50
      + ts_rank(tsv)×10 + log(star_cnt+1) + packaged_bonus(20) + hit_bonus
```

- **英文**：name 精确/前缀/trigram + `tsv_en` 全文；
- **中文**：MVP 用 `zh_desc` 上的 pg_trgm（CJK 双字组合下 trigram 实测够用）+ simple tsv；
  P2 可选 `zhparser` / `pg_jieba`（顺便 dogfooding）；P3 上 `pgvector` 语义搜索
  （"我想做地理围栏" → postgis / h3）——搜索框成为自家目录的活广告；
- 返回：name、按 locale 选择的 desc、category、打包状态、安装徽章；目标 < 30ms；
- 同一个 `/api/search` 同时喂 ⌘K 面板和 `/search` 结果页（无 JS 也能用）。

## 数据质量注意事项（设计必须兼容）

| 问题 | 事实 | 对策 |
|---|---|---|
| star 数失真 | postgis 主仓在 osgeo gitea → `star_cnt=9` | 榜单按 `gh_repo` 数据 + 人工白名单修正；详情页标注数据来源 |
| 未打包扩展字段缺失 | 1,102 个 `state='n/a'`（无版本/包/矩阵） | 详情页模板按数据存在性分区渲染，缺失区块整块隐藏 |
| 无 commit 频次曲线 | `activity_json` 只有最近 commit/release | P2 若要 sparkline 需扩展抓取；先用"最近活跃"徽章表达新鲜度 |
| 文档覆盖 531/1633 | 只有打包扩展有长文档 | 未打包扩展详情页以元数据 + GitHub 链接为主体 |
| `doc` 外键指向 `extension` | universe 全量出页时 join 要走 name | 详情页查询以 `universe` 为主表 LEFT JOIN `doc` |

## 数据还能长出什么（增值清单，按 价值 × 成本 排序）

**Now（P1–P2，纯现有数据）**
1. **中文详情页**：531 篇 zh_doc 上线即发布（今天没发布！）；
2. **Trending / 热门榜**：`hit_daily` + `gh_repo` 新鲜度；
3. **依赖关系图**：requires / require_by 已是现成的边表——详情页画局部子图；
4. **云厂商专页** `/vendor/aws` 等：91 个专有扩展 + kernel 归属，回答"RDS 上有什么独家货"；
5. **新增/更新 RSS**：`mtime` + `last_release_date` 直接生成 feed；
6. **统计仪表盘** `/stats`：分类/许可证/语言/仓库分布，全是一条 GROUP BY。

**Next（P2–P3，少量新代码）**
7. **README 徽章服务** `/badge/<name>.svg`："Available on PG 14–18 · EL/Debian/Ubuntu"
   ——扩展作者贴进 README，反链 + 品牌曝光的增长飞轮；
8. **公共 JSON API + `llms.txt` + 每页 `.md` 端点**：让 AI 助手直接消费目录；
9. **MCP Server**（`pgext mcp` 或 server 内置 `/mcp`）：Claude/Codex 用户直接问
   "帮我查 u24 arm 上 pg18 的 timescaledb"——2026 年的差异化入口；
10. **对比页** `/compare/pgvector-vs-vchord`：同类扩展并排（SEO 长尾金矿）。

**Later（需要新数据抓取）**
11. 云厂商**支持矩阵**（RDS/Aurora/AlloyDB 支持哪些通用扩展 × 版本）——高价值，需爬各家文档；
12. commit 活跃度 sparkline、contributor 数、下载量估算。

## 顺手要修的几件小事

1. `AGENTS.md`（CLAUDE.md）中描述的 `pgext.extension_all` 已经演化为 `pgext.universe`，
   文档与实际 schema 不一致——随本次工作顺手更新；
2. **`universe`、`gh_repo` 两张表不在 `db/schema.sql` 里**（12 张表未含它们，属 live-DB 演化产物）。
   网站将把 `universe` 当主表，必须把 DDL 补进 `schema.sql`，保证 `pgext schema` 能完整重建
   （`db/universe.csv` 已随仓库版本化，缺的只是建表语句；`gh_repo` 由 `bin/github_activity.py` 回填）；
3. `bin/gen-ext.py`、`bin/ext-index.py`、`bin/list-*.py` 等已被 Go `gen` 取代的遗留
   Python 生成器，随 Hugo 退役一并清理。
