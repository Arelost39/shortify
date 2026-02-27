package logger

import (
    "go.uber.org/zap"
    "sync"
)

var (
    instance *zap.Logger
    once sync.Once
)

func Init() {
    once.Do(func() {
        l, err := zap.NewProduction()
        if err != nil {
            panic(err)
        }
        instance = l
    })
}

// возвращает указатель на уже инициализированный логгер
func Logger() *zap.Logger {
    if instance == nil {
        Init()
    }
    return instance
}

// при завершении, чтобы не потерять логи
func Sync() error { return instance.Sync() }
