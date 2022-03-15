/**
    @Author:     ZonzeeLi
    @Project:    chat_demo
    @CreateDate: 3/14/2022
    @UpdateDate: 3/15/2022
    @Note:       MongoDb数据库操作
**/

package chat

import (
	"chat_demo/global"
	"chat_demo/model/chat"
	"context"
	"encoding/json"
	"sort"
	"time"

	"go.uber.org/zap"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func InsertMsg(database string, id string, content string, read uint, expire int64) (err error) {
	collection := global.Chat_MONGODB.Database(database).Collection(id)
	comment := chat.Trainer{
		Content:   content,
		StartTime: time.Now().Unix(),
		EndTime:   time.Now().Unix() + expire,
		Read:      read,
	}
	_, err = collection.InsertOne(context.TODO(), comment)
	return
}

func FindHistory(database string, sendId string, id string, time int64, pageSize int) (results []chat.Result) {
	var resultsID []chat.Trainer
	var resultsSendID []chat.Trainer
	sendIdCollection := global.Chat_MONGODB.Database(database).Collection(sendId)
	idCollection := global.Chat_MONGODB.Database(database).Collection(id)
	// 如果不知道该使用什么context，可以通过context.TODO() 产生context
	sendIdTimeCursor, err := sendIdCollection.Find(context.TODO(), bson.D{{}},
		options.Find().SetSort(bson.D{{"startTime", -1}}), options.Find().SetLimit(int64(pageSize)))
	if err != nil {
		global.Chat_LOG.Error("[MongoDb] find sendID failed ", zap.Error(err))
	}
	idTimeCursor, err := idCollection.Find(context.TODO(), bson.D{{}},
		options.Find().SetSort(bson.D{{"startTime", -1}}), options.Find().SetLimit(int64(pageSize)))
	if err != nil {
		global.Chat_LOG.Error("[MongoDb] find ID failed ", zap.Error(err))
	}
	err = sendIdTimeCursor.All(context.TODO(), &resultsSendID) // sendId 对面发过来的
	if err != nil {
		global.Chat_LOG.Error("[MongoDb] decode sendID results failed ", zap.Error(err))
	}
	err = idTimeCursor.All(context.TODO(), &resultsID) // Id 发给对面的
	if err != nil {
		global.Chat_LOG.Error("[MongoDb] decode ID results failed ", zap.Error(err))
	}
	results = AppendAndSort(resultsID, resultsSendID, id, sendId)
	return
}

func FindUnreadMsg(database string, sendId string, id string) (results []chat.Result) {
	// 查询双方未读消息
	var resultsID []chat.Trainer
	var resultsSendID []chat.Trainer
	sendIdCollection := global.Chat_MONGODB.Database(database).Collection(sendId)
	idCollection := global.Chat_MONGODB.Database(database).Collection(id)
	filters := bson.D{{"read", bson.M{"$eq": 0}}}
	sendIdCursor, err := sendIdCollection.Find(context.TODO(), filters, options.Find().SetSort(bson.D{{"startTime", 1}}))
	if err != nil {
		global.Chat_LOG.Error("[MongoDb] find sendId unread msg failed", zap.Error(err))
		return
	}
	var unReads []chat.Trainer
	err = sendIdCursor.All(context.TODO(), &unReads)
	if err != nil {
		global.Chat_LOG.Error("[MongoDb] decode sendId unread msg failed", zap.Error(err))
	}
	if len(unReads) > 0 {
		timeFilter := bson.M{
			"startTime": bson.M{
				"$gte": unReads[0].StartTime,
			},
		}
		sendIdTimeCursor, err := sendIdCollection.Find(context.TODO(), timeFilter)
		if err != nil {
			global.Chat_LOG.Error("[MongoDb] find sendId unread msg by gte starttime failed", zap.Error(err))
		}
		idTimeCursor, err := idCollection.Find(context.TODO(), timeFilter)
		if err != nil {
			global.Chat_LOG.Error("[MongoDb] find id unread msg before starttime failed", zap.Error(err))
		}
		err = sendIdTimeCursor.All(context.TODO(), &resultsSendID)
		if err != nil {
			global.Chat_LOG.Error("[MongoDb] decode sendId unread msg before starttime failed", zap.Error(err))
		}
		err = idTimeCursor.All(context.TODO(), &resultsID)
		if err != nil {
			global.Chat_LOG.Error("[MongoDb] decode id unread msg before starttime failed", zap.Error(err))
		}
		results = AppendAndSort(resultsID, resultsSendID, id, sendId)

	} else {
		return
	}
	overTimeFilter := bson.D{
		{"$and", bson.A{
			bson.D{{"endTime", bson.M{"&lt": time.Now().Unix()}}},
			bson.D{{"read", bson.M{"$eq": 1}}},
		}},
	}
	// 删掉过期的历史记录
	_, err = sendIdCollection.DeleteMany(context.TODO(), overTimeFilter)
	if err != nil {
		global.Chat_LOG.Error("[MongoDb] delete sendId expired msg failed", zap.Error(err))
	}
	_, err = idCollection.DeleteMany(context.TODO(), overTimeFilter)
	if err != nil {
		global.Chat_LOG.Error("[MongoDb] delete id expired msg failed", zap.Error(err))
	}
	// 将己方未读的消息所有的维度设置为已读
	_, err = sendIdCollection.UpdateMany(context.TODO(), filters, bson.M{"$set": bson.M{"read": 1}})
	if err != nil {
		global.Chat_LOG.Error("[MongoDb] update sendId unread failed", zap.Error(err))
	}
	_, err = sendIdCollection.UpdateMany(context.TODO(), filters, bson.M{"$set": bson.M{"endTime": time.Now().Unix() + int64(3*chat.Month)}})
	if err != nil {
		global.Chat_LOG.Error("[MongoDb] update id unread failed", zap.Error(err))
	}
	return
}

func AppendAndSort(resultsID, resultsSendID []chat.Trainer, id, sendId string) (results []chat.Result) {
	for _, r := range resultsID {
		sendSort := chat.SendSortMsg{
			Content:  r.Content,
			Read:     r.Read,
			CreateAt: r.StartTime,
		}
		msg, _ := json.Marshal(sendSort)
		result := chat.Result{
			StartTime: r.StartTime,
			Content:   sendSort,
			//Msg:       msg,
			Msg:  string(msg),
			From: id, // 可以传入username
		}
		results = append(results, result)
	}
	for _, r := range resultsSendID {
		sendSort := chat.SendSortMsg{
			Content:  r.Content,
			Read:     r.Read,
			CreateAt: r.StartTime,
		}
		msg, _ := json.Marshal(sendSort)
		result := chat.Result{
			StartTime: r.StartTime,
			Content:   sendSort,
			//Msg:       msg,
			Msg:  string(msg),
			From: sendId,
		}
		results = append(results, result)
	}
	// 最后进行排序
	sort.Slice(results, func(i, j int) bool { return results[i].StartTime < results[j].StartTime })
	return results
}
