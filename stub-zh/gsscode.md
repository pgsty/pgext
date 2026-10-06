## 用法

来源：

- [README](https://api.pgxn.org/src/gsscode/gsscode-1.1.3/README.md)
- [Control file / 控制文件](https://api.pgxn.org/src/gsscode/gsscode-1.1.3/gsscode.control)
- [SQL](https://api.pgxn.org/src/gsscode/gsscode-1.1.3/gsscode--1.1.3.sql)

`gsscode` 将九字符的英国 ONS/GSS 地理代码压入 32 位。1.1.3 包含修正后的前缀搜索语义：前缀匹配并非相等关系，不能占用 B-tree 的相等策略位置。

### 核心工作流

```sql
CREATE EXTENSION gsscode;
CREATE TABLE areas (code gsscode PRIMARY KEY, label text);
INSERT INTO areas VALUES ('E01000001', 'Example area');
SELECT code, country(code), gss_type(code), area(code)
FROM areas
WHERE code >= gsscode_range_lower('E01')
  AND code < gsscode_range_upper('E01');
```

### 匹配与登记表

`%` 和 `!%` 仍是布尔前缀过滤器，也支持数组形式，但自身不再通过 B-tree 索引加速前缀匹配。应使用 `gsscode_range_lower` 和 `gsscode_range_upper` 表达真正的左闭右开区间，非法前缀长度会报错。

`is_valid_gss`、`country`、`gss_type` 和 `area` 检查词法形式或打包后的组成部分；`description` 与 `type_info` 查询内部 `gsscode_types` 登记表，`isnan` 识别保留的区域代码。还提供兼容文本的正则表达式与 `left`，但它们不会自动使用基础类型的 B-tree 索引。

### 升级与边界

1.1.3 将 `is_valid(text)` 改名为 `is_valid_gss(text)`，避免与其他扩展冲突。升级后需要修改 SQL 调用处；扩展升级会保留函数身份及依赖。1.1.2 还增加了与文本的相等比较及从文本赋值的类型转换。

已有 1.0.0 安装在更新文件后应运行 `ALTER EXTENSION gsscode UPDATE`，移除不正确的运算符族条目。只替换共享库不能完成这一目录修复。

输入格式合法不代表 ONS 实际分配了该代码。登记表刷新由 `gsscode_ons_refresh` 提供，它的控制版本仍为 1.0.0。核心扩展可重定位，无需预加载；发行版未声明 PostgreSQL 主版本矩阵。
