package main

import (
    "spatial-id/common/logger"

)

func main() {
    message := "LogDir_Test: %s"
    values := "exist"

    logger.Debug(message, values)
}
