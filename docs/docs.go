// Package docs GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import (
	"bytes"
	"encoding/json"
	"strings"
	"text/template"

	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "API Support",
            "url": "http://www.swagger.io/support",
            "email": "support@swagger.io"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/categories": {
            "get": {
                "description": "Get a list of Category.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Category"
                ],
                "summary": "Get all Category.",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Category"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "Creating a new Category.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Category"
                ],
                "summary": "Create New Category.",
                "parameters": [
                    {
                        "description": "the body to create a new Category",
                        "name": "Body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controllers.categoryInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Category"
                        }
                    }
                }
            }
        },
        "/categories/{id}": {
            "get": {
                "description": "Get an Category by id.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Category"
                ],
                "summary": "Get Category by Id.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Category id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Category"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete a Category by id.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Category"
                ],
                "summary": "Delete one Category.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Category id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "boolean"
                            }
                        }
                    }
                }
            },
            "patch": {
                "description": "Update Category by id.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Category"
                ],
                "summary": "Update Category.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Category id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "the body to update category",
                        "name": "Body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controllers.categoryInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Category"
                        }
                    }
                }
            }
        },
        "/customer/login": {
            "post": {
                "description": "Logging in to get jwt token to access customer api.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth Customer"
                ],
                "summary": "Login as as customer.",
                "parameters": [
                    {
                        "description": "the body to login a customer",
                        "name": "Body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controllers.LoginInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        },
        "/customer/register": {
            "post": {
                "description": "registering a customer from public access.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth Customer"
                ],
                "summary": "Register a Customer.",
                "parameters": [
                    {
                        "description": "the body to register a customer",
                        "name": "Body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controllers.RegisterInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        },
        "/customer/{id}": {
            "patch": {
                "description": "Update customer password by id.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth Customer"
                ],
                "summary": "Update Customer Password.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Customer id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "the body to update customer password",
                        "name": "Body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controllers.ChangePasswordInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        },
        "/order": {
            "get": {
                "description": "Get a list of Order.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Order"
                ],
                "summary": "Get all Order.",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Order"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "Creating a new Order.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Order"
                ],
                "summary": "Create New Order.",
                "parameters": [
                    {
                        "description": "the body to create a new Order",
                        "name": "Body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controllers.orderInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Order"
                        }
                    }
                }
            }
        },
        "/order/{id}": {
            "get": {
                "description": "Get a Order by id.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Order"
                ],
                "summary": "Get Order by Id.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Order id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Order"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete a Order by id.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Order"
                ],
                "summary": "Delete one Order.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Order id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "boolean"
                            }
                        }
                    }
                }
            },
            "patch": {
                "description": "Update Order Status by id.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Order"
                ],
                "summary": "Update Order Status.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Order id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "the body to update Order Status",
                        "name": "Body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controllers.orderInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Order"
                        }
                    }
                }
            }
        },
        "/products": {
            "get": {
                "description": "Get a list of Product.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Product"
                ],
                "summary": "Get all Product.",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Product"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "Creating a new Product.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Product"
                ],
                "summary": "Create New Product.",
                "parameters": [
                    {
                        "description": "the body to create a new Product",
                        "name": "Body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controllers.productInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Product"
                        }
                    }
                }
            }
        },
        "/products/rating": {
            "post": {
                "description": "Creating a new Rating.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Rating"
                ],
                "summary": "Create New Rating.",
                "parameters": [
                    {
                        "description": "the body to create a new Rating",
                        "name": "Body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controllers.productRatingInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.ProductRating"
                        }
                    }
                }
            }
        },
        "/products/rating/{id}": {
            "get": {
                "description": "Get a Rating by id.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Rating"
                ],
                "summary": "Get Rating by Id.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Rating id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.ProductRating"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete a Rating by id.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Rating"
                ],
                "summary": "Delete one Rating.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Rating id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "boolean"
                            }
                        }
                    }
                }
            }
        },
        "/products/search": {
            "get": {
                "description": "Search a list of Product.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Product"
                ],
                "summary": "Search Product.",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Product"
                            }
                        }
                    }
                }
            }
        },
        "/products/{id}": {
            "get": {
                "description": "Get a Product by id.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Product"
                ],
                "summary": "Get Product by Id.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Product id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Product"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete a Product by id.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Product"
                ],
                "summary": "Delete one Product.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Product id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "boolean"
                            }
                        }
                    }
                }
            },
            "patch": {
                "description": "Update Product by id.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Product"
                ],
                "summary": "Update Product.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Product id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "the body to update Product",
                        "name": "Body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controllers.productInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Product"
                        }
                    }
                }
            }
        },
        "/products/{id}/category": {
            "get": {
                "description": "Get a Product by category id.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Product"
                ],
                "summary": "Get Product by Category Id.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Product id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Product"
                        }
                    }
                }
            }
        },
        "/products/{id}/rating": {
            "get": {
                "description": "Get a Rating by Product id.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Rating"
                ],
                "summary": "Get Rating by Product Id.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Rating id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.ProductRating"
                        }
                    }
                }
            }
        },
        "/produts/ratings/{id}": {
            "patch": {
                "description": "Update Rating by id.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Rating"
                ],
                "summary": "Update Rating.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Rating id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "the body to update Rating",
                        "name": "Body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controllers.productRatingInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.ProductRating"
                        }
                    }
                }
            }
        },
        "/seller/login": {
            "post": {
                "description": "Logging in to get jwt token to access seller api.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth Seller"
                ],
                "summary": "Login as as seller.",
                "parameters": [
                    {
                        "description": "the body to login a seller",
                        "name": "Body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controllers.LoginInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        },
        "/seller/register": {
            "post": {
                "description": "registering a seller from public access.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth Seller"
                ],
                "summary": "Register a seller.",
                "parameters": [
                    {
                        "description": "the body to register a seller",
                        "name": "Body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controllers.RegisterInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        },
        "/seller/{id}": {
            "patch": {
                "description": "Update seller password by id.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth Seller"
                ],
                "summary": "Update Seller Password.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Seller id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "the body to update seller password",
                        "name": "Body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controllers.ChangePasswordInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "controllers.ChangePasswordInput": {
            "type": "object",
            "required": [
                "new_password",
                "new_password_confirm"
            ],
            "properties": {
                "new_password": {
                    "type": "string"
                },
                "new_password_confirm": {
                    "type": "string"
                }
            }
        },
        "controllers.LoginInput": {
            "type": "object",
            "required": [
                "password",
                "username"
            ],
            "properties": {
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "controllers.RegisterInput": {
            "type": "object",
            "required": [
                "email",
                "password",
                "username"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "controllers.categoryInput": {
            "type": "object",
            "properties": {
                "category": {
                    "type": "string"
                }
            }
        },
        "controllers.orderDetailInput": {
            "type": "object",
            "properties": {
                "product_id": {
                    "type": "integer"
                },
                "qty": {
                    "type": "integer"
                }
            }
        },
        "controllers.orderInput": {
            "type": "object",
            "properties": {
                "customer_id": {
                    "type": "integer"
                },
                "product_detail_id": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/controllers.orderDetailInput"
                    }
                },
                "seller_id": {
                    "type": "integer"
                },
                "total_price": {
                    "type": "integer"
                }
            }
        },
        "controllers.productInput": {
            "type": "object",
            "properties": {
                "category_id": {
                    "type": "integer"
                },
                "description": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "price": {
                    "type": "integer"
                },
                "seller_id": {
                    "type": "integer"
                },
                "stock": {
                    "type": "integer"
                }
            }
        },
        "controllers.productRatingInput": {
            "type": "object",
            "properties": {
                "comment": {
                    "type": "string"
                },
                "customer_id": {
                    "type": "integer"
                },
                "product_id": {
                    "type": "integer"
                },
                "rating": {
                    "type": "integer"
                }
            }
        },
        "models.Category": {
            "type": "object",
            "properties": {
                "category": {
                    "type": "string"
                },
                "created_at": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "models.Order": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "customer_id": {
                    "type": "integer"
                },
                "id": {
                    "type": "integer"
                },
                "invoice": {
                    "type": "string"
                },
                "order_status": {
                    "type": "string"
                },
                "seller_id": {
                    "type": "integer"
                },
                "total_price": {
                    "type": "integer"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "models.Product": {
            "type": "object",
            "properties": {
                "category_id": {
                    "type": "integer"
                },
                "created_at": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "price": {
                    "type": "integer"
                },
                "seller_id": {
                    "type": "integer"
                },
                "stock": {
                    "type": "integer"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "models.ProductRating": {
            "type": "object",
            "properties": {
                "comment": {
                    "type": "string"
                },
                "created_at": {
                    "type": "string"
                },
                "customer_id": {
                    "type": "integer"
                },
                "id": {
                    "type": "integer"
                },
                "product_id": {
                    "type": "integer"
                },
                "rating": {
                    "type": "integer"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "",
	Host:        "",
	BasePath:    "",
	Schemes:     []string{},
	Title:       "",
	Description: "",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
		"escape": func(v interface{}) string {
			// escape tabs
			str := strings.Replace(v.(string), "\t", "\\t", -1)
			// replace " with \", and if that results in \\", replace that with \\\"
			str = strings.Replace(str, "\"", "\\\"", -1)
			return strings.Replace(str, "\\\\\"", "\\\\\\\"", -1)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}
