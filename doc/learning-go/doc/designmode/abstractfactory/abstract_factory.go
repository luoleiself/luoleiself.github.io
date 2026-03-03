package main

import (
	"fmt"
)

func main() {
	fmt.Println("抽象工厂模式: 提供一个创建一系列相关或相互依赖对象的接口, 无须指定他们具体的结构体, 由实现此接口的结构体决定需要创建的对象.")
	fmt.Println("优点:")
	fmt.Println("\t1. 当需要产品族时, 抽象工厂可以保证客户始终只使用一个产品的产品族")
	fmt.Println("\t2. 增强了程序的可扩展性, 对于新产品族的增加, 只需要实现一个新的具体工厂即可, 符合OCP(开闭原则)")
	fmt.Println("缺点:")
	fmt.Println("\t1. 规定了所有可能被创建的产品集合, 产品族中扩展新的产品困难, 需要修改抽象工厂的接口")
	fmt.Println("\t2. 增加了系统的抽象性和理解难度")
	fmt.Println(`
// 重点, 创建所有可能被创建的产品集合接口, 所有产品的工厂结构体都实现此接口, 对扩展不易(如果增加展示源码则需要修改此处)
type abstractCourseFactory interface {
    createVideo() abstractCourseFactory
    createNote() abstractCourseFactory
}`)
	fmt.Println("------------------------------------")

	fmt.Println(`
type JavaVideo struct{}

func (j *JavaVideo) record() {
    fmt.Println("录制 java 视频")
}

type JavaNote struct{}

func (j *JavaNote) edit() {
    fmt.Println("编写 java 笔记")
}

// 结构体 JavaCourseFactory 工厂  实现 abstractCourseFactory
type JavaCourseFactory struct{}

func (j *JavaCourseFactory) createVideo() *JavaVideo {
    // 接口为引用类型, 如果使用链式调用方法, 此处则需要返回指针类型
    return &JavaVideo{}
}
func (j *JavaCourseFactory) createNote() *JavaNote {
    // 接口为引用类型, 如果使用链式调用方法, 此处则需要返回指针类型
    return &JavaNote{}
}
var jcf = JavaCourseFactory{}
jcf.createVideo().record()
jcf.createNote().edit()`)
	fmt.Println("---------------------")
	var jcf = JavaCourseFactory{}
	jcf.createVideo().record()
	jcf.createNote().edit()
	fmt.Println("------------------------------------")

	fmt.Println(`
type PythonVideo struct{}

func (p *PythonVideo) record() {
    fmt.Println("录制 python 视频")
}

type PythonNote struct{}

func (p *PythonNote) edit() {
    fmt.Println("编写 python 笔记")
}

// 结构体 PythonCourseFactory 工厂  实现 abstractCourseFactory
type PythonCourseFactory struct{}

func (p *PythonCourseFactory) createVideo() PythonVideo {
    return PythonVideo{}
}

func (p *PythonCourseFactory) createNote() PythonNote {
    return PythonNote{}
}
var pcf = PythonCourseFactory{}
var pv = pcf.createVideo()
pv.record()
var pn = pcf.createNote()
pn.edit()`)
	fmt.Println("---------------------")
	var pcf = PythonCourseFactory{}
	var pv = pcf.createVideo()
	pv.record()
	var pn = pcf.createNote()
	pn.edit()
	fmt.Println("------------------------------------")
}

type IVideo interface {
	record()
}
type INote interface {
	edit()
}

// 重点, 创建所有可能被创建的产品集合接口, 所有产品的工厂结构体都实现此接口, 对扩展不易
type abstractCourseFactory interface {
	createVideo() abstractCourseFactory
	createNote() abstractCourseFactory
}

type JavaVideo struct{}

func (j *JavaVideo) record() {
	fmt.Println("录制 java 视频")
}

type JavaNote struct{}

func (j *JavaNote) edit() {
	fmt.Println("编写 java 笔记")
}

// 结构体 JavaCourseFactory 工厂  实现 abstractCourseFactory
type JavaCourseFactory struct{}

func (j *JavaCourseFactory) createVideo() *JavaVideo {
	// 接口为引用类型, 如果使用链式调用方法, 此处则需要返回指针类型
	return &JavaVideo{}
}
func (j *JavaCourseFactory) createNote() *JavaNote {
	// 接口为引用类型, 如果使用链式调用方法, 此处则需要返回指针类型
	return &JavaNote{}
}

type PythonVideo struct{}

func (p *PythonVideo) record() {
	fmt.Println("录制 python 视频")
}

type PythonNote struct{}

func (p *PythonNote) edit() {
	fmt.Println("编写 python 笔记")
}

// 结构体 PythonCourseFactory 工厂  实现 abstractCourseFactory
type PythonCourseFactory struct{}

func (p *PythonCourseFactory) createVideo() PythonVideo {
	return PythonVideo{}
}

func (p *PythonCourseFactory) createNote() PythonNote {
	return PythonNote{}
}
