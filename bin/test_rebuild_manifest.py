#!/usr/bin/env python3
import importlib.util
import pathlib
import sys
import unittest
from unittest import mock


sys.dont_write_bytecode = True
MODULE_PATH = pathlib.Path(__file__).with_name("rebuild_manifest.py")
spec = importlib.util.spec_from_file_location("rebuild_manifest", MODULE_PATH)
manifest = importlib.util.module_from_spec(spec)
spec.loader.exec_module(manifest)

ALL_PG = "{18,17,16,15,14}"


def catalog_row(name, package=None, **changes):
    row = {
        "name": name,
        "pkg": package or name,
        "repo": "PGDG",
        "rpm_repo": "PGDG",
        "deb_repo": "PGDG",
        "rpm_pg": ALL_PG,
        "deb_pg": ALL_PG,
        "pg_ver": ALL_PG,
        "version": "1.3.1",
        "rpm_ver": "1.3.1",
        "deb_ver": "1.3.1",
    }
    row.update(changes)
    return row


def mobility_rows():
    return [catalog_row("mobilitydb"), catalog_row("mobilitydb_datagen", "mobilitydb")]


def official_mobility_packages():
    # Native PGDG samples: Jammy 1.2 lacks DataGen; other platforms have 1.3.
    # The official Jammy PG18 binary is absent on both architectures.
    packages = {}
    for platform in manifest.DEB_PLATFORMS:
        for pg in range(14, 19):
            if platform.startswith("u22.") and pg == 18:
                continue
            version = "1.2.0-2.pgdg22.04+1" if platform.startswith("u22.") else "1.3.0-1.pgdg12+1"
            packages[("deb", "mobilitydb", platform, pg)] = {version}
    return packages


class PGDGSupplementPolicyTest(unittest.TestCase):
    def test_h3_pigsty_build_uses_catalog_extension_and_pinned_core_sources(self):
        row = catalog_row(
            "h3", "pg_h3", id="1530", lead="t", lead_ext="h3",
            source="h3-pg-4.5.0.tar.gz h3-4.5.0.tar.gz",
            repo="PIGSTY", rpm_repo="PIGSTY", deb_repo="PIGSTY",
        )
        sources, conflicts = manifest.source_list(
            "rpm", "h3-pg_$v", "h3-pg", [row], [14, 15, 16, 17, 18], "el9.x86_64"
        )
        self.assertEqual(sources, ["h3-pg-4.5.0.tar.gz", "h3-4.5.0.tar.gz"])
        self.assertEqual(conflicts, [])

    def test_h3_companion_does_not_invent_gaps_in_lead_only_matrix(self):
        rows = [catalog_row("h3", "pg_h3"), catalog_row("h3_postgis", "pg_h3")]
        # All 28 covered PGDG RPMs contain h3_postgis.control in native samples.
        real_gaps = {("el8.x86_64", 17), ("el8.x86_64", 18)}
        packages = {
            ("rpm", "pg_h3", platform, pg): {"4.2.3-1PGDG.rhel9"}
            for platform in manifest.RPM_PLATFORMS
            for pg in range(14, 19)
            if (platform, pg) not in real_gaps
        }
        selected = {
            (platform, pg)
            for platform in manifest.RPM_PLATFORMS
            for pg in range(14, 19)
            if manifest.pgdg_build_reasons("rpm", rows, platform, pg, packages)
        }
        self.assertEqual(selected, real_gaps)

    def test_mobility_current_eight_companion_gaps_and_eight_security_exceptions(self):
        packages = official_mobility_packages()
        selected, companions, security, missing_packages = set(), set(), set(), set()
        for platform in manifest.DEB_PLATFORMS:
            for pg in range(14, 19):
                reasons = manifest.pgdg_build_reasons("deb", mobility_rows(), platform, pg, packages)
                coordinate = (platform, pg)
                if reasons:
                    selected.add(coordinate)
                if "missing_companion:mobilitydb_datagen" in reasons:
                    companions.add(coordinate)
                if "security:CVE-2026-102639" in reasons:
                    security.add(coordinate)
                if "missing_package:mobilitydb" in reasons:
                    missing_packages.add(coordinate)
        self.assertEqual(companions, {
            (platform, pg) for platform in ("u22.x86_64", "u22.aarch64") for pg in range(14, 18)
        })
        self.assertEqual(security, {
            (platform, 18) for platform in manifest.DEB_PLATFORMS if not platform.startswith("u22.")
        })
        self.assertEqual(missing_packages, {("u22.x86_64", 18), ("u22.aarch64", 18)})
        self.assertEqual(len(selected), 18)

    def test_datagen_gap_ends_when_released_pgdg_13_bundle_is_available(self):
        key = ("deb", "mobilitydb", "u22.x86_64", 14)
        packages = {key: {"1.2.0-2.pgdg22.04+1", "1.3.0~rc1-1.pgdg22.04+1"}}
        self.assertIn("missing_companion:mobilitydb_datagen",
                      manifest.pgdg_build_reasons("deb", mobility_rows(), key[2], 14, packages))
        packages[key].add("1.3.0-1.pgdg22.04+1")
        self.assertEqual(manifest.pgdg_build_reasons("deb", mobility_rows(), key[2], 14, packages), ())

    def test_security_exception_ends_with_fixed_package_even_when_old_package_remains(self):
        packages = official_mobility_packages()
        for platform in manifest.DEB_PLATFORMS:
            if platform.startswith("u22."):
                continue
            key = ("deb", "mobilitydb", platform, 18)
            packages[key].add("1.3.1-1.pgdg12+1")
            self.assertEqual(manifest.pgdg_build_reasons("deb", mobility_rows(), platform, 18, packages), ())

    def test_unknown_and_prerelease_packages_do_not_end_security_exception(self):
        key = ("deb", "mobilitydb", "d12.x86_64", 18)
        for candidate in ("unknown", "1.3.1~rc1-1.pgdg12+1", "1.3.1rc1-1.pgdg12+1",
                          "1.3.1-rc1", "1.3.1~beta1-1.pgdg12+1", "1.4.0~alpha1-1.pgdg12+1"):
            with self.subTest(candidate=candidate):
                packages = {key: {"1.3.0-1.pgdg12+1", candidate}}
                self.assertEqual(manifest.pgdg_build_reasons("deb", mobility_rows(), key[2], 18, packages),
                                 ("security:CVE-2026-102639",))

    def test_fixed_version_uses_numeric_release_and_not_epoch_or_lexical_order(self):
        key = ("deb", "mobilitydb", "d12.x86_64", 18)
        for fixed in ("1.3.10-1.pgdg12+1", "1.10.0-1.pgdg12+1", "1.3.1+dfsg-1.pgdg12+1"):
            with self.subTest(fixed=fixed):
                self.assertEqual(manifest.pgdg_build_reasons(
                    "deb", mobility_rows(), key[2], 18, {key: {fixed}}), ())
        self.assertIn("security:CVE-2026-102639", manifest.pgdg_build_reasons(
            "deb", mobility_rows(), key[2], 18, {key: {"2:1.3.0-1.pgdg12+1"}}))

    def test_security_policy_does_not_expand_to_old_pg_or_rpm(self):
        for fmt, platform, pg in (("deb", "d12.x86_64", 17), ("rpm", "el9.x86_64", 18)):
            with self.subTest(fmt=fmt, pg=pg):
                packages = {(fmt, "mobilitydb", platform, pg): {"1.3.0-1.pgdg12+1"}}
                self.assertEqual(manifest.pgdg_build_reasons(fmt, mobility_rows(), platform, pg, packages), ())

    def test_mobility_12_only_is_a_companion_gap_without_unproven_cve_claim(self):
        key = ("deb", "mobilitydb", "u22.x86_64", 18)
        packages = {key: {"1.2.0-2.pgdg22.04+1"}}
        reasons = manifest.pgdg_build_reasons("deb", mobility_rows(), key[2], 18, packages)
        self.assertEqual(reasons, ("missing_companion:mobilitydb_datagen",))
        self.assertEqual(manifest.supplement_origin("PGDG", reasons), "catalog_pgdg_gap")

    def test_only_rows_declaring_this_pg_are_required(self):
        key = ("deb", "mobilitydb", "u22.x86_64", 14)
        for declarations in ({"deb_pg": "{18}"}, {"pg_ver": "{18}"}):
            with self.subTest(declarations=declarations):
                rows = [catalog_row("mobilitydb"),
                        catalog_row("mobilitydb_datagen", "mobilitydb", **declarations)]
                self.assertEqual(manifest.pgdg_build_reasons(
                    "deb", rows, key[2], 14, {key: {"1.2.0-2.pgdg22.04+1"}}), ())

    def test_no_declared_row_does_not_schedule_a_job(self):
        rows = [catalog_row("h3", "pg_h3", pg_ver="{18}", rpm_pg="{18}")]
        self.assertEqual(manifest.pgdg_build_reasons("rpm", rows, "el8.x86_64", 14, {}), ())

    def test_pgdg_vector_086_prevents_general_local_087_override(self):
        rows = [catalog_row("vector", "pgvector", version="0.8.7", deb_ver="0.8.7")]
        packages = {("deb", "pgvector", "u22.x86_64", 18): {"0.8.6-1.pgdg22.04+2"}}
        self.assertEqual(manifest.pgdg_build_reasons("deb", rows, "u22.x86_64", 18, packages), ())

    def test_overall_pgdg_label_does_not_override_deb_pigsty_primary(self):
        rows = [catalog_row("topn", repo="PGDG", deb_repo="PIGSTY")]
        packages = {("deb", "topn", "u22.x86_64", 18): {"2.7.1-1.pgdg22.04+1"}}
        self.assertEqual(manifest.pgdg_build_reasons("deb", rows, "u22.x86_64", 18, packages),
                         ("format_primary",))

    def test_format_specific_package_keys_do_not_cross_cover(self):
        packages = {("rpm", "mobilitydb", "u22.x86_64", 14): {"1.3.1-1PGDG"}}
        self.assertEqual(manifest.pgdg_build_reasons("deb", mobility_rows(), "u22.x86_64", 14, packages),
                         ("missing_package:mobilitydb",))

    def test_loader_preserves_raw_versions_and_lead_package_identity(self):
        facts = ("deb|mobilitydb|u22.x86_64|14|1.2.0-2.pgdg22.04+1\n"
                 "deb|mobilitydb|d12.x86_64|18|1.3.1~rc1-1.pgdg12+1\n"
                 "rpm|pg_h3|el9.x86_64|18|4.2.3-1PGDG.rhel9")
        with mock.patch.object(manifest, "run_text", side_effect=[facts, "deb.pgdg|before/after"]):
            packages, stamps = manifest.load_pgdg_state()
        self.assertEqual(packages[("deb", "mobilitydb", "d12.x86_64", 18)], {"1.3.1~rc1-1.pgdg12+1"})
        self.assertEqual(packages[("rpm", "pg_h3", "el9.x86_64", 18)], {"4.2.3-1PGDG.rhel9"})
        self.assertEqual(stamps, {"deb.pgdg": "before/after"})

    def test_security_manifest_metadata_does_not_claim_gap_only(self):
        reasons = manifest.pgdg_build_reasons(
            "deb", mobility_rows(), "d12.x86_64", 18,
            {("deb", "mobilitydb", "d12.x86_64", 18): {"1.3.0-1.pgdg12+1"}})
        self.assertEqual(manifest.supplement_origin("PGDG", reasons), "catalog_pgdg_security_supplement")
        policy = manifest.release_policy("deb", "mobilitydb", "d12.x86_64", mobility_rows(), True, reasons)
        self.assertIn("CVE-2026-102639", policy)
        self.assertIn("PGDG stable >=1.3.1", policy)
        self.assertNotIn("PGDG-gap only", policy)
        self.assertEqual(manifest.supplement_origin("PGDG", ("missing_package:pg_h3",)), "catalog_pgdg_gap")
        self.assertEqual(manifest.supplement_origin("PIGSTY", ()), "catalog_pigsty")


if __name__ == "__main__":
    unittest.main()
