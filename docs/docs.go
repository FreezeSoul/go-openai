// Code generated by swaggo/swag. DO NOT EDIT
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
        "/geek/GetCourseType": {
            "get": {
                "description": "查看大课程的类别",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "geek"
                ],
                "summary": "课程大分类",
                "parameters": [
                    {
                        "type": "string",
                        "description": "token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/geek/GetGeekArticle": {
            "post": {
                "description": "查看课程下的章节",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "geek"
                ],
                "summary": "章节查看",
                "parameters": [
                    {
                        "type": "string",
                        "description": "token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "参数",
                        "name": "GeekCourseModel",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/geek.GeekCourseModel"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/geek/GetGeekCourse": {
            "post": {
                "description": "查看大分类下的课程",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "geek"
                ],
                "summary": "课程查看",
                "parameters": [
                    {
                        "type": "string",
                        "description": "token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "参数",
                        "name": "GeekCourseModel",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/geek.GeekCourseModel"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/geek/getArticleContent": {
            "get": {
                "description": "查看章节下的内容",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "geek"
                ],
                "summary": "章节内容查看",
                "parameters": [
                    {
                        "type": "string",
                        "description": "token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "参数",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/log/cmd": {
            "get": {
                "description": "执行命令",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "log"
                ],
                "summary": "根据参数执行命令",
                "responses": {}
            }
        },
        "/log/getListByVersion": {
            "get": {
                "description": "根据版本查询日志",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "log"
                ],
                "summary": "CICD日志",
                "parameters": [
                    {
                        "type": "string",
                        "description": "版本号",
                        "name": "version",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "项目",
                        "name": "project",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/openai/getChatCompletions": {
            "post": {
                "description": "暂时不支持上下文",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "GPT"
                ],
                "summary": "GPT-3.5模型聊天对话(支持openai和azure两种接口)",
                "responses": {}
            }
        },
        "/openai/getCompletions": {
            "post": {
                "description": "不支持stream模式",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "GPT"
                ],
                "summary": "GPT-3.0聊天对话模式(暂时准备等过段时间注释掉了 直接可以使用3.5)",
                "responses": {}
            }
        },
        "/openai/getEmbeddings": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "GPT"
                ],
                "summary": "数据转换，转换为向量数据",
                "responses": {}
            }
        },
        "/openai/getImageGenerations": {
            "post": {
                "description": "暂时写死只能生成一张",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "GPT"
                ],
                "summary": "根据文字描述生成图片",
                "responses": {}
            }
        },
        "/openai/getModels": {
            "post": {
                "description": "列出所有模型",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "GPT"
                ],
                "summary": "openai 所有开放的模型model",
                "responses": {}
            }
        },
        "/openai/getSpeechToText": {
            "post": {
                "description": "限制最大为25M",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "GPT"
                ],
                "summary": "根据上传的语音转换为文字",
                "responses": {}
            }
        },
        "/openai/getUsage": {
            "post": {
                "description": "当前用户的总额度和使用额度",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "GPT"
                ],
                "summary": "用户授信额度使用",
                "responses": {}
            }
        },
        "/user/add": {
            "post": {
                "description": "添加一个用户",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "添加一个用户",
                "responses": {}
            }
        },
        "/user/delete/{id}": {
            "get": {
                "description": "根据传递的id查找来删除用户",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "删除用户",
                "responses": {}
            }
        },
        "/user/list": {
            "get": {
                "description": "根据传分页信息查询用户列表",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "用户列表",
                "responses": {}
            }
        },
        "/user/list/{name}": {
            "get": {
                "description": "根据传递的name查找来用户，可能返回一个或多个用户数据",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "查询用户",
                "responses": {}
            }
        },
        "/user/login": {
            "post": {
                "description": "根据用户的账号和密码",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "登录",
                "parameters": [
                    {
                        "description": "User information",
                        "name": "loginModel",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/user.LoginModel"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/user/update": {
            "post": {
                "description": "根据传递的用户信息进行更新，必须要传递已存在的id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "修改用户",
                "responses": {}
            }
        },
        "/user/{id}": {
            "get": {
                "description": "根据传递的name查找来用户，可能返回一个或多个用户数据",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "查询用户",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "int valid",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {}
            }
        }
    },
    "definitions": {
        "geek.GeekCourseModel": {
            "type": "object",
            "properties": {
                "limit": {
                    "description": "条数",
                    "type": "integer"
                },
                "page": {
                    "description": "页数",
                    "type": "integer"
                },
                "typeId": {
                    "description": "类型",
                    "type": "string"
                }
            }
        },
        "user.LoginModel": {
            "type": "object",
            "properties": {
                "account": {
                    "description": "账号",
                    "type": "string"
                },
                "password": {
                    "description": "密码",
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "0.0.1",
	Host:             "",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "CICD自动化部署 API",
	Description:      "CICD time",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
