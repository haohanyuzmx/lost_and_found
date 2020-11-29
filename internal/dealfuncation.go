package internal

import "github.com/gin-gonic/gin"

const (
	Success      = 10000
	ParamWrong   = 20000
	ServiceError = 20001
	DataError    = 50000
	Repeat       = 50001
)

var (
	UnKnownError = ResponseMsg{
		Status: 404,
		Info:   "Unknown error",
	}
	AuthorizedError = ResponseMsg{
		Status: 403,
		Info:   "Unauthorized",
	}
)

type ResponseMsg struct {
	Status int         `json:"status"`
	Info   string      `json:"info"`
	Data   interface{} `json:"data,omitempty"`
}

func getRespMsgByCode(respCode int, data interface{}) (respMsg ResponseMsg) {
	switch respCode {
	case Success:
		return ResponseMsg{
			Status: Success,
			Info:   "success",
			Data:   data,
		}
	case ParamWrong:
		return ResponseMsg{
			Status: ParamWrong,
			Info:   "param wrong",
			Data:   data,
		}
	case ServiceError:
		return ResponseMsg{
			Status: ServiceError,
			Info:   "service error",
			Data:   data,
		}
	case DataError:
		return ResponseMsg{
			Status: DataError,
			Info:   "data error",
			Data:   data,
		}
	case Repeat:
		return ResponseMsg{
			Status: Repeat,
			Info:   "repeat",
			Data:   data,
		}

	default:
		respMsg = UnKnownError
	}
	return
}

func definedResp(status int, c *gin.Context, resp ResponseMsg) {
	c.JSON(status, resp)
}
func definedSuccess(c *gin.Context) {
	c.JSON(200, getRespMsgByCode(Success, nil))
}
