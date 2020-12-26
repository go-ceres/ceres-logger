// Copyright 2020 Go-Ceres
// Author https://github.com/go-ceres/go-ceres
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package CeresLogger

import (
	CeresConfig "github.com/go-ceres/ceres-config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Config struct {
	Debug        bool          // 是否开启debug模式，默认false
	Level        string        // 日志等级
	Stdout       bool          // 是否在控制台打印，默认，true
	Rotate       bool          // 是否文件日志输出
	RotateConfig *RotateConfig // 日志输出配置文件
	Fields       []zap.Field   // 初始化字段
	AddCaller    bool          // 是否打印调用者信息，默认，true
	CallerSkip   int           // 表示输出当前栈帧，默认，1
	autoLevelKey string        // 日志等级监听key
	Core         zapcore.Core
}

// 获取一个默认的配置
func defaultConfig() *Config {
	return &Config{
		Stdout: true,
	}
}

// RawConfig 根据key构建配置
func RawConfig(key string) *Config {
	conf := defaultConfig()
	if err := CeresConfig.Get(key).Scan(conf); err != nil {
		panic(err)
	}
	return conf
}

// SanConfig 根据name构建配置
func SanConfig(name string) *Config {
	return RawConfig("ceres.logger." + name)
}

// 创建logger
func (c Config) Build() Logger {
	logger := newLogger(&c)
	if c.autoLevelKey != "" {
		logger.AutoLevel(c.autoLevelKey)
	}
	return logger
}
