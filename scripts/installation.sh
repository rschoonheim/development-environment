# Include functions used for installation.
#
source scripts/functions/logging.sh
source scripts/functions/filesystem.sh

# The installation must be done as root
#
if [ "$EUID" -ne 0 ]; then
  error "Please run as root"
  exit
fi


# Setup the filesystem for the development environment
#
paths=(
  "/etc/rschoonheim"
  "/etc/rschoonheim/development"
  "/etc/rschoonheim/development/binaries"
)

for path in "${paths[@]}"; do
  if ! path_exists "$path"; then
    mkdir "$path"
    info "Created $path"
  fi
done