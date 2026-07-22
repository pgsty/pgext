## 用法

来源：

- [Official README](https://github.com/ZhengYang/couchdb_fdw/blob/df93c5ec00034f5582ca135895ef4b427896746f/README.md)
- [Extension control file](https://github.com/ZhengYang/couchdb_fdw/blob/df93c5ec00034f5582ca135895ef4b427896746f/couchdb_fdw.control)
- [FDW implementation](https://github.com/ZhengYang/couchdb_fdw/blob/df93c5ec00034f5582ca135895ef4b427896746f/couchdb_fdw.c)

couchdb_fdw 是一个年代久远、处于实验状态的只读外部数据包装器，用于把 CouchDB 文档暴露为 PostgreSQL 行。上游称其为 beta、列出了已知缺陷，并表示只在 PostgreSQL 9.1 时代的软件及 macOS 10.6 上测试过。它应被视为历史源码，而非生产集成方案。

### 核心流程

安装所需的 libcurl 和 YAJL 后，定义服务器、凭据和一个外部表。`_doc` 是返回整个文档 JSON 的保留列名；其他列映射顶层 JSON 属性。

```sql
CREATE EXTENSION couchdb_fdw;

CREATE SERVER couch
  FOREIGN DATA WRAPPER couchdb_fdw
  OPTIONS (address '127.0.0.1', port '5984');

CREATE USER MAPPING FOR CURRENT_USER
  SERVER couch
  OPTIONS (username 'reader', password 'secret');

CREATE FOREIGN TABLE couch_docs (
  _id text,
  _rev text,
  title text,
  _doc text
)
SERVER couch
OPTIONS (database 'example');

SELECT _id, title FROM couch_docs WHERE _id = 'document-id';
```

### 重要对象

- `couchdb_fdw` 是外部数据包装器。
- `address` 和 `port` 是服务器选项；默认值分别为 127.0.0.1 和 5984。
- `username` 和 `password` 是用户映射选项。
- `database` 是外部表选项，默认值为 test。
- `_id`、`_rev` 和 `_doc` 是特殊文档列。

### 安全与查询限制

实现会把用户名和密码拼进明文 HTTP URL，且不支持 TLS 或现代认证。只有 `_id` 等值条件有实际下推效果，还可附带 `_rev`；无条件扫描会先取得全部文档 ID，再逐一获取文档。它只映射顶层属性且不支持写入。凭据会留在 PostgreSQL 系统目录中，远程 I/O 为同步操作，失败可中断后端查询。不要通过这份过时代码暴露敏感凭据或不可信 CouchDB 端点。
