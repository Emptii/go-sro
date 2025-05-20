#!/bin/bash

# Download Silkroad binaries

# Extract them to game_client_binaries/

#
[ -L /etc/go-sro/Data ] || ln -s "$(pwd game_client_binaries/Silkroad/Data.pk2)" /etc/go-sro/Data
