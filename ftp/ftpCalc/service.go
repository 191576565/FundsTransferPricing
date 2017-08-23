// Copyright 2016 huangzhanwei. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package ftpCalc

import (
	"net"
)

//计算引擎接口函数
//首先open打开一个rpc连接
//
//Write方法，向这个rpc指定的地方写入信息
//
//Readf方法，从这个rpc指定地方读取信息
//
//Ioctl方法，用于更新rpc指定地方的属性
//
//Close方法，用于关闭rpc连接
type FtpServ interface {
	Open(string) (net.Conn, error)
	Write(dst interface{}) (int, error)
	Read(dst interface{}) (int, error)
	Ioctl(dst interface{}) (int, error)
	Close() error
}
