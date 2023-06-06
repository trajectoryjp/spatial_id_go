package main

import (
    "spatial-id/common/logger"

)

func main() {
    message := "LogFile_Test: %s"
    values := "exist"

    logger.Debug(message, values)
}
