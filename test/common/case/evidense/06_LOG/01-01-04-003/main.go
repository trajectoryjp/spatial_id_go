package main

import (
    "spatial-id/common/logger"

)

func main() {
    message := "LogMessage_Test: %s, %d, %s, %g"

    logger.Debug(message)
}
