# Agree Backend Engineer Coding Challenge

Me tome la libertad de realizar el proyecto en el lenguaje Go debiado a que la posicion requiere particularmente el conocimiento de dicho. He de resaltar el hecho de que es la primera API que realizo en mi vida desde cero.

### Estrategia

Mi objetivo fue demostrar que pese a no tener experiencia haciendo APIs podia armar una en poco tiempo y en un lenguaje completamente desconocido para mi.

Para eso me centré únicamente en realizar el CRUD. Un desafío fue la arquitectura, la cual intenté que fuera lo mas CleanCode posible. Otro mas fue el setup de la Base de Datos MySQL. La construccion de las tablas fue sencillas pero la configuracion de acceso por usuario:contraseña me demoró mucho.

Una tarea que nunca había realizado era la creación de una instancia en AWS. Utilicé una EC2 tipo T2 micro con Ubuntu 64 bits.
Dentro de la instancia se encuentran instalados Go y mysql-server. El método de despliegue es demasiado rústico. Consiste en conectarse via SSH a la instancia y luego ejecutar git pull, seguido de go build.

Adicionalmente configuré un servicio de systemd para que reinicie el webserver en caso de falla.

### Defectos que no he podido corregir por falta de tiempo

1. Validaciones de seguridad informatica
2. Validaciones de los datos enum en ciertos campos como por ejemplo el PokemonType.
3. Manejo de errores mejorado.
4. Distintas configuraciones para los diferentes ambientes DEV/PROD.
5. CI/CD con Jenkins y Github.
6. Naming de variables y casing de las mismas.
7. Reemplazar consultas a la BD por Store Procedures

### Tareas pendientes

1. Paginacion.
2. Filtros.
3. Autenticacion.
4. Unit tests y e2e
5. Endpoint de carga de imagen y utilizacion de AWS para el almacenamiento de la misma.
6. Dockerizacion y deploy del container en AWS.


### Accesos

#### Endpoint principal

http://ec2-3-145-30-65.us-east-2.compute.amazonaws.com:3000/pokemon

#### Swagger

http://ec2-3-145-30-65.us-east-2.compute.amazonaws.com:3000/swagger/index.html



