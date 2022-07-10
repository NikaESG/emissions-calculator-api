## API 可以考虑的方向

1. 可以入参是两部分：

- 表达式
- 表达式中变量的内容

通过括号进行优先级

experssion: ((A+B)*C)
variable: A: xxxx, B: xxx, C: xx

2. 创建每一级的表达式，但是每一个表达式必须简单

同样的需要创建((A+B)*C)

- 创建 Name1： A+B
- 创建 Name2： Name1*C

