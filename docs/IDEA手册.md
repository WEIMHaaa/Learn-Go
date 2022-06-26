1. 添加模板
- File > Settings > Editor > File and Code Templates > Includes > File Header 添加如下模板：
```
/**
  *@author          weimenghua
  *@time            ${YEAR}-${MONTH}-${DAY} ${TIME}
  *@description
  */
```
- File > Settings > Editor > File and Code Templates > Files > 添加如下模板 (Extension: sh)
```
#!/bin/bash

#-----------------------------------------------------------------------------
# author: wmh
# 脚本说明:
#-----------------------------------------------------------------------------
```


2. 快捷键
- 左移：tab+shift
- 右移：tab
- 上移：shift+alt+向上方向键
- 下移：shift+alt+向下方向键


3. 插件
- File Expander：查看jar包
- GitToolBox：查看项目提交情况
- Translation：翻译
- arthas idea：Java在线诊断工具
- Search In Repository：查看jar包的maven/gradle的坐标
- Zoolytic：查看zookeeper节点数据, 合适：ip:port(一般2181)
- Git Commit Template：commit规范
- RestfulToolkit：RestfulToolkit提供了一套 RESTful 服务开发辅助工具集, 根据 URI 的部分信息来查找对应的Controller 中方法
