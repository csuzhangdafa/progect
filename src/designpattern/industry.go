/*
工厂模式：
1.简单工厂
2.工厂方法
3.抽象工厂

简单工厂模式：
定义一个工厂类，它可以根据参数的不同返回不同类的实例，被创建的实例通常都具有共同的父类。
因为在简单工厂模式用于创建实例的方法是静态的方法，因此简单工厂模式又被称为静态工厂方法模式，它属于类创建型模式。

优点：
1.工厂类含有必要的判断逻辑，可以决定在什么时候创建哪一个产品类的实例
客户端可以免除直接创建产品对象的责任，而仅仅“消费”产品
简单工厂模式通过这种做法实现了对责任的分割，它提供了专门的工厂类用于创建对象
2.客户端无须知道所创建的具体产品类的类名，只需要知道具体产品类所对应的参数即可
3.通过引入配置文件，可以在不修改任何客户端代码的情况下更换和增加新的具体产品类

缺点：
1.由于工厂类集中了所有产品创建逻辑，一旦不能正常工作，整个系统都要受到影响
2.使用简单工厂模式将会增加系统中类的个数，在一定程序上增加了系统的复杂度和理解难度。
3.系统扩展困难，一旦添加新产品就不得不修改工厂逻辑
4.简单工厂模式由于使用了静态工厂方法，造成工厂角色无法形成基于继承的等级结构。

适用范围：
工厂类负责创建的对象比较少，客户只知道传入了工厂类的参数，对于始何创建对象（逻辑）不关心

type apple struct{}

type FruitFactory interface {
	Fruit() string
}

func (*apple) Fruit() string {
	return "我是苹果，我很好吃"
}

type banana struct{}

func (*banana) Fruit() string {
	return "我是香蕉，我最好吃了"
}

func getFruit(t string) FruitFactory {
	switch t {
	case "apple":
		return &apple{}
	case "banana":
		return &banana{}
	}

	return nil
}

func main() {
	f := getFruit("apple")
	fmt.Println(f.Fruit())
}

工厂方法模式
实现了“工厂”概念的面向对象设计模式，它也是处理在不指定对象具体类型的情况下创建对象的问题
工厂方法模式的实质是“定义一个创建对象的接口，但让实现这个接口的类来决定实例化哪个类
工厂方法让类的实例化推迟到子类中进行

优点：
1.一个调用者想创建一个对象，只要知道其名称就可以了。
2.扩展性高，如果想增加一个产品，只要扩展一个工厂类就可以。
3.屏蔽产品的具体实现，调用者只关心产品的接口。

缺点：
每次增加一个产品时，都需要增加一个具体类和对象实现工厂，使得系统中类的个数成倍增加，
在一定程度上增加了系统的复杂度，同时也增加了系统具体类的依赖。

适用范围
当对象的创建逻辑比较复杂，不只是简单的 new 一下就可以，而是要组合其他类对象，做各种初始化操作的时候，推荐使用工厂方法模式
将复杂的创建逻辑拆分到多个工厂类中，让每个工厂类都不至于过于复杂

type Fruit interface{
	Fruit() string
}

type appleFactory struct{}

func (*appleFactory) Fruit() string{
	return "apple"
}

type bananaFactory struct{}

func (*bananaFactory) Fruit() string {
	return "banana"
}

func main() {
	apple := appleFactory{}
	fmt.Println(apple.Fruit())

	banana := bananaFactory{}
	fmt.Println(banana.Fruit())
}

抽象工厂模式
提供一个创建一系列相关或相互依赖对象的接口，而无需指定它们具体的类。

优点：
1.可以在类的内部对产品族进行约束。所谓的产品族，一般或多或少的都存在一定的关联，
抽象工厂模式就可以在类内部对产品族的关联关系进行定义和描述，而不必专门引入一个新的类来进行管理。
2.易于交换产品系列，由于具体的工场类，在使用的时候只需要在应用中初始化一次，所以改变工厂就很简单，
只需要改变具体地工厂就能使用对应的配置信息。
3.它让具体地创建过程和客户端分离，客户端通过他们的抽象接口操作实例，产品的具体类名也和具体工厂分离，
不会出现在客户端代码中。

缺点
抽象工厂模式在于难于应付“新对象”的需求变动。难以支持新种类的产品。难以扩展抽象工厂以生产新种类的产品。
这是因为抽象工厂几乎确定了可以被创建的产品集合，支持新种类的产品就需要扩展该工厂接口，这将涉及抽象工厂类及其所有子类的改变。

适用范围
1.一个系统不应当依赖于产品类实例如何被创建、组合和表达的细节，这对于所有形态的工厂模式都是重要的。
2.这个系统的产品有多于一个的产品族，而系统只消费其中某一族的产品。
3.同属于同一个产品族的产品是在一起使用的，这一约束必须在系统的设计中体现出来。
4.系统提供一个产品类的库，所有的产品以同样的接口出现，从而使客户端不依赖于实现。

package main

import "fmt"

type FruitInterface interface {
	ChooseApple() ProductInterface
	ChooseBanana() ProductInterface
}

type ProductInterface interface {
	Fruit()
}

type HainanApple struct {
}

func (h HainanApple) Fruit() {
	fmt.Println("我是苹果，来自海南")
}

type HainanBanana struct {
}

func (h HainanBanana) Fruit() {
	fmt.Println("我是香蕉，来自海南")
}

type WuhanApple struct {
}

func (w WuhanApple) Fruit() {
	fmt.Println("我是苹果，来自武汉")
}

type WuhanBanana struct {
}

func (w WuhanBanana) Fruit() {
	fmt.Println("我是香蕉，来自武汉")
}

type WuhanFruitFactory struct {
}

func (w WuhanFruitFactory) ChooseApple() ProductInterface {
	return WuhanApple{} //WuhanApple是一个实现了ProductInterface接口类型的结构体
}

func (w WuhanFruitFactory) ChooseBanana() ProductInterface {
	return WuhanBanana{}
}

type HainanFruitFactory struct {
}

func (gd HainanFruitFactory) ChooseApple() ProductInterface {
	return HainanApple{}
}

func (gd HainanFruitFactory) ChooseBanana() ProductInterface {
	return HainanBanana{}
}

func main() {
	f := WuhanFruitFactory{}
	b := f.ChooseApple()
	b.Fruit()
}

*/

package main
