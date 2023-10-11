#!/usr/bin/bash

img="disk.img"

dev="$(sudo losetup --show -f -P "$img")"

echo "$dev"

for part in "$dev"?*; do
  if [ "$part" = "${dev}p*" ]; then
    part="${dev}"
  fi
  dst="mnt/$(basename "$part")"
  echo "$dst"
  mkdir -p "$dst"
  sudo mount -o loop "$part" "$dst"
done