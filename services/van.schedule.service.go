package services

import (
	"context"
	"van_thailand_server/models"
	"van_thailand_server/repositories"

	"go.mongodb.org/mongo-driver/mongo"
)

func GetVanSchedule(ctx context.Context, targetScheduleId string) *models.ReturnScheduleStruct {
	result := repositories.GetVanSchedule(ctx, targetScheduleId)
	return result
}

func GetVanSchedules(ctx context.Context, targetVanId string) []*models.ReturnScheduleStruct {
	result := repositories.GetVanSchedules(ctx, targetVanId)
	return result
}

func CreateVanSchedule(ctx context.Context, schedule *models.ScheduleStruct) *mongo.InsertOneResult {
	result := repositories.CreateSchedule(ctx, schedule)
	return result
}

func UpdateSchedule(ctx context.Context, scheduleId string, schedule *models.ScheduleStruct) int {
	result := repositories.UpdateSchedule(ctx, scheduleId, schedule)
	return result
}

func DeleteSchedule(ctx context.Context, scheduleId string) int {
	result := repositories.DeleteSchedule(ctx, scheduleId)
	return result
}
