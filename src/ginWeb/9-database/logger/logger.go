package logger

import (
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"time"
)

// @Description
// @Author lianggengguang
// @Date 2023/6/18

var Log = logrus.New()

type logFileWriter struct {
	file     *os.File
	logPath  string
	fileDate string
	appName  string
}

func (p *logFileWriter) newLog() (n int, err error) {

	err = os.MkdirAll(fmt.Sprintf("%s/%s", p.logPath, p.fileDate), os.ModePerm)
	if err != nil {
		return 0, err
	}
	fileName := fmt.Sprintf("%s/%s/%s-%s.log", p.logPath, p.fileDate, p.appName, p.fileDate)
	p.file, err = os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE|os.O_APPEND|os.O_SYNC, 0600)
	if err != nil {
		return 0, err
	}
	return 0, nil
}

//实现IO的write接口
func (p *logFileWriter) Write(data []byte) (n int, err error) {

	if p == nil {
		return 0, errors.New("logFileWriter is nil")
	}
	if p.file == nil {
		return 0, errors.New("logFile not opened")
	}
	//判断是否需要切换日期
	fileDate := time.Now().Format("2006-01-02")
	if p.fileDate != fileDate {
		p.file.Close()
		p.fileDate = fileDate
		_, err := p.newLog()
		if err != nil {
			return 0, err
		}
	}
	n, e := p.file.Write(data)
	return n, e
}

func initCfg(filePath string, appName string, logLevel logrus.Level) *logrus.Logger {

	log := logrus.New()
	//设置在输出日志中添加文件名和方法信息
	//log.SetReportCaller(true)
	//设置输出格式为json
	formatter := logrus.TextFormatter{
		FullTimestamp:   true,                  //表示展示日期
		TimestampFormat: "2006-01-02 15:04:04", //日期的格式
		ForceColors:     true,                  //控制台输出的日志带有颜色
	}
	log.SetFormatter(&formatter)
	//设置日志基本为debug
	log.SetLevel(logLevel)
	//日志打印到控制台和文件  Stderr是无缓冲 每个输出都会立即flush 符合作为日志的需要
	fileDate := time.Now().Format("2006-01-02")
	fileWriter := logFileWriter{nil, filePath, fileDate, appName}
	_, err := fileWriter.newLog()
	if err != nil {
		logrus.Error(err)
		return log
	}
	//在文件和控制台都打印
	log.SetOutput(io.MultiWriter(fileWriter.file, os.Stderr))
	return log
}

func init() {

	//初始化配置
	Log = initCfg("./log/", "myApp", logrus.InfoLevel)
	Log.Info("Initialize Log...")
}
