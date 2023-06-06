package main

import (
    "spatial-id/common/logger"

)

func main() {
    message := "LogDir_Test: %s"
    values := "not exist"
    logger.Debug(message, values)
}
