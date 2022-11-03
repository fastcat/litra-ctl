#!/bin/sh

set -xeu
udevadm control --reload-rules
udevadm trigger
