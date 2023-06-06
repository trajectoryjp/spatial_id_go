package main

import (
    "spatial-id/common/logger"

)

func main() {
    message := "LogLevel_Test: %s"
    values := "Debug"

    logger.Debug(message, values)
}
