#!/bin/bash

: ${USE_SUDO:="true"}
: ${BACKEND_BINARY_NAME:="iptables-viz-backend"}
: ${BACKEND_INSTALL_DIR:="/usr/local/bin"}
: ${BACKEND_SERVICE_NAME:="iptables-viz-backend"}
: ${FRONTEND_BINARY_NAME:="iptables-viz-frontend"}
: ${FRONTEND_INSTALL_DIR:="/etc/iptables-viz"}
: ${FRONTEND_SERVICE_NAME:="iptables-viz-frontend"}
: ${SERVICE_DIR:="/etc/systemd/system"}

# runs the given command as root (detects if we are root already)
run_as_root() {
    if [ $EUID -ne 0 ] && [ "$USE_SUDO" == "true" ]; then
        sudo "${@}"
    else
        "${@}"
    fi
}

# checks if the main systemd service is running or not
check_if_main_systemd_is_running() {
  if [[ "$(systemctl is-active iptables-viz)" = "active" ]]; then
    return 0
  else
    return 1
  fi
}

# checks if the main systemd service is enabled or not
check_if_main_systemd_is_enabled() {
  if [[ "$(systemctl is-enabled iptables-viz)" = "enabled" ]]; then
    return 0
  else
    return 1
  fi
}

# checks if the backend systemd service is enabled or not
check_if_backend_systemd_is_enabled() {
  if [[ "$(systemctl is-enabled ${BACKEND_SERVICE_NAME}.service)" = "enabled" ]]; then
    return 0
  else
    return 1
  fi
}

# checks if the frontend systemd service is enabled or not
check_if_frontend_systemd_is_enabled() {
  if [[ "$(systemctl is-enabled ${FRONTEND_SERVICE_NAME}.service)" = "enabled" ]]; then
    return 0
  else
    return 1
  fi
}

# checks if the main systemd unit file already exists or not
is_main_unit_file_exists() {
  if [[ -f "${SERVICE_DIR}/iptables-viz.service" ]]; then
    return 0
  else
    return 1
  fi
}

# checks if the frontend systemd unit file already exists or not
is_frontend_unit_file_exists() {
  if [[ -f "${SERVICE_DIR}/${FRONTEND_SERVICE_NAME}.service" ]]; then
    return 0
  else
    return 1
  fi
}

# checks if the backend systemd unit file already exists or not
is_backend_unit_file_exists() {
  if [[ -f "${SERVICE_DIR}/${BACKEND_SERVICE_NAME}.service" ]]; then
    return 0
  else
    return 1
  fi
}

# checks if the frontend install directory already exists or not
check_if_frontend_install_dir_exists() {
  if [[ -d "${FRONTEND_INSTALL_DIR}" ]]; then
    return 0
  else
    return 1
  fi
}

# checks if the backend binary file already exists or not
is_backend_exists() {
  if [[ -f "${BACKEND_INSTALL_DIR}/${BACKEND_BINARY_NAME}" ]]; then
    return 0
  else
    return 1
  fi
}


if check_if_main_systemd_is_running; then 
  run_as_root systemctl stop iptables-viz.service
fi
if check_if_main_systemd_is_enabled; then
  run_as_root systemctl disable iptables-viz
fi
if check_if_backend_systemd_is_enabled; then
  run_as_root systemctl disable iptables-viz-backend
fi
if check_if_frontend_systemd_is_enabled; then 
  run_as_root systemctl disable iptables-viz-frontend
fi
if check_if_frontend_install_dir_exists; then
  run_as_root rm -rf /etc/iptables-viz
  echo "Removed iptables frontend binary"
fi
if is_backend_unit_file_exists; then
  run_as_root rm /etc/systemd/system/iptables-viz-backend.service
fi
if is_frontend_unit_file_exists; then
  run_as_root rm /etc/systemd/system/iptables-viz-frontend.service
fi
if is_main_unit_file_exists; then
  run_as_root rm /etc/systemd/system/iptables-viz.service
fi
if is_backend_exists; then
  run_as_root rm /usr/local/bin/iptables-viz-backend
  echo "Removed iptables backend binary"
fi