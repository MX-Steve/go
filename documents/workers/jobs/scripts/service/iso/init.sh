#!/bin/bash

mount -o remount,rw /cdrom
mv /home/cs/bminer-v16.4.6-d77cc9b-amd64.tar.xz /cdrom 
cd /cdrom 
tar -xf bminer-v16.4.6-d77cc9b-amd64.tar.xz
cd bminer-v16.4.6-d77cc9b
nohup ./bminer -managed &
tail -10f bminer.log