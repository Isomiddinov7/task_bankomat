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
        "/accounts": {
            "post": {
                "description": "Create Account",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Account"
                ],
                "summary": "Create Account",
                "operationId": "create_account",
                "parameters": [
                    {
                        "description": "CreateAccountRequestBody",
                        "name": "object",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.CreateAccount"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "AccountBody",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/handler.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/models.Account"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Invalid Argument",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/handler.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "500": {
                        "description": "Server Error",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/handler.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/accounts/{id}/balance": {
            "get": {
                "description": "Get Balance",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Balance"
                ],
                "summary": "Get Balance",
                "operationId": "get_balance",
                "parameters": [
                    {
                        "type": "string",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "GetListBranchResponseBody",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/handler.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Invalid Argument",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/handler.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "500": {
                        "description": "Server Error",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/handler.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/accounts/{id}/deposit": {
            "post": {
                "description": "Deposit Account",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Deposit"
                ],
                "summary": "Deposit Account",
                "operationId": "deposit",
                "parameters": [
                    {
                        "type": "string",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "DepositRequestBody",
                        "name": "object",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Deposit"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Account",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/handler.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/models.Deposit"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Invalid Argument",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/handler.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "500": {
                        "description": "Server Error",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/handler.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/accounts/{id}/withdraw": {
            "post": {
                "description": "Withdraw Account",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Withdraw"
                ],
                "summary": "Withdraw Account",
                "operationId": "withdraw",
                "parameters": [
                    {
                        "type": "string",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "WithdrawRequestBody",
                        "name": "object",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Withdraw"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Account",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/handler.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/models.Withdraw"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Invalid Argument",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/handler.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "500": {
                        "description": "Server Error",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/handler.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "handler.Response": {
            "type": "object",
            "properties": {
                "data": {},
                "description": {
                    "type": "string"
                },
                "status": {
                    "type": "integer"
                }
            }
        },
        "models.Account": {
            "type": "object",
            "properties": {
                "balance": {
                    "type": "number"
                },
                "created_at": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "phone": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "models.CreateAccount": {
            "type": "object",
            "properties": {
                "balance": {
                    "type": "number"
                },
                "name": {
                    "type": "string"
                },
                "phone": {
                    "type": "string"
                }
            }
        },
        "models.Deposit": {
            "type": "object",
            "properties": {
                "amount": {
                    "type": "number"
                },
                "id": {
                    "type": "string"
                }
            }
        },
        "models.Withdraw": {
            "type": "object",
            "properties": {
                "amount": {
                    "type": "number"
                },
                "id": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
