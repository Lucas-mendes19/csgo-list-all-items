package logger

import (
	"fmt"
	"io"
	"os"

	log "github.com/sirupsen/logrus"
)

func Setup() (*log.Logger, *os.File, error) {
	logger := log.New()

	file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0o666)
	if err != nil {
		return nil, nil, fmt.Errorf("erro ao abrir arquivo de log: %w", err)
	}

	multiWriter := io.MultiWriter(
		file,
		os.Stdout,
	)

	logger.SetOutput(multiWriter)
	logger.SetFormatter(&log.TextFormatter{
		FullTimestamp:   true,
		ForceColors:     true,
		TimestampFormat: "2006-01-02 15:04:05",
	})

	logger.AddHook(&fileHook{file})

	logger.Info("Aplicação iniciada")
	return logger, file, nil
}

type fileHook struct {
	file *os.File
}

func (h *fileHook) Levels() []log.Level {
	return log.AllLevels
}

func (h *fileHook) Fire(entry *log.Entry) error {
	formatter := &log.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	}
	line, err := formatter.Format(entry)
	if err != nil {
		return err
	}
	
	_, err = h.file.Write(line)
	return err
}