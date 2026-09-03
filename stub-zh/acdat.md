## 用法

来源：

- [官方 README v0.1.0](https://github.com/Vonng/ac/blob/v0.1.0/README.md)
- [扩展控制文件](https://github.com/Vonng/ac/blob/v0.1.0/acdat.control)
- [版本化安装 SQL](https://github.com/Vonng/ac/blob/v0.1.0/sql/acdat--0.1.0.sql)
- [官方使用指南](https://github.com/Vonng/ac/blob/v0.1.0/docs/usage.md)
- [可执行 SQL 演示](https://github.com/Vonng/ac/blob/v0.1.0/examples/demo.sql)

`acdat` 0.1.0 将大规模精确字面量模式词典编译为不可变的 Aho-Corasick Double-Array machine，再对每个 `text` 或 `bytea` 值执行一次扫描以完成匹配或替换。它适合策略规则、失陷指标、实体名称、脱敏别名等稳定且会被反复使用的词典。

### 核心流程

创建扩展、编译词典，并在大量输入上复用生成的 `acdat.machine` 值：

```sql
CREATE EXTENSION acdat;

WITH machine AS (
    SELECT acdat.compile(
        ARRAY['he', 'she', 'his', 'hers'],
        ARRAY[1, 2, 3, 4]::bigint[]
    ) AS value
)
SELECT acdat.contains('ushers', value) AS matched,
       acdat.info(value)->>'pattern_count' AS patterns
FROM machine;
```

生产词典的源规则应保存在应用自有表中。`acdat.compile()` 的聚合重载可以直接从模式、ID、替换值和优先级记录构建确定性的 machine；一次编译后即可扫描大量输入。

### 匹配与替换

`acdat.contains()` 在首次命中后停止。`acdat.matches()` 返回 `acdat.hit` 行，其中包含模式 ID、字节与字符坐标以及优先级。`acdat.replace()` 执行字面量、非递归替换：

```sql
WITH machine AS (
    SELECT acdat.compile(
        ARRAY['病毒', '特征码', '病毒特征码'],
        ARRAY[10, 11, 12]::bigint[],
        ARRAY['[VIRUS]', '[SIGNATURE]', '[IOC]'],
        ARRAY[20, 20, 5]::integer[]
    ) AS value
)
SELECT *
FROM acdat.matches('发现病毒特征码', (SELECT value FROM machine), 'all_overlapping');

SELECT acdat.replace(
    'aaa',
    acdat.compile(
        ARRAY['a', 'aa', 'aaa'],
        ARRAY[1, 2, 3]::bigint[],
        ARRAY['[x]', '[yy]', '[zzz]']
    ),
    'leftmost_longest'
);
```

匹配策略包括 `all_overlapping`、`leftmost_longest` 和 `leftmost_priority`。替换只能使用非重叠策略。可通过 `acdat.info()` 检查已编译 machine，并在移动或校验产物时使用导出、验证、导入和指纹函数。

`acdat.matches()` 的 `max_matches` 默认值为 10000，`acdat.replace()` 的 `max_output_bytes` 默认值为 268435456。对于不可信或高命中输入，应设置更严格的上限，确保命中枚举和替换输出有界。

### 托管词典

可选的目录层用于发布不可变、内容寻址的 build，并以原子方式选择一个活动 build。其控制函数使用 `SECURITY INVOKER`，且不向 `PUBLIC` 授予执行权限：

```sql
WITH machine AS (
    SELECT acdat.compile(pattern, pattern_id)
    FROM app_keyword
    WHERE enabled
), published AS (
    SELECT acdat.publish('moderation', 1, machine) AS build_id
    FROM machine
)
SELECT acdat.activate('moderation', build_id)
FROM published;

SELECT name, version, build_id, machine
FROM acdat.active_machine
WHERE name = 'moderation';
```

应用表始终是事实源。逻辑备份包含目录元数据和活动 machine 载荷，但不会包含所有历史产物，因此应保留重建已退休或非活动版本所需的源模式。

### 兼容性与安全

0.1.0 已在 PostgreSQL 14 至 18 上测试，不需要预加载或重启服务器，没有外部扩展依赖，也不定义 GUC。控制文件将 schema 固定为 `acdat`，并设置 `relocatable = false` 和 `trusted = false`，因此 `CREATE EXTENSION` 需要超级用户。

ACDAT 索引的是模式词典，而不是文档表：扫描已有大表时仍需读取候选行。匹配是精确且区分大小写的；扩展不提供正则表达式、模糊匹配、分词、自动大小写折叠、Unicode 规范化或文档侧索引。文本引擎支持 UTF-8 和单字节服务器编码，二进制数据应使用 bytea 接口。需要反复反向查询时，应将 `(document_id, pattern_id)` 命中物化到应用表中。

编译格式具备自描述和校验和，导入的产物会在使用前接受验证。卸载前应清点依赖：`DROP EXTENSION acdat` 会删除托管词典状态，而添加 `CASCADE` 还可能删除依赖 `acdat.machine` 的用户列或其他对象。
