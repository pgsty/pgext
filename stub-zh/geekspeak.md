## 用法

来源：

- [官方 README](https://github.com/ttfkam/pg_geekspeak/blob/8d176e395f7f8a98a3c4dd2d48a1ee1e359b4e47/README.md)
- [官方扩展 SQL](https://github.com/ttfkam/pg_geekspeak/blob/8d176e395f7f8a98a3c4dd2d48a1ee1e359b4e47/geekspeak--1.0.0.sql)
- [官方扩展控制文件](https://github.com/ttfkam/pg_geekspeak/blob/8d176e395f7f8a98a3c4dd2d48a1ee1e359b4e47/geekspeak.control)

`geekspeak` 会安装 GeekSpeak 播客网站使用的完整应用模式：节目、人员、会话、新闻条目、节目片段、认证辅助函数、搜索和基于文件系统的媒体清单。它是面向特定应用的数据定义，并非通用播客工具包；安装会在所选模式中创建大量全局命名对象。

### 前置条件与安装

版本 `1.0.0` 依赖 `btree_gist`、`isn`、`multicorn`、`pgcrypto` 和 `plpgsql`。其 SQL 会向两个既有角色授权并定义 Multicorn 文件系统表，因此安装前应有意识地创建或映射这些角色：

```sql
CREATE ROLE geekspeak_org NOLOGIN;
CREATE ROLE geekspeak_audit NOLOGIN;
CREATE EXTENSION geekspeak CASCADE;
```

模式把 `/var/www/geekspeak.org/media` 硬编码为 `episode_media_fdt` 和 `episode_misc_fdt` 的根目录。刷新物化媒体视图前，PostgreSQL 服务账号和 Multicorn Python 环境必须能够访问该目录。

### 主要对象组

- `episodes`、`participants`、`people`、`locations`、`headlines`、`bits` 和 `bit_templates` 保存节目及编辑数据。
- `passwords`、`acls`、`sessions` 和 `active_sessions` 实现网站的密码、角色与持久会话模型。
- `episode_media_fdt`、`episode_misc_fdt` 和 `episode_media` 通过 `multicorn.fsfdw.FilesystemFdw` 暴露并缓存文件。
- `login()`、`logout()`、`register()`、`recover()`、`confirm()` 和 `update_password()` 实现应用认证流程。
- `episode_as_json()`、`headlines_as_json()`、`bits_as_json()` 及相关辅助函数塑造 API 输出。
- `geekspeak.session_duration` 控制会话期限，默认为一个月。

### 读取示例

准备好应用数据与媒体后，可通过扩展提供的辅助函数和视图读取：

```sql
SELECT episode_num(current_date) AS current_episode;
SELECT * FROM active_sessions;
REFRESH MATERIALIZED VIEW episode_media;
SELECT * FROM episode_media ORDER BY num, filename;
```

### 运维边界

在非空模式中安装前，应审阅完整 SQL：它会创建表、触发器、服务器、外部表、物化视图、固定所有者的函数、授权、绝对路径以及带历史背景的应用假设。密码哈希、会话过期、IP 跟踪、ACL、文件系统暴露和 API 函数共同构成安全边界；未经现代安全审计不得直接复用。还应针对实际目标版本测试这套 PostgreSQL 10 时代的 SQL 和 Multicorn 包装器，并为该应用模式独立规划迁移、备份、媒体权限与卸载行为。
