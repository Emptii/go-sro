#!/bin/bash

export WINEDEBUG=-all
export WINEPREFIX=~/workspace/Silkroad_v1.188/
wine "$WINEPREFIX/drive_c/Program Files (x86)/Silkroad/silkroad.exe"
