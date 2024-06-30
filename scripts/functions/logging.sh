# Write generic log
#
# @param $1 The log level
# @param $2 The log message
#
function log() {
  # Log level
  local level=$1

  # Log message
  local message=$2

  # Timestamp
  local timestamp=$(date +"%Y-%m-%d %H:%M:%S")

  # Print log message
  echo "[$timestamp] [$level] $message"
}

# Print error message and exit
#
# @param $1 The error message
#
function error() {
  local red=$(tput setaf 1)
  local reset=$(tput sgr0)

  log "${red}ERROR${reset}" "$1"
  exit 1
}

# Print info message
#
# @param $1 The info message
#
function info() {
  local blue=$(tput setaf 4)
  local reset=$(tput sgr0)

  log "${blue}INFO${reset}" "$1"
}
