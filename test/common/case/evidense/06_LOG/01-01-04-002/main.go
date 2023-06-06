package main

import (
    "spatial-id/common/logger"

)

func main() {
    message := "LogMessage_Test: %s, %d, %s, %g"
    values := "int"

    logger.Debug(message, values, 10, "float", 10.12345)
}
