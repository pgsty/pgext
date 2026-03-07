
## 用法

实际的 PDF 解析由 [poppler](https://poppler.freedesktop.org) 完成。

该扩展允许以符合 ACID 事务的方式处理 PDF 文件。
常见的替代方案依赖于外部脚本或服务，容易导致数据摄取流程变得脆弱，并使原始数据出现不同步的问题。

- [在 Postgres 中对 PDF 进行全文搜索](https://tselai.com/full-text-search-pdf-postgres)
- [pgpdf：Postgres 的 pdf 类型](https://tselai.com/pgpdf-pdf-type-postgres)

下载一些 PDF 文件。

```sh
wget https://wiki.postgresql.org/images/e/ea/PostgreSQL_Introduction.pdf -O /tmp/pgintro.pdf
wget https://pdfobject.com/pdf/sample.pdf -O /tmp/sample.pdf
```

通过将 `text` 类型的文件路径或 `bytea` 列进行类型转换，即可创建 `pdf` 类型。

```sql
CREATE EXTENSION pgpdf;
SELECT '/tmp/pgintro.pdf'::pdf;
```

```sql
                                       pdf
----------------------------------------------------------------------------------
 PostgreSQL Introduction                                                         +
 Digoal.Zhou                                                                     +
 7/20/2011Catalog                                                                +
  PostgreSQL Origin
```

如果文件系统中没有 PDF 文件，但已将其内容存储在 `bytea` 列中，可以直接将其转换为 `pdf` 类型。

```sql
SELECT pg_read_binary_file('/tmp/pgintro.pdf')::bytea::pdf;
```

--------

## 示例

创建一个包含 `pdf` 列的表：

```sql
CREATE TABLE pdfs(name text primary key, doc pdf);

INSERT INTO pdfs VALUES ('pgintro', '/tmp/pgintro.pdf');
INSERT INTO pdfs VALUES ('pgintro', '/tmp/sample.pdf');
```

解析和验证会自动进行。
文件只会从磁盘读取一次！

> [!NOTE]
> 文件路径必须是 `postgres` 进程/用户可访问的！
> 这与运行 psql 的用户不同。
> 如果不确定此处含义，请咨询 DBA！

### 字符串函数和运算符

标准的 PostgreSQL [字符串函数和运算符](https://www.postgresql.org/docs/17/functions-string.html)均可正常使用：

```sql
SELECT 'Below is the PDF we received ' || '/tmp/pgintro.pdf'::pdf;
```

```sql
SELECT upper('/tmp/pgintro.pdf'::pdf::text);
```

```sql
SELECT name
FROM pdfs
WHERE doc::text LIKE '%Postgres%';
```

### 全文搜索（FTS）

由于 `pdf` 文件可以像普通文本一样操作，因此也支持全文搜索（FTS）。

```sql
SELECT '/tmp/pgintro.pdf'::pdf::text @@ to_tsquery('postgres');
```

```sql
 ?column?
----------
 t
(1 row)
```

```sql
SELECT '/tmp/pgintro.pdf'::pdf::text @@ to_tsquery('oracle');
```

```sql
 ?column?
----------
 f
(1 row)
```

### 使用 `pg_trgm` 计算文档相似度

可以使用 [pg_trgm](https://postgresql.org/docs/17/interactive/pgtrgm.html) 计算两个文档之间的相似度：

```sql
CREATE EXTENSION pg_trgm;

SELECT similarity('/tmp/pgintro.pdf'::pdf::text, '/tmp/sample.pdf'::pdf::text);
```

### 元数据

以下函数可用：

- `pdf_title(pdf) → text`
- `pdf_author(pdf) → text`
- `pdf_num_pages(pdf) → integer`

  文档的总页数
- `pdf_page(pdf, integer) → text`

  获取第 i 页的文本内容
- `pdf_creator(pdf) → text`
- `pdf_keywords(pdf) → text`
- `pdf_metadata(pdf) → text`
- `pdf_version(pdf) → text`
- `pdf_subject(pdf) → text`
- `pdf_creation(pdf) → timestamp`
- `pdf_modification(pdf) → timestamp`

```sql
SELECT pdf_title('/tmp/pgintro.pdf');
```

```sql
        pdf_title
-------------------------
 PostgreSQL Introduction
(1 row)
```

```sql
SELECT pdf_author('/tmp/pgintro.pdf');
```

```sql
 pdf_author
------------
 周正中
(1 row)
```

获取部分页面

```sql
SELECT pdf_num_pages('/tmp/pgintro.pdf');
```

```sql
 pdf_num_pages
---------------
            24
(1 row)
```

```sql
SELECT pdf_page('/tmp/pgintro.pdf', 1);
```

```sql
           pdf_page
------------------------------
 Catalog                     +
  PostgreSQL Origin         +
  Layout                    +
  Features                  +
  Enterprise Class Attribute+
  Case
(1 row)
```

```sql
SELECT pdf_subject('/tmp/pgintro.pdf');
```

```sql
 pdf_subject
-------------

(1 row)
```

```sql
SELECT pdf_creation('/tmp/pgintro.pdf');
```

```sql
       pdf_creation
--------------------------
 Wed Jul 20 11:13:37 2011
(1 row)
```

```sql
SELECT pdf_modification('/tmp/pgintro.pdf');
```

```sql
     pdf_modification
--------------------------
 Wed Jul 20 11:13:37 2011
(1 row)
```

```sql
SELECT pdf_creator('/tmp/pgintro.pdf');
```

```sql
            pdf_creator
------------------------------------
 Microsoft® Office PowerPoint® 2007
(1 row)
```

```sql
SELECT pdf_metadata('/tmp/pgintro.pdf');
```

```sql
 pdf_metadata
--------------

(1 row)
```

```sql
SELECT pdf_version('/tmp/pgintro.pdf');
```

```sql
 pdf_version
-------------
 PDF-1.5
(1 row)
```

## 安装

安装 [poppler](https://poppler.freedesktop.org) 依赖

**Linux**
```
sudo apt install -y libpoppler-glib-dev pkg-config
```

**Homebrew/MacOS**

```
brew install poppler pkgconf
```

```
cd /tmp
git clone https://github.com/Florents-Tselai/pgpdf.git
cd pgpdf
make
make install # 可能需要 sudo
```

安装完成后，在数据库会话中执行：

```sql
CREATE EXTENSION pgpdf;
```

### Docker

通过以下命令获取 [Docker 镜像](https://hub.docker.com/r/florents/pgpdf)：

```sh
docker pull florents/pgpdf:pg17
```

此镜像在 [Postgres 镜像](https://hub.docker.com/_/postgres) 基础上添加了 pgpdf（将 `17` 替换为所需的 Postgres 服务器版本，运行方式相同）。

在容器中运行该镜像。

```sh
docker run --name pgpdf -p 5432:5432 -e POSTGRES_PASSWORD=pass florents/pgpdf:pg17
```

通过另一个终端连接到正在运行的服务器（容器）。

```sh
PGPASSWORD=pass psql -h localhost -p 5432 -U postgres
```

> [!WARNING]
> 将任意二进制数据（PDF）读入数据库可能存在安全风险。
> 请仅对可信的文件使用此功能。
