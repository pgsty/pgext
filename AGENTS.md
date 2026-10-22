# Repository Guidelines

## OINK build baseline

As of 2026-10-18, this site uses OINK v1.3.2 and Hugo Extended 0.168.0.
Keep CI and local builds on this baseline. Do not enable `--panicOnWarning`
or `HUGO_PANICONWARNING`; keep warnings visible and review them. The upstream
LibSass deprecation warning is accepted. No external Sass compiler is required.

## Review Execution

The user revoked automatic Claude / Claude Code / Opus review on 2026-10-01 because the account is unavailable. Do not invoke, retry, or probe Claude based on earlier conversations, run ledgers, or legacy instructions; resume only if the user explicitly requests it again. The current agent performs source verification, metadata and documentation review, tests, and acceptance checks. Claude review is not a completion requirement, and its absence must not block routine discovery, intake, or propagation or produce repeated account-error notifications. All existing quality checks below remain required.

## Project Structure & Module Organization

`cmd/` exposes the Cobra command surface, while `cli/` carries the data-fetch, parse, and recap pipelines. `content/` and `layouts/` drive the Hugo front end, and `db/` plus `bin/` store CSV exports and Python generators that feed the site. Treat `public/` and any other build artefacts as disposable—rebuild instead of editing them.

## Build, Test, and Development Commands

- `make dev` (alias `make d`): run `hugo serve` for local previews at http://localhost:1313.
- `make build`: generate the static site with `hugo build`; use before publishing.
- `go run .`: execute the CLI entrypoint (`pgext`) for tasks like `pgext reload`.
- `go run . serve`: run the self-contained PGEXT.CLOUD web app + JSON query API at http://localhost:8432 (assets embedded from `server/web/`, data queried live from the `pgext` schema via `--db`/`$PGURL`; see `server/README.md`).
- `go test ./...`: run Go unit tests across `cli/` (add `-cover` when needed).
- `make dump` / `make load`: sync PostgreSQL metadata using `PGURL`; confirm the target database first.

## Coding Style & Naming Conventions

Format Go code with `gofmt` (Go 1.27.3) and keep imports sorted via `goimports`. Use descriptive exported names for shared structs (e.g., `BinaryPackage`) and lower-case verb/noun pairs for Cobra commands (`reload`, `status`). Stick with `logrus` for logging to stay compatible with the root setup. Hugo Markdown files remain lower-case with front matter matching existing sections, and Python helpers should follow snake_case and PEP 8.

## Testing Guidelines

Author table-driven Go tests like `cli/generator_test.go`, placing `_test.go` files next to the code they cover. Run `go test ./...` before opening a PR, and lean on `go test ./cli -run CleanVersion -v` while iterating. Mock external services or isolate pure functions to keep tests deterministic, and add lightweight schema checks for generated CSV or JSON when touching the build scripts.

## Commit & Pull Request Guidelines

Use short, imperative commit subjects (`add hugo hextra theme`) and avoid bundling unrelated site, CLI, and data updates. Pull requests should explain motivation, highlight the main changes, and list the validation commands you ran (`make build`, `go test ./...`). Include links or screenshots whenever Hugo UI output changes.

## Data & Environment Notes

Export `PGURL` before running database-dependent commands, ideally using a scoped role. For current local `pgext` metadata and Universe curation work, use the `data` database on the local PostgreSQL socket, for example `PGUSER=postgres psql -h /tmp -p 5432 -d data`; the schema is `pgext`. Do not assume a bare `postgres:///` connection is correct, because it can land in the `postgres` database where the `pgext` schema may not exist. Review tracked CSV diffs in `db/` for unexpected churn, and activate the project virtualenv (`source ~/.venv/bin/activate`) when invoking Python generators.

### Package Sources and Gap Builds

- `repo`, `rpm_repo`, and `deb_repo` describe the principal package supplier. Keep PGDG as the source when it supplies most supported coordinates; Pigsty may fill missing OS/architecture/PostgreSQL combinations without taking over that classification. Preserve the existing curated overall `repo` label, including PGDG-dominant entries with a Pigsty DEB side; use `MIXED` where the catalog already distinguishes principal RPM and DEB suppliers. Do not replace this curation with a formula based on recipe ownership or a single batch.
- A local SPEC, DEB recipe, successful build, or newer candidate version proves build capability, not a change of supplier. `extra.rpm` / `extra.deb` and Universe `rpm_build` / `deb_build` record that capability independently. `rpm_pg` / `deb_pg` cover the supported package range across suppliers, not just the majors built in the current run.
- Before importing a batch, use current PGDG metadata and the gap policy in `bin/rebuild_manifest.py` to select publication coordinates. Preserve full build outputs separately. For PGDG sides, publish Pigsty packages for missing coordinates; a version lead alone does not authorize publishing replacements. Record functional exceptions such as TimescaleDB TSL and narrowly scoped security fixes explicitly; end a security exception once PGDG supplies the fixed release. Verify companion extensions in multi-extension packages rather than assuming that a primary package supplies every family member.

## Local `pgext` Schema Table Guide

The local working database is `data`, schema `pgext`. The main curated extension universe is `pgext.universe`; the packaged baseline catalog remains `pgext.extension`. `pgext.extension_all` is an obsolete name and must not be used as an intake target.

Core catalog and package pipeline:
- `pgext.extension`: baseline curated catalog for packaged Pigsty/PGDG extensions.
- `pgext.universe`: canonical ecosystem-wide catalog for packaged, source-only, provider-only, and discovered PostgreSQL extensions, with one row per canonical extension name.
- `pgext.doc`: one row per `pgext.universe` extension, containing English Markdown from `stub/` and Chinese Markdown from `stub-zh/`; `bin/load-stub` synchronizes the complete bilingual corpus atomically.
- `pgext.category`, `pgext.pg`, `pgext.os`, `pgext.status`: category, PostgreSQL version, OS target, and lifecycle status dimension tables.
- `pgext.repository`, `pgext.repo_data`, `pgext.apt`, `pgext.dnf`, `pgext.bin`, `pgext.pkg`: repository configuration, raw repository metadata cache, parsed APT/RPM package metadata, normalized binary package rows, and final package availability matrix.
- `pgext.changelog`: per-package release history distilled from `content/release/{rpm,deb}{,.zh}.md` by `bin/changelog.py` into `db/changelog.csv` (aligned batch dates, normalized pkg names, merged RPM/DEB sides, bilingual notes).
- `pgext.counter`: per-extension visit/download counters; `pgext serve` accumulates increments in memory and flushes them additively every minute.

Reporting views:
- `pgext.active_pg`, `pgext.active_os`: active PostgreSQL and OS targets.
- `pgext.pkg_pgdg`, `pgext.pkg_pigsty`: what-if variants of `pgext.pkg` as if only that repository organization existed (AVAIL recomputed from `pgext.bin` per org, real-world N/A preserved, everything else MISS); `pgext serve` inlines the same logic for `/api/v1/matrix?repo=pgdg|pigsty` and the `/repo/PGDG`, `/repo/PIGSTY` matrix pages.
- `pgext.pig`: website/reporting view joining extension metadata with package availability.
- `pgext.summary`: high-level totals for website/reporting output.

## End-to-End New Extension Intake SOP

Use this workflow whenever the user asks to “新收录一个扩展”, add or index a new extension, put a project into `pgext.universe`, or when an explicitly authorized discovery run finds an unlisted extension. This is an end-to-end workflow: every accepted new extension must finish Universe curation, bilingual stub creation, and `pgext.doc` synchronization. It is broader than a stub-only request and explicitly authorizes both `stub/` and `stub-zh/` for the accepted extension.

### Completion Contract

An extension is `DONE` only when all of these are true:

1. Authoritative upstream evidence proves that the candidate is a real PostgreSQL extension and identifies its canonical extension name.
2. A reviewed row exists in local `data.pgext.universe`, and `db/universe.csv` contains the exact exported row with no unrelated generated-data churn.
3. `stub/<extension>.md` and `stub-zh/<extension>.md` both exist, satisfy the repository document contract, and have been checked against version-specific official sources.
4. `bin/load-stub --dry-run` succeeds, the complete bilingual corpus is loaded atomically into local `data.pgext.doc`, and the target row plus global cardinality/metadata/content checks pass.

Universe intake, CSV export, stub source validation, `pgext.doc` loading, generated-site builds, package builds, commits, pushes, and publication are distinct layers. Do not claim or perform the later build/repository/publication layers unless the user separately requests them.

### Phase 0: Establish Scope, Database, and Baseline

- Start with the current branch, `git status --short --branch`, and the exact in-scope candidate list. Preserve all existing user changes and do not rewrite unrelated Universe rows or stubs.
- Connect explicitly with `PGUSER=postgres psql -X -h /tmp -p 5432 -d data`; confirm `current_database() = 'data'`, `to_regclass('pgext.universe')`, and `to_regclass('pgext.doc')` before mutation.
- Inspect the current schema rather than relying on an older prompt or migration. Use `db/schema.sql`, `information_schema.columns`, table comments, constraints, and several comparable rows as the field contract.
- Search for existing rows by case-insensitive `name`, `pkg`, repository URL, homepage, control URL, and PGXN URL. Resolve aliases, renamed projects, forks, and multi-extension distributions before treating a discovery as new.
- Before modifying the database, export `pgext.universe ORDER BY id` to a temporary file and compare it with tracked `db/universe.csv`. Record or resolve pre-existing drift instead of attributing it to the new row.

### Phase 1: Prove Admission and Canonical Identity

- Require extension evidence from the exact release or source revision: a control file or PGXN `META.json`, versioned extension SQL, and documented `CREATE EXTENSION` support. A PostgreSQL-related client, driver, proxy, migration tool, benchmark, operator, or standalone utility is not enough.
- Inspect authoritative material in this order: control and versioned SQL files; the extension-specific subdirectory in a monorepo; official README/documentation; tagged release notes or changelog; PGXN metadata; official vendor documentation for provider-only extensions; source/API metadata for activity and repository facts.
- For a multi-extension distribution, create one Universe row per actual extension name. Put the distribution/project handle in `pkg`, choose the primary extension as `lead_ext`, and keep the family together during review.
- Headless, `need_ddl=false`, or provider-only entries still require an authoritative extension identity and enablement boundary. Do not fabricate an open-source repository, SQL surface, control setting, compatibility claim, or `CREATE EXTENSION` example.
- If the authoritative sources conflict on name, version, license, compatibility, or operational requirements, stop the affected item as `NEEDS_REVIEW`, `BLOCKED_NAME`, or `BLOCKED_SOURCE`; do not guess a row merely to complete the batch.

### Phase 2: Curate the `pgext.universe` Row

Use an explicit column list and fill every field from evidence or a deliberate catalog decision:

- Identity: allocate the next stable `id` under a table lock; store the control/`CREATE EXTENSION` name in `name`, the source distribution in `pkg`, the family primary in `lead_ext`, and one valid `pgext.category` value in `category`.
- Status: set `packaged=true` only when the extension belongs to the approved packaged baseline represented by `pgext.extension`; an upstream tarball or binary release alone does not make it packaged here. Set `lifecycle` and `kind` (`standard`, `preload`, `puresql`, or `headless`) from current upstream behavior.
- Classification: verify `kernel`, `vendor`, `contrib`, primary implementation `lang`, SPDX-style `license`, and concise stable `tags`. Use `NULL` for unknown facts rather than an attractive guess.
- Version and sources: record the extension/control version and stable public `url`, `repo_url`, `home_url`, `doc_url`, `license_url`, `control_url`, `author_url`, `cargo_url`, and `pgxn_url` as applicable. Prefer tag-specific file URLs for version-specific evidence and check that every stored URL resolves.
- Runtime: derive `has_bin`, `has_lib`, `need_ddl`, `need_load`, `trusted`, `relocatable`, `libs`, `schemas`, and `pg_ver` from the control file, build metadata, source, and official documentation. Do not infer preload, restart, privileges, or PostgreSQL-major support from project category or language.
- Relationships: populate `requires` and `see_also` only with canonical Universe names. Treat `required_by` as a derived reverse-dependency cache; run `SELECT pgext.refresh_deps()` when forward dependencies change, then inspect every resulting `required_by` delta.
- Packaging: leave `tarball`, `pgrx_ver`, RPM, and DEB fields null/false for a source-only intake unless current package/build evidence justifies them. Do not turn intake into a package-availability claim.
- Activity: collect `stars`, subscriber/watch count, forks, `repo_created_at`, `last_commit`, `last_release`, `last_active`, and `checked_at` from current upstream APIs and release metadata. Normalize activity dates to UTC calendar dates.
- Description: write concise factual `en_desc` and natural `zh_desc`; keep `comment` single-line and reserve it for material version conflicts, support boundaries, preload/restart needs, security caveats, or manual curation decisions. Keep `extra` small and stable; never store evidence dumps or workflow notes there. Set `mtime` to the catalog-change date.

Perform the insert in a transaction with `LOCK TABLE pgext.universe IN SHARE ROW EXCLUSIVE MODE`, keep the self/family `lead_ext` foreign key valid, and re-check uniqueness, category/family references, dependency names, required non-null fields, and row count before commit. After commit, query the full row vertically and compare it field-by-field with the evidence notes.

Export through the confirmed local socket using the Make-safe percent-encoded URI:

```bash
make PGURL='postgresql://postgres@%2Ftmp:5432/data' dump
```

`make dump` rewrites several tracked CSVs. Confirm that the intended result is exactly the new/updated Universe row and that other CSVs did not drift. Never discard unrelated user changes to make the diff look clean.

### Phase 3: Create the English and Chinese Stubs

- Create `stub/<name>.md` and `stub-zh/<name>.md` for every newly accepted canonical name. Read any existing files before deciding whether this is a create, material refresh, or no-op.
- Follow the full Extension Stub Documentation Workflow below. The English file must use exactly one `## Usage`; the Chinese file must use exactly one `## 用法`. `Sources:` or `来源：` plus official linked bullets must be the first content after the H2.
- Use the same version-specific official URL multiset in both languages. Keep fenced code blocks and inline identifiers synchronized; translate the explanation into natural technical Chinese without translating SQL identifiers, GUCs, package names, extension names, operators, paths, or configuration values.
- Cover the smallest real core workflow, important user-facing objects, configuration and privileges, dependencies, preload/restart behavior, PostgreSQL compatibility, maintenance implications, and failure/security boundaries. Omit unsupported claims, generic build instructions, badges, marketing benchmarks, and unrelated project history.
- For `need_ddl=false` or provider-only entries, document the actual enablement path and do not add `CREATE EXTENSION` merely for consistency with other stubs.

### Phase 4: Validate and Load `pgext.doc`

- Run the compatible Phase B pair lint when a bound run directory/manifest exists; otherwise perform the equivalent focused `lint_pair` checks from `bin/stub_phaseb.py`. The pair must have matching source URLs, code fences, inline identifiers, and heading levels, with no front matter, H1, placeholder, unclosed fence, or language inversion.
- Check every official source URL, inspect the final stub diffs, run `git diff --check`, and run the loader tests:

```bash
python3 -m unittest -v bin/test_stub_doc.py
```

- Query `pgext.universe` and `pgext.doc` counts before loading. The loader synchronizes the entire Universe corpus, not only the new row: extra stub files merely warn and are excluded, but a missing or empty English/Chinese file for any Universe entry blocks the transaction.
- Always run the complete transactional dry-run first, against the explicit local database:

```bash
bin/load-stub --dry-run --pgurl 'host=/tmp port=5432 dbname=data user=postgres'
```

- Only after dry-run and source/diff checks pass, perform the authorized local load:

```bash
bin/load-stub --pgurl 'host=/tmp port=5432 dbname=data user=postgres'
```

- Postflight must prove: `pgext.doc` count equals `pgext.universe`; no `en_doc` or `zh_doc` is null/blank; the target `id`, `ext`, `pkg`, and URL metadata match Universe; there are no catalog foreign-key remnants; and the stored English/Chinese bodies exactly equal the two files, preferably by both direct comparison and content hash.
- Report source validation, Universe/CSV diff, stub lint, dry-run, committed `pgext.doc` load, tests, and any blocked candidates separately. Do not call a row `DONE` if only research, drafting, or dry-run completed.

### Daily Discovery Mode

Use this mode only for an explicitly authorized recurring discovery task. Each run must browse live sources and use local Universe state for deduplication.

- Check PGXN recent releases/distributions, PostgreSQL.org news and community announcements, the latest Postgres Weekly issue/archive, GitHub repository/code/release search, and general web search. Rotate focused queries such as `"PostgreSQL extension"`, `"Postgres extension"`, `"CREATE EXTENSION"`, `"pg extension"`, `filename:*.control PostgreSQL`, and recent release variants so the run is not limited to one search ranking.
- Cover at least the interval since the last successful run with a 24-hour overlap; when no prior run state is available, inspect the last seven days. Record source URLs and publication/release dates. Search results are discovery hints only—admission still requires exact upstream control/META/SQL and official documentation evidence.
- Deduplicate every candidate against `pgext.universe` by canonical name, package, repository, homepage, control URL, and PGXN URL before doing deep work. Existing entries, version-only updates, non-extension tools, mirrors, and forks without a distinct extension identity are not new intake.
- For multiple candidates, create a dated ledger under `tmp/` using the statuses in the Batch Execution and Tracking section. Process every clearly admissible candidate through Phases 0–4; leave ambiguous candidates blocked with attempted URLs and a concrete reason rather than inserting speculative metadata.
- If no new extension qualifies, make no database or file changes and report the time window, sources checked, candidate names, and why each was already present or rejected. If candidates are accepted, report exact rows/files/counts and all validation layers.
- A discovery run may update the local database and working tree through this SOP, but it must not commit, push, build packages, regenerate/publish the site, or publish repositories unless the automation prompt or user explicitly authorizes those additional layers.

### Daily Universe Propagation

The user authorized recurring database propagation on 2026-09-21. After local discovery/intake validation, synchronize the reviewed `data.pgext.universe` snapshot to all five destinations below. This database synchronization and the sites' extension search indexes are part of the daily task; they do not authorize code deployment, site/package builds, commits, pushes, or copying other catalog tables.

| Destination | Connection / runtime |
| --- | --- |
| PGEXT production | `ssh pg`, PostgreSQL socket `/var/run/postgresql:5432`, database `pgext`; running `pgext` service |
| PGWeb local | `/tmp:5432`, database `pgweb`; `/Users/vonng/pgsty/pgweb` |
| PGWeb production | `ssh pg`, database `pgweb`; `/data/app/pgsql.cc`, service `pgsql.cc` |
| Center local | `/tmp:5432`, database `center`; `/Users/vonng/pgsty/pg.center` |
| Center production | `ssh pg`, database `center`; `/data/app/pg.center`, service `pg.center` |

- Local PGEXT uses the authoritative `data` database; do not invent a separate local `pgext` target. Reconfirm service checkout, database identity, schema/column contract, and connection at each run. The retained `/data/app/pgweb` predecessor is not an active destination. Use the Center service's `/etc/pg.center.env` and OS user `center` for its management commands; never log secrets.
- Export one fixed source snapshot through the explicit local socket, and verify it matches `db/universe.csv` and the completed local intake. Inspect every target even when discovery finds no changes, so an earlier incomplete propagation is retried. Equal targets are no-ops.
- Before a changed target is written, take a dedicated custom-format table dump in the current run directory, check the archive, restore its schema/data into an unused local temporary database, and compare the restored rows with the target preflight. Preserve the original table OID, owner, ACL, indexes, constraints, dependencies, and identity sequence.
- For the four sites, use their existing `tools/ext/sync_catalog.py` / `pgweb.ext.sync.import_snapshot` implementation, wrapped in a transaction that locks `pgext.universe` and verifies the target has not changed since backup. For PGEXT, use an explicit-column staged upsert with equivalent backup-drift and full-result guards. Run dry-run before apply. Update only changed rows, preserve source IDs, and never move an identity sequence backwards.
- Do not automatically delete target-only rows, renumber name/ID conflicts, migrate mismatched schemas, or overwrite concurrent target changes. Record the exact discrepancy and stop the affected target; preserve completed-target evidence for retry.
- After import, verify all rows and fields against the fixed source snapshot and compare normalized content hashes. Refresh only `manage.py index_docs --extensions` on each changed site. Verify extension search names equal Universe and non-extension index contents are unchanged. Check local application responses and public PGEXT API / PGWeb / Center detail pages after their ordinary caches refresh; do not restart services for catalog changes.
- Synchronization covers `pgext.universe` and its derived site search entries. Do not copy `pgext.doc`, `pgext.extension`, `pgext.category`, package tables, counters, or other application content as a side effect. Local bilingual documentation still follows the intake SOP above.
- Record source coverage/local acceptance and per-target propagation separately under `tmp/`. Advance the discovery watermark only after complete source coverage and local acceptance; mark propagation complete only after every destination and application check passes. A local-only success is not full completion. Retry failed destinations against the latest validated source on the next run, and notify on a new intake/material update, completed drift repair, failure, or required decision; otherwise remain quiet.
- The audited reference run is `tmp/universe_sync_20260921T024439Z/` (source export, reusable run helpers, five backups, rollback trials, full comparisons, and application checks). When adapting it, use a new run directory and current row count/schema/cluster identity; do not reuse its old snapshot, target baselines, or completed-state files.

## Extension Stub Documentation Workflow

Use this workflow whenever the user asks to create, refresh, complete, or audit extension usage documentation under `stub/`.

The versioned operational SOP is [docs/stub-documentation-sop.md](docs/stub-documentation-sop.md). It covers bilingual format review, the shared Markdown checks, and safe generation. The 2026-10-04 CI regression was caused by fixing only generated website pages while the invalid source stubs remained; subsequent `pgext gen io` copied the original errors back. All such repairs must reach the source files in `stub/` and, when in scope, `stub-zh/`.

### Scope and Naming

- Process only the extensions named by the user. Do not opportunistically scan or rewrite unrelated stubs unless a full-catalog refresh was explicitly requested.
- Resolve the canonical PostgreSQL extension name before writing. Use the user-supplied target list plus repository metadata (`db/extension.csv` for the packaged baseline and the current universe metadata for broader catalog entries). The canonical extension name, not necessarily the package name, determines `stub/<extension>.md`; record package-to-extension mappings instead of guessing.
- English output goes only to `stub/<extension>.md` unless the user explicitly requests Chinese synchronization, invokes the End-to-End New Extension Intake SOP above, or requests generated site pages. Do not edit `stub-zh/`, `content/e/`, catalog CSV/SQL, or package metadata as an unrelated side effect.
- Read an existing stub before refreshing it. For a version update, change only facts or examples with a material user-facing upstream delta. If there is no material change, leave the file intact and report `reviewed upstream <version>, no material stub delta`.
- Distinguish upstream feature changes from Pigsty packaging, rebuild, platform-matrix, or catalog-only changes.

### Research and Evidence

- Start with local metadata only to identify the canonical name, package mapping, version, repository, homepage, documentation, and control-file locations. Published stub prose must cite public upstream sources, not local CSV files or relative links such as `../db/extension.csv`.
- Inspect authoritative upstream material in this order: control and extension SQL files; the official README or relevant monorepo subdirectory; official documentation and project homepage; release notes/changelog for the target version; PGXN metadata; official vendor documentation for provider-only extensions.
- Treat catalog descriptions and comments as discovery hints, not sufficient evidence. Do not use third-party blogs, reposts, SEO pages, or AI summaries when an official source exists, and never invent an API, compatibility claim, configuration requirement, or version.
- In a multi-extension repository, verify the specific control/SQL subproject for the requested extension; do not copy repository-wide capabilities into every stub.
- Only document `shared_preload_libraries`, `LOAD`, restart behavior, dependencies, privileges, supported PostgreSQL versions, and GUCs when an authoritative source confirms them. For headless or `need_ddl=false` entries, do not fabricate `CREATE EXTENSION`. For vendor-only entries, document the actual platform and enablement boundary without inventing an open-source repository or SQL surface.
- When catalog and upstream versions differ, state the distinction only when it helps the reader and record it in the task note; do not silently change catalog data.

### Required English Stub Shape

Every English stub must satisfy the current `bin/stub_phaseb.py` document contract:

```markdown
## Usage

Sources:

- [Official README](https://...)
- [Official documentation](https://...)

Concise summary of what the extension does and its primary use cases.

### Core Workflow

...
```

- No front matter and no H1. The entire file has exactly one H2, spelled `## Usage`; all remaining headings are H3 or lower and all prose lives beneath that H2.
- `Sources:` must be the first content after `## Usage`, followed by a Markdown bullet list containing traceable official links. Prefer version/tag-specific links when documenting version-specific facts.
- Mention the canonical extension name verbatim. Open with a concise summary: what it does, when to use it, and the most important operational boundary.
- Cover the core workflow and use cases, then provide the smallest real SQL/configuration examples that make the feature understandable. Do not force SQL or `CREATE EXTENSION` examples where they do not apply.
- Add a compact, scannable index of important user-facing objects when applicable: types, functions, procedures, operators, index/access methods, views, GUCs, and helper commands. Explain key parameters or return behavior instead of dumping generated DDL.
- Include confirmed prerequisites, dependencies, preload/restart requirements, privileges, compatibility constraints, maintenance implications, and caveats that materially affect use.
- Keep package names, extension names, SQL, identifiers, operators, GUCs, and configuration values exactly spelled. Close all code fences and remove `TODO`, `TBD`, placeholders, and model disclaimers.
- Every opening code fence must carry an appropriate language: `sql`, `bash`, `ini`, `json`, `yaml`, or `text`. Use `text` for plans, result tables, terminal transcripts, CLI help, and signatures. Review both language files and preserve code bodies; do not blindly label every fence `sql`. Closing fences carry no language and must match the opening marker.
- Put blank lines before and after headings, use one marker space, remove stray prose trailing whitespace, and end files with a newline. Intentional two-space Markdown hard breaks and whitespace inside code bodies are preserved.
- Summarize rather than copy. Omit badges, CI instructions, sponsor/contribution text, routine source-build steps, verbose installation material, full changelogs, benchmarks presented as guarantees, and unrelated project history.

### Batch Execution and Tracking

- For a multi-extension batch, create a run-specific ledger under `tmp/`, for example `tmp/stub_refresh_<date>_master.md`, with at least: `package | extension | old | new | action | stub_path | status | owner | source_urls | note`.
- Use statuses `TODO`, `DOING`, `DONE`, `NEEDS_REVIEW`, `BLOCKED_NAME`, `BLOCKED_SOURCE`, and `DEFER_LARGE`. A blocked item must record attempted URLs and a concrete reason. Use `DEFER_LARGE` sparingly; complexity alone is not permission to skip a requested extension.
- For a scoped formatting audit, use `FORMAT_FIXED` and `NO_FORMAT_CHANGE` in the ledger. Record both source paths, issue types/lines, before/after hashes, and code-body/identifier preservation. These statuses do not claim upstream fact verification or full compliance of legacy document structure.
- When parallelizing a large batch, split ownership by disjoint stub files and keep extensions from the same monorepo together. Each writer owns only its assigned stubs and shard ledger; the coordinating agent merges status and performs final QA.
- Do not mark an item `DONE` merely because sources were found or a draft was written. `DONE` requires the target file to exist, the facts to be checked against upstream, the source URLs and version/caveat notes to be recorded, and the format checks to pass.

### Validation and Handoff

- Inspect the final diff for every in-scope stub and confirm no out-of-scope file was changed by this workflow.
- Run `go run . gen lint stub/<name>.md stub-zh/<name>.md` for the affected pair and `make check-stubs` for the complete bilingual corpus. These checks are required in addition to the existing pair/technical lint. Go CI executes the same shared validator. Invalid stubs cause site generators to return the source path and line number before writing the corresponding page; IO/CC batch generation preflights the selected corpus before any site writes.
- A formatting fix must repair the source, not just `content/e/` or another site's generated `content/ext/e/`. Keep the entire bilingual pair in one batch where possible. For generator/validator changes, add a regression test proving that an invalid source is rejected and an existing valid output remains intact, then run `go test ./...`.
- When generation is authorized, write to an explicit staging `--output` first and run the target site's strict Markdown, Hugo, rendered-content, and internal-link checks. Source lint does not replace target-site acceptance. Keep source repair, DB load, generation, commit/push, deployment, and public verification distinct in the handoff.
- Check: no front matter/H1; exactly one `## Usage`; `Sources:` plus linked bullets immediately after it; canonical name present; balanced code fences; no placeholders; core use cases/examples and confirmed caveats covered.
- When a compatible Phase B run directory and manifest exist, run `python3 bin/stub_phaseb.py lint --run-dir <run-dir> --manifest <manifest> --name <extension>` and use `check` for the final batch gate. Otherwise perform equivalent focused checks and run `git diff --check`.
- Report completed, unchanged-after-review, blocked, and deferred items separately. State source/diff validation separately from any database load, generated-site build, commit, or publication; none of those later layers is implied by a stub edit.

## Git 内容与大文件规则

用户于 2026-10-10 明确的通用仓库规则：

- 仓库默认保存源代码、配置和必要的轻量文档、元数据。构建产物、可执行文件、软件包、容器/虚拟机镜像、下载压缩包、数据库导出和大体积原始数据默认不进入 Git；文本格式的数据大文件同样需要检查。
- 在工作树内下载或生成上述文件之前，先选择仓库外的制品/数据位置，或补上合理的 `.gitignore` 规则并验证实际路径已被忽略，防止暂存和代理恢复/差异快照将它们写入 Git 对象库。
- 暂存、提交前检查新增文件的大小、类型和 staged 文件列表，只加入本次需要的源文件；不得用强制添加绕过忽略规则。优先保留来源清单、校验值及可复现的下载/构建方法。
- 对二进制素材、数据集或生成文件是否应纳入版本控制拿不准时，先问用户，再暂存或提交；确需例外时也先取得用户明确指示。
- 本规则不授权删除已有文件、stash、恢复快照或改写 Git 历史。忽略规则不会清除已存在的 Git 对象，历史清理需要单独明确范围。
