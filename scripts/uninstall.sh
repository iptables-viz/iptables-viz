#!/bin/bash

: "${USE_SUDO:="true"}"
: "${BACKEND_BINARY_NAME:="iptables-viz-backend"}"
: "${BACKEND_INSTALL_DIR:="/usr/local/bin"}"
: "${BACKEND_SERVICE_NAME:="iptables-viz-backend"}"
: "${FRONTEND_BINARY_NAME:="iptables-viz-frontend"}"
: "${FRONTEND_INSTALL_DIR:="/etc/iptables-viz"}"
: "${FRONTEND_SERVICE_NAME:="iptables-viz-frontend"}"
: "${SERVICE_DIR:="/etc/systemd/system"}"

HAS_SYSTEMD="$(type "systemctl" &> /dev/null && echo true || echo false)"

# verifySupported checks that the os/arch combination is supported for
# binary builds, as well whether or not necessary tools are present
verify_supported() {
  if [ "${HAS_SYSTEMD}" != "true" ]; then
    echo "Systemd is required for uninstallation"
    exit 1
  fi
}

# runs the given command as root (detects if we are root already)
run_as_root() {
    if [ $EUID -ne 0 ] && [ "$USE_SUDO" == "true" ]; then
        sudo "${@}"
    else
        "${@}"
    fi
}


verify_supported
if [[ "$(systemctl is-active iptables-viz)" = "active" ]]; then
  run_as_root systemctl stop iptables-viz.service
fi
if [[ "$(systemctl is-enabled iptables-viz)" = "enabled" ]]; then
  run_as_root systemctl disable iptables-viz
fi
if [[ "$(systemctl is-enabled "${BACKEND_SERVICE_NAME}".service)" = "enabled" ]]; then
  run_as_root systemctl disable iptables-viz-backend
fi
if [[ "$(systemctl is-enabled "${FRONTEND_SERVICE_NAME}".service)" = "enabled" ]]; then
  run_as_root systemctl disable iptables-viz-frontend
fi
if [[ -d "${FRONTEND_INSTALL_DIR}" ]]; then
  run_as_root rm -rf /etc/iptables-viz
  echo "Removed iptables frontend binary"
fi
if [[ -f "${SERVICE_DIR}/${BACKEND_SERVICE_NAME}.service" ]]; then
  run_as_root rm /etc/systemd/system/iptables-viz-backend.service
fi
if [[ -f "${SERVICE_DIR}/${FRONTEND_SERVICE_NAME}.service" ]]; then
  run_as_root rm /etc/systemd/system/iptables-viz-frontend.service
fi
if [[ -f "${SERVICE_DIR}/iptables-viz.service" ]]; then
  run_as_root rm /etc/systemd/system/iptables-viz.service
fi
if [[ -f "${BACKEND_INSTALL_DIR}/${BACKEND_BINARY_NAME}" ]]; then
  run_as_root rm /usr/local/bin/iptables-viz-backend
  echo "Removed iptables backend binary"
fi
