# Function to determine if a path exists.
#
# @param $1 The path to check
#
# @return 0 if the path exists, 1 otherwise
#
function path_exists() {
  if [ -d "$1" ]; then
    return 0
  else
    return 1
  fi
}
