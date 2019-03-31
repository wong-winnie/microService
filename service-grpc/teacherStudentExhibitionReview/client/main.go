package main

import (
	"fmt"
	"context"
	"google.golang.org/grpc"
	pb "gitee.com/winnie_gss/microService/service-grpc/teacherStudentExhibitionReview/proto"
	"time"
	"gitee.com/winnie_gss/microService/service-log"
)

//发送消息通知
func SendActivityReviewMessage(client pb.ActivityReviewClient, ctx context.Context) {
	userList := []int64{111111, 222222, 333333, 444444, 555555, 666666, 777777, 888888}
	data := &pb.SendActivityReviewMessageRequest{
		MessageActivityId: 15,
		MessageSendId:     999999,
		MessageReceiveId:  userList,
		MessageType:       pb.MESSAGE_TYPE_RESULT_MESSAGE_TYPE_RESULT_RECEIVE,
		MessageTitle:      "春游",
		MessageContent:    "让我们来一次说走就走的旅行",
		MessageTime:       int32(time.Now().Unix()),
	}

	response, err := client.SendActivityReviewMessage(ctx, data)
	service_log.ErrorLog(err)

	fmt.Println(response)
}

//获取消息通知
func GetActivityReviewMessage(client pb.ActivityReviewClient, ctx context.Context) {
	data := &pb.GetActivityReviewMessageRequest{
		MessageReceiveId: 200275,
		MessageType:      pb.MESSAGE_TYPE_RESULT_MESSAGE_TYPE_RESULT_RECEIVE,
	}

	response, err := client.GetActivityReviewMessage(ctx, data)
	service_log.ErrorLog(err)

	fmt.Println(response)
}

//阅读消息通知
func ReadActivityReviewMessage(client pb.ActivityReviewClient, ctx context.Context) {
	data := &pb.ReadActivityReviewMessageRequest{
		MessageId: 70,
	}

	response, err := client.ReadActivityReviewMessage(ctx, data)
	service_log.ErrorLog(err)

	fmt.Println(response)
}

//消息查阅情况
func CheckActivityReviewMessage(client pb.ActivityReviewClient, ctx context.Context) {
	data := &pb.CheckActivityReviewMessageRequest{
		MessageParentId: 269,
	}

	response, err := client.CheckActivityReviewMessage(ctx, data)
	service_log.ErrorLog(err)

	fmt.Println(response)
}

func main() {
	service_log.InitLog()

	conn, err := grpc.Dial("192.168.3.177:50100", grpc.WithInsecure())
	service_log.ErrorLog(err)
	defer conn.Close()

	client := pb.NewActivityReviewClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	//TODO : 发送消息通知
	//SendActivityReviewMessage(client, ctx)

	//TODO : 获取消息通知
	//GetActivityReviewMessage(client, ctx)

	//TODO : 阅读消息通知
	//ReadActivityReviewMessage(client, ctx)

	//TODO : 消息查阅情况
	CheckActivityReviewMessage(client, ctx)
}
