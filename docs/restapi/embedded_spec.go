// Code generated by go-swagger; DO NOT EDIT.

package restapi

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
)

var (
	// SwaggerJSON embedded version of the swagger document used at generation time
	SwaggerJSON json.RawMessage
	// FlatSwaggerJSON embedded flattened version of the swagger document used at generation time
	FlatSwaggerJSON json.RawMessage
)

func init() {
	SwaggerJSON = json.RawMessage([]byte(`{
  "schemes": [
    "https",
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "This is the oficial documentation for the EasyPoll API. If you have any problems or requests, please contact us on GitHub.",
    "title": "EasyPoll",
    "contact": {
      "email": "noreply@easypoll.com.br"
    },
    "version": "1.0.0"
  },
  "basePath": "/v1",
  "paths": {
    "/login": {
      "post": {
        "description": "Authenticates an user and returns a token to be used in requests",
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "auth"
        ],
        "summary": "Login to the application",
        "operationId": "loginUser",
        "parameters": [
          {
            "description": "Account credentials",
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "type": "object",
              "required": [
                "email",
                "password"
              ],
              "properties": {
                "email": {
                  "description": "Email of the account",
                  "type": "string"
                },
                "password": {
                  "description": "Password of the account",
                  "type": "string",
                  "format": "password"
                }
              }
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Login successful",
            "schema": {
              "allOf": [
                {
                  "$ref": "#/definitions/Token"
                },
                {
                  "type": "object",
                  "properties": {
                    "user": {
                      "$ref": "#/definitions/User"
                    }
                  }
                }
              ]
            }
          },
          "400": {
            "description": "Incorrect password or invalid input",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          },
          "404": {
            "description": "Email not registered",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          },
          "422": {
            "description": "Missing required fields"
          }
        }
      }
    },
    "/polls": {
      "get": {
        "description": "Gets a list with all the polls in the database",
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "poll"
        ],
        "summary": "List all polls",
        "operationId": "getPolls",
        "responses": {
          "200": {
            "description": "Poll list fetched successfully",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/Poll"
              }
            }
          }
        }
      },
      "post": {
        "description": "Inserts a new poll in the database",
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "poll"
        ],
        "summary": "Create a new poll",
        "operationId": "createPoll",
        "parameters": [
          {
            "description": "Poll that will be created",
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "required": [
                "question",
                "answers",
                "multipleAnswers"
              ],
              "allOf": [
                {
                  "$ref": "#/definitions/PollRequest"
                }
              ]
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Poll created",
            "schema": {
              "$ref": "#/definitions/Poll"
            }
          },
          "400": {
            "description": "Invalid input"
          },
          "422": {
            "description": "Missing required fields"
          }
        }
      }
    },
    "/polls/{id}": {
      "get": {
        "description": "Gets a single poll by id",
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "poll"
        ],
        "summary": "Display a single poll by id",
        "operationId": "getPollById",
        "parameters": [
          {
            "type": "integer",
            "description": "The id of the poll",
            "name": "id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Poll fetched successfully",
            "schema": {
              "$ref": "#/definitions/Poll"
            }
          },
          "404": {
            "description": "Resource not found"
          }
        }
      },
      "put": {
        "description": "Updates the whole poll object with a new one\n\nUnspecified optional fields will be counted as zero-value and will be overwritten\n",
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "poll"
        ],
        "summary": "Update a poll by id",
        "operationId": "updatePollById",
        "parameters": [
          {
            "type": "integer",
            "description": "The id of the poll",
            "name": "id",
            "in": "path",
            "required": true
          },
          {
            "description": "New poll data",
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "required": [
                "question",
                "answers",
                "multipleAnswers"
              ],
              "allOf": [
                {
                  "$ref": "#/definitions/PollRequest"
                }
              ]
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Poll updated",
            "schema": {
              "$ref": "#/definitions/Poll"
            }
          },
          "400": {
            "description": "Invalid input"
          },
          "404": {
            "description": "Resource not found"
          },
          "422": {
            "description": "Missing required fields"
          }
        }
      },
      "delete": {
        "description": "Removes a poll from the database by id",
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "poll"
        ],
        "summary": "Delete a poll by id",
        "operationId": "deletePollById",
        "parameters": [
          {
            "type": "integer",
            "description": "The id of the poll",
            "name": "id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "204": {
            "description": "Poll deleted"
          },
          "404": {
            "description": "Resource not found"
          }
        }
      },
      "patch": {
        "description": "Updates (part of) a poll properties, all fields are optional\n\nUnspecified fields will be ignored and won't be updated\n",
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "poll"
        ],
        "summary": "Update (part of) a poll by id",
        "operationId": "patchPollById",
        "parameters": [
          {
            "type": "integer",
            "description": "The id of the poll",
            "name": "id",
            "in": "path",
            "required": true
          },
          {
            "description": "Properties of the poll that will be updated",
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/PollRequest"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Poll updated",
            "schema": {
              "$ref": "#/definitions/Poll"
            }
          },
          "400": {
            "description": "Invalid input"
          },
          "404": {
            "description": "Resource not found"
          }
        }
      }
    },
    "/polls/{id}/votes": {
      "post": {
        "description": "Adds one vote to answers of a poll, the request body is an array with integers, each number is the index of the answer\n\nSending [1, 3, 2] will add one vote to the 2nd, 4th and 3rd answers of the poll\n\nYou can only send more than one value in the array if the poll accepts multiple answers\n\nRepeated numbers will be counted as one\n",
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "vote"
        ],
        "summary": "Votes in answers of a poll",
        "operationId": "addVotePoll",
        "parameters": [
          {
            "type": "integer",
            "description": "The id of the poll",
            "name": "id",
            "in": "path",
            "required": true
          },
          {
            "description": "Array with answers to be voted",
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "type": "array",
              "items": {
                "type": "integer"
              }
            }
          }
        ],
        "responses": {
          "204": {
            "description": "Vote(s) added"
          },
          "400": {
            "description": "Invalid input",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          },
          "404": {
            "description": "Resource not found"
          },
          "422": {
            "description": "Missing required fields"
          }
        }
      }
    },
    "/register": {
      "post": {
        "description": "Create a new account and returns a token to be used in requests",
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "auth"
        ],
        "summary": "Create a new account",
        "operationId": "registerUser",
        "parameters": [
          {
            "description": "Account details",
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "type": "object",
              "required": [
                "name",
                "email",
                "password"
              ],
              "properties": {
                "email": {
                  "description": "Email of the account",
                  "type": "string"
                },
                "name": {
                  "description": "Name of the user",
                  "type": "string"
                },
                "password": {
                  "description": "Password of the account",
                  "type": "string",
                  "format": "password"
                }
              }
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Account created",
            "schema": {
              "allOf": [
                {
                  "$ref": "#/definitions/Token"
                },
                {
                  "type": "object",
                  "properties": {
                    "user": {
                      "$ref": "#/definitions/User"
                    }
                  }
                }
              ]
            }
          },
          "400": {
            "description": "Email already registered or invalid input"
          },
          "422": {
            "description": "Missing required fields"
          }
        }
      }
    },
    "/user": {
      "get": {
        "security": [
          {
            "Bearer": []
          }
        ],
        "description": "Get the authenticated user account details",
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "auth"
        ],
        "summary": "Get the logged user info",
        "operationId": "getAuthenticatedUser",
        "responses": {
          "200": {
            "description": "Return the account of the authenticated user",
            "schema": {
              "allOf": [
                {
                  "$ref": "#/definitions/User"
                }
              ]
            }
          },
          "401": {
            "description": "Invalid token"
          }
        }
      }
    }
  },
  "definitions": {
    "Answer": {
      "type": "object",
      "properties": {
        "title": {
          "type": "string",
          "example": "Blue"
        },
        "votes": {
          "type": "integer",
          "x-omitempty": false
        }
      }
    },
    "Error": {
      "type": "object",
      "properties": {
        "code": {
          "type": "integer",
          "example": 622
        },
        "message": {
          "type": "string",
          "example": "error message"
        }
      }
    },
    "Poll": {
      "type": "object",
      "properties": {
        "answers": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/Answer"
          }
        },
        "createdAt": {
          "type": "string",
          "format": "date-time"
        },
        "id": {
          "type": "integer"
        },
        "multipleAnswers": {
          "type": "boolean"
        },
        "question": {
          "type": "string",
          "example": "What's your favorite color?"
        },
        "votes": {
          "type": "integer",
          "x-omitempty": false
        }
      }
    },
    "PollRequest": {
      "type": "object",
      "properties": {
        "answers": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/Answer"
          }
        },
        "multipleAnswers": {
          "type": "boolean"
        },
        "question": {
          "type": "string",
          "example": "What's your favorite color?"
        }
      }
    },
    "Token": {
      "type": "object",
      "properties": {
        "token": {
          "type": "string",
          "example": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1aWQiOjEsImlhdCI6MTUzMzY2MTY0Mn0.w_qI2tv8imANq_ys_ZKvLtrGItd7hrfRjifJoYSJzfo"
        }
      }
    },
    "User": {
      "type": "object",
      "properties": {
        "email": {
          "description": "Email of the user",
          "type": "string",
          "example": "john@doe.com"
        },
        "name": {
          "description": "Name of the user",
          "type": "string",
          "example": "John Doe"
        },
        "password": {
          "description": "Password of the user",
          "type": "string",
          "format": "password",
          "example": "secret"
        }
      }
    }
  },
  "securityDefinitions": {
    "Bearer": {
      "type": "apiKey",
      "name": "Authorization",
      "in": "header"
    }
  },
  "tags": [
    {
      "description": "Operations about authentication",
      "name": "auth"
    },
    {
      "description": "Operations to get, create, update and delete polls",
      "name": "poll"
    },
    {
      "description": "Operations to vote",
      "name": "vote"
    }
  ]
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "schemes": [
    "https",
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "This is the oficial documentation for the EasyPoll API. If you have any problems or requests, please contact us on GitHub.",
    "title": "EasyPoll",
    "contact": {
      "email": "noreply@easypoll.com.br"
    },
    "version": "1.0.0"
  },
  "basePath": "/v1",
  "paths": {
    "/login": {
      "post": {
        "description": "Authenticates an user and returns a token to be used in requests",
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "auth"
        ],
        "summary": "Login to the application",
        "operationId": "loginUser",
        "parameters": [
          {
            "description": "Account credentials",
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "type": "object",
              "required": [
                "email",
                "password"
              ],
              "properties": {
                "email": {
                  "description": "Email of the account",
                  "type": "string"
                },
                "password": {
                  "description": "Password of the account",
                  "type": "string",
                  "format": "password"
                }
              }
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Login successful",
            "schema": {
              "allOf": [
                {
                  "$ref": "#/definitions/Token"
                },
                {
                  "type": "object",
                  "properties": {
                    "user": {
                      "$ref": "#/definitions/User"
                    }
                  }
                }
              ]
            }
          },
          "400": {
            "description": "Incorrect password or invalid input",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          },
          "404": {
            "description": "Email not registered",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          },
          "422": {
            "description": "Missing required fields"
          }
        }
      }
    },
    "/polls": {
      "get": {
        "description": "Gets a list with all the polls in the database",
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "poll"
        ],
        "summary": "List all polls",
        "operationId": "getPolls",
        "responses": {
          "200": {
            "description": "Poll list fetched successfully",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/Poll"
              }
            }
          }
        }
      },
      "post": {
        "description": "Inserts a new poll in the database",
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "poll"
        ],
        "summary": "Create a new poll",
        "operationId": "createPoll",
        "parameters": [
          {
            "description": "Poll that will be created",
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "required": [
                "question",
                "answers",
                "multipleAnswers"
              ],
              "allOf": [
                {
                  "$ref": "#/definitions/PollRequest"
                }
              ]
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Poll created",
            "schema": {
              "$ref": "#/definitions/Poll"
            }
          },
          "400": {
            "description": "Invalid input"
          },
          "422": {
            "description": "Missing required fields"
          }
        }
      }
    },
    "/polls/{id}": {
      "get": {
        "description": "Gets a single poll by id",
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "poll"
        ],
        "summary": "Display a single poll by id",
        "operationId": "getPollById",
        "parameters": [
          {
            "type": "integer",
            "description": "The id of the poll",
            "name": "id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Poll fetched successfully",
            "schema": {
              "$ref": "#/definitions/Poll"
            }
          },
          "404": {
            "description": "Resource not found"
          }
        }
      },
      "put": {
        "description": "Updates the whole poll object with a new one\n\nUnspecified optional fields will be counted as zero-value and will be overwritten\n",
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "poll"
        ],
        "summary": "Update a poll by id",
        "operationId": "updatePollById",
        "parameters": [
          {
            "type": "integer",
            "description": "The id of the poll",
            "name": "id",
            "in": "path",
            "required": true
          },
          {
            "description": "New poll data",
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "required": [
                "question",
                "answers",
                "multipleAnswers"
              ],
              "allOf": [
                {
                  "$ref": "#/definitions/PollRequest"
                }
              ]
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Poll updated",
            "schema": {
              "$ref": "#/definitions/Poll"
            }
          },
          "400": {
            "description": "Invalid input"
          },
          "404": {
            "description": "Resource not found"
          },
          "422": {
            "description": "Missing required fields"
          }
        }
      },
      "delete": {
        "description": "Removes a poll from the database by id",
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "poll"
        ],
        "summary": "Delete a poll by id",
        "operationId": "deletePollById",
        "parameters": [
          {
            "type": "integer",
            "description": "The id of the poll",
            "name": "id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "204": {
            "description": "Poll deleted"
          },
          "404": {
            "description": "Resource not found"
          }
        }
      },
      "patch": {
        "description": "Updates (part of) a poll properties, all fields are optional\n\nUnspecified fields will be ignored and won't be updated\n",
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "poll"
        ],
        "summary": "Update (part of) a poll by id",
        "operationId": "patchPollById",
        "parameters": [
          {
            "type": "integer",
            "description": "The id of the poll",
            "name": "id",
            "in": "path",
            "required": true
          },
          {
            "description": "Properties of the poll that will be updated",
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/PollRequest"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Poll updated",
            "schema": {
              "$ref": "#/definitions/Poll"
            }
          },
          "400": {
            "description": "Invalid input"
          },
          "404": {
            "description": "Resource not found"
          }
        }
      }
    },
    "/polls/{id}/votes": {
      "post": {
        "description": "Adds one vote to answers of a poll, the request body is an array with integers, each number is the index of the answer\n\nSending [1, 3, 2] will add one vote to the 2nd, 4th and 3rd answers of the poll\n\nYou can only send more than one value in the array if the poll accepts multiple answers\n\nRepeated numbers will be counted as one\n",
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "vote"
        ],
        "summary": "Votes in answers of a poll",
        "operationId": "addVotePoll",
        "parameters": [
          {
            "type": "integer",
            "description": "The id of the poll",
            "name": "id",
            "in": "path",
            "required": true
          },
          {
            "description": "Array with answers to be voted",
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "type": "array",
              "items": {
                "type": "integer"
              }
            }
          }
        ],
        "responses": {
          "204": {
            "description": "Vote(s) added"
          },
          "400": {
            "description": "Invalid input",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          },
          "404": {
            "description": "Resource not found"
          },
          "422": {
            "description": "Missing required fields"
          }
        }
      }
    },
    "/register": {
      "post": {
        "description": "Create a new account and returns a token to be used in requests",
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "auth"
        ],
        "summary": "Create a new account",
        "operationId": "registerUser",
        "parameters": [
          {
            "description": "Account details",
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "type": "object",
              "required": [
                "name",
                "email",
                "password"
              ],
              "properties": {
                "email": {
                  "description": "Email of the account",
                  "type": "string"
                },
                "name": {
                  "description": "Name of the user",
                  "type": "string"
                },
                "password": {
                  "description": "Password of the account",
                  "type": "string",
                  "format": "password"
                }
              }
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Account created",
            "schema": {
              "allOf": [
                {
                  "$ref": "#/definitions/Token"
                },
                {
                  "type": "object",
                  "properties": {
                    "user": {
                      "$ref": "#/definitions/User"
                    }
                  }
                }
              ]
            }
          },
          "400": {
            "description": "Email already registered or invalid input"
          },
          "422": {
            "description": "Missing required fields"
          }
        }
      }
    },
    "/user": {
      "get": {
        "security": [
          {
            "Bearer": []
          }
        ],
        "description": "Get the authenticated user account details",
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "auth"
        ],
        "summary": "Get the logged user info",
        "operationId": "getAuthenticatedUser",
        "responses": {
          "200": {
            "description": "Return the account of the authenticated user",
            "schema": {
              "allOf": [
                {
                  "$ref": "#/definitions/User"
                }
              ]
            }
          },
          "401": {
            "description": "Invalid token"
          }
        }
      }
    }
  },
  "definitions": {
    "Answer": {
      "type": "object",
      "properties": {
        "title": {
          "type": "string",
          "example": "Blue"
        },
        "votes": {
          "type": "integer",
          "x-omitempty": false
        }
      }
    },
    "Error": {
      "type": "object",
      "properties": {
        "code": {
          "type": "integer",
          "example": 622
        },
        "message": {
          "type": "string",
          "example": "error message"
        }
      }
    },
    "Poll": {
      "type": "object",
      "properties": {
        "answers": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/Answer"
          }
        },
        "createdAt": {
          "type": "string",
          "format": "date-time"
        },
        "id": {
          "type": "integer"
        },
        "multipleAnswers": {
          "type": "boolean"
        },
        "question": {
          "type": "string",
          "example": "What's your favorite color?"
        },
        "votes": {
          "type": "integer",
          "x-omitempty": false
        }
      }
    },
    "PollRequest": {
      "type": "object",
      "properties": {
        "answers": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/Answer"
          }
        },
        "multipleAnswers": {
          "type": "boolean"
        },
        "question": {
          "type": "string",
          "example": "What's your favorite color?"
        }
      }
    },
    "Token": {
      "type": "object",
      "properties": {
        "token": {
          "type": "string",
          "example": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1aWQiOjEsImlhdCI6MTUzMzY2MTY0Mn0.w_qI2tv8imANq_ys_ZKvLtrGItd7hrfRjifJoYSJzfo"
        }
      }
    },
    "User": {
      "type": "object",
      "properties": {
        "email": {
          "description": "Email of the user",
          "type": "string",
          "example": "john@doe.com"
        },
        "name": {
          "description": "Name of the user",
          "type": "string",
          "example": "John Doe"
        },
        "password": {
          "description": "Password of the user",
          "type": "string",
          "format": "password",
          "example": "secret"
        }
      }
    }
  },
  "securityDefinitions": {
    "Bearer": {
      "type": "apiKey",
      "name": "Authorization",
      "in": "header"
    }
  },
  "tags": [
    {
      "description": "Operations about authentication",
      "name": "auth"
    },
    {
      "description": "Operations to get, create, update and delete polls",
      "name": "poll"
    },
    {
      "description": "Operations to vote",
      "name": "vote"
    }
  ]
}`))
}
