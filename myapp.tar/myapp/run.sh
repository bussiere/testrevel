#!/bin/sh
SCRIPTPATH=`dirname "$0"`
chmod u+x "$SCRIPTPATH/myapp.exe"
"$SCRIPTPATH/myapp.exe" -importPath myapp -srcPath "$SCRIPTPATH/src" -runMode prod
