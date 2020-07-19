package controller

import ( 
	"fmt"
	"github.com/valyala/fasthttp"
)

//Interface2String : Allow to covert interface values to string one
func Interface2String(val interface{}) string {
	return fmt.Sprintf("%v", val)
}

//GetParameters : search for given paramenter into entry url and return its values as hash table
func GetParameters(ctx *fasthttp.RequestCtx, parameters []string) map[string]string {

	parametersList := make(map[string]string)

	for i := 0; i < len(parameters); i++ {
		value := Interface2String(ctx.UserValue(parameters[i]))
		if value != "" {
			parametersList[parameters[i]] = value
		}
	}

	return parametersList
}