## Cosas que son causantes de reentrega y que se deben corregir:

- Si hago un GET /courses y no hay courses obtengo {data: null}, cuando deberia obtener {data: []}. Si bien parece una trivialidad, OpenAPI funciona como un contrato, y si no se cumple, aquel que consuma tu API va a tener problemas con su front/servicio

## Cosas importantes pero no descalificantes, no es necesario corregirlas:

- Solo testeas el happy path para cada endpoint. En esta API chiquita quizas no es de mucha importancia, pero a medida que se complejiza van teniendo mas relevancia, testear de forma no automatizada no solo se hace tedioso sino dificil de mantener. Tene en cuenta que en el TP grupal les piden un minimo de coverage, guarda con esto en el futuro
- La docu de las funciones no dice mucho

## Aspectos positivos:

- Bien aplicada la arquitectura en capas
- EL codigo es claro
- Se validan los tipos de los parametros url y de body. En caso de ser invalidos se envia un error con la descripcion pertinente
