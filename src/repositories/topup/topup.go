package topup

import (
	"context"
	"time"

	configuration "golang-structure/src/configs"

	topup_model "golang-structure/src/models/topup"

	db "golang-structure/src/database"

	"go.mongodb.org/mongo-driver/bson"
)

func GetAll(limit int, offset int, from_date time.Time, to_date time.Time) ([]*topup_model.TopupModel, error) {
	var result []*topup_model.TopupModel
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	collection := db.MgDB.Db.Collection(configuration.Config.GetString("mongo.collections.topup"))
	//  Stages
	great_than_stage := bson.M{"$match": bson.M{"topup_time": bson.M{"$gt": from_date}}}
	less_than_stage := bson.M{"$match": bson.M{"topup_time": bson.M{"$lt": to_date}}}
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
	collection := db.MgDB.Db.Collection(configuration.Config.GetString("mongo.collections.topup"))
	match_stage := bson.M{"topup_time": bson.M{"$gt": from_date, "$lt": to_date}}
	total, err := collection.CountDocuments(ctx, match_stage)
	if err != nil {
		return 0, err
	}
	return total, nil
}

func GetTotalAmount(from_date time.Time, to_date time.Time) (int, error) {
	result := []*topup_model.TopupGmvModel{}
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	collection := db.MgDB.Db.Collection(configuration.Config.GetString("mongo.collections.topup"))
	great_than_stage := bson.M{"$match": bson.M{"topup_time": bson.M{"$gt": from_date}}}
	less_than_stage := bson.M{"$match": bson.M{"topup_time": bson.M{"$lt": to_date}}}
	group_stage := bson.M{"$group": bson.M{"_id": "", "amount": bson.M{"$sum": "$amount"}}}
	pipeline := []bson.M{great_than_stage, less_than_stage, group_stage}
	cursor, err := collection.Aggregate(ctx, pipeline)
	if err != nil {
		return 0, err
	}
	defer cursor.Close(ctx)
	find_err := cursor.All(ctx, &result)
	if find_err != nil {
		return 0, err
	}
	if len(result) > 0 {
		return result[0].GMV, nil
	} else {
		return 0, nil
	}
}
