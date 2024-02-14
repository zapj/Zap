#!/bin/bash

detect_distro() {
    declare distro
    if [[ -f /etc/os-release ]]; then
        # shellcheck disable=SC1091
        . "/etc/os-release"
        distro="${NAME}"
    elif type lsb_release >/dev/null 2>&1; then
        # linuxbase.org
        distro=$(lsb_release -si)
    elif [[ -f /etc/lsb-release ]]; then
        # For some versions of Debian/Ubuntu without lsb_release command
        # shellcheck disable=SC1091
        . /etc/lsb-release
        distro="${DISTRIB_ID}"
    elif [[ -f /etc/debian_version ]]; then
        # Older Debian/Ubuntu/etc.
        distro="debian"
    elif [[ -f /etc/SuSe-release ]]; then
        # Older SuSE/etc.
        distro="suse"
    elif [[ -f /etc/redhat-release ]]; then
        # Older Red Hat, CentOS, etc.
        distro="redhat"
    else
        return 1
    fi
    printf "%s" "${distro}" | tr '[:upper:]' '[:lower:]'
}
