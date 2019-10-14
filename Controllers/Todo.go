package Controllers

import (
	"awesomeProject/TODO APP/Models"
	"github.com/gin-gonic/gin"
	"net/http"
)
//エラーがあれば404、なければ全todo表示
func GetTodos(c *gin.Context) {
	var todo []Models.Todo
	err := Models.GetAllTodos(&todo)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	}else{
		c.JSON(http.StatusOK, todo)
	}
}
//bind: IPアドレスとドメイン名の対応付け
//create、dbに入れるときbind
//createしたtodo示？
func CreateATodo(c *gin.Context){
	var todo Models.Todo
	c.BindJSON(&todo)
	err := Models.CreateATodo(&todo)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	}else{
		c.JSON(http.StatusOK, todo)
	}
}
//id探す？
//idのtodo表示
func GetATodo(c *gin.Context) {
	id := c.Params.ByName("id")
	var todo Models.Todo
	err := Models.GetATodo(&todo, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	}else{
		c.JSON(http.StatusOK, todo)
	}
}
//id探す
//getatodoとupdateatodoのどちらのエラーも検証
//updateのときbind
func UpdateATodo(c *gin.Context){
	var todo Models.Todo                        //update前のtodo?
	id := c.Params.ByName("id")
	err := Models.GetATodo(&todo, id)
	if err != nil {
		c.JSON(http.StatusNotFound, todo)
	}
	c.BindJSON(&todo)                          //ここからupdate後のtodo?
	err = Models.UpdateATodo(&todo, id)
	if err != nil{
		c.AbortWithStatus(http.StatusNotFound)
	}else{
		c.JSON(http.StatusOK, todo)
	}
}
//idのtodoさがす。
//消せたらメッセージ表示
func DeleteATodo(c *gin.Context) {
	var todo Models.Todo
	id := c.Params.ByName("id")
	err := Models.DeleteATodo(&todo, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	}else {
		c.JSON(http.StatusOK, gin.H{"id:" + id: "deleted"})
	}
}