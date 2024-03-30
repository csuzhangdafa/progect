//抽象工厂模式的一般流程如下：

//定义抽象产品接口：
//首先，需要定义一组抽象产品接口，这些接口描述了产品的共同特点，例如创建按钮、文本框等。

//实现具体产品类：
//然后，针对每个抽象产品接口，创建具体的产品类，这些类负责实现不同操作系统或不同风格下的具体产品功能。

//定义抽象工厂接口：
//接着，需要定义抽象工厂接口，该接口包含多个方法，用于创建一组相关的产品。

//实现具体工厂类：
//创建具体工厂类，实现抽象工厂接口中定义的方法。每个具体工厂类负责创建特定系列的具体产品。

//客户端调用：
//最后，客户端通过抽象工厂接口来获取具体产品，而不需要关心具体的产品是如何创建的。
//客户端代码应该依赖于抽象工厂接口，而不是具体的产品类。
//总的来说，抽象工厂模式通过封装一组相关产品的创建，实现了产品系列的独立性和系统的扩展性，
//同时遵循了依赖倒转原则和开闭原则。


package main

import (
	"fmt"
)
//假设我们有一个跨平台应用程序，需要在 Windows 和 macOS 系统上显示按钮和文本框。
//我们将使用抽象工厂模式来实现这个需求。首先，定义抽象产品接口 Button 和 TextBox
// Abstract Product: Button interface
type Button interface {
    Paint()
}

// Abstract Product: TextBox interface
type TextBox interface {
    Display()
}


//然后，定义具体产品类在 Windows 和 macOS 下的实现

// Concrete Product: WindowsButton
type WindowsButton struct{}

func (w *WindowsButton) Paint() {
    fmt.Println("Painting a Windows button")
}

// Concrete Product: WindowsTextBox
type WindowsTextBox struct{}

func (w *WindowsTextBox) Display() {
    fmt.Println("Displaying a Windows text box")
}

// Concrete Product: MacOSButton
type MacOSButton struct{}

func (m *MacOSButton) Paint() {
    fmt.Println("Painting a macOS button")
}

// Concrete Product: MacOSTextBox
type MacOSTextBox struct{}

func (m *MacOSTextBox) Display() {
    fmt.Println("Displaying a macOS text box")
}


//接下来，定义抽象工厂接口 GUIFactory
// Abstract Factory: GUIFactory interface
type GUIFactory interface {
    CreateButton() Button
    CreateTextBox() TextBox
}


//然后，实现具体工厂类 WindowsFactory 和 MacOSFactory
// Concrete Factory: WindowsFactory
type WindowsFactory struct{}

func (w *WindowsFactory) CreateButton() Button {
    return &WindowsButton{}
}

func (w *WindowsFactory) CreateTextBox() TextBox {
    return &WindowsTextBox{}
}

// Concrete Factory: MacOSFactory
type MacOSFactory struct{}

func (m *MacOSFactory) CreateButton() Button {
    return &MacOSButton{}
}

func (m *MacOSFactory) CreateTextBox() TextBox {
    return &MacOSTextBox{}
}

//最后，客户端代码可以通过抽象工厂来创建不同操作系统下的按钮和文本框
func main() {
    windowsFactory := &WindowsFactory{}
    windowsButton := windowsFactory.CreateButton()
    windowsTextBox := windowsFactory.CreateTextBox()

    macosFactory := &MacOSFactory{}
    macosButton := macosFactory.CreateButton()
    macosTextBox := macosFactory.CreateTextBox()

    // 使用 Windows 界面元素
    windowsButton.Paint()
    windowsTextBox.Display()

    // 使用 macOS 界面元素
    macosButton.Paint()
    macosTextBox.Display()
}