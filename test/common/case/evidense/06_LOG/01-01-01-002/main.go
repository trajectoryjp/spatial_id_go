package main

import (
    "spatial-id/common/logger"

)

func main() {
    message := "LogLevel_Test: %s"
    values := "Info"

    logger.Debug(message, values)
}
