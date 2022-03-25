package mobilecard

import (
	"context"

	mobile_card_model "golang-structure/src/models/mobile_card"

	db "golang-structure/src/database"

	configuration "golang-structure/src/configs"

	"time"

	"go.mongodb.org/mongo-driver/bson"
)

func GetAll(limit int, offset int, from_date time.Time, to_date time.Time) ([]*mobile_card_model.MobileCardModel, error) {
	var result []*mobile_card_model.MobileCardModel
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	collection := db.MgDB.Db.Collection(configuration.Config.GetString("mongo.collections.mobile_card"))
	//  Stages
	great_than_stage := bson.M{"$match": bson.M{"exchange_time": bson.M{"$gt": from_date}}}
	less_than_stage := bson.M{"$match": bson.M{"exchange_time": bson.M{"$lt": to_date}}}
	order_stage := bson.M{"$sort": bson.M{"create_time": 1}}
	skip_stage := bson.M{"$skip": offset}
	limit_stage := bson.M{"$limit": limit}
	pipeline := []bson.M{great_than_stage, less_than_stage, order_stage, skip_stage, limit_stage}
	// Truy vấn dữ liệu
	cursor, err := collection.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	find_err := cursor.All(ctx, &result)
	if find_err != nil {
		return nil, find_err
	}
	return result, nil
}

func CountBeetwenDate(from_date time.Time, to_date time.Time) (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	collection := db.MgDB.Db.Collection(configuration.Config.GetString("mongo.collections.mobile_card"))
	match_stage := bson.M{"exchange_time": bson.M{"$gt": from_date, "$lt": to_date}}
	total, err := collection.CountDocuments(ctx, match_stage)
	if err != nil {
		return 0, err
	}
	return total, nil
}

func GetTotalPriceOfBrand(from_date time.Time, to_date time.Time) ([]*mobile_card_model.MobileCardGMVModel, error) {
	result := []*mobile_card_model.MobileCardGMVModel{}
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	collection := db.MgDB.Db.Collection(configuration.Config.GetString("mongo.collections.mobile_card"))
	great_than_stage := bson.M{"$match": bson.M{"exchange_time": bson.M{"$gt": from_date}}}
	less_than_stage := bson.M{"$match": bson.M{"exchange_time": bson.M{"$lt": to_date}}}
	group_stage := bson.M{"$group": bson.M{"_id": "$brand", "price": bson.M{"$sum": "$price"}}}
	pipeline := []bson.M{great_than_stage, less_than_stage, group_stage}
	cursor, err := collection.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	find_err := cursor.All(ctx, &result)
	if find_err != nil {
		return nil, find_err
	}
	return result, nil
}
