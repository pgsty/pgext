## 用法

来源：

- [pg_opencc 0.1.0 README](https://github.com/pgsty/pg_opencc/blob/v0.1.0/README.md)
- [扩展控制文件](https://github.com/pgsty/pg_opencc/blob/v0.1.0/opencc.control)
- [SQL API](https://github.com/pgsty/pg_opencc/blob/v0.1.0/sql/opencc--0.1.0.sql)
- [发布说明](https://github.com/pgsty/pg_opencc/releases/tag/v0.1.0)

`opencc` 由 `pg_opencc` 项目提供，用于简体中文、繁体中文、台湾与香港用语以及日文新字体之间的转换。0.1.0 版本将 OpenCC 1.4.2 及其词典嵌入共享库，运行时无需系统 OpenCC 或外部词典文件。

### 启用与转换

支持 PostgreSQL 14–18，要求数据库使用 `UTF8` 编码。这是受信任扩展，拥有数据库 `CREATE` 权限的用户即可安装。对象位于固定的 `opencc` 模式中，扩展不可迁移模式。基本使用无需预加载或重启服务器。

```sql
CREATE EXTENSION opencc;

SELECT opencc.s2twp('鼠标里面的硅二极管坏了，导致光标分辨率降低。');
SELECT opencc.convert('繁體中文', 't2s');
SELECT * FROM opencc.configs;
```

`opencc.convert(input, config)` 用于明确指定转换配置。`opencc.s2t()`、`opencc.t2s()` 等快捷函数与各个配置一一对应：

| 配置 | 转换用途 |
| --- | --- |
| `s2t`, `t2s` | 简体中文与繁体中文互转 |
| `s2tw`, `tw2s`, `s2hk`, `hk2s` | 台湾或香港字形转换 |
| `s2twp`, `tw2sp` | 台湾地区词汇转换 |
| `s2hkp`, `hk2sp` | 香港地区词汇转换，仍在完善中 |
| `t2tw`, `tw2t`, `t2hk`, `hk2t` | 繁体中文不同地区字形互转 |
| `t2jp`, `jp2t` | 日文旧字体与新字体互转，实验性功能 |

### 地区转换与结果检查

`opencc.localize(input, target, source)` 接受 `zh-Hans`、`zh-Hant`、`zh-CN`、`zh-TW`、`zh-HK` 和 `ja`。默认的 `source => 'auto'` 可以判断文字类别，但无法区分台湾和香港词汇。需要准确转换地区用语时，应明确指定来源地区。

```sql
SELECT opencc.localize('滑鼠裡面的矽二極體壞了', 'zh-CN', source => 'zh-TW');
SELECT opencc.detect('中文'), opencc.script_stats('中文');
SELECT * FROM opencc.routes;
SELECT * FROM opencc.ambiguities('大战文丑的时候，他的头发很干燥', 's2t');
SELECT opencc.candidates('发', 's2t');
SELECT * FROM opencc.explain('数据库连接池', 's2twp');
```

| 对象 | 用途 |
| --- | --- |
| `opencc.detect()` | 判断文字类别，包括混合文字与无法区分简繁的文字 |
| `opencc.script_stats()` | 统计汉字、简体字、繁体字和假名 |
| `opencc.configs`, `opencc.routes` | 检查配置与地区转换链 |
| `opencc.ambiguities()`, `opencc.candidates()` | 检查歧义片段和词典候选项 |
| `opencc.explain()` | 检查各阶段的转换结果 |
| `opencc.version()` | 返回扩展版本、引擎版本和词典摘要 |
| `opencc.dependents()` | 查找依赖对象并给出建议的重建语句 |

### 运行与升级

转换函数在同一扩展版本内具有 `IMMUTABLE STRICT PARALLEL SAFE` 属性，可用于表达式索引和存储生成列。PostgreSQL 18 的生成列若调用自定义转换函数，必须明确指定 `STORED`。

未来扩展版本更新词典时，转换结果可能变化。升级时应检查 `opencc.version()`，并审阅 `opencc.dependents()` 返回的语句。依赖检查无法发现封装在其他自定义函数中的调用。

```sql
SELECT * FROM opencc.version();
SELECT * FROM opencc.dependents();
```

每个后端按需初始化转换器。也可选择在服务器启动时预加载，使后端共享相关内存；修改 `shared_preload_libraries` 需要重启：

```ini
shared_preload_libraries = 'opencc'
opencc.preload = 's2t,t2s,s2twp,tw2sp'
```

若已有预加载列表，将 `opencc` 加入其中。`opencc.preload` 也接受 `all`。删除扩展后会保留空模式。
