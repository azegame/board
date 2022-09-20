package main

import (
    "board/controller"  
)


func main() {
    router := controller.GetRouter()
    router.Run()
}



