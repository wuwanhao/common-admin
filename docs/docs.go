// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/captcha": {
            "get": {
                "description": "验证码接口",
                "produces": [
                    "application/json"
                ],
                "summary": "验证码接口",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/result.Result"
                        }
                    }
                }
            }
        },
        "/api/login": {
            "post": {
                "description": "用户登录接口",
                "produces": [
                    "application/json"
                ],
                "summary": "用户登录接口",
                "parameters": [
                    {
                        "description": "data",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/entity.LoginDto"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/result.Result"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "entity.LoginDto": {
            "type": "object",
            "required": [
                "idKey",
                "image",
                "password",
                "username"
            ],
            "properties": {
                "idKey": {
                    "description": "UUID",
                    "type": "string"
                },
                "image": {
                    "description": "验证码",
                    "type": "string",
                    "maxLength": 6,
                    "minLength": 4
                },
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "result.Result": {
            "type": "object",
            "properties": {
                "code": {
                    "description": "状态码",
                    "type": "integer"
                },
                "data": {
                    "description": "返回的数据"
                },
                "message": {
                    "description": "提示信息",
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "通用后台管理系统",
	Description:      "后台管理系统API接口文档",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
