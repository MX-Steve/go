#!/bin/sh

source /worker.env

cp /bminer-current/libcratercl.so ./usr/lib/x86_64-linux-gnu/
MINER=/bminer-current/bminer
$MINER -managed | logger -p local1.info &