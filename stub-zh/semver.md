

## 用法

> [semver: 语义版本号数据类型](https://github.com/theory/pg-semver)

`semver` 扩展提供了实现[语义版本号 2.0.0](https://semver.org/spec/v2.0.0.html)的数据类型。

```sql
CREATE EXTENSION semver;

SELECT '1.2.1'::semver;
SELECT '1.2.0'::semver > '1.2.0-b1'::semver;  -- true（预发布版 < 正式版）
```

### 运算符

| 运算符 | 说明 | 示例 | 结果 |
|----------|-------------|---------|--------|
| `=` | 等于 | `'1.2.0'::semver = '1.2.00'::semver` | `t` |
| `<>` | 不等于 | `'1.2.0'::semver <> '1.2.00'::semver` | `f` |
| `<` | 小于 | `'3.4.0-b1'::semver < '3.4.0'::semver` | `t` |
| `<=` | 小于等于 | `'3.4.0-b1'::semver <= '3.4.0'::semver` | `t` |
| `>` | 大于 | `'3.4.0-b1'::semver > '3.4.0'::semver` | `f` |
| `>=` | 大于等于 | `'3.4.0-b1'::semver >= '3.4.0'::semver` | `f` |

### 函数

| 函数 | 说明 | 示例 | 结果 |
|----------|-------------|---------|--------|
| `to_semver(text)` | 宽松解析 | `to_semver('1.0')` | `1.0.0` |
| `is_semver(text)` | 验证格式 | `is_semver('1.2.0')` | `true` |
| `semver(text)` | 严格转换 | `semver('1.2.1')` | `1.2.1` |
| `get_semver_major(semver)` | 主版本号 | `get_semver_major('4.2.0')` | `4` |
| `get_semver_minor(semver)` | 次版本号 | `get_semver_minor('4.2.0')` | `2` |
| `get_semver_patch(semver)` | 补丁版本号 | `get_semver_patch('4.2.0')` | `0` |
| `get_semver_prerelease(semver)` | 预发布部分 | `get_semver_prerelease('2.1.0-b2+bfb13')` | `b2` |

支持从 `text`、`numeric`、`real`、`double precision`、`integer`、`bigint`、`smallint` 类型的转换。

### 范围类型

`semverrange` 类型支持标准范围运算符：

```sql
SELECT '1.0.5'::semver <@ '[1.0.0, 2.0.0)'::semverrange;  -- true
```

### 聚合函数

支持 `MIN(semver)` 和 `MAX(semver)`。可使用 Btree 和 Hash 索引。
