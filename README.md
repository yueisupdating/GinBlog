- router: 把url转为对应的api接口
- 数据模型: 进行数据库操作，返回数据库错误码
- api: router和数据库之间的中间层,返回报文


- cookie: 服务器创建的记录客户端某些信息的数据结构，保存在浏览器
- session: cookie中包含session_id字段,服务器根据session_id映射到相应的session。一般使用redis集群共享session
- Token: 用户名+密码发给服务器得到token,服务器解析token实现映射到用户
- JWT:head(服务器加密算法)+ payload(过期时间等非敏感信息)+sign(前面二者的签名)
