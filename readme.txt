    由于电脑内存硬件不足，下载了GOLANG工具（与 android studio很像），但一直卡顿不能运行，所以所学的go语言程序都用cmd运行测试。

    学习了两三天go语言，以下是学习任务进度。

    1、先配置完成go语言的环境。
    2、go基础语法 数据类型 变量 常量简单过了一遍：helloworld.go   basic_structure.go
    3、函数 和 切片 ： text_function.go   slice.go
    4、结构体 与 Map集合（在json解析传输中很重要，数据映射在结构体）: sturcture.go   map_create_check.go  map_delete.go
    5.接口（所有的具有共性的方法定义在一起）：interface.go  
    6.单元测试：
        （最简单的）在sum.go中只写入测试函数，在sum_test.go中运用sum.go的函数 与实际结果进行比较 （go test -v sum_test.go sum.go）
      (带有github包的)：path.go  path_test.go  意思同上(大概是文件没有放在gopath路径中，报错)
    7. json解析（定义sturct，json.Unmarshal([]byte(str), &s) ）：json.go

    关于GIN WEB由于时间和电脑的问题，这边用我的毕业设计的一部分代码做参考，是基于Android studio开发的，在android_json文件夹中，包含了res界面activity_ .xml ， bean类 Photo.java , 适配器adapter_mine_showtime.java ， 主代码（含有json解析）Mine_showtime.java  Mine_showtime_jilu.java ，数据url MyConfig.java。  showphoto.png是这部分代码展示
    .mp4是我的毕业设计项目展示，有兴趣可以看看！ https://pan.baidu.com/s/1DpZyb57YLgZ7hSB33L80QA