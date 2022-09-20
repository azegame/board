package controller

import (
    "fmt"
    "strconv"
    
    "github.com/gin-gonic/gin"

    "board/model/repository"
)


func indexPage(c *gin.Context) {
	allThreads, err := repository.GetAllThreads()
	if err != nil {
		fmt.Println(err)
	}

    c.HTML(200, "index.html", gin.H{
    	"threads": allThreads
    })
}


func createThread(c *gin.Context) {
	title := c.PostForm("title")
	username := c.PostForm("username")
	message := c.PostForm("message")

	allThreads, error := repository.GetAllThreads()
	if error != nil {
		fmt.Println(error)
	}

	if title == "" || username == "" || message == "" {
		c.HTML(200, "index.html", gin.H{
			"threads": allThreads,
			"error": "未入力項目があります",
		})
		return
	}

	err := repository.InsertThreadInfo(title)

	//スレ名を一意にするためのハンドリング
	if err != nil {
    	c.HTML(200, "index.html", gin.H{
    		"threads": allThreads,
    		"error": err,
    	})
		return
	}

	thread, err := repository.GetThreadByTitle(title)
	if err != nil {
		fmt.Println(err)
	}

	threadId := thread.ThreadId
	threadTitle := thread.Title

	resNo, err := repository.GetResponseNo(threadId)
	if err != nil {
		fmt.Println(err)
	}

	err = repository.InsertResponseInfo(threadId, username, message, resNo)
	if err != nil {
		fmt.Println(err)
	}

	responses, err := repository.GetResponse(threadId)
	if err != nil {
		fmt.Println(err)
	}

	resQty, err := repository.CountResponse(threadId)
	if err != nil {
		fmt.Println(err)
	}

    c.HTML(200, "thread.html", gin.H{
    	"title": threadTitle,
    	"responses": responses,
    	"threadId": threadId,
    	"qty": resQty,
    })
}


func getThreadPage(c *gin.Context) {
	threadId, _ := strconv.Atoi(c.Param("ThreadId"))

	thread, err := repository.GetThreadByThreadId(threadId)
	if err != nil {
		fmt.Println(err)
	}

	threadTitle := thread.Title
	responses, err := repository.GetResponse(threadId)
	if err != nil {
		fmt.Println(err)
	}

	resQty, err := repository.CountResponse(threadId)
	if err != nil {
		fmt.Println(err)
	}

	c.HTML(200, "thread.html", gin.H{
		"title": threadTitle,
		"responses": responses,
		"threadId": threadId,
		"qty": resQty,
	})
}


func postResponse(c *gin.Context) {
	threadId, _ := strconv.Atoi(c.PostForm("threadid"))
	username := c.PostForm("username")
	message := c.PostForm("message")

	thread, err := repository.GetThreadByThreadId(threadId)
	if err != nil {
		fmt.Println(err)
	}

	threadTitle := thread.Title

	if username == "" || message == "" {
		responses, err := repository.GetResponse(threadId)
		if err != nil {
			fmt.Println(err)
		}
	
		resQty, err := repository.CountResponse(threadId)
		if err != nil {
			fmt.Println(err)
		}
	
		c.HTML(200, "thread.html", gin.H{
			"title": threadTitle,
			"responses": responses,
			"threadId": threadId,
			"qty": resQty,
			"error": "未入力項目があります",
		})
		return
	}

	resNo, err := repository.GetResponseNo(threadId)
	if err != nil {
		fmt.Println(err)
	}

	err = repository.InsertResponseInfo(threadId, username, message, resNo)
	if err != nil {
		fmt.Println(err)
	}

	responses, err := repository.GetResponse(threadId)
	if err != nil {
		fmt.Println(err)
	}

    c.HTML(200, "thread.html", gin.H{
    	"title": threadTitle,
    	"responses": responses,
    	"threadId": threadId,
    	"qty": resNo,
    })
}



