#!/bin/bash

: ${USE_SUDO:="true"}
: ${BACKEND_BINARY_NAME:="iptables-viz-backend"}
: ${BACKEND_INSTALL_DIR:="/usr/local/bin"}
: ${BACKEND_SERVICE_NAME:="iptables-viz-backend"}
: ${FRONTEND_BINARY_NAME:="iptables-viz-frontend"}
: ${FRONTEND_INSTALL_DIR:="/etc/iptables-viz"}
: ${FRONTEND_SERVICE_NAME:="iptables-viz-frontend"}
: ${SERVICE_DIR:="/etc/systemd/system"}

HAS_CURL="$(type "curl" &> /dev/null && echo true || echo false)"
HAS_WGET="$(type "wget" &> /dev/null && echo true || echo false)"
HAS_SYSTEMD="$(type "systemctl" &> /dev/null && echo true || echo false)"
HAS_JC="$(type "jc" &> /dev/null && echo true || echo false)"
HAS_SERVE="$(type "serve" &> /dev/null && echo true || echo false)"

# discovers the operating system for this system
init_os() {
    OS=$(uname | tr '[:upper:]' '[:lower:]')
}

# discovers the architecture for this system
init_arch() {
    ARCH=$(uname -m)
    case $ARCH in
    armv5*) ARCH="arm" ;;
    armv6*) ARCH="arm" ;;
    armv7*) ARCH="arm" ;;
    aarch64) ARCH="arm64" ;;
    x86) ARCH="386" ;;
    x86_64) ARCH="amd64" ;;
    i686) ARCH="386" ;;
    i386) ARCH="386" ;;
    esac
}

# checks if the main systemd service is running or not
check_if_main_systemd_is_running() {
  if [[ "$(systemctl is-active iptables-viz)" = "active" ]]; then
    echo "iptables-viz is already running"
    return 0
  else
    return 1
  fi
}

# checks if the main systemd service is enabled or not
check_if_main_systemd_is_enabled() {
  if [[ "$(systemctl is-enabled iptables-viz)" = "enabled" ]]; then
    echo "iptables-viz is already enabled"
    return 0
  else
    return 1
  fi
}

# checks if the backend systemd service is enabled or not
check_if_backend_systemd_is_enabled() {
  if [[ "$(systemctl is-enabled ${BACKEND_SERVICE_NAME})" = "enabled" ]]; then
    echo "$BACKEND_SERVICE_NAME is already enabled"
    return 0
  else
    return 1
  fi
}

# checks if the frontend systemd service is enabled or not
check_if_frontend_systemd_is_enabled() {
  if [[ "$(systemctl is-enabled ${FRONTEND_SERVICE_NAME})" = "enabled" ]]; then
    echo "$FRONTEND_SERVICE_NAME is already enabled"
    return 0
  else
    return 1
  fi
}

# verifySupported checks that the os/arch combination is supported for
# binary builds, as well whether or not necessary tools are present
verify_supported() {
  local supported="linux-386\nlinux-amd64\nlinux-arm\nlinux-arm64"
  if ! echo "${supported}" | grep -q "${OS}-${ARCH}"; then
    echo "No prebuilt binary for ${OS}-${ARCH}."
    exit 1
  fi

  if [ "${HAS_SYSTEMD}" != "true" ]; then
    echo "Systemd is required for installation"
    exit 1
  fi

  if [ "${HAS_CURL}" != "true" ] && [ "${HAS_WGET}" != "true" ]; then
    echo "Either curl or wget is required"
    exit 1
  fi

  if [ "${HAS_JC}" != "true" ]; then
    echo "jc is required for installation"
    exit 1
  fi

  if [ "${HAS_SERVE}" != "true" ]; then
    echo "serve is required for installation"
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

# create the frontend directory which will be referred by the systemd file
create_frontend_dir() {
  run_as_root mkdir "$FRONTEND_INSTALL_DIR"
  echo "$FRONTEND_INSTALL_DIR directory created"
}

# checks if the main systemd unit file already exists or not
is_main_unit_file_exists() {
  if [[ -f "${SERVICE_DIR}/iptables-viz.service" ]]; then
    echo "$SERVICE_DIR/iptables-viz.service already exists."
    return 0
  else
    return 1
  fi
}

# checks if the frontend systemd unit file already exists or not
is_frontend_unit_file_exists() {
  if [[ -f "${SERVICE_DIR}/${FRONTEND_SERVICE_NAME}.service" ]]; then
    echo "$SERVICE_DIR/$FRONTEND_SERVICE_NAME.service already exists."
    return 0
  else
    return 1
  fi
}

# checks if the backend systemd unit file already exists or not
is_backend_unit_file_exists() {
  if [[ -f "${SERVICE_DIR}/${BACKEND_SERVICE_NAME}.service" ]]; then
    echo "$SERVICE_DIR/$BACKEND_SERVICE_NAME.service already exists."
    return 0
  else
    return 1
  fi
}

# checks if the frontend install directory already exists or not
check_if_frontend_install_dir_exists() {
  if [[ -d "${FRONTEND_INSTALL_DIR}" ]]; then
    echo "$FRONTEND_INSTALL_DIR already exists."
    return 0
  else
    return 1
  fi
}

# checks if the backend binary file already exists or not
is_backend_exists() {
  if [[ -f "${BACKEND_INSTALL_DIR}/${BACKEND_BINARY_NAME}" ]]; then
    echo "$BACKEND_BINARY_NAME already exists."
    return 0
  else
    return 1
  fi
}

# checks if the frontend binary file already exists or not
is_frontend_exists() {
  if [[ -f "${FRONTEND_INSTALL_DIR}/${FRONTEND_BINARY_NAME}" ]]; then
    echo "$FRONTEND_BINARY_NAME already exists."
    return 0
  else
    return 1
  fi
}

# downloads the backend binary and copies it to the backend installation directory
download_backend() {
    echo "Downloading backend"
    BACKEND_BINARY_NAME="iptables-viz-backend"
    TAG="master" 
    BACKEND_DIST="iptables-viz-backend-$ARCH-$TAG.tar.gz"
    BIN_DOWNLOAD_URL="https://firebasestorage.googleapis.com/v0/b/neelanjan-manna.appspot.com/o/demo%2F$BACKEND_DIST?alt=media&token=b7563a58-2640-4206-bda5-dbcbee0e4af1"
    BACKEND_TMP_ROOT="$(mktemp -dt iptables-viz-backend-installer-XXXXXX)"
    BACKEND_TMP_FILE="$BACKEND_TMP_ROOT/$BACKEND_DIST"
    echo "Downloading $BIN_DOWNLOAD_URL"
    if [ "${HAS_CURL}" == "true" ]; then
        curl -SsL "$BIN_DOWNLOAD_URL" -o "$BACKEND_TMP_FILE"
    elif [ "${HAS_WGET}" == "true" ]; then
        wget -q -O "$BACKEND_TMP_FILE" "$BIN_DOWNLOAD_URL"
    fi
    BACKEND_TMP="$BACKEND_TMP_ROOT/$BACKEND_BINARY_NAME"
    mkdir -p "$BACKEND_TMP"
    tar xvf "$BACKEND_TMP_FILE" -C "$BACKEND_TMP"
    BACKEND_TMP_BIN="$BACKEND_TMP/$BACKEND_BINARY_NAME"
    echo "Preparing to install $BACKEND_BINARY_NAME into ${BACKEND_AGENT_INSTALL_DIR}"
    run_as_root cp "$BACKEND_TMP_BIN" "$BACKEND_INSTALL_DIR/$BACKEND_BINARY_NAME"
    echo "$BACKEND_BINARY_NAME installed into $BACKEND_INSTALL_DIR/$BACKEND_BINARY_NAME"
}

# downloads the frontend binary and copies it to the frontend installation directory
download_frontend() {
    echo "Downloading frontend"
    FRONTEND_BINARY_NAME="iptables-viz-frontend"
    TAG="master"
    FRONTEND_DIST="iptables-viz-frontend-$ARCH-$TAG.tar.gz"
    BIN_DOWNLOAD_URL="https://firebasestorage.googleapis.com/v0/b/neelanjan-manna.appspot.com/o/demo%2F$FRONTEND_DIST?alt=media&token=85a3439f-1915-469a-80b0-2f6e3b600914"
    FRONTEND_TMP_ROOT="$(mktemp -dt iptables-viz-frontend-installer-XXXXXX)"
    FRONTEND_TMP_FILE="$FRONTEND_TMP_ROOT/$FRONTEND_DIST"
    echo "Downloading $BIN_DOWNLOAD_URL"
    if [ "${HAS_CURL}" == "true" ]; then
        curl -SsL "$BIN_DOWNLOAD_URL" -o "$FRONTEND_TMP_FILE"
    elif [ "${HAS_WGET}" == "true" ]; then
        wget -q -O "$FRONTEND_TMP_FILE" "$BIN_DOWNLOAD_URL"
    fi
    FRONTEND_TMP="$FRONTEND_TMP_ROOT/$FRONTEND_BINARY_NAME"
    mkdir -p "$FRONTEND_TMP"
    tar xvf "$FRONTEND_TMP_FILE" -C "$FRONTEND_TMP"
    FRONTEND_TMP_BIN="$FRONTEND_TMP/$FRONTEND_BINARY_NAME"
    echo "Preparing to install $FRONTEND_BINARY_NAME into ${FRONTEND_AGENT_INSTALL_DIR}"
    run_as_root cp -r "$FRONTEND_TMP_BIN" "$FRONTEND_INSTALL_DIR/$FRONTEND_BINARY_NAME"
    echo "$FRONTEND_BINARY_NAME installed into $FRONTEND_INSTALL_DIR/$FRONTEND_BINARY_NAME"
}

# creates the main systemd file
create_systemd_file_main() {
    echo "creating main systemd file"
    run_as_root printf "[Unit]
Description=Oneshot service for iptables-viz

[Service]
# The dummy program will exit
Type=oneshot
# Execute a dummy program
ExecStart=/bin/true
# This service shall be considered active after start
RemainAfterExit=yes

[Install]
# Components of this application should be started at boot time
WantedBy=multi-user.target
" > "$SERVICE_DIR/iptables-viz.service"
}

# creates the backend systemd file
create_systemd_file_backend() {
    echo "creating backend systemd file"
    run_as_root printf "[Unit]
Description=Iptables Visualization Go Service
ConditionPathExists=$BACKEND_INSTALL_DIR/$BACKEND_SERVICE_NAME
PartOf=iptables-viz.service
After=iptables-viz.service

[Service]
Type=simple
User=root
Group=root
ExecStart=$BACKEND_INSTALL_DIR/$BACKEND_SERVICE_NAME --platform linux
Restart=on-failure
RestartSec=10
StandardOutput=syslog
StandardError=syslog
SyslogIdentifier=$BACKEND_SERVICE_NAME

[Install]
WantedBy=iptables-viz.service
" > "$SERVICE_DIR/$BACKEND_SERVICE_NAME.service"
}

# creates the frontend systemd file
create_systemd_file_frontend() {
    echo "creating frontend systemd file"
    run_as_root printf "[Unit]
Description=Iptables Visualization Frontend App
ConditionPathExists=$FRONTEND_INSTALL_DIR/$FRONTEND_SERVICE_NAME
PartOf=iptables-viz.service
After=$BACKEND_SERVICE_NAME.service
After=iptables-viz.service

[Service]
Type=simple
User=root
Group=root
ExecStart=/usr/local/bin/serve -s $FRONTEND_INSTALL_DIR/$FRONTEND_SERVICE_NAME
Restart=on-failure
RestartSec=10
StandardOutput=syslog
StandardError=syslog
SyslogIdentifier=$FRONTEND_SERVICE_NAME

[Install]
WantedBy=iptables-viz.service
" > "$SERVICE_DIR/$FRONTEND_SERVICE_NAME.service"
}

# starts the systemd service
run_systemd_file() {
    run_as_root systemctl daemon-reload
    if ! check_if_main_systemd_is_enabled || ! check_if_backend_systemd_is_enabled || ! check_if_frontend_systemd_is_enabled; then
      run_as_root systemctl enable iptables-viz "$BACKEND_SERVICE_NAME" "$FRONTEND_SERVICE_NAME"
    else
      echo "iptables viz service is already enabled"
    fi
    if ! check_if_main_systemd_is_running; then
      run_as_root systemctl start iptables-viz
    else
      echo "iptables viz service is already active"
    fi
}

# cleans up the tmp directory where the binaries were downloaded temporarily
cleanup() {
  if [[ -d "${BACKEND_TMP_ROOT:-}" ]]; then
    rm -rf "$BACKEND_TMP_ROOT"
  fi
  if [[ -d "${FRONTEND_AGENT_TMP_ROOT:-}" ]]; then
    rm -rf "$FRONTEND_TMP_ROOT"
  fi
}

init_arch
init_os
verify_supported
if ! is_backend_exists & ! is_frontend_exists; then
  download_backend
  if ! check_if_frontend_install_dir_exists; then
    create_frontend_dir
  fi
  download_frontend
  if ! is_main_unit_file_exists; then
    create_systemd_file_main
  fi
  if ! is_backend_unit_file_exists; then
    create_systemd_file_backend
  fi
  if ! is_frontend_unit_file_exists; then
    create_systemd_file_frontend
  fi
  run_systemd_file
fi
cleanup
