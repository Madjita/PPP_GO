package main

import (
    "fmt"
)


//------------------------------------------------------------------
//Enum перечисление
// значения локальных ключей для каждой из фигур
type key int
const(
     RECTANGLE key = 0
     TRIANGLE      = 1
	 CIRCL 		   = 2
)
//------------------------------------------------------------------

//------------------------------------------------------------------
//Указатель на любой тип данных
// используется универсальный указатель
type Figure interface{
}
//------------------------------------------------------------------

//------------------------------------------------------------------
// Использование параметрического обобщения, построенного на
// основе косвенного альтернативного связывания специализаций.
// Структура, обобщающая все имеющиеся фигуры
type Shape struct
{
	k key;	   // ключ
	ptr Figure // подключается любая специализация
}

//------------------------------------------------------------------
//Специализации
type Rectangle struct
{
	a float64
	b float64
}
type Triangle struct
{
	a float64
	b float64
	c float64
}

type Circl struct
{
	r float64
}
//------------------------------------------------------------------

//------------------------------------------------------------------
// Инициализация прямоугольника
func Create_rectangle( a,b float64) *Rectangle {
	item := new(Rectangle)
	item.a = a
	item.b = b
	return  item
}
//------------------------------------------------------------------
// Инициализация треугольника
func Create_triangle( a,b,c float64) *Triangle {
	item := new(Triangle)
	item.a = a
	item.b = b
	item.c = c
	return  item
}
//------------------------------------------------------------------
// Инициализация круга
func Create_circl( r float64) *Circl {
	item := new(Circl)
	item.r = r
	return  item
}
//------------------------------------------------------------------

//------------------------------------------------------------------
// Динамическое создание обобщенного прямоугольника
func Create_shape_rectangle( a,b float64) *Shape {
    s := new(Shape)
	s.k = RECTANGLE;
    s.ptr = Create_rectangle(a,b)
    return s;
}
//------------------------------------------------------------------
// Динамическое создание обобщенного треугольника
func Create_shape_triangle(a,b,c float64) *Shape {
    s := new(Shape)
	s.k = TRIANGLE;
    s.ptr = Create_triangle(a,b,c)
    return s;
}
//------------------------------------------------------------------
// Динамическое создание обобщенного круга
func Create_shape_circl(r float64) *Shape {
    s := new(Shape)
	s.k = CIRCL;
    s.ptr = Create_circl(r)
    return s;
}

//------------------------------------------------------------------
// Дополнительное переопределение процедуры вывода какая фигу используется
// обобщенной фигуры, сделанное для сокрытия ее реального вида.
// Может отсутствовать.
func Out(fp *Shape) string{

	listPoint := []func() string{OutRectangle,OutTriangle,OutCircl}
	ans := listPoint[fp.k]();
	return ans
}

//------------------------------------------------------------------
// Дополнительное переопределение процедуры вычисления площади
// обобщенной фигуры, сделанное для сокрытия ее реального вида.
// Может отсутствовать.
func Square(item *Shape) {

	listPoint := []func(item *Shape) float64{square_rectangle_of_shape,square_triangle_of_shape,square_circl_of_shape}

	ans := listPoint[item.k](item)
	name := Out(item)

    fmt.Printf("\nSquare %s = %f\n",name,ans) 
}

//------------------------------------------------------------------
// Дополнительное переопределение процедуры вычисления периметра
// обобщенной фигуры, сделанное для сокрытия ее реального вида.
// Может отсутствовать.
func Perimeter(item *Shape) {

	listPoint := []func(item *Shape) float64{perimeter_rectangle_of_shape,perimeter_triangle_of_shape,perimeter_circl_of_shape}

	ans := listPoint[item.k](item)
	name := Out(item)

    fmt.Printf("\nPerimeter %s = %f\n",name,ans) 
}

//------------------------------------------------------------------
//Главная функция Создает 3 объекта обобщения:
// Прямоугольник, треугольник, круг
// 1) Выводи на экран созданные специализации
// Out(...) -- Обобщенная процедура возвращает имя специализации созданные чере обобщение 
// 2) Выыодит на экран расчитанную площадь каждой специализации
// Square(...) -- Обобщенная процедура расчитывающая чере обобщение площадь специализации
// 2) Выыодит на экран расчитанную площадь каждой специализации
// Perimeter(...) -- Обобщенная процедура расчитывающая чере обобщение периметр специализации
func main() {

	rec := Create_shape_rectangle(1,2)
	tri := Create_shape_triangle(1,2,3)
	cir := Create_shape_circl(4)
	
	fmt.Printf("\n=====Create type data=====\n")
	fmt.Printf("\n%s, %s, %s\n",Out(rec),Out(tri),Out(cir))
	
	fmt.Printf("\n=====Next Square=====\n") 

	Square(rec);
	Square(tri);
	Square(cir);

	fmt.Printf("\n=====Next Perimetr=====\n") 

	Perimeter(rec);
	Perimeter(tri);
	Perimeter(cir);

	fmt.Printf("\n=====Exit=====\n") 
}


////------------------------------------------------------------------
//Сигнатуры требуемых функций вывода имени специализации
func OutRectangle() string {
    return "Rectangle"
}

func OutTriangle() string {
	return "Triangle"
}

func OutCircl() string {
	return "Circl"
}

//------------------------------------------------------------------
// Сигнатуры требуемых функций
func square_rectangle(item *Rectangle) float64 {
	s := item.a * item.b
    return s
}

func square_triangle(item *Triangle) float64 {
	s := 1/2*(item.a + item.b + item.c)
    return s
}

func square_circl(item *Circl) float64 {
	s := 3.14*item.r
    return s
}

//------------------------------------------------------------------
// Обработчикики специализаций

// предназначенный для вычисления площади
// прямоугольника. Используется как элемент параметрического массива
// в процедуре вычисления площади обобщенной фигуры.
func square_rectangle_of_shape(item *Shape) float64 {
    return square_rectangle(item.ptr.(*Rectangle))
}

//предназначенный для вычисления площади
// треугольника. Используется как элемент параметрического массива
// в процедуре вычисления площади обобщенной фигуры.
func square_triangle_of_shape(item *Shape) float64 {
	return square_triangle(item.ptr.(*Triangle))
}

//предназначенный для вычисления площади
// круга. Используется как элемент параметрического массива
// в процедуре вычисления площади обобщенной фигуры.
func square_circl_of_shape(item *Shape) float64 {
	return square_circl(item.ptr.(*Circl))
}

//------------------------------------------------------------------
// Сигнатуры требуемых функций
func perimeter_rectangle(item *Rectangle) float64 {
	perimeter := 2*(item.a * item.b)
    return perimeter
}

func perimeter_triangle(item *Triangle) float64 {
	perimeter := item.a + item.b + item.c
    return perimeter
}

func perimeter_circl(item *Circl) float64 {
	perimeter := 2*3.14*item.r
    return perimeter
}
//------------------------------------------------------------------
// Обработчикики специализаций

// предназначенный для вычисления периметра
// прямоугольника. Используется как элемент параметрического массива
// в процедуре вычисления периметра обобщенной фигуры.
func perimeter_rectangle_of_shape(item *Shape) float64 {
    return perimeter_rectangle(item.ptr.(*Rectangle))
}

// предназначенный для вычисления периметра
// треугольника. Используется как элемент параметрического массива
// в процедуре вычисления периметра обобщенной фигуры.
func perimeter_triangle_of_shape(item *Shape) float64 {
	return square_triangle(item.ptr.(*Triangle))
}

// предназначенный для вычисления периметра
// круга. Используется как элемент параметрического массива
// в процедуре вычисления периметра обобщенной фигуры.
func perimeter_circl_of_shape(item *Shape) float64 {
	return perimeter_circl(item.ptr.(*Circl))
}
/////////////////////////////////////////