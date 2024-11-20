set global local_infile=true;

delete from `DB_Shard`.refpackageitem;
delete from `DB_Shard`.refpricepolicyofitem;
delete from `DB_Shard`.refshopgoods;
delete from `DB_Shard`.refshopgroup;
delete from `DB_Shard`.refshoptab;

load data infile '/data/Media/server_dep/silkroad/conv/refpackageitem.txt' into table `DB_Shard`.refpackageitem character set utf8 FIELDS TERMINATED BY '\t' LINES TERMINATED BY '\n';
load data infile '/data/Media/server_dep/silkroad/conv/refpricepolicyofitem.txt' into table `DB_Shard`.refpricepolicyofitem character set utf8 FIELDS TERMINATED BY '\t' LINES TERMINATED BY '\n';
load data infile '/data/Media/server_dep/silkroad/conv/refshopgoods.txt' into table `DB_Shard`.refshopgoods character set utf8 FIELDS TERMINATED BY '\t' LINES TERMINATED BY '\n';
load data infile '/data/Media/server_dep/silkroad/conv/refshopgroup.txt' into table `DB_Shard`.refshopgroup character set utf8 FIELDS TERMINATED BY '\t' LINES TERMINATED BY '\n';
load data infile '/data/Media/server_dep/silkroad/conv/refshoptab.txt' into table `DB_Shard`.refshoptab character set utf8 FIELDS TERMINATED BY '\t' LINES TERMINATED BY '\n';

/* now the data that was never inserted by the original project: */
load data infile '/data/Media/server_dep/silkroad/conv/itemdata.csv' into table `DB_Shard`.refitemdata character set utf8 FIELDS TERMINATED BY ',' LINES TERMINATED BY '\n' IGNORE 1 ROWS;
