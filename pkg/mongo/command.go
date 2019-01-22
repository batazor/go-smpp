package mongo

// message ====================================================================

//func (c *session) addMessage(cmd Document) {
//	c.client.Database(MONGO_DATABASE_SMPP).Collection(*cmd.ID).InsertOne(nil, cmd.Item)
//	//
//	//if err != nil {
//	//	log.Error("Insert message error: ", err)
//	//	//cmd.Response <- Response{
//	//	//	error: err,
//	//	//	payload: nil,
//	//	//}
//	//	return
//	//}
//
//	//cmd.Response <- Response{
//	//	error: nil,
//	//	payload: res.InsertedID,
//	//}
//}

//func (c *session) updateMessage(cmd Document) {
//	filter := bson.NewDocument(bson.EC.String("_id", *cmd.ID))
//
//	c.client.Database(MONGO_DATABASE_SMPP).Collection(*cmd.ID).UpdateOne(nil, filter, cmd.Item)
//
//	//if err != nil {
//	//	log.Error("Update message error: ", err)
//	//	cmd.Callback("", err)
//	//	return
//	//}
//	//
//	//cmd.Callback(res.UpsertedID, nil)
//}

//func (c *session) deleteMessage(cmd Document) {
//	filter := bson.NewDocument(bson.EC.String("_id", *cmd.ID))
//
//	c.client.Database(MONGO_DATABASE_SMPP).Collection(*cmd.ID).DeleteOne(nil, filter)
//	//
//	//if err != nil {
//	//	log.Error("Delete message error: ", err)
//	//	cmd.Callback("", err)
//	//	return
//	//}
//	//
//	//cmd.Callback(res.DeletedCount, nil)
//}

//func (c *session) statusMessage(cmd Document) {
//	filter := bson.NewDocument(bson.EC.String("_id", *cmd.ID))
//
//	c.client.Database(MONGO_DATABASE_SMPP).Collection(*cmd.ID).Find(nil, filter)
//	//
//	//if err != nil {
//	//	log.Error("Get all message error: ", err)
//	//	cmd.Callback("", err)
//	//	return
//	//}
//
//	//cmd.Callback(res, nil)
//}

// worker =====================================================================

//func (c *session) addWorker(cmd Document) {
//	c.client.Database(MONGO_DATABASE_WORKER).Collection(*cmd.ID).InsertOne(nil, cmd.Item)
//	//
//	//if err != nil {
//	//	log.Error("Insert mongodb error: ", err)
//	//	cmd.Callback("", err)
//	//	return
//	//}
//	//
//	//cmd.Callback(res.InsertedID, nil)
//}

//func (c *session) updateWorker(cmd Document) {
//	filter := bson.NewDocument(bson.EC.String("_id", *cmd.ID))
//
//	c.client.Database(MONGO_DATABASE_WORKER).Collection(*cmd.ID).UpdateOne(nil, filter, cmd.Item)
//	//
//	//if err != nil {
//	//	log.Error("Insert mongodb error: ", err)
//	//	cmd.Callback("", err)
//	//	return
//	//}
//	//
//	//cmd.Callback(res.UpsertedID, nil)
//}

//func (c *session) deleteWorker(cmd Document) {
//	filter := bson.NewDocument(bson.EC.String("_id", *cmd.ID))
//
//	c.client.Database(MONGO_DATABASE_WORKER).Collection(*cmd.ID).DeleteOne(nil, filter)
//	//
//	//if err != nil {
//	//	log.Error("Insert mongodb error: ", err)
//	//	cmd.Callback("", err)
//	//	return
//	//}
//	//
//	//cmd.Callback(res.DeletedCount, nil)
//}
