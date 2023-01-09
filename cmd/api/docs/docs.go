// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "API Support",
            "url": "sethukumarj.com",
            "email": "sethukumarj.76@gmail.com"
        },
        "license": {
            "name": "MIT",
            "url": "https://opensource.org/licenses/MIT"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/admin/approveevent": {
            "patch": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Admin"
                ],
                "summary": "approves the event for admin",
                "operationId": "approves event",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Event Name : ",
                        "name": "title",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    }
                }
            }
        },
        "/admin/listEvents": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Admin"
                ],
                "summary": "list all upcoming events for admin",
                "operationId": "list all upcoming events",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Page number: ",
                        "name": "page",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Page capacity : ",
                        "name": "pagesize",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "boolean",
                        "description": "List event by approved non approved : ",
                        "name": "approved",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    }
                }
            }
        },
        "/admin/listUsers": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Admin"
                ],
                "summary": "list all active users for admin",
                "operationId": "list all active users",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Page number: ",
                        "name": "page",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Page capacity : ",
                        "name": "pagesize",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    }
                }
            }
        },
        "/admin/login": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Admin",
                    "Admin"
                ],
                "summary": "Login for Admin",
                "operationId": "Admin Login",
                "parameters": [
                    {
                        "type": "string",
                        "description": "admin email: ",
                        "name": "email",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "admin password: ",
                        "name": "password",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    }
                }
            }
        },
        "/admin/signup": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Admin",
                    "Admin"
                ],
                "summary": "SignUp for Admin",
                "operationId": "SignUp authentication",
                "parameters": [
                    {
                        "description": "admin signup with username, phonenumber email ,password",
                        "name": "RegisterAdmin",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.Admins"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    }
                }
            }
        },
        "/admin/token/refresh": {
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Admin",
                    "Admin"
                ],
                "summary": "Refresh token for admin",
                "operationId": "Admin RefreshToken",
                "parameters": [
                    {
                        "type": "string",
                        "description": "token string: ",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    }
                }
            }
        },
        "/admin/vipuser": {
            "patch": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Admin"
                ],
                "summary": "makes the user vip",
                "operationId": "make vip user",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User Name : ",
                        "name": "username",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    }
                }
            }
        },
        "/event/approved": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Event"
                ],
                "summary": "list all approved upcoming events",
                "operationId": "list all approved events",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Page number: ",
                        "name": "page",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Page capacity : ",
                        "name": "pagesize",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    }
                }
            }
        },
        "/event/delete": {
            "delete": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Event"
                ],
                "summary": "delete event",
                "operationId": "Delete event",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Title: ",
                        "name": "title",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    }
                }
            }
        },
        "/event/geteventbytitle": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Event"
                ],
                "summary": "delete event",
                "operationId": "Get event by id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Title: ",
                        "name": "title",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    }
                }
            }
        },
        "/event/update": {
            "patch": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "update event",
                "operationId": "Update event",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Title: ",
                        "name": "title",
                        "in": "query",
                        "required": true
                    },
                    {
                        "description": "update event with new body",
                        "name": "Updateevent",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.Users"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    }
                }
            }
        },
        "/user/event/create": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Create event",
                "operationId": "Create event",
                "parameters": [
                    {
                        "type": "string",
                        "description": "organizerName: ",
                        "name": "userName",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "Create event",
                        "name": "CreateEvent",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.Events"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    }
                }
            }
        },
        "/user/login": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User",
                    "User"
                ],
                "summary": "Login for users",
                "operationId": "User Login",
                "parameters": [
                    {
                        "description": "userlogin: ",
                        "name": "UserLogin",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.Users"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    }
                }
            }
        },
        "/user/signup": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "SignUp for users",
                "operationId": "User SignUp",
                "parameters": [
                    {
                        "description": "user signup with username, phonenumber email ,password",
                        "name": "RegisterUser",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.Users"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    }
                }
            }
        },
        "/user/token/refresh": {
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User",
                    "User"
                ],
                "summary": "Refresh token for users",
                "operationId": "User RefreshToken",
                "parameters": [
                    {
                        "type": "string",
                        "description": "token string: ",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    }
                }
            }
        },
        "/user/verify/account": {
            "delete": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Verify account",
                "operationId": "Verify account",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Email: ",
                        "name": "Email",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "code: ",
                        "name": "Code",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "domain.Admins": {
            "type": "object",
            "required": [
                "adminname",
                "email",
                "password"
            ],
            "properties": {
                "adminid": {
                    "type": "integer"
                },
                "adminname": {
                    "type": "string",
                    "maxLength": 50,
                    "minLength": 2
                },
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "phonenumber": {
                    "type": "string"
                },
                "verification": {
                    "type": "boolean"
                }
            }
        },
        "domain.Events": {
            "type": "object",
            "required": [
                "eventdate",
                "eventpic",
                "organizername",
                "title"
            ],
            "properties": {
                "applicationclosingdate": {
                    "type": "string"
                },
                "applicationlink": {
                    "type": "string"
                },
                "approved": {
                    "type": "boolean"
                },
                "archived": {
                    "type": "boolean"
                },
                "createdat": {
                    "type": "string"
                },
                "cusatonly": {
                    "type": "boolean"
                },
                "eventdate": {
                    "type": "string"
                },
                "eventid": {
                    "type": "integer"
                },
                "eventpic": {
                    "type": "string"
                },
                "location": {
                    "type": "string"
                },
                "longdiscription": {
                    "type": "string"
                },
                "maxapplications": {
                    "type": "integer"
                },
                "online": {
                    "type": "boolean"
                },
                "organizername": {
                    "type": "string"
                },
                "paid": {
                    "type": "boolean"
                },
                "sex": {
                    "type": "string"
                },
                "shortdiscription": {
                    "type": "string"
                },
                "subevents": {
                    "type": "string"
                },
                "title": {
                    "type": "string",
                    "maxLength": 50,
                    "minLength": 2
                },
                "websitelink": {
                    "type": "string"
                }
            }
        },
        "domain.Users": {
            "type": "object",
            "required": [
                "email",
                "firstname",
                "lastname",
                "password",
                "username"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "eventid": {
                    "type": "integer"
                },
                "firstname": {
                    "type": "string",
                    "maxLength": 50,
                    "minLength": 2
                },
                "lastname": {
                    "type": "string",
                    "maxLength": 50,
                    "minLength": 1
                },
                "password": {
                    "type": "string"
                },
                "phonenumber": {
                    "type": "string"
                },
                "profile": {
                    "type": "string"
                },
                "userid": {
                    "type": "integer"
                },
                "username": {
                    "type": "string",
                    "maxLength": 50,
                    "minLength": 2
                },
                "verification": {
                    "type": "boolean"
                },
                "vip": {
                    "type": "boolean"
                }
            }
        },
        "response.Response": {
            "type": "object",
            "properties": {
                "data": {},
                "errors": {},
                "message": {
                    "type": "string"
                },
                "status": {
                    "type": "boolean"
                }
            }
        }
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        },
        "BearerAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:3000",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Go + Gin Radar API",
	Description:      "This is an Events Radar project. You can visit the GitHub repository at https://github.com/SethukumarJ/Events_Radar_Developement",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
