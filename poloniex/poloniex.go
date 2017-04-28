package poloniex

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"
)

const (
	API_BASE = "https://poloniex.com/"
	DEFAULT_HTTPCLIENT_TIMEOUT = 30
)
