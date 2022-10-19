# Prueba de conocimiento - Amaris Consulting

<br />
<p align="center">
  <a href="#">
    <img src="https://www.amaris.com/wp-content/themes/amaris-v2/dist/images/amaris-logo-pink.svg" alt="Logo" width="50%">
  </a>

  <h3 align="center">En esta prueba se observará mis habilidades técnicas en el language de programación en Golang</h3>
</p>
<hr>

## Compilation

<hr>

### Automatic

#### Requirements

1 - Realizar una funcion que reciba una cadena de texto, que contenga una lista de nombres separados por comas ejemplo "Luis,Camilo,Andres,Laura" y devuelva dos parametros: una array de strings con los nombres ordenados alfabeticamente y un entero indicando la cantidad de nombres

2 - Realizar una funcion que reciba un numero entero "id" de un pokemon y devuelva su nombre. Para esto la funcion debe consultar una api de pokemon enviarle el id y obtener el campo "nombre" el cual va a retornar. URL = https://pokeapi.co/api/v2/pokemon-form/{{id}}

3 - Se dice que dos cadenas X y Y son amigas si existen dos cadenas no vacías u y v tales que X = uv e Y = vu. Por ejemplo, "tokyo" y "kyoto" son amigas, siendo u = “to” y v = “kyo”.
Escriba una funcion que reciba como entrada dos cadenas X e Y, e imprima si X e Y son amigas.


#### List of commands available

```bash
# Check commands available
make build -- Checked

make run -- Start

make test -- testing

make up -- docker up

make down -- docker down

```
## Endpoint

```bash
Url: /api/v1/orderText 

Method: POST

Payload:
{
    "data": "Luis,Camilo,Andres,Laura"
}

Response: 
{
    "Data": [
        "Andres",
        "Camilo",
        "Laura",
        "Luis"
    ]
}
```

```bash
Url: /api/v1/getPokemon/{id} 

Method: Get

Response: 
{
    "data": {
        "form_name": "",
        "form_names": [],
        "form_order": 1,
        "id": 1,
        "is_battle_only": false,
        "is_default": true,
        "is_mega": false,
        "name": "bulbasaur",
        "names": [],
        "order": 1,
        "pokemon": {
            "name": "bulbasaur",
            "url": "https://pokeapi.co/api/v2/pokemon/1/"
        },
        "sprites": {
            "back_default": "https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/back/1.png",
            "back_female": null,
            "back_shiny": "https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/back/shiny/1.png",
            "back_shiny_female": null,
            "front_default": "https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/1.png",
            "front_female": null,
            "front_shiny": "https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/shiny/1.png",
            "front_shiny_female": null
        },
        "types": [
            {
                "slot": 1,
                "type": {
                    "name": "grass",
                    "url": "https://pokeapi.co/api/v2/type/12/"
                }
            },
            {
                "slot": 2,
                "type": {
                    "name": "poison",
                    "url": "https://pokeapi.co/api/v2/type/4/"
                }
            }
        ],
        "version_group": {
            "name": "red-blue",
            "url": "https://pokeapi.co/api/v2/version-group/1/"
        }
    }
}
```

```bash
Url: /api/v1/friendChains 

Method: POST

Payload:
{
    "x": "TOKYO",
    "y": "KYOTO"
}

Response: 
{
     "data": "Las cadenas 'TOKYO' y 'KYOTO' son amigas"
}
```

```bash
Error Response:
{
    "error": {},
    "message": "no exist Pokemon ID",
    "code": 400
}

```
## Authors

- Victor Elias Barrera Florez <vbarrera@outlook.com>