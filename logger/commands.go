// Copyright 2013, David Fisher. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// author: David Fisher <ddf1991@gmail.com>
// based on previous package by: Cong Ding <dinggnu@gmail.com>
package logger

import (
	"fmt"
)

var BasicFormat = "%s [%9s] %s- %s\n"
var BasicFields = []string{"time", "priority", "prefix", "message"}
var RichFormat = "%s [%9s] %d %s - %s:%s:%d - %s\n"
var RichFields = []string{"full_time", "priority", "seq", "prefix", "filename", "funcname", "lineno", "message"}

// This function has an unusual name to aid in finding it while walking the
// stack. We need to do some dead reckoning from this function to access the
// caller's stack, so there is a consistent call depth above this function.
func (logger *Logger) Log(priority Priority, message string) {
	fields := logger.fieldValues()
	fields["priority"] = priority
	fields["message"] = message
	for _, sink := range logger.sinks {
		sink.Log(fields)
	}
}

func (logger *Logger) Logf(priority Priority, format string, v ...interface{}) {
	logger.Log(priority, fmt.Sprintf(format, v...))
}


func (logger *Logger) Emergency(message string) {
	logger.Log(PriEmerg, message)
}
func (logger *Logger) Emergencyf(format string, v...interface{}) {
	logger.Log(PriEmerg, fmt.Sprintf(format, v...))
}

func (logger *Logger) Alert(message string) {
	logger.Log(PriAlert, message)
}
func (logger *Logger) Alertf(format string, v...interface{}) {
	logger.Log(PriAlert, fmt.Sprintf(format, v...))
}

func (logger *Logger) Critical(message string) {
	logger.Log(PriCrit, message)
}
func (logger *Logger) Criticalf(format string, v...interface{}) {
	logger.Log(PriCrit, fmt.Sprintf(format, v...))
}

func (logger *Logger) Error(message string) {
	logger.Log(PriErr, message)
}
func (logger *Logger) Errorf(format string, v...interface{}) {
	logger.Log(PriErr, fmt.Sprintf(format, v...))
}

func (logger *Logger) Warning(message string) {
	logger.Log(PriWarning, message)
}
func (logger *Logger) Warningf(format string, v...interface{}) {
	logger.Log(PriWarning, fmt.Sprintf(format, v...))
}

func (logger *Logger) Notice(message string) {
	logger.Log(PriNotice, message)
}
func (logger *Logger) Noticef(format string, v...interface{}) {
	logger.Log(PriNotice, fmt.Sprintf(format, v...))
}

func (logger *Logger) Info(message string) {
	logger.Log(PriInfo, message)
}
func (logger *Logger) Infof(format string, v...interface{}) {
	logger.Log(PriInfo, fmt.Sprintf(format, v...))
}

func (logger *Logger) Debug(message string) {
	logger.Log(PriDebug, message)
}
func (logger *Logger) Debugf(format string, v...interface{}) {
	logger.Log(PriDebug, fmt.Sprintf(format, v...))
}


func Emergency(message string) {
	defaultLogger.Log(PriEmerg, message)
}
func Emergencyf(format string, v...interface{}) {
	defaultLogger.Log(PriEmerg, fmt.Sprintf(format, v...))
}

func Alert(message string) {
	defaultLogger.Log(PriAlert, message)
}
func Alertf(format string, v...interface{}) {
	defaultLogger.Log(PriAlert, fmt.Sprintf(format, v...))
}

func Critical(message string) {
	defaultLogger.Log(PriCrit, message)
}
func Criticalf(format string, v...interface{}) {
	defaultLogger.Log(PriCrit, fmt.Sprintf(format, v...))
}

func Error(message string) {
	defaultLogger.Log(PriErr, message)
}
func Errorf(format string, v...interface{}) {
	defaultLogger.Log(PriErr, fmt.Sprintf(format, v...))
}

func Warning(message string) {
	defaultLogger.Log(PriWarning, message)
}
func Warningf(format string, v...interface{}) {
	defaultLogger.Log(PriWarning, fmt.Sprintf(format, v...))
}

func Notice(message string) {
	defaultLogger.Log(PriNotice, message)
}
func Noticef(format string, v...interface{}) {
	defaultLogger.Log(PriNotice, fmt.Sprintf(format, v...))
}

func Info(message string) {
	defaultLogger.Log(PriInfo, message)
}
func Infof(format string, v...interface{}) {
	defaultLogger.Log(PriInfo, fmt.Sprintf(format, v...))
}

func Debug(message string) {
	defaultLogger.Log(PriDebug, message)
}
func Debugf(format string, v...interface{}) {
	defaultLogger.Log(PriDebug, fmt.Sprintf(format, v...))
}
