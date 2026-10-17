---
title: 上手
description: 快速上手 pig，PostgreSQL 包管理器
icon: fa-solid fa-play
weight: 200
breadcrumb: false
---

本教程安装 PostgreSQL 18 和 pgvector，启动本地测试实例，并在数据库中启用扩展。

准备一台[受支持的 Linux 测试主机](/zh/pig/intro/#linux-兼容性)，确保能联网并已安装 `curl`。你需要 root 或 `sudo` 权限。按顺序执行以下命令。

## 安装 PIG {#安装}

```bash
curl -fsSL https://repo.pigsty.cc/pig | bash
pig version
```

中国大陆以外可使用 `https://repo.pigsty.io/pig`。安装脚本会选择与主机匹配的 RPM 或 DEB 包。其他方式见[安装](/zh/pig/install/)。

## 检查主机 {#检查环境}

```bash
pig status
```

查看操作系统代码、CPU 架构和已检测到的 PostgreSQL 安装。此命令也会测试软件仓库的连通性。

## 添加仓库 {#添加仓库}

添加 PGDG 与 Pigsty PostgreSQL 仓库，并刷新软件包元数据：

```bash
sudo pig repo add pgsql --update
```

需要中国镜像时，在命令后加上 `--mirror`。

## 安装 PostgreSQL 与 pgvector {#安装-pg}

```bash
sudo pig install -v 18 pg18 vector
```

检查包管理器列出的变更并确认。

`pg18` 选择 PostgreSQL 18 软件包，`vector` 选择 pgvector。`-v 18` 指定扩展使用的 PostgreSQL 大版本。

检查已安装的扩展文件：

```bash
pig ext status -v 18
```

## 启动测试实例 {#start-a-test-instance}

本教程使用 `/pg/data-demo` 和 `15432` 端口，以避开 `5432` 端口上的默认集群。请使用尚未占用的目录和端口。

创建空的数据目录，并将所有者设为操作系统用户 `postgres`：

```bash
sudo install -d -m 700 -o postgres -g postgres /pg/data-demo
sudo pig pg init -v 18 -D /pg/data-demo
```

在 `15432` 端口启动实例：

```bash
sudo pig pg start -v 18 -D /pg/data-demo -O '-p 15432'
sudo pig pg status -v 18 -D /pg/data-demo
```

## 启用并测试 pgvector {#enable-and-test-pgvector}

连接测试实例：

```bash
sudo pig pg psql -v 18 -D /pg/data-demo
```

PIG 从指定数据目录读取实例的连接信息。在 `psql` 提示符下执行：

```sql
CREATE EXTENSION vector;
SELECT extname, extversion FROM pg_extension WHERE extname = 'vector';
SELECT '[1,2,3]'::vector <-> '[4,5,6]'::vector AS distance;
```

第一个查询返回已安装的 pgvector 版本，第二个查询返回约为 `5.196` 的距离值。输入 `\q` 退出 `psql`。

`CREATE EXTENSION` 在当前连接的数据库中启用扩展。其他数据库需要 pgvector 时，也要分别执行此语句。

完成测试后停止实例：

```bash
sudo pig pg stop -v 18 -D /pg/data-demo
```

数据保留在 `/pg/data-demo` 中。再次执行启动命令即可继续使用。服务配置与其他操作见 [PostgreSQL 管理](/zh/pig/cmd/pg/)。

## 查找更多扩展 {#列出扩展}

```bash
pig ext list                 # 列出扩展目录
pig ext list duckdb          # 按名称或描述搜索
pig ext info pg_duckdb       # 查看元数据和启用要求
pig ext avail pg_duckdb      # 查看软件包可用矩阵
```

安装其他扩展前，用 `info` 查看启用要求，用 `avail` 检查软件包是否可用。

## 选择软件包与版本 {#别名翻译}

需要其他 PostgreSQL 大版本时，替换上述命令中的 `18`。
软件包组、版本选择和已有安装的处理方式见[扩展安装](/zh/pig/cmd/ext/#ext-add)。

## 检查与更新 {#显示扩展}

`pig ext status -v 18` 列出已安装的扩展文件。要查看数据库中已启用的扩展，在 `psql` 中执行 `\dx`。
需要升级软件包及数据库中的扩展对象时，参阅[扩展更新](/zh/pig/cmd/ext/#ext-update)。
