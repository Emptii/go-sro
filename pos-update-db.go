UpdateCharacterPosition string = "UPDATE `SRO_SHARD`.`CHAR` SET POS_X = ?, POS_Y = ?, POS_Z = ?, FK_REGION = ? WHERE REF_OBJ_ID = ?"

// Update character position in db
conn := db.OpenConnAccount()
defer conn.Close()

queryHandle, err := conn.Prepare(UpdateCharacterPosition)
db.CheckError(err)
defer queryHandle.Close() // Close the prepared statement

logrus.Infof("updating character position in db...")
_, err = queryHandle.Exec(pPos.Offset.X, pPos.Offset.Y, pPos.Offset.Z, pPos.Region.ID, player.RefObjectID)
logrus.Infof("done updating character position in db")
db.CheckError(err)
logrus.Infof(player.CharName)
