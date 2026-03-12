


## Usage

> [omni_containers: Docker container management](https://docs.omnigres.org/omni_containers/intro/)

The `omni_containers` extension manages Docker container lifecycle from PostgreSQL. Requires Docker Engine API access (defaults to `unix:///var/run/docker.sock`).

### Create and Start a Container

```sql
SELECT * FROM omni_containers.docker_container_create(
    'busybox',
    cmd  => 'echo Hello world',
    pull => true
);
```

Parameters: `image` (required), `cmd`, `attach` (default `db.omni`), `start` (default true), `wait` (default false), `pull` (default false), `options` (jsonb).

### Container Logs

```sql
SELECT * FROM omni_containers.docker_container_logs('container_id');
```

Supports `stdout`, `stderr`, `since`, `until`, `timestamps`, `tail` parameters.

### Execute Command in Container

```sql
SELECT * FROM omni_containers.docker_container_exec('container_id', 'ls -la');
```

### Stop and Inspect

```sql
SELECT * FROM omni_containers.docker_container_stop('container_id');
SELECT info->'State'->'ExitCode' FROM omni_containers.docker_container_inspect('container_id') AS info;
```

### List Docker Images

```sql
SELECT repo_tags FROM omni_containers.docker_images;
```

### Full Workflow Example

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
