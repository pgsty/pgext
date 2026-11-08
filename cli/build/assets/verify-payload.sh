#!/usr/bin/env bash
# Verify all binary and debug payloads from one isolated PGEXT build.
set -euo pipefail
family=$1
output=$2
no_debug=$3
scratch=$(mktemp -d)
trap 'rm -rf "$scratch"' EXIT
die() { echo "$*" >&2; exit 1; }
declare -A roots=()
while IFS= read -r -d '' package; do
    if [[ $family == rpm ]]; then
        name=$(rpm -qp --qf '%{NAME}' "$package")
    else
        name=$(dpkg-deb -f "$package" Package)
    fi
    [[ $name =~ ^[A-Za-z0-9][A-Za-z0-9_.+~-]*$ ]] || die 'invalid package name'
    root="$scratch/$name"
    mkdir "$root"
    if [[ $family == rpm ]]; then
        (cd "$root" && rpm2cpio "$package" | cpio -idm --quiet)
    else
        dpkg-deb -x "$package" "$root"
    fi
    roots[$name]=$root
done < <(find "$output" -type f \( -name '*.rpm' -o -name '*.deb' -o -name '*.ddeb' \) -print0)
for name in "${!roots[@]}"; do
    root=${roots[$name]}
    [[ -n $(find "$root" -type f -size +0c -print -quit) ]] || die "empty payload: $name"
    case "$name" in *-debuginfo|*-debugsource|*-dbgsym) continue;; esac
    native=0
    while IFS= read -r -d '' payload; do
        file -b "$payload" | grep -q '^ELF ' || continue
        # Object-only devel packages have no linked ELF build ID to pair.
        readelf -h "$payload" | grep -Eq 'Type:.*(DYN|EXEC)' || continue
        native=$((native + 1))
        [[ $no_debug == false ]] || continue
        id=$(readelf -n "$payload" | sed -n 's/.*Build ID:[[:space:]]*//p')
        [[ $id =~ ^([a-f0-9]{2})+$ ]] || die "missing Build ID: $name/$payload"
        if [[ $family == rpm ]]; then
            debug_name="$name-debuginfo"
        else
            debug_name="$name-dbgsym"
        fi
        # RPM may put companion symbols in the source package's debuginfo.
        candidates=()
        if [[ -n ${roots[$debug_name]+x} ]]; then
            candidates+=("${roots[$debug_name]}")
        elif [[ $family == rpm ]]; then
            for other in "${!roots[@]}"; do
                [[ $other == *-debuginfo ]] && candidates+=("${roots[$other]}")
            done
        fi
        [[ ${#candidates[@]} -gt 0 ]] || die "missing debug package for $name"
        matched=false
        while IFS= read -r -d '' debug; do
            debug_id=$(readelf -n "$debug" 2>/dev/null | sed -n 's/.*Build ID:[[:space:]]*//p') || continue
            [[ $debug_id == "$id" ]] || continue
            sections=$(readelf -SW "$debug")
            [[ $sections == *'.debug_info'* ]] || die "debug package has no DWARF: $debug"
            matched=true
            break
        done < <(find "${candidates[@]}" -type f -name '*.debug' -print0)
        [[ $matched == true ]] || die "no matching DWARF for $name Build ID $id"
    done < <(find "$root" -type f -size +0c -print0)
    if [[ $family == rpm && $native -gt 0 && $no_debug == false ]]; then
        found=false
        for other in "${!roots[@]}"; do
            if [[ $other == *-debugsource && -n $(find "${roots[$other]}" -type f -size +0c -print -quit) ]]; then found=true; fi
        done
        [[ $found == true ]] || die "missing debugsource for $name"
    fi
    echo "verified payload=$name native_elf=$native"
done
