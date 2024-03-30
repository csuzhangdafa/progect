package severce
import(
	"fmt"
)

type student struct{
	Name string
	score float64
}


//构建一个函数来访问私有属性
func NewStudent(n string , s float64) *student{
	return &student{
		Name : n,
		score : s,
	}
}

//简单工厂模式


//如果score字段首字母小写，则，在其他包不可以直接方法，我们可以提供一个方法来访问结构体私有属性
//提供一个公有的方法
func (s *student) Getscore () float64{
    return s.score
}


// Shape 接口定义了形状对象需要实现的方法
type Shape interface {
    Draw()
}

// Circle 结构体表示圆形对象
type Circle struct{}

// Draw 是 Circle 结构体实现的 Shape 接口方法
func (c Circle) Draw() {
    fmt.Println("Draw a circle")
}

// Square 结构体表示正方形对象
type Square struct{}

// Draw 是 Square 结构体实现的 Shape 接口方法
func (s Square) Draw() {
    fmt.Println("Draw a square")
}

// ShapeFactory 是一个工厂函数，根据传入的参数创建不同类型的形状对象
func ShapeFactory(shapeType string) Shape {
    if shapeType == "circle" {
        return Circle{}
    } else if shapeType == "square" {
        return Square{}
    }
    return nil
}

func main() {
    // 使用工厂函数创建圆形对象
    circle := ShapeFactory("circle")
    circle.Draw()

    // 使用工厂函数创建正方形对象
    square := ShapeFactory("square")
    square.Draw()
}