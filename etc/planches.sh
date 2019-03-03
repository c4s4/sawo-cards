#!/bin/sh

set -e

montage -mode concatenate -tile 4x2 \
    ../img/perso.png \
    ../img/background.png \
    ../img/pnj.png \
    ../img/joker.png \
    ../img/extras.png \
    ../img/intrigue.png \
    ../img/lieu.png \
    ../img/template.png \
    planche1.png

montage -mode concatenate -tile 2x2 \
    ../img/perso.png \
    ../img/background.png \
    ../img/perso.png \
    ../img/background.png \
    planche2.png
