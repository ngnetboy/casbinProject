# casbinProject
casbin的使用示例
使用了 gin + gorm 框架
casbin 的 policy 使用 gorm-adapter 存储到数据库中。
项目初始化的时候已经插入了两个用户 —— 用户名：admin； role：admin； 密码：123456；  用户名：anonymous； role：anonymous； 密码：123456
初始化一条 policy —— "admin" "/*" "(GET)|(POST)|(DELETE)"
示例中提供了操作policy的接口，可以获取policy的所有信息，按照指定 role 删除 policy。
