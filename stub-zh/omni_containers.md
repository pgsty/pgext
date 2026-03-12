


## 用法

> [omni_containers: Docker 容器管理](https://docs.omnigres.org/omni_containers/intro/)

`omni_containers` 扩展从 PostgreSQL 管理 Docker 容器生命周期。需要 Docker Engine API 访问权限（默认 `unix:///var/run/docker.sock`）。

### 创建并启动容器

```sql
SELECT * FROM omni_containers.docker_container_create(
    'busybox',
    cmd  => 'echo Hello world',
    pull => true
);
```

参数：`image`（必需）、`cmd`、`attach`（默认 `db.omni`）、`start`（默认 true）、`wait`（默认 false）、`pull`（默认 false）、`options`（jsonb）。

### 容器日志

```sql
SELECT * FROM omni_containers.docker_container_logs('container_id');
```

支持 `stdout`、`stderr`、`since`、`until`、`timestamps`、`tail` 参数。

### 在容器中执行命令

```sql
SELECT * FROM omni_containers.docker_container_exec('container_id', 'ls -la');
```

### 停止和检查

```sql
SELECT * FROM omni_containers.docker_container_stop('container_id');
SELECT info->'State'->'ExitCode' FROM omni_containers.docker_container_inspect('container_id') AS info;
```

### 列出 Docker 镜像

```sql
SELECT repo_tags FROM omni_containers.docker_images;
```

### 完整工作流示例

```sql
WITH container AS (
    SELECT * FROM omni_containers.docker_container_create(
        'ghcr.io/yrashk/psql',
        cmd  => $$psql -c 'create table test ();'$$,
        wait => true, pull => true) AS id
)
SELECT
    omni_containers.docker_container_inspect(id)->'State'->'ExitCode' AS exit_code,
    omni_containers.docker_container_logs(id) AS logs
FROM container;
```
