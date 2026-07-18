## Usage

Sources:

- [Crunchy Bridge Container Apps documentation](https://docs.crunchybridge.com/container-apps)
- [Crunchy Bridge product page](https://www.crunchydata.com/products/crunchy-bridge)

`pgpodman` is the provider extension behind Crunchy Bridge Container Apps. It launches and manages container images on the managed database host from SQL. The feature is currently public beta, and its extension version is controlled by the service rather than published as a portable package.

### Install and inspect

Running a container beside the database is privileged. Crunchy Bridge requires installation by a superuser in the `postgres` database:

```sql
CREATE EXTENSION pgpodman;

SELECT pgpodman_version();
SELECT * FROM list_containers(true);
SELECT * FROM list_images();
```

The API includes `run_container()`, `start_container()`, `stop_container()`, `container_status()`, `container_logs()`, `remove_container()`, and image-management functions. Container options are passed as a command string, so review every image digest, port, mount, environment variable, and privilege before execution.

### Security and service limits

`run_container()` can execute arbitrary public container images next to PostgreSQL. Treat it as host-level code execution with access to the database network boundary, not as an ordinary SQL helper. Pin trusted images by digest, minimize container privileges and outbound access, apply provider firewall controls, and never place long-lived secrets in SQL text, audit trails, statement logs, container arguments, or image layers. Container logs may also contain database or application data.

During the beta, only public ports `5433` through `5442` are exposed, and container filesystem use is limited to 2 GB. After maintenance, failover, or resize, the service checks `pgpodman_status` about every 10 minutes and restarts containers expected to be running; this is not an instantaneous availability guarantee. Stopping a container does not remove it, and removal requires it to be stopped first. Include image cleanup, port ownership, backup scope, failover recovery time, monitoring, and provider support escalation in the operating plan.
