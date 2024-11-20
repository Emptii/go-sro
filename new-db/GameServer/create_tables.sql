#
# WORLD DATA TABLES
#

source /docker-entrypoint-initdb.d/GameServer/WorldData/refcharacterdata.sql
source /docker-entrypoint-initdb.d/GameServer/WorldData/refcharacter_quality.sql
source /docker-entrypoint-initdb.d/GameServer/WorldData/refguildpromotepolicy.sql
source /docker-entrypoint-initdb.d/GameServer/WorldData/refitemdata.sql
source /docker-entrypoint-initdb.d/GameServer/WorldData/refleveldata.sql
source /docker-entrypoint-initdb.d/GameServer/WorldData/reflevelgold.sql
source /docker-entrypoint-initdb.d/GameServer/WorldData/refmagicoption.sql
source /docker-entrypoint-initdb.d/GameServer/WorldData/refmagicoption_values.sql
source /docker-entrypoint-initdb.d/GameServer/WorldData/refmappingshopgroup.sql
source /docker-entrypoint-initdb.d/GameServer/WorldData/refmappingshopwithtab.sql
source /docker-entrypoint-initdb.d/GameServer/WorldData/refnpcpos.sql
source /docker-entrypoint-initdb.d/GameServer/WorldData/refoptionalteleport.sql
source /docker-entrypoint-initdb.d/GameServer/WorldData/refpackageitem.sql
source /docker-entrypoint-initdb.d/GameServer/WorldData/refpricepolicyofitem.sql
source /docker-entrypoint-initdb.d/GameServer/WorldData/refquantityofitem.sql
source /docker-entrypoint-initdb.d/GameServer/WorldData/refregioncode.sql
source /docker-entrypoint-initdb.d/GameServer/WorldData/refshopgoods.sql
source /docker-entrypoint-initdb.d/GameServer/WorldData/refshopgroup.sql
source /docker-entrypoint-initdb.d/GameServer/WorldData/refshoptab.sql
source /docker-entrypoint-initdb.d/GameServer/WorldData/refsiegefortress.sql
source /docker-entrypoint-initdb.d/GameServer/WorldData/refskilldata.sql
source /docker-entrypoint-initdb.d/GameServer/WorldData/refcharactertask.sql
source /docker-entrypoint-initdb.d/GameServer/WorldData/refteleportbuilding.sql
source /docker-entrypoint-initdb.d/GameServer/WorldData/refteleportdata.sql
source /docker-entrypoint-initdb.d/GameServer/WorldData/refteleportlink.sql

#
# GAME DATA TABLES
#  

source /docker-entrypoint-initdb.d/GameServer/ContentData/character.sql
source /docker-entrypoint-initdb.d/GameServer/ContentData/character_block.sql
source /docker-entrypoint-initdb.d/GameServer/ContentData/character_completed_quest.sql
source /docker-entrypoint-initdb.d/GameServer/ContentData/character_hotkey.sql
source /docker-entrypoint-initdb.d/GameServer/ContentData/character_item.sql
source /docker-entrypoint-initdb.d/GameServer/ContentData/character_item_magicoption.sql
source /docker-entrypoint-initdb.d/GameServer/ContentData/character_job.sql
source /docker-entrypoint-initdb.d/GameServer/ContentData/character_mastery.sql
source /docker-entrypoint-initdb.d/GameServer/ContentData/character_pending_quest.sql
source /docker-entrypoint-initdb.d/GameServer/ContentData/character_skill.sql
source /docker-entrypoint-initdb.d/GameServer/ContentData/friends.sql
source /docker-entrypoint-initdb.d/GameServer/ContentData/guild.sql
source /docker-entrypoint-initdb.d/GameServer/ContentData/guild_alliance.sql
source /docker-entrypoint-initdb.d/GameServer/ContentData/guild_member.sql
source /docker-entrypoint-initdb.d/GameServer/ContentData/item_magicoption.sql

