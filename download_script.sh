#!/bin/bash

d=`date -d "+5 hours" +%d%m%Y` # Indian time zone - 30 minutes. 

echo "Downloading today's paper : $d"

go run get_todays_newspapers.go

echo "Converting Anandabajar to pdf"
convert $(ls -d anandabajar/$d/*) anandabajar/$d/$d.pdf

echo "Converting Telegraph to pdf"
convert $(ls -d telegraph/$d/*) telegraph/$d/$d.pdf

# echo "Deleting the Images"
# rm $(ls -d $d/*|grep -v pdf)

echo "Opening todays Anandabajar newspaper"
evince anandabajar/$d/$d.pdf

echo "Opening todays Telegraph newspaper"
evince telegraph/$d/$d.pdf
