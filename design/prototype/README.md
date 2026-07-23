# PGEXT.CLOUD 重设计原型

单文件静态原型：把 `pgext.universe` 里全部 **1,633 个扩展**装进一个自包含的 HTML（约 790 KiB，零依赖、零外部请求），
先用静态方式把新版 PGEXT.CLOUD 的交互与视觉完整跑通，后续再由 `pgext server` 动态化。

```bash
open design/prototype/index.html        # 直接本地打开即可，无需构建
```

## 设计概念：目录即数据库（The Catalog IS a Database）

整站的视觉与文案都从"这是一份用数据库供给的目录"出发，两个签名元素：

1. **宇宙点阵（Universe Field）**——首页 hero 下方的全宽 canvas 点阵，
   每一个点就是一个扩展，按 16 分类着色、按分类分组排布，1,633 个点一屏尽收。
   悬停显示名称与描述，点击直达详情页；已打包扩展实色、未打包半透明。
2. **实时 SQL 读出（Live SQL Readout）**——搜索与筛选状态实时编译成一条真实可执行的
   `SELECT * FROM pgext.universe WHERE …`，行尾配 psql 风格的 `(N rows)` 计数，点击复制。
   搜索框本身即 `pgext=#` 提示符，并支持 `cat:GIS lang:rust pg:18` 操作符。

配套的语气细节：品牌角标是 `\dx`，404 页是 psql 报错
（`ERROR:  extension "xxx" does not exist` + `HINT`），页脚以 `\q` 收尾。

## 视觉系统

- **字体三声部**：衬线展示体（Charter/Georgia/宋体——"年鉴"气质，用于大标题与扩展副标）
  + 等宽数据体（所有标识符、版本、SQL、徽章）+ 系统无衬线正文。全部系统字体，无外部加载。
- **双主题**：墨蓝深色（默认跟随系统）/ 纸白浅色，`data-theme` 覆盖。
- **16 分类色板**：按 TIME→ETL 的固定顺序分配 16 个色相，
  亮暗两套均通过 dataviz 校验器（明度带、色度下限、CVD 相邻对分离、表面对比度四项检查；
  CVD 处于 8–12 floor 区间，站内所有分类标记均伴随文字标签作为次级编码，合规）。
  色值见 `style.css` 中 `--c-TIME` … `--c-ETL`。

## 页面与路由（hash router）

| 路由 | 内容 |
|---|---|
| `#/` | Landing：hero 数字叙事 → 宇宙点阵 + 图例 → 搜索台（SQL 读出）→ 筛选条 → 1,633 卡片墙（卡片/行两种密度，懒加载分页） |
| `#/ext/{name}` | 扩展详情：全部 1,633 个都有页面。安装标签页（SQL/pig/yum/apt）、可用性矩阵（EXT/RPM/DEB × PG14–18）、依赖与关联（requires/被依赖反推/see also/同包）、规格表侧栏、双语用法文档、同类推荐 |
| `#/cat/{CODE}` | 分类页：16 分类迷你光谱导航 + 精选大卡（优先带文档的旗舰）+ 全量表格 |
| `#/browse` | 多维索引枢纽：分类/许可证/语言/仓库/云厂商/PG 版本/形态 7 个维度卡 |
| `#/dim/{key}` | 维度表：取值 × 条形 × 计数 × 示例，每行都是即时筛选入口 |
| `#/about` | 关于 + 路线图（Matrix / Galaxy / pgext server） |

**通用索引能力**的实现方式：所有维度页与徽章都路由到同一个带参 Landing
（`#/?license=MIT`、`#/?cat=GIS&q=vector`……），一套 faceted 引擎服务所有入口。

## 双语

- 全站 UI 文案 `I18N` 字典（en/zh），导航栏一键切换，localStorage 持久化，默认英文；
  URL 加 `?lang=zh` 可指定语言（便于分享）。
- 数据层自带 `en_desc`/`zh_desc`；分类中英文描述来自 `pgext.category`。
- 12 个演示扩展内置 stub/stub-zh 的完整双语用法文档（postgis、pgrouting、ogr_fdw、
  timescaledb、pg_cron、pg_partman、vector、citus、pg_search、pgmq、age、pg_stat_statements），
  其余扩展显示占位说明——正式版 555 篇全量供给。

## 数据管线

`data/` 下是数据库快照（2026-07-10），刷新方式：

```bash
# ext.json —— 1,633 行精简列（位置数组编码，GitHub URL 前缀压缩）
PGUSER=postgres psql -h /tmp -p 5432 -d data -XAtc "select json_agg(json_build_array(
  name, nullif(pkg,name), nullif(lead_ext,name), category, (state='available')::int, url, license, lang, repo, version,
  (lead::int)|(contrib::int<<1)|(has_bin::int<<2)|(has_lib::int<<3)|(need_ddl::int<<4)|(need_load::int<<5)|(trusted::int<<6)|(relocatable::int<<7),
  case when schemas='{}' then null else schemas end, pg_ver,
  case when requires='{}' then null else requires end,
  case when see_also='{}' then null else see_also end,
  nullif(rpm_ver,version), nullif(rpm_repo,repo), rpm_pkg, case when rpm_pg=pg_ver then null else rpm_pg end,
  nullif(deb_ver,version), nullif(deb_repo,repo), deb_pkg, case when deb_pg=pg_ver then null else deb_pg end,
  ext_type, nullif(ext_kernel,''), nullif(ext_vendor,''), star_cnt, last_commit_date,
  en_desc, zh_desc) order by coalesce(star_cnt,-1) desc, name) from pgext.universe;" > data/ext.json

# cat.json —— 16 分类元数据
PGUSER=postgres psql -h /tmp -p 5432 -d data -XAtc \
  "select json_agg(json_build_array(id,name,zh_desc,en_desc) order by id) from pgext.category;" > data/cat.json

# docs.json —— 12 个演示扩展的双语文档（stub/ + stub-zh/ 原文打包）
# 见 build 历史，或直接沿用现有快照

python3 build.py     # → index.html（本地/部署用）+ artifact.html（Artifact 用）
```

注意：`require_by` 不入库快照，由前端从 `requires` 反推；`pg_ver` 数组在解码层统一转数字。

## 已知边界（原型阶段刻意为之）

- 详情页"用法文档"仅 12 篇内置；其余占位。
- 可用性矩阵按 RPM/DEB 家族聚合展示，未展开到 OS×架构粒度（那是 Phase 2 Matrix 的主场）。
- 无真实 URL 路由与 SEO（hash router）；`pgext server` 阶段换 SSR 即可复用全部模板逻辑。
- 数据全量内嵌（约 620 KiB JSON），动态化后改为 API 供给。
