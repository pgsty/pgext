## 用法

来源：

- [固定提交的上游 README](https://github.com/dylex/pguri/blob/6b80c2f61b1de04faa3f2e6c1f768e9a8997db30/README)
- [固定提交的 1.0 版安装 SQL](https://github.com/dylex/pguri/blob/6b80c2f61b1de04faa3f2e6c1f768e9a8997db30/pguri--1.0.sql)
- [固定提交的 C 解析器实现](https://github.com/dylex/pguri/blob/6b80c2f61b1de04faa3f2e6c1f768e9a8997db30/pguri.c)
- [固定提交的扩展 control 文件](https://github.com/dylex/pguri/blob/6b80c2f61b1de04faa3f2e6c1f768e9a8997db30/pguri.control)

`pguri` `1.0` 增加经过解析的 `domainname` 和 `uri` 类型，用于保存浏览历史与策略模式。它把 URI 拆分为 scheme、认证信息、domain、port、path、query 和 fragment 字段，提供 domain/URI 包含关系，并包含 URI 感知的全文搜索解析器与配置。

### 核心流程

```sql
CREATE EXTENSION pguri;

SELECT ('https://user@example.com:8443/a/b?q=one#top'::uri).domain;
SELECT 'foo.com'::domainname @> 'www.foo.com'::domainname;
SELECT 'foo.com'::uri @> 'https://www.foo.com/path?q=one'::uri;

SELECT to_tsvector(
  'uri'::regconfig,
  'https://www.example.com/a%20path?q=one#top'
);
```

URI 解析器接受部分值，因此按照扩展的组件与前缀规则，`foo.com` 这样的策略值可以包含更具体的 URI。

### 重要对象

- `domainname`：类似可排序 text 的类型，具有 B-tree 和 hash operator class、比较/拼接操作符、`domainname_parts()`、`domainname_parents()` 以及包含操作符 `@>` 和 `<@`。
- `uri`：带有 `domain`、`port`、`path`、`query`、`auth`、`scheme` 和 `fragment` 字段的解析类型，并提供包含操作符 `@>` 与 `<@`。
- `uri(text)` 与 `text(uri)`：面向现有 text 列的显式转换函数；已复核 SQL 没有安装对应 cast。
- 名为 `uri` 的全文搜索 parser、template、dictionary 和 configuration：切分 URI 组件，并对选定的 domain、path、query 与 fragment token 进行 percent decode。

### 运维说明

包含关系是应用策略逻辑，不等同于 RFC URI 等价性：它比较选定组件，并对 path 或 query 使用前缀匹配。应针对应用预期测试末尾斜线、转义字符、默认端口、大小写、国际化域名、IPv6 literal 和 opaque scheme；需要无损往返时应保留原始文本。

安装 SQL 包含明确的 catalog-update workaround，用于替换已经创建的组合 `uri` 类型的输入/输出函数。C 代码还依赖 PostgreSQL 后端全文搜索与 tuple 内部接口。应把它视为版本敏感的旧代码，在准确的 PostgreSQL 构建上运行往返和索引测试；未经验证，不要假定它在大版本之间具有二进制或 dump/restore 兼容性。
