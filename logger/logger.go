package logger

import "fmt"

//	"log"

const (
	//go-logger version
	_VER string = "1.0.3"
)

type LEVEL int32
type UNIT int64
type _ROLLTYPE int //dailyRolling ,rollingFile

const _DATEFORMAT = "2006-01-02"

var logLevel LEVEL = 1

const (
	_       = iota
	KB UNIT = 1 << (iota * 10)
	MB
	GB
	TB
)

const (
	ALL LEVEL = iota
	DEBUG
	INFO
	WARN
	ERROR
	FATAL
	OFF
)

const (
	_DAILY _ROLLTYPE = iota
	_ROLLFILE
)

func SetConsole(isConsole bool) {
	defaultlog.setConsole(isConsole)
}

func SetLevel(_level LEVEL) {
	defaultlog.setLevel(_level)
}

func SetFormat(logFormat string) {
	defaultlog.setFormat(logFormat)
}

func SetRollingFile(fileDir, fileName string, maxNumber int32, maxSize int64, _unit UNIT) {
	//	maxFileCount = maxNumber
	//	maxFileSize = maxSize * int64(_unit)
	//	RollingFile = true
	//	dailyRolling = false
	//	mkdirlog(fileDir)
	//	logObj = &_FILE{dir: fileDir, filename: fileName, isCover: false, mu: new(sync.RWMutex)}
	//	logObj.mu.Lock()
	//	defer logObj.mu.Unlock()
	//	for i := 1; i <= int(maxNumber); i++ {
	//		if isExist(fileDir + "/" + fileName + "." + strconv.Itoa(i)) {
	//			logObj._suffix = i
	//		} else {
	//			break
	//		}
	//	}
	//	if !logObj.isMustRename() {
	//		logObj.logfile, _ = os.OpenFile(fileDir+"/"+fileName, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)
	//		logObj.lg = log.New(logObj.logfile, "", log.Ldate|log.Ltime|log.Lshortfile)
	//	} else {
	//		logObj.rename()
	//	}
	//	go fileMonitor()
	defaultlog.setRollingFile(fileDir, fileName, maxNumber, maxSize, _unit)
}

func SetRollingDaily(fileDir, fileName string) {
	//	RollingFile = false
	//	dailyRolling = true
	//	t, _ := time.Parse(_DATEFORMAT, time.Now().Format(_DATEFORMAT))
	//	mkdirlog(fileDir)
	//	logObj = &_FILE{dir: fileDir, filename: fileName, _date: &t, isCover: false, mu: new(sync.RWMutex)}
	//	logObj.mu.Lock()
	//	defer logObj.mu.Unlock()
	//	if !logObj.isMustRename() {
	//		logObj.logfile, _ = os.OpenFile(fileDir+"/"+fileName, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)
	//		logObj.lg = log.New(logObj.logfile, "", log.Ldate|log.Ltime|log.Lshortfile)
	//	} else {
	//		logObj.rename()
	//	}
	defaultlog.setRollingDaily(fileDir, fileName)
}

//func console(s ...interface{}) {
//	if consoleAppender {
//		_, file, line, _ := runtime.Caller(2)
//		short := file
//		for i := len(file) - 1; i > 0; i-- {
//			if file[i] == '/' {
//				short = file[i+1:]
//				break
//			}
//		}
//		file = short
//		log.Println(file, strconv.Itoa(line), s)
//	}
//}

//func catchError() {
//	if err := recover(); err != nil {
//		log.Println("err", err)
//	}
//}

func Debug(v ...interface{}) {
	//	if dailyRolling {
	//		fileCheck()
	//	}
	//	defer catchError()
	//	if logObj != nil {
	//		logObj.mu.RLock()
	//		defer logObj.mu.RUnlock()
	//	}
	//	if logLevel <= DEBUG {
	//		if logObj != nil {
	//			logObj.lg.Output(2, fmt.Sprintln("debug", v))
	//		}
	//		console("debug", v)
	//	}
	defaultlog.debug(v...)
}
func Debugf(format string, v ...interface{}) {
	defaultlog.debug(fmt.Sprintf(format, v...))
}
func Info(v ...interface{}) {
	//	if dailyRolling {
	//		fileCheck()
	//	}
	//	defer catchError()
	//	if logObj != nil {
	//		logObj.mu.RLock()
	//		defer logObj.mu.RUnlock()
	//	}
	//	if logLevel <= INFO {
	//		if logObj != nil {
	//			if format == "" {
	//				logObj.lg.Output(2, fmt.Sprintln("info", v))
	//			} else {
	//				logObj.lg.Output(2, fmt.Sprintf(format, v...))
	//			}
	//		}
	//		console("info", v)
	//	}
	defaultlog.info(v...)
}
func Infof(format string, v ...interface{}) {
	defaultlog.info(fmt.Sprintf(format, v...))
}
func Warn(v ...interface{}) {
	//	if dailyRolling {
	//		fileCheck()
	//	}
	//	defer catchError()
	//	if logObj != nil {
	//		logObj.mu.RLock()
	//		defer logObj.mu.RUnlock()
	//	}
	//	if logLevel <= WARN {
	//		if logObj != nil {
	//			logObj.lg.Output(2, fmt.Sprintln("warn", v))
	//		}
	//		console("warn", v)
	//	}
	defaultlog.warn(v...)
}
func Warnf(format string, v ...interface{}) {
	defaultlog.warn(fmt.Sprintf(format, v...))
}
func Error(v ...interface{}) {
	//	if dailyRolling {
	//		fileCheck()
	//	}
	//	defer catchError()
	//	if logObj != nil {
	//		logObj.mu.RLock()
	//		defer logObj.mu.RUnlock()
	//	}
	//	if logLevel <= ERROR {
	//		if logObj != nil {
	//			logObj.lg.Output(2, fmt.Sprintln("error", v))
	//		}
	//		console("error", v)
	//	}
	defaultlog.error(v...)
}
func Errorf(format string, v ...interface{}) {
	defaultlog.error(fmt.Sprintf(format, v...))
}
func Fatal(v ...interface{}) {
	//	if dailyRolling {
	//		fileCheck()
	//	}
	//	defer catchError()
	//	if logObj != nil {
	//		logObj.mu.RLock()
	//		defer logObj.mu.RUnlock()
	//	}
	//	if logLevel <= FATAL {
	//		if logObj != nil {
	//			logObj.lg.Output(2, fmt.Sprintln("fatal", v))
	//		}
	//		console("fatal", v)
	//	}
	defaultlog.fatal(v...)
}
func Fatalf(format string, v ...interface{}) {
	defaultlog.fatal(fmt.Sprintf(format, v...))
}

func SetLevelFile(level LEVEL, dir, fileName string) {
	defaultlog.setLevelFile(level, dir, fileName)
}

//func isMustRename() bool {
//	if dailyRolling {
//		t, _ := time.Parse(_DATEFORMAT, time.Now().Format(_DATEFORMAT))
//		if t.After(*f._date) {
//			return true
//		}
//	} else {
//		if maxFileCount > 1 {
//			if fileSize(f.dir+"/"+f.filename) >= maxFileSize {
//				return true
//			}
//		}
//	}
//	return false
//}

//func rename() {
//	if dailyRolling {
//		fn := f.dir + "/" + f.filename + "." + f._date.Format(_DATEFORMAT)
//		if !isExist(fn) && f.isMustRename() {
//			if f.logfile != nil {
//				f.logfile.Close()
//			}
//			err := os.Rename(f.dir+"/"+f.filename, fn)
//			if err != nil {
//				f.lg.Println("rename err", err.Error())
//			}
//			t, _ := time.Parse(_DATEFORMAT, time.Now().Format(_DATEFORMAT))
//			f._date = &t
//			f.logfile, _ = os.Create(f.dir + "/" + f.filename)
//			f.lg = log.New(logObj.logfile, "\n", log.Ldate|log.Ltime|log.Lshortfile)
//		}
//	} else {
//		f.coverNextOne()
//	}
//}

//func nextSuffix() int {
//	return int(f._suffix%int(maxFileCount) + 1)
//}

//func coverNextOne() {
//	f._suffix = f.nextSuffix()
//	if f.logfile != nil {
//		f.logfile.Close()
//	}
//	if isExist(f.dir + "/" + f.filename + "." + strconv.Itoa(int(f._suffix))) {
//		os.Remove(f.dir + "/" + f.filename + "." + strconv.Itoa(int(f._suffix)))
//	}
//	os.Rename(f.dir+"/"+f.filename, f.dir+"/"+f.filename+"."+strconv.Itoa(int(f._suffix)))
//	f.logfile, _ = os.Create(f.dir + "/" + f.filename)
//	f.lg = log.New(logObj.logfile, "\n", log.Ldate|log.Ltime|log.Lshortfile)
//}

//func fileMonitor() {
//	timer := time.NewTicker(1 * time.Second)
//	for {
//		select {
//		case <-timer.C:
//			fileCheck()
//		}
//	}
//}

//func fileCheck() {
//	defer func() {
//		if err := recover(); err != nil {
//			log.Println(err)
//		}
//	}()
//	if logObj != nil && logObj.isMustRename() {
//		logObj.mu.Lock()
//		defer logObj.mu.Unlock()
//		logObj.rename()
//	}
//}
