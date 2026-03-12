


## Usage

> [omni_service: Service management](https://docs.omnigres.org/omni_service/management/)

The `omni_service` extension provides unified management for database services like `omni_httpd`. It is a templated extension.

### Setup

```sql
SELECT omni_service.instantiate(schema => 'omni_service');
```

### List Services

```sql
SELECT * FROM omni_service.current_services;
```

| Column  | Type           | Description                                  |
|:--------|:---------------|:---------------------------------------------|
| `name`  | text           | Unique service identifier                    |
| `info`  | jsonb          | Service metadata and configuration           |
| `state` | service_state  | `not-running`, `running`, or `stopped`       |

### Start and Stop Services

```sql
SELECT omni_service.service_start('service_name');
SELECT omni_service.service_stop('service_name');

-- Stop all running services:
SELECT omni_service.service_stop(name)
FROM omni_service.current_services
WHERE state = 'running';
```

### Operation Log

The `service_operations` table records all service actions with `name`, `postmaster_start_time`, `operation` (`start`/`stop`), and `applied_at`.
