1 先创建数据库k8s_test；然后导入sql
2 flow.properties文件需要通过configmap挂载（选择配置文件类型），挂载的容器路径到opt/app/config/flow.properties
更改默认的svcname为实际环境创建的mysql的svcname,mysql的用户名和密码更改为实际环境下mysql服务的用户名和密码
3 demo访问 路径/itemapp,页面输入名词、价格与数据库进行交互