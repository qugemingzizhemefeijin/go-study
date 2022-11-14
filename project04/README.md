### Go Web 编程 - 编程源码

知识点1：

当导入一个包时，该包下的文件里所有`init()`函数都会被执行，然而，有些时候我们并不需要把整个包都导入进来，仅仅是是希望它执行`init()`函数而已。这个时候就可以使用`import _`引用该包。即使用`import _`包路径只是引用该包，仅仅是为了调用`init()`函数，所以无法通过包名来调用包中的其他函数。

知识点2：

| 字段 | 需要调用的方法或需要访问的字段 | 键值对来源 | 内容类型 |
| --- | --- | --- | --- |
| Form | ParseForm方法 | URL，表单 | URL编码 |
| PostForm | Form字段 | 表单 | Url编码 |
| MultipartForm | PareMultipartForm方法 | 表单 | Multipart编码 |
| FormValue | 无 | URL，表单 | URL编码 |
| PostFormValue | 无 | 表单 | URL编码 |
