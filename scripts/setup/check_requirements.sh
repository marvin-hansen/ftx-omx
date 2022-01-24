# Copyright (c) 2021-2022. Marvin Hansen | marvin.hansen@gmail.com

# bin/bash
set -o errexit
set -o nounset
set -o pipefail

# "---------------------------------------------------------"
# "-                                                       -"
# "-  Checks for all dependencies required by make         -"
# "-                                                       -"
# "---------------------------------------------------------"

echo ""
echo "==============================="
echo "Check for all OMX dependencies!"
echo "==============================="
echo ""

command bash --version >/dev/null 2>&1 || {
  # command sudo apt-get -qqq -y install curl
   echo "Please install bash"
   exit
}
echo "* Bash installed"

command make --version >/dev/null 2>&1 || {
   echo "Please install curl"
   echo "Debian / Ubuntu Linux: "
   echo "sudo apt-get -qqq -y install make"
   echo "Mac/homebrew: "
   echo "brew install make"
   exit
}
echo "* Make installed"

command clang --version >/dev/null 2>&1 || {
   echo "Please install clang"
   echo "Debian / Ubuntu Linux: "
   echo "sudo apt-get -qqq -y install clang"
   echo "Mac: "
   echo "xcode-select —install"
   exit
}
echo "* Clang installed"

command echo "export CC=clang" >>~/.bashrc
echo "* Bash configured clang as CC for Bazel"

command curl --version >/dev/null 2>&1 || {
   echo "Please install curl"
   echo "Debian / Ubuntu Linux: "
   echo "sudo apt-get -qqq -y install curl"
   exit
}
echo "* Curl installed"

command wget --version >/dev/null 2>&1 || {
     echo "Please install curl"
     echo "Debian / Ubuntu Linux: "
     echo "sudo apt-get -qqq -y install wget"
     echo "Mac/homebrew: "
     echo "brew install wget"
     exit
}
echo "* WGet installed"


command docker --version >/dev/null 2>&1 || {
  # command sudo apt-get -qqq -y install docker.io
  echo "Please install Docker"
  echo "https://docs.docker.com/engine/install/"
  exit
}
echo "* Docker installed"

command bazel version >/dev/null 2>&1 || {
      echo "Please install Bazelisk to manage Bazel installations"
      echo "Bazelisk Download:"
      echo "https://github.com/bazelbuild/bazelisk"
      echo "Install instructions for Mac & Linux"
      echo "https://stackoverflow.com/questions/65656165/how-do-you-install-bazel-using-bazelisk"
      exit
}
echo "* Bazel installed"

echo ""
echo "==============================="
echo "All OMX dependencies installed!"
echo "==============================="
echo ""
