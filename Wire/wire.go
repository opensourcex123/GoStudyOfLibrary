//+build wireinject

package main

import (
	"github.com/google/wire"
)

// wire是 Google 开源的一个依赖注入工具。它是一个代码生成器，并不是一个框架。
// 我们只需要在一个特殊的go文件中告诉wire类型之间的依赖关系，它会自动帮我们生成代码，帮助我们创建指定类型的对象，并组装它的依赖。

func InitMission(name string) Mission {
	wire.Build(NewMonster, NewPlayer, NewMission)
	return Mission{}
}
