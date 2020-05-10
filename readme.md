### 云音乐歌曲、评论搜索API
使用Go、Gin、Elasticsearch开发的一个网易云音乐歌曲、评论搜索API，可以在web上点击发请求哦，能查到数据，不过当然不是全部的数据啦。

### 目前的接口
* 根据关键词搜索歌手
* 根据歌手id获取歌曲列表
* 根据歌曲id获取评论列表
* 通过关键词搜索歌曲
* 通过关键词搜索评论
* 获取热门歌曲，按评论数排行
* 获取热门评论，按点赞数排行

[这里](http://47.99.131.182/swagger/index.html#)是swagger的web页面，可以在上面随便点击一个接口，点右边的Try it out, 然后输入参数，一般是页码和关键字，点击Execute来发请求，就可以看到接口的结果了。

最下面那张截图是获取热门评论的接口的结果，字段like_count是点赞量，可以看到排名第一的评论点赞量是一百五十多万了。

Enjoy it.


![](https://i.loli.net/2020/05/10/3PFqmcJxawsijE1.png)

![](https://i.loli.net/2020/05/10/HOTuj87wnpohLqV.png)

