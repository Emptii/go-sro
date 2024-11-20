#!/bin/bash

# first extract the Media.pk2 directory into this directory here (so {pwd}/Media)

rm -r Media/server_dep/silkroad/conv/
mkdir Media/server_dep/silkroad/conv/

cat ./Media/server_dep/silkroad/textdata/characterdata_*.txt > ./Media/server_dep/silkroad/conv/characterdata_full_original.txt
iconv -f utf-16le -t utf-8 -c ./Media/server_dep/silkroad/conv/characterdata_full_original.txt -o ./Media/server_dep/silkroad/conv/characterdata_full.txt

cat ./Media/server_dep/silkroad/textdata/itemdata_*.txt > ./Media/server_dep/silkroad/conv/itemdata_full_original.txt
iconv -f utf-16le -t utf-8 -c ./Media/server_dep/silkroad/conv/itemdata_full_original.txt -o ./Media/server_dep/silkroad/conv/itemdata_full.txt

cat ./Media/server_dep/silkroad/textdata/skilldata_*.txt > ./Media/server_dep/silkroad/conv/skilldata_full_original.txt
iconv -f utf-16le -t utf-8 -c ./Media/server_dep/silkroad/conv/skilldata_full_original.txt -o ./Media/server_dep/silkroad/conv/skilldata_full.txt


cp ./Media/server_dep/silkroad/textdata/refpackageitem.txt ./Media/server_dep/silkroad/conv/refpackageitem_original.txt  
iconv -f utf-16le -t utf-8 -c ./Media/server_dep/silkroad/conv/refpackageitem_original.txt -o ./Media/server_dep/silkroad/conv/refpackageitem.txt

cp ./Media/server_dep/silkroad/textdata/refpricepolicyofitem.txt ./Media/server_dep/silkroad/conv/refpricepolicyofitem_original.txt
iconv -f utf-16le -t utf-8 -c ./Media/server_dep/silkroad/conv/refpricepolicyofitem_original.txt -o ./Media/server_dep/silkroad/conv/refpricepolicyofitem.txt

cp ./Media/server_dep/silkroad/textdata/refshopgoods.txt ./Media/server_dep/silkroad/conv/refshopgoods_original.txt
iconv -f utf-16le -t utf-8 -c ./Media/server_dep/silkroad/conv/refshopgoods_original.txt -o ./Media/server_dep/silkroad/conv/refshopgoods.txt

cp ./Media/server_dep/silkroad/textdata/refshopgroup.txt ./Media/server_dep/silkroad/conv/refshopgroup_original.txt
iconv -f utf-16le -t utf-8 -c ./Media/server_dep/silkroad/conv/refshopgroup_original.txt -o ./Media/server_dep/silkroad/conv/refshopgroup.txt

cp ./Media/server_dep/silkroad/textdata/refshoptab.txt ./Media/server_dep/silkroad/conv/refshoptab_original.txt
iconv -f utf-16le -t utf-8 -c ./Media/server_dep/silkroad/conv/refshoptab_original.txt -o ./Media/server_dep/silkroad/conv/refshoptab.txt

# Removing BOM
sed -i '1s/^\xEF\xBB\xBF//' ./Media/server_dep/silkroad/conv/*.txt
