package main

import (
	"net"
	"google.golang.org/grpc"
	pb "gitee.com/winnie_gss/microService/service-grpc/teacherStudentExhibitionReview/proto"
	"google.golang.org/grpc/reflection"
	"context"
	"gitee.com/winnie_gss/microService/service-log"
	"fmt"
	"gitee.com/winnie_gss/microService/service-sqlx"
	"errors"
	"time"
	"strconv"
)

const tableMessage = "ca_activity_message"
const defaultPage = 1
const defaultPageSize = 10

//[protoc --go_out=plugins=grpc:. ./activity_review.proto]
type server struct{}

type messageItem struct {
	MessageId         int32  `db:"messageId"`
	MessageActivityId int32  `db:"messageActivityId"`
	MessageSendId     int64  `db:"messageSendId"`
	MessageReceiveId  int64  `db:"messageReceiveId"`
	MessageType       int32  `db:"messageType"`
	MessageTitle      string `db:"messageTitle"`
	MessageContent    string `db:"messageContent"`
	MessageTime       int32  `db:"messageTime"`
	MessageIsRead     int32  `db:"messageIsRead"`
	MessageReadTime   int32  `db:"messageReadTime"`
}

type checkItem struct {
	MessageReceiveId int64 `db:"messageReceiveId"`
	MessageIsRead    int32 `db:"messageIsRead"`
	MessageReadTime  int32 `db:"messageReadTime"`
}

//TODO : 发送消息通知
func (s *server) SendActivityReviewMessage(ctx context.Context, request *pb.SendActivityReviewMessageRequest) (response *pb.CommonResponse, err error) {

	receiveId := request.MessageReceiveId

	data := ""
	var row int64 = 0
	for _, v := range receiveId {
		row++
		data += fmt.Sprintf("(%d, %d, %d, %d, '%s', '%s', %d, %d),", request.MessageActivityId, request.MessageSendId, v, request.MessageType, request.MessageTitle, request.MessageContent, request.MessageTime, request.MessageParentId)
	}

	dataRune := []rune(data)

	sql := fmt.Sprintf("insert into %s (messageActivityId, messageSendId, messageReceiveId, messageType, messageTitle, messageContent, messageTime, MessageParentId) values %s", tableMessage, string(dataRune[:len(dataRune)-1]))

	result := service_sqlx.UpdateDate(sql)
	lastInsertId, err := result.LastInsertId()

	if rowAffected, err := result.RowsAffected(); err != nil || rowAffected != row {
		service_log.ErrorLog(err)
		service_log.ErrorLog(errors.New(fmt.Sprintf("预计增加%d条, 实际增加%d条", row, rowAffected)))
		return &pb.CommonResponse{Result: pb.RPC_CALL_RESULT_RPC_CALL_RESULT_ERROR}, nil
	}

	return &pb.CommonResponse{Result: pb.RPC_CALL_RESULT_RPC_CALL_RESULT_SUCCESS, Msg: strconv.FormatInt(lastInsertId, 10)}, nil
}

//TODO : 获取消息通知
func (s *server) GetActivityReviewMessage(ctx context.Context, request *pb.GetActivityReviewMessageRequest) (response *pb.GetActivityReviewMessageResponse, err error) {

	page := request.Page
	if page == 0 {
		page = defaultPage
	}

	pageSize := request.PageSize
	if pageSize == 0 {
		pageSize = defaultPageSize
	}

	sql := ""
	if request.MessageType == pb.MESSAGE_TYPE_RESULT_MESSAGE_TYPE_RESULT_ALL {
		sql = fmt.Sprintf("select messageId, messageActivityId, messageSendId, messageReceiveId, messageType, messageTitle, messageContent, messageTime, messageIsRead, messageReadTime from %s where messageReceiveId=%d order by messageTime desc limit %d, %d", tableMessage, request.MessageReceiveId, (page-1)*pageSize, pageSize)
	} else {
		sql = fmt.Sprintf("select messageId, messageActivityId, messageSendId, messageReceiveId, messageType, messageTitle, messageContent, messageTime, messageIsRead, messageReadTime from %s where messageReceiveId=%d and messageType=%d order by messageTime desc limit %d, %d", tableMessage, request.MessageReceiveId, request.MessageType, (page-1)*pageSize, pageSize)
	}

	var data1 []messageItem
	var data2 []*pb.MessageItem
	service_sqlx.SelectAll(sql, &data1)

	var mType pb.MESSAGE_TYPE_RESULT
	var mIsRead pb.MESSAGE_TYPE_READ
	for _, v := range data1 {
		if v.MessageType == 1 {
			mType = pb.MESSAGE_TYPE_RESULT_MESSAGE_TYPE_RESULT_SEND
		} else {
			mType = pb.MESSAGE_TYPE_RESULT_MESSAGE_TYPE_RESULT_RECEIVE
		}

		if v.MessageIsRead == 0 {
			mIsRead = pb.MESSAGE_TYPE_READ_MESSAGE_TYPE_READ_NO
		} else {
			mIsRead = pb.MESSAGE_TYPE_READ_MESSAGE_TYPE_READ_YES
		}

		list := &pb.MessageItem{
			MessageId:         v.MessageId,
			MessageActivityId: v.MessageActivityId,
			MessageSendId:     v.MessageSendId,
			MessageReceiveId:  v.MessageReceiveId,
			MessageType:       mType,
			MessageTitle:      v.MessageTitle,
			MessageContent:    v.MessageContent,
			MessageTime:       v.MessageTime,
			MessageIsRead:     mIsRead,
			MessageReadTime:   v.MessageReadTime,
		}
		data2 = append(data2, list)
	}

	var maxCount int32
	sqlMaxCount := ""
	if request.MessageType == pb.MESSAGE_TYPE_RESULT_MESSAGE_TYPE_RESULT_ALL {
		sqlMaxCount = fmt.Sprintf("select count(*) from %s where messageReceiveId=%d", tableMessage, request.MessageReceiveId)
	} else {
		sqlMaxCount = fmt.Sprintf("select count(*) from %s where messageReceiveId=%d and messageType=%d", tableMessage, request.MessageReceiveId, request.GetMessageType())
	}
	service_sqlx.SelectOne(sqlMaxCount, &maxCount)

	var noReadCount int32
	sqlNoReadCount := fmt.Sprintf("select count(*) from %s where messageReceiveId=%d and messageIsRead=%d and messageType=2", tableMessage, request.MessageReceiveId, pb.MESSAGE_TYPE_READ_MESSAGE_TYPE_READ_NO)
	service_sqlx.SelectOne(sqlNoReadCount, &noReadCount)

	return &pb.GetActivityReviewMessageResponse{Result: pb.RPC_CALL_RESULT_RPC_CALL_RESULT_SUCCESS, MaxCount: maxCount, NoReadCount: noReadCount, List: data2}, nil
}

//TODO : 阅读消息通知
func (s *server) ReadActivityReviewMessage(ctx context.Context, request *pb.ReadActivityReviewMessageRequest) (response *pb.CommonResponse, err error) {

	sql := fmt.Sprintf("update %s set messageIsRead=%d, messageReadTime=%d where messageId=%d", tableMessage, pb.MESSAGE_TYPE_READ_MESSAGE_TYPE_READ_YES, time.Now().Unix(), request.MessageId)
	result := service_sqlx.UpdateDate(sql)

	if n, err := result.RowsAffected(); n != 1 || err != nil {
		service_log.ErrorLog(errors.New(fmt.Sprintf("影响行数:%d", n)))
		service_log.ErrorLog(err)
		return &pb.CommonResponse{Result: pb.RPC_CALL_RESULT_RPC_CALL_RESULT_ERROR}, nil
	}

	return &pb.CommonResponse{Result: pb.RPC_CALL_RESULT_RPC_CALL_RESULT_SUCCESS}, err
}

//TODO : 消息查阅情况
func (s *server) CheckActivityReviewMessage(ctx context.Context, request *pb.CheckActivityReviewMessageRequest) (response *pb.CheckActivityReviewMessageResponse, err error) {

	page := request.Page
	if page == 0 {
		page = defaultPage
	}

	pageSize := request.PageSize
	if pageSize == 0 {
		pageSize = defaultPageSize
	}

	sql := fmt.Sprintf("select messageReceiveId, messageIsRead, messageReadTime from %s where messageParentId=%d order by messageReceiveId limit %d, %d", tableMessage, request.MessageParentId, (page-1)*pageSize, pageSize)

	var data1 []checkItem
	var data2 []*pb.CheckItem
	service_sqlx.SelectAll(sql, &data1)

	var mIsRead pb.MESSAGE_TYPE_READ
	for _, v := range data1 {
		if v.MessageIsRead == 0 {
			mIsRead = pb.MESSAGE_TYPE_READ_MESSAGE_TYPE_READ_NO
		} else {
			mIsRead = pb.MESSAGE_TYPE_READ_MESSAGE_TYPE_READ_YES
		}

		list := &pb.CheckItem{
			MessageReceiveId: v.MessageReceiveId,
			MessageIsRead:    mIsRead,
			MessageReadTime:  v.MessageReadTime,
		}
		data2 = append(data2, list)
	}

	var receiveIdList1 []struct{ MessageReceiveId int64 `db:"messageReceiveId"` }
	var receiveIdList2 []int64
	sqlMaxCount := fmt.Sprintf("select messageReceiveId from %s where messageParentId=%d", tableMessage, request.MessageParentId)
	service_sqlx.SelectAll(sqlMaxCount, &receiveIdList1)

	var maxCount int32 = 0
	for _, v := range receiveIdList1 {
		maxCount ++
		receiveIdList2 = append(receiveIdList2, v.MessageReceiveId)
	}

	var mTime int32
	sqlTime := fmt.Sprintf("select messageTime from %s where messageId=%d", tableMessage, request.MessageParentId)
	service_sqlx.SelectOne(sqlTime, &mTime)

	return &pb.CheckActivityReviewMessageResponse{Result: pb.RPC_CALL_RESULT_RPC_CALL_RESULT_SUCCESS, MaxCount: maxCount, MessageReceiveId: receiveIdList2, MessageTime: mTime, List: data2}, nil
}

func main() {

	service_log.InitLog()
	service_sqlx.InitDb()

	listen, err := net.Listen("tcp", ":50100")
	service_log.ErrorLog(err)

	s := grpc.NewServer()
	pb.RegisterActivityReviewServer(s, &server{})
	reflection.Register(s)

	if err := s.Serve(listen); err != nil {
		service_log.ErrorLog(err)
	}
}
