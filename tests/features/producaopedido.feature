Feature: API Producao Pedido

  Scenario Outline: ProducaoPedido creation
    Given Parameter ID: <id>
    When request POST /producaoPedido
    Then statusCode should be <statusCode>

    Examples:
        | id | statusCode |
        | 21  | 201       |
        | 1  | 500        |

  Scenario Outline: ProducaoPedido update
    Given Parameter ID: <id>
    When request PUT /producaoPedido with status "<status>"
    Then statusCode should be <statusCode>

    Examples:
        | id | statusCode | status |
        | 1  | 200       | CRIADO |
        | 2  | 422        | ERRO  |

  Scenario Outline: ProducaoPedido GET
    Given Parameter ID: <id>
    When request GET /producaoPedido by id
    Then statusCode should be <statusCode>

    Examples:
        | id | statusCode |
        | 1  | 200       |
        | 50  | 404        |

  Scenario Outline: Get by status
    Given Parameter status "<status>"
    When request GET /producaoPedido by "<status>"
    Then statusCode should be <statusCode>

    Examples:
        | status | statusCode |
        | CRIADO  | 200       |
        | ERRO  | 422        |

  Scenario: Get all
    When request GET /producaoPedido
    Then statusCode should be 200

  Scenario: Get by ID and validate JSON
    Given Parameter ID: 2
    When request GET /producaoPedido by id
    Then statusCode should be 200
    And response should match json:
    """
    {
    "PedidoID": 2,
    "Status": "CRIADO",
    "CreatedAt": "2023-07-21T13:00:00-03:00",
    "UpdatedAt": "0001-01-01T00:00:00Z",
    "DeletedAt": null
  }
    """

  Scenario Outline: Should get healthcheck
    When request GET /healthcheck
    Then statusCode should be <statusCode>

    Examples:
    | statusCode |
    | 200       |



