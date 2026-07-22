## 用法

来源：

- [官方 database.dev 软件包页面](https://database.dev/raminder/countries)
- [官方 database.dev 安装指南](https://database.dev/docs/install-a-package)

`countries` 是一个 database.dev 可信语言扩展，会创建名为 `countries` 的小型查找表。尽管软件包描述另有表述，发布的 `0.0.4` 版 SQL 实际只包含六行，因此不能把它当作完整或基于标准的国家数据集。

### 核心流程

先安装 `pg_tle` 前置扩展，再使用 dbdev CLI 为指定软件包版本生成迁移；审核生成的 SQL 后，通过常规迁移流程应用。

```bash
dbdev add -o ./migrations -s public -v 0.0.4 package -n "raminder@countries"
```

迁移创建扩展后，务必先检查表内容再使用。

```sql
TABLE public.countries ORDER BY name;
```

发布的 `0.0.4` 安装脚本只定义一列，并写入 Afghanistan、Albania、Algeria、Zambia、Zimbabwe 和 India 六个值。

```sql
SELECT count(*) AS row_count
FROM public.countries;
```

### 重要对象

- `public.countries`：仅含 `name text PRIMARY KEY` 一列的表。
- `pg_extension_config_dump`：安装脚本把 `public.countries` 注册为扩展配置数据，使其内容可进入扩展感知的转储。

### 运维说明

该软件包通过 database.dev 分发并依赖 `pg_tle`，不是传统的原生扩展归档。即使生成软件包时指定其他目标模式，`0.0.4` 版仍把名称硬编码为 `public.countries`。随附的六个名称没有 ISO 编码、地区、本地化名称、生命周期状态或更新机制。应用若需要权威国家列表，应改用持续维护、源自 ISO 的数据集。编辑这些行时应按应用数据处理，并在变更种子数据前测试转储、恢复与扩展升级行为。
