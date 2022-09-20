package main

import (
	"flag"
	"fmt"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"os"
)

const ( // 수정필수
	tgmToken = "xxxxxxxxxxxxxI56Glt40zX2T-lyo" // 텔레그램 토큰
	logPath  = "./"                            // 로그가 기록될 경로 // 로그파일명
)

const logFileName = "telegram_sendlog.log"

type MsgFormat struct {
	targetID int64
	sendMSG  string
	tgCursor *tgbotapi.BotAPI
}

func (msg *MsgFormat) tgConnect() {
	tgCursor, err := tgbotapi.NewBotAPI(tgmToken)
	if err != nil {
		log.Panic(err)
	}
	msg.tgCursor = tgCursor
	msg.tgCursor.Debug = false
	me, err := tgCursor.GetMe()
	fmt.Println(me)
	if err != nil {
		return
	}
	if err != nil {
		log.Panic(err)
	}
}

func (msg *MsgFormat) msgSender() {
	sendResult := tgbotapi.NewMessage(msg.targetID, msg.sendMSG)
	msg.tgCursor.Send(sendResult)
}

func main() {
	fpLog, err := os.OpenFile(logPath+logFileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("Error : make " + logPath + logFileName)
		log.Panic(err)
	}
	defer fpLog.Close()
	myLogger := log.New(fpLog, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	reciever := flag.Int64("r", 0, "SET msg reciever")
	message := flag.String("m", "", "SET message")
	flag.Parse()
	if *reciever == 0 || *message == "" {
		fmt.Println("- 옵션이 부족함")
		os.Exit(0)
	}
	var mf MsgFormat
	mf.targetID = *reciever
	mf.sendMSG = *message
	mf.tgConnect()
	mf.msgSender()
	myLogger.Println(*message, "발송완료")
}
