## 6.2

在声明方法时，类型本身的指针是不允许出现在接收器中的
```
type P *int
func (P) f() {}  // invalid receiver type P (pointer or interface type)
```

不能通过一个无法取到地址的接收器来调用指针方法，比如临时变量的地址无法获取
> Point{1, 2}.ScaleBy(2) // compile error: can't take address of Point literal


