#!/usr/bin/env python3
import importlib.util
import json
import pathlib
import tempfile
import unittest
from unittest import mock


MODULE_PATH = pathlib.Path(__file__).with_name("github_activity.py")
spec = importlib.util.spec_from_file_location("github_activity", MODULE_PATH)
github_activity = importlib.util.module_from_spec(spec)
spec.loader.exec_module(github_activity)


class GitHubActivityParsingTest(unittest.TestCase):
    def test_normalize_github_repo_urls(self):
        cases = {
            "https://github.com/PGVECTOR/pgvector.git": ("PGVECTOR", "pgvector", "https://github.com/pgvector/pgvector"),
            "http://github.com/pgvector/pgvector/issues?q=x": ("pgvector", "pgvector", "https://github.com/pgvector/pgvector"),
            "git@github.com:Timescale/timescaledb.git": ("Timescale", "timescaledb", "https://github.com/timescale/timescaledb"),
            "https://github.com/Snowflake-Labs/pg_lake/tree/main/pg_extension_base": (
                "Snowflake-Labs",
                "pg_lake",
                "https://github.com/snowflake-labs/pg_lake",
            ),
        }
        for raw, expected in cases.items():
            with self.subTest(raw=raw):
                self.assertEqual(github_activity.normalize_github_repo(raw), expected)

    def test_live_manifest_round_trip_and_extension_resolution(self):
        live = [
            {
                "url_norm": "https://github.com/example/shared",
                "repo_owner": "example",
                "repo_name": "shared",
                "api_url": "https://api.github.com/repos/example/shared",
                "extension_ids": [10, 20],
                "extension_names": ["alpha", "beta-ext"],
                "extension_count": 2,
            }
        ]
        rows = github_activity.manifest_rows_from_live(live)
        self.assertEqual(
            github_activity.resolve_extension_urls(rows, ["BETA-EXT"]),
            {"https://github.com/example/shared"},
        )
        with tempfile.TemporaryDirectory() as tmpdir:
            path = github_activity.write_manifest(
                pathlib.Path(tmpdir) / "manifest.csv", rows
            )
            _, loaded = github_activity.load_manifest(path)
        self.assertEqual(loaded[0]["extension_ids_list"], [10, 20])
        self.assertEqual(loaded[0]["extension_names_list"], ["alpha", "beta-ext"])

    def test_extension_resolution_rejects_non_github_extension(self):
        with self.assertRaises(github_activity.ManifestError):
            github_activity.resolve_extension_urls([], ["not-on-github"])

    def test_target_mapping_upsert_writes_only_selected_repository(self):
        rows = [
            {
                "url_norm": f"https://github.com/example/{name}",
                "repo_owner": "example",
                "repo_name": name,
                "api_url": f"https://api.github.com/repos/example/{name}",
                "extension_ids": f"{{{index}}}",
                "extension_names": f"{{{name}}}",
                "extension_count": "1",
            }
            for index, name in enumerate(("one", "two"), start=1)
        ]
        selected = "https://github.com/example/two"
        result = {
            "staged": 2,
            "selected": 1,
            "missing_before": 1,
            "mapping_changed": 1,
            "upserted": 1,
            "states": [{"url_norm": selected, "status": "pending"}],
        }
        with mock.patch.object(
            github_activity, "run_psql_script_json", return_value=result
        ) as run_sql:
            self.assertEqual(
                github_activity.upsert_target_mappings(
                    rows, selected_urls=[selected]
                ),
                result,
            )
        sql = run_sql.call_args.args[0]
        selected_csv = sql.split(
            "COPY q_selected_url FROM STDIN WITH (FORMAT csv, HEADER true);\n",
            1,
        )[1].split("\\.\n", 1)[0]
        self.assertEqual(selected_csv, f"url_norm\n{selected}\n")

    def test_target_mapping_upsert_rejects_unknown_or_empty_selection(self):
        row = {
            "url_norm": "https://github.com/example/one",
            "repo_owner": "example",
            "repo_name": "one",
            "api_url": "https://api.github.com/repos/example/one",
            "extension_ids": "{1}",
            "extension_names": "{one}",
            "extension_count": "1",
        }
        with self.assertRaises(ValueError):
            github_activity.upsert_target_mappings([row], selected_urls=[])
        with self.assertRaises(ValueError):
            github_activity.upsert_target_mappings(
                [row], selected_urls=["https://github.com/example/missing"]
            )

    def test_parse_rest_commit_uses_committer_date_and_sha_url(self):
        payload = [
            {
                "sha": "abc123",
                "html_url": "https://github.com/o/r/commit/abc123",
                "commit": {
                    "author": {"date": "2026-01-01T00:00:00Z"},
                    "committer": {"date": "2026-01-02T03:04:05Z"},
                },
            }
        ]

        commit = github_activity.parse_commit_response(payload)

        self.assertEqual(commit["last_commit_at"], "2026-01-02T03:04:05Z")
        self.assertEqual(commit["last_commit_sha"], "abc123")
        self.assertEqual(commit["last_commit_html_url"], "https://github.com/o/r/commit/abc123")

    def test_rate_limit_check_handles_list_json_payloads(self):
        response = {
            "status": 200,
            "headers": {"x-ratelimit-remaining": "4999"},
            "data": [{"sha": "abc123"}],
            "text": "[{}]",
        }

        self.assertFalse(github_activity.is_rate_limited(response))

    def test_graphql_release_and_annotated_tag_dates_are_kept_separate(self):
        payload = {
            "data": {
                "repository": {
                    "releases": {
                        "nodes": [
                            {
                                "tagName": "v1.0.0",
                                "publishedAt": None,
                                "createdAt": "2025-12-01T00:00:00Z",
                            }
                        ]
                    },
                    "refs": {
                        "nodes": [
                            {
                                "name": "v1.1.0",
                                "target": {
                                    "__typename": "Tag",
                                    "oid": "tag-object",
                                    "tagger": {"date": "2026-01-02T00:00:00Z"},
                                    "target": {
                                        "__typename": "Commit",
                                        "oid": "commit-object",
                                        "committedDate": "2026-01-01T00:00:00Z",
                                    },
                                },
                            }
                        ]
                    },
                }
            }
        }

        activity = github_activity.parse_release_tag_response(payload)

        self.assertEqual(activity["latest_release_tag"], "v1.0.0")
        self.assertEqual(activity["last_release_at"], "2025-12-01T00:00:00Z")
        self.assertEqual(activity["latest_tag_name"], "v1.1.0")
        self.assertEqual(activity["latest_tag_at"], "2026-01-02T00:00:00Z")
        self.assertEqual(activity["latest_tag_date_source"], "tagger.date")
        self.assertEqual(activity["latest_tag_commit_sha"], "commit-object")
        self.assertEqual(activity["latest_tag_commit_at"], "2026-01-01T00:00:00Z")
        self.assertEqual(activity["last_release_or_tag_at"], "2026-01-02T00:00:00Z")
        self.assertEqual(activity["last_release_or_tag_source"], "tag")

    def test_graphql_lightweight_tag_uses_commit_date(self):
        payload = {
            "data": {
                "repository": {
                    "releases": {"nodes": []},
                    "refs": {
                        "nodes": [
                            {
                                "name": "v2",
                                "target": {
                                    "__typename": "Commit",
                                    "oid": "commit-sha",
                                    "committedDate": "2024-05-06T07:08:09Z",
                                },
                            }
                        ]
                    },
                }
            }
        }

        activity = github_activity.parse_release_tag_response(payload)

        self.assertEqual(activity["latest_tag_at"], "2024-05-06T07:08:09Z")
        self.assertEqual(activity["latest_tag_date_source"], "commit.committedDate")
        self.assertEqual(activity["latest_tag_commit_sha"], "commit-sha")
        self.assertEqual(activity["last_release_or_tag_source"], "tag")

    def test_last_update_uses_max_of_commit_release_and_tag(self):
        row = github_activity.merge_activity_dates(
            {
                "last_commit_at": "2026-02-01T00:00:00Z",
                "last_release_or_tag_at": "2026-01-01T00:00:00Z",
                "last_release_or_tag_source": "release",
                "last_release_at": "2025-12-01T00:00:00Z",
                "latest_tag_at": "2026-01-01T00:00:00Z",
            }
        )

        self.assertEqual(row["last_update_at"], "2026-02-01T00:00:00Z")
        self.assertEqual(row["last_update_source"], "commit")
        self.assertEqual(row["last_commit_date"], "2026-02-01T00:00:00Z")
        self.assertEqual(row["last_release_date"], "2025-12-01T00:00:00Z")
        self.assertEqual(row["last_tag_date"], "2026-01-01T00:00:00Z")
        self.assertEqual(row["last_update_date"], "2026-02-01T00:00:00Z")

    def test_successful_fetch_keeps_complete_raw_evidence(self):
        class FakeClient:
            def rest_repo(self, owner, repo):
                return {
                    "status": 200,
                    "headers": {"etag": '"etag"', "x-ratelimit-remaining": "4999"},
                    "request_url": "https://api.github.com/repos/old/repo",
                    "url": "https://api.github.com/repositories/212439157",
                    "request_attempts": 1,
                    "data": {
                        "full_name": "example/repo",
                        "html_url": "https://github.com/example/repo",
                        "stargazers_count": 1,
                        "forks_count": 2,
                        "watchers_count": 1,
                        "subscribers_count": 3,
                        "pushed_at": "2026-01-01T00:00:00Z",
                        "updated_at": "2026-01-02T00:00:00Z",
                        "default_branch": "main",
                        "archived": False,
                    },
                }

            def rest_latest_commit(self, owner, repo, default_branch):
                url = "https://api.github.com/repos/example/repo/commits?per_page=1&sha=main"
                return {
                    "status": 200,
                    "headers": {"x-ratelimit-remaining": "4998"},
                    "request_url": url,
                    "url": url,
                    "request_attempts": 1,
                    "data": [
                        {
                            "sha": "abc",
                            "html_url": "https://github.com/example/repo/commit/abc",
                            "commit": {"committer": {"date": "2026-01-03T00:00:00Z"}},
                        }
                    ],
                }

            def release_tag_activity(self, owner, repo):
                return {
                    "status": 200,
                    "headers": {},
                    "request_url": "https://api.github.com/graphql",
                    "url": "https://api.github.com/graphql",
                    "request_attempts": 1,
                    "data": {
                        "data": {
                            "repository": {
                                "nameWithOwner": "example/repo",
                                "isArchived": False,
                                "releases": {"nodes": []},
                                "refs": {"nodes": []},
                            },
                            "rateLimit": {
                                "remaining": 4998,
                                "resetAt": "2026-01-04T00:00:00Z",
                            },
                        }
                    },
                }

        budget = github_activity.RateBudget(request_retries=2)
        budget.core_remaining = 5000
        budget.core_reset = "2026-07-15T05:00:00Z"
        budget.graphql_remaining = 5000
        budget.graphql_reset = "2026-07-15T05:00:00Z"
        record = github_activity.fetch_activity(
            {
                "url_norm": "https://github.com/old/repo",
                "repo_owner": "old",
                "repo_name": "repo",
                "api_url": "https://api.github.com/repos/old/repo",
            },
            FakeClient(),
            budget,
        )
        audit = json.loads(record["activity_json"])
        self.assertEqual(record["status"], "fetched")
        self.assertIsNotNone(github_activity.parse_iso(audit["collected_at"]))
        self.assertEqual(audit["repo"]["status"], 200)
        self.assertEqual(audit["latest_commit"]["status"], 200)
        self.assertEqual(audit["release_tag"]["status"], 200)
        self.assertEqual(audit["repo"]["data"], json.loads(record["api_json"]))
        self.assertTrue(audit["redirect"]["detected"])
        self.assertEqual(record["api_url"], "https://api.github.com/repos/example/repo")
        self.assertEqual(
            audit["redirect"]["final_api_url"],
            "https://api.github.com/repositories/212439157",
        )
        self.assertIsInstance(record["rate_limit_remaining"], int)
        self.assertIsNotNone(github_activity.parse_iso(record["rate_limit_reset"]))

    def test_current_null_default_branch_does_not_reuse_cached_branch(self):
        class EmptyClient:
            def rest_repo(self, owner, repo):
                return {
                    "status": 200,
                    "headers": {
                        "etag": '"etag"',
                        "x-ratelimit-remaining": "4999",
                        "x-ratelimit-reset": "1784102400",
                    },
                    "request_url": "https://api.github.com/repos/example/empty",
                    "url": "https://api.github.com/repos/example/empty",
                    "request_attempts": 1,
                    "data": {
                        "full_name": "Example/Empty",
                        "html_url": "https://github.com/Example/Empty",
                        "stargazers_count": 0,
                        "forks_count": 0,
                        "watchers_count": 0,
                        "subscribers_count": 0,
                        "pushed_at": None,
                        "updated_at": "2026-01-02T00:00:00Z",
                        "default_branch": None,
                        "archived": False,
                    },
                }

            def rest_latest_commit(self, owner, repo, default_branch):
                raise AssertionError("stale cached branch must not be queried")

            def release_tag_activity(self, owner, repo):
                return {
                    "status": 200,
                    "headers": {},
                    "request_url": "https://api.github.com/graphql",
                    "url": "https://api.github.com/graphql",
                    "request_attempts": 1,
                    "data": {
                        "data": {
                            "repository": {
                                "nameWithOwner": "Example/Empty",
                                "isArchived": False,
                                "releases": {"nodes": []},
                                "refs": {"nodes": []},
                            },
                            "rateLimit": {
                                "remaining": 4998,
                                "resetAt": "2026-07-15T05:00:00Z",
                            },
                        }
                    },
                }

        budget = github_activity.RateBudget(request_retries=2)
        budget.core_remaining = 5000
        budget.core_reset = "2026-07-15T05:00:00Z"
        budget.graphql_remaining = 5000
        budget.graphql_reset = "2026-07-15T05:00:00Z"
        record = github_activity.fetch_activity(
            {
                "url_norm": "https://github.com/example/empty",
                "repo_owner": "example",
                "repo_name": "empty",
                "api_url": "https://api.github.com/repos/example/empty",
                "default_branch": "stale-main",
                "api_default_branch": "stale-main",
            },
            EmptyClient(),
            budget,
        )
        audit = json.loads(record["activity_json"])
        self.assertEqual(record["status"], "fetched")
        self.assertIsNone(record["default_branch"])
        self.assertEqual(record["api_url"], "https://api.github.com/repos/example/empty")
        self.assertFalse(audit["redirect"]["detected"])
        self.assertIsNone(audit["latest_commit"])
        self.assertEqual(audit["warnings"][0]["kind"], "missing_default_branch")


class GitHubActivityResumeTest(unittest.TestCase):
    manifest_sha = "a" * 64
    urls = [
        "https://github.com/example/one",
        "https://github.com/example/two",
        "https://github.com/example/three",
    ]

    def make_plan_ledger(self, path, planned_urls=None, statuses=None, limit=0):
        planned_urls = list(planned_urls or self.urls)
        plan = github_activity.build_run_plan(
            self.manifest_sha,
            statuses or [],
            [],
            limit,
            planned_urls,
        )
        checksum = github_activity.plan_sha256(plan)
        github_activity.append_ledger(
            path,
            {
                "event": "run_plan",
                "manifest_sha256": self.manifest_sha,
                "plan_sha256": checksum,
                "plan": plan,
            },
            durable=True,
        )
        return plan, checksum

    def append_checkpoint(self, path, checksum, urls, result=None):
        github_activity.append_ledger(
            path,
            {
                "event": "batch_applied",
                "manifest_sha256": self.manifest_sha,
                "plan_sha256": checksum,
                "urls": list(urls),
                "result": result
                or {"staged": len(urls), "success": len(urls), "failed": 0},
            },
            durable=True,
        )

    def test_attempt_only_is_not_applied_and_valid_checkpoint_is(self):
        with tempfile.TemporaryDirectory() as tmpdir:
            ledger = pathlib.Path(tmpdir) / "run.jsonl"
            plan, checksum = self.make_plan_ledger(ledger)
            github_activity.append_ledger(
                ledger,
                {
                    "event": "attempt",
                    "manifest_sha256": self.manifest_sha,
                    "plan_sha256": checksum,
                    "record": {"url_norm": self.urls[0]},
                },
            )
            loaded_plan, loaded_checksum, applied = github_activity.load_resume_state(
                ledger, self.manifest_sha, set(self.urls)
            )
            self.assertEqual(loaded_plan, plan)
            self.assertEqual(loaded_checksum, checksum)
            self.assertEqual(applied, set())

            self.append_checkpoint(ledger, checksum, self.urls[:2])
            _, _, applied = github_activity.load_resume_state(
                ledger, self.manifest_sha, set(self.urls)
            )
            self.assertEqual(applied, set(self.urls[:2]))

    def test_checkpoint_validation_rejects_manifest_plan_urls_and_counts(self):
        bad_events = [
            {"manifest_sha256": "b" * 64},
            {"plan_sha256": "b" * 64},
            {"urls": [self.urls[0], self.urls[0]]},
            {"urls": ["https://github.com/not/in-plan"]},
            {"result": {"staged": 2, "success": 1, "failed": 0}},
        ]
        for override in bad_events:
            with self.subTest(override=override), tempfile.TemporaryDirectory() as tmpdir:
                ledger = pathlib.Path(tmpdir) / "run.jsonl"
                _, checksum = self.make_plan_ledger(ledger)
                event = {
                    "event": "batch_applied",
                    "manifest_sha256": self.manifest_sha,
                    "plan_sha256": checksum,
                    "urls": [self.urls[0]],
                    "result": {"staged": 1, "success": 1, "failed": 0},
                }
                event.update(override)
                github_activity.append_ledger(ledger, event, durable=True)
                with self.assertRaises(github_activity.ManifestError):
                    github_activity.load_resume_state(
                        ledger, self.manifest_sha, set(self.urls)
                    )

    def test_repeated_checkpoint_url_is_rejected(self):
        with tempfile.TemporaryDirectory() as tmpdir:
            ledger = pathlib.Path(tmpdir) / "run.jsonl"
            _, checksum = self.make_plan_ledger(ledger)
            self.append_checkpoint(ledger, checksum, [self.urls[0]])
            self.append_checkpoint(ledger, checksum, [self.urls[0]])
            with self.assertRaises(github_activity.ManifestError):
                github_activity.load_resume_state(
                    ledger, self.manifest_sha, set(self.urls)
                )

    def test_torn_final_record_is_truncated_before_future_appends(self):
        with tempfile.TemporaryDirectory() as tmpdir:
            ledger = pathlib.Path(tmpdir) / "run.jsonl"
            plan, checksum = self.make_plan_ledger(ledger)
            self.append_checkpoint(ledger, checksum, [self.urls[0]])
            with ledger.open("ab") as handle:
                handle.write(b'{"event":"batch_applied"')

            loaded_plan, _, applied = github_activity.load_resume_state(
                ledger, self.manifest_sha, set(self.urls)
            )
            self.assertEqual(loaded_plan, plan)
            self.assertEqual(applied, {self.urls[0]})
            github_activity.append_ledger(ledger, {"event": "safe_stop"}, durable=True)
            github_activity.append_ledger(ledger, {"event": "run_start"}, durable=True)
            events = github_activity.read_ledger_events(ledger)
            self.assertEqual(
                [event["event"] for _, event in events],
                ["run_plan", "batch_applied", "safe_stop", "run_start"],
            )

    def test_valid_final_json_without_newline_is_delimited(self):
        with tempfile.TemporaryDirectory() as tmpdir:
            ledger = pathlib.Path(tmpdir) / "run.jsonl"
            plan = github_activity.build_run_plan(
                self.manifest_sha, [], [], 0, self.urls
            )
            checksum = github_activity.plan_sha256(plan)
            event = {
                "event": "run_plan",
                "manifest_sha256": self.manifest_sha,
                "plan_sha256": checksum,
                "plan": plan,
            }
            ledger.write_text(json.dumps(event, separators=(",", ":")), encoding="utf-8")
            github_activity.load_resume_state(ledger, self.manifest_sha, set(self.urls))
            self.assertTrue(ledger.read_bytes().endswith(b"\n"))
            github_activity.append_ledger(ledger, {"event": "run_start"}, durable=True)
            self.assertEqual(len(github_activity.read_ledger_events(ledger)), 2)

    def test_newline_terminated_corrupt_record_fails_closed(self):
        with tempfile.TemporaryDirectory() as tmpdir:
            ledger = pathlib.Path(tmpdir) / "run.jsonl"
            self.make_plan_ledger(ledger)
            with ledger.open("ab") as handle:
                handle.write(b"not-json\n")
            with self.assertRaises(github_activity.ManifestError):
                github_activity.load_resume_state(
                    ledger, self.manifest_sha, set(self.urls)
                )

    def test_missing_resume_ledger_and_nonempty_new_ledger_are_rejected(self):
        with tempfile.TemporaryDirectory() as tmpdir:
            missing = pathlib.Path(tmpdir) / "missing.jsonl"
            with self.assertRaises(github_activity.ManifestError):
                github_activity.load_resume_state(
                    missing, self.manifest_sha, set(self.urls)
                )

            manifest = pathlib.Path(tmpdir) / "manifest.csv"
            manifest.write_text("placeholder", encoding="utf-8")
            existing = pathlib.Path(tmpdir) / "existing.jsonl"
            existing.write_text("{}\n", encoding="utf-8")
            rows = [{"url_norm": self.urls[0]}]
            with mock.patch.object(
                github_activity, "load_manifest", return_value=(manifest, rows)
            ), mock.patch.object(github_activity, "database_preflight") as preflight:
                with self.assertRaises(github_activity.ManifestError):
                    github_activity.main_checked(
                        ["--manifest", str(manifest), "--ledger", str(existing)]
                    )
                preflight.assert_not_called()

    def test_resume_rejects_selection_flags_and_production_requires_ledger(self):
        with self.assertRaises(github_activity.ManifestError):
            github_activity.main_checked(["--manifest", "unused", "--resume", "--ledger", "x", "--status", "error"])
        with self.assertRaises(github_activity.ManifestError):
            github_activity.main_checked(["--manifest", "unused"])

    def test_apply_failure_has_no_checkpoint_and_success_has_one(self):
        with tempfile.TemporaryDirectory() as tmpdir:
            ledger = pathlib.Path(tmpdir) / "run.jsonl"
            _, checksum = self.make_plan_ledger(ledger)
            records = [{"url_norm": self.urls[0]}]
            with mock.patch.object(
                github_activity,
                "apply_activity_updates",
                side_effect=RuntimeError("database failed"),
            ):
                with self.assertRaises(RuntimeError):
                    github_activity.apply_and_checkpoint_batch(
                        records,
                        None,
                        ledger,
                        pathlib.Path("manifest.csv"),
                        self.manifest_sha,
                        checksum,
                        self.urls,
                    )
            self.assertEqual(
                [event["event"] for _, event in github_activity.read_ledger_events(ledger)],
                ["run_plan"],
            )

            with mock.patch.object(
                github_activity,
                "apply_activity_updates",
                return_value={"staged": 1, "success": 1, "failed": 0},
            ):
                github_activity.apply_and_checkpoint_batch(
                    records,
                    None,
                    ledger,
                    pathlib.Path("manifest.csv"),
                    self.manifest_sha,
                    checksum,
                    self.urls,
                )
            events = github_activity.read_ledger_events(ledger)
            self.assertEqual([event["event"] for _, event in events], ["run_plan", "batch_applied"])

    def test_checkpoint_failure_causes_conservative_replay(self):
        with tempfile.TemporaryDirectory() as tmpdir:
            ledger = pathlib.Path(tmpdir) / "run.jsonl"
            _, checksum = self.make_plan_ledger(ledger)
            records = [{"url_norm": self.urls[0]}]
            real_append = github_activity.append_ledger

            def fail_checkpoint(path, payload, durable=False):
                if payload.get("event") == "batch_applied":
                    raise OSError("disk full")
                return real_append(path, payload, durable)

            with mock.patch.object(
                github_activity,
                "apply_activity_updates",
                return_value={"staged": 1, "success": 1, "failed": 0},
            ), mock.patch.object(github_activity, "append_ledger", side_effect=fail_checkpoint):
                with self.assertRaises(OSError):
                    github_activity.apply_and_checkpoint_batch(
                        records,
                        None,
                        ledger,
                        pathlib.Path("manifest.csv"),
                        self.manifest_sha,
                        checksum,
                        self.urls,
                    )
            _, _, applied = github_activity.load_resume_state(
                ledger, self.manifest_sha, set(self.urls)
            )
            self.assertEqual(applied, set())


class GitHubActivityMainBatchingTest(unittest.TestCase):
    def run_mocked_main(
        self,
        row_count,
        fetch_side_effect=None,
        rate_preflight_side_effect=None,
        summary_side_effect=None,
    ):
        tempdir = tempfile.TemporaryDirectory()
        self.addCleanup(tempdir.cleanup)
        root = pathlib.Path(tempdir.name)
        manifest = root / "manifest.csv"
        manifest.write_text("stable manifest bytes\n", encoding="utf-8")
        ledger = root / "run.jsonl"
        rows = [
            {
                "url_norm": f"https://github.com/example/repo-{index:02d}",
                "repo_owner": "example",
                "repo_name": f"repo-{index:02d}",
                "api_url": f"https://api.github.com/repos/example/repo-{index:02d}",
            }
            for index in range(row_count)
        ]
        states = [
            {
                "url_norm": row["url_norm"],
                "status": "fetched",
                "default_branch": "main",
                "api_default_branch": "main",
            }
            for row in rows
        ]
        budget = github_activity.RateBudget(request_retries=2)
        budget.core_remaining = 5000
        budget.graphql_remaining = 5000

        if fetch_side_effect is None:
            def fetch_side_effect(target, _client, _budget):
                return {
                    "url_norm": target["url_norm"],
                    "status": "fetched",
                    "success": True,
                }

        applied_batch_sizes = []

        def apply(records, dsn=None):
            applied_batch_sizes.append(len(records))
            return {"staged": len(records), "success": len(records), "failed": 0}

        rate_preflight_patch = (
            mock.patch.object(
                github_activity,
                "preflight_rate_limits",
                side_effect=rate_preflight_side_effect,
            )
            if rate_preflight_side_effect is not None
            else mock.patch.object(
                github_activity, "preflight_rate_limits", return_value=budget
            )
        )
        summary_patch = (
            mock.patch.object(
                github_activity, "fetch_summary", side_effect=summary_side_effect
            )
            if summary_side_effect is not None
            else mock.patch.object(github_activity, "fetch_summary", return_value={})
        )
        patches = [
            mock.patch.object(github_activity, "load_manifest", return_value=(manifest, rows)),
            mock.patch.object(
                github_activity,
                "database_preflight",
                return_value=(
                    {
                        "target_count": row_count,
                        "backup": {"rows": 1342, "fingerprint": "expected"},
                    },
                    {row["url_norm"]: state for row, state in zip(rows, states)},
                ),
            ),
            mock.patch.object(
                github_activity,
                "upsert_target_mappings",
                return_value={"staged": row_count, "missing_before": 0, "states": states},
            ),
            mock.patch.object(github_activity, "discover_token", return_value=("token", "test")),
            rate_preflight_patch,
            mock.patch.object(github_activity, "fetch_activity", side_effect=fetch_side_effect),
            mock.patch.object(github_activity, "apply_activity_updates", side_effect=apply),
            summary_patch,
            mock.patch.object(github_activity, "print_summary"),
            mock.patch.object(github_activity.time, "sleep"),
        ]
        mocks = [patch.start() for patch in patches]
        self.addCleanup(lambda: [patch.stop() for patch in reversed(patches)])
        result = github_activity.main_checked(
            [
                "--manifest",
                str(manifest),
                "--ledger",
                str(ledger),
                "--batch-size",
                "25",
                "--min-delay",
                "0",
            ]
        )
        return result, ledger, applied_batch_sizes

    def test_full_and_final_partial_batches_are_checkpointed(self):
        result, ledger, applied_batch_sizes = self.run_mocked_main(26)
        self.assertEqual(result, 0)
        self.assertEqual(applied_batch_sizes, [25, 1])
        events = github_activity.read_ledger_events(ledger)
        batches = [event for _, event in events if event.get("event") == "batch_applied"]
        self.assertEqual([len(event["urls"]) for event in batches], [25, 1])

    def test_safe_stop_flushes_partial_batch_before_durable_stop(self):
        calls = 0

        def fetch(target, _client, _budget):
            nonlocal calls
            calls += 1
            if calls == 2:
                raise github_activity.SafeStop("reserve reached")
            return {"url_norm": target["url_norm"], "status": "fetched", "success": True}

        result, ledger, applied_batch_sizes = self.run_mocked_main(3, fetch_side_effect=fetch)
        self.assertEqual(result, 3)
        self.assertEqual(applied_batch_sizes, [1])
        event_names = [event["event"] for _, event in github_activity.read_ledger_events(ledger)]
        self.assertLess(event_names.index("batch_applied"), event_names.index("safe_stop"))

    def test_rate_limit_preflight_stop_is_durable_and_resumable(self):
        result, ledger, applied_batch_sizes = self.run_mocked_main(
            3,
            rate_preflight_side_effect=github_activity.SafeStop(
                "GraphQL reserve would be crossed; reset=2026-07-15T12:00:00Z"
            ),
        )
        self.assertEqual(result, 3)
        self.assertEqual(applied_batch_sizes, [])
        events = [event for _, event in github_activity.read_ledger_events(ledger)]
        self.assertEqual(events[-1]["event"], "safe_stop")
        self.assertEqual(events[-1]["stage"], "rate_limit_preflight")
        self.assertTrue(events[-1]["plan_sha256"])

    def test_summary_failure_does_not_mask_durable_safe_stop(self):
        calls = 0

        def fetch(target, _client, _budget):
            nonlocal calls
            calls += 1
            if calls == 2:
                raise github_activity.SafeStop("reserve reached")
            return {"url_norm": target["url_norm"], "status": "fetched", "success": True}

        result, ledger, applied_batch_sizes = self.run_mocked_main(
            3,
            fetch_side_effect=fetch,
            summary_side_effect=RuntimeError("summary unavailable"),
        )
        self.assertEqual(result, 3)
        self.assertEqual(applied_batch_sizes, [1])
        self.assertEqual(
            [event["event"] for _, event in github_activity.read_ledger_events(ledger)][-1],
            "safe_stop",
        )


if __name__ == "__main__":
    unittest.main()
