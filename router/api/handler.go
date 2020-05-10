package api

import (
	"github.com/goalong/hot-comment/router/e"
	"net/http"
	"strconv"
)
import "github.com/gin-gonic/gin"
import "github.com/goalong/hot-comment/search"

type Page struct {
	PageNum int `form:"page_num" binding:"required,max=30,min=1"`
	//PageSize int `form:"page_size" binding:"required,max=50,min=1"`
}

func GetPage(c *gin.Context) (pageNum, pageSize, code int) {
	code = e.SUCCESS
	var page Page
	err := c.ShouldBindQuery(&page)
	if err != nil {
		code = e.INVALID_PAGE
	}
	return page.PageNum, 20, code

}

// @Summary 热门评论，按点赞数排行
// @Produce  json
// @Param page_num query int true "页码"
// @Success 200 {string} string "{"code":200,"data":{},"msg":""}"
// @Router /api/hot_comments [get]
func GetHotComments(c *gin.Context) {
	data := make(map[string]interface{})
	pageNum, pageSize, code := GetPage(c)
	if code != e.SUCCESS {
		c.JSON(http.StatusOK, gin.H{
			"code": code,
			"msg":  e.GetMsg(code),
			"data": data,
		})
		return

	}

	data["items"], code = search.GetCommentsByLikeCount(pageNum, pageSize)

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  "",
		"data": data,
	})
}

// @Summary 获取热门歌曲，按评论数排行
// @Produce  json
// @Param page_num query int true "页码"
// @Success 200 {string} string "{"code":200,"data":{},"msg":""}"
// @Router /api/hot_songs [get]
func GetHotSongs(c *gin.Context) {
	data := make(map[string]interface{})
	pageNum, pageSize, code := GetPage(c)
	if code != e.SUCCESS {
		c.JSON(http.StatusOK, gin.H{
			"code": code,
			"msg":  e.GetMsg(code),
			"data": data,
		})
		return
	}

	data["items"], code = search.GetSongsByCommentCount(pageNum, pageSize)

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  "",
		"data": data,
	})
}

// @Summary 按关键词搜索歌曲
// @Produce  json
// @Param page_num query int true "页码"
// @Param keyword query string true "关键词"
// @Success 200 {string} string "{"code":200,"data":{},"msg":""}"
// @Router /api/search/songs [get]
func SearchSongs(c *gin.Context) {
	data := make(map[string]interface{})
	pageNum, pageSize, code := GetPage(c)
	if code != e.SUCCESS {
		c.JSON(http.StatusOK, gin.H{
			"code": code,
			"msg":  e.GetMsg(code),
			"data": data,
		})
		return
	}

	keyword := c.Query("keyword")
	data["items"], code = search.SearchSong(keyword, pageNum, pageSize)

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  "",
		"data": data,
	})

}

// @Summary 按关键词搜索评论
// @Produce  json
// @Param page_num query int true "页码"
// @Param keyword query string true "关键词"
// @Success 200 {string} string "{"code":200,"data":{},"msg":""}"
// @Router /api/search/comments [get]
func SearchComments(c *gin.Context) {
	data := make(map[string]interface{})
	pageNum, pageSize, code := GetPage(c)
	if code != e.SUCCESS {
		c.JSON(http.StatusOK, gin.H{
			"code": code,
			"msg":  e.GetMsg(code),
			"data": data,
		})
		return
	}

	keyword := c.Query("keyword")
	data["items"], code = search.SearchComment(keyword, pageNum, pageSize)

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  "",
		"data": data,
	})

}

// @Summary 搜索歌手
// @Produce  json
// @Param page_num query int true "页码"
// @Param keyword query string true "关键词"
// @Success 200 {string} string "{"code":200,"data":{},"msg":""}"
// @Router /api/search/artists [get]
func SearchArtists(c *gin.Context) {
	data := make(map[string]interface{})
	pageNum, pageSize, code := GetPage(c)
	if code != e.SUCCESS {
		c.JSON(http.StatusOK, gin.H{
			"code": code,
			"msg":  e.GetMsg(code),
			"data": data,
		})
		return
	}

	keyword := c.Query("keyword")
	data["items"], code = search.SearchArtist(keyword, pageNum, pageSize)

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  "",
		"data": data,
	})

}

// @Summary 根据歌手ID获取歌曲列表
// @Produce  json
// @Param page_num query int true "页码"
// @Param artist_id query string true "歌手ID"
// @Success 200 {string} string "{"code":200,"data":{},"msg":""}"
// @Router /api/songs [get]
func GetSongsByArtistId(c *gin.Context) {
	data := make(map[string]interface{})
	pageNum, pageSize, code := GetPage(c)
	artistId, err := strconv.Atoi(c.Query("artist_id"))
	if err != nil {
		code = e.INVALID_PARAMS
	}
	if code != e.SUCCESS {
		c.JSON(http.StatusOK, gin.H{
			"code": code,
			"msg":  e.GetMsg(code),
			"data": data,
		})
		return

	}
	data["items"], code = search.GetSongsByArtistId(artistId, pageNum, pageSize)

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  "",
		"data": data,
	})

}

// @Summary 根据歌曲ID获取评论列表
// @Produce  json
// @Param page_num query int true "页码"
// @Param song_id query string true "歌曲ID"
// @Success 200 {string} string "{"code":200,"data":{},"msg":""}"
// @Router /api/comments [get]
func GetCommentsBySongId(c *gin.Context) {
	data := make(map[string]interface{})
	pageNum, pageSize, code := GetPage(c)
	songId, err := strconv.Atoi(c.Query("song_id"))
	if err != nil {
		code = e.INVALID_PARAMS
	}
	if code != e.SUCCESS {
		c.JSON(http.StatusOK, gin.H{
			"code": code,
			"msg":  e.GetMsg(code),
			"data": data,
		})
		return
	}
	data["items"], code = search.GetCommentsBySongId(songId, pageNum, pageSize)
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  "",
		"data": data,
	})

}