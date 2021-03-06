package fixtures

var LinkObject = `{
  "id": "link_nMNejK7BT8oGbvO4",
  "object": "link",
  "username": "183917137",
  "link_token": "link_nMNejK7BT8oGbvO4_token_GLtktZX5SKphRtJFe_yJTDWT",
  "mode": "test",
  "active": true,
  "status": "active",
  "holder_type": "individual",
  "created_at": "2020-04-22T21:10:19.254Z",
  "institution": {
      "country": "cl",
      "id": "cl_banco_de_chile",
      "name": "Banco de Chile"
  },
  "accounts": [
      {
          "id": "acc_Z6AwnGn4idL7DPj4",
          "object": "account",
          "name": "Cuenta Corriente",
          "official_name": "Cuenta Corriente Moneda Local",
          "number": "9530516286",
          "holder_id": "134910798",
          "holder_name": "Jon Snow",
          "type": "checking_account",
          "currency": "CLP",
          "balance": {
            "available": 7010510,
            "current": 7010510,
            "limit": 7510510
          }
      },
      {
          "id": "acc_BO381oEATXonG6bj",
          "object": "account",
          "name": "Línea de Crédito",
          "official_name": "Linea De Credito Personas",
          "number": "19534121467",
          "holder_id": "134910798",
          "holder_name": "Jon Snow",
          "type": "line_of_credit",
          "currency": "CLP",
          "balance": {
            "available": 500000,
            "current": 500000,
            "limit": 500000
          }
      }
  ]
}`
