# RabbitGetter
Este ejecutable recupera los PaymentHubIds especificados de una cola de RabbitMQ, y genera directamente una consulta lista para utilizarse en Kibana, formateada de la siguiente manera: "ID" OR "ID" OR ...

Aparte, se agruparán los IDs por tipos de x-exception-message, para así saber cuantos tipos de errores hay.

En ningún caso este programa provocará cambios en la cola, simplemente leerá los N primeros mensajes de la cola y hará la lógica de extracción y transformación.


## Instalar Go
Versión recomendable >=1.20

Para instalar en MAC se debe añadir la variable de entorno del ejecutable Go:

`open -e ~/.zshrc`

`export PATH="$PATH:/usr/local/go/bin"`

Cerrar terminal, volver a abrirlo y confirmar que funciona:

`go version`

## Configurar variables de entorno

Adicionalmente debemos de añadir las siguientes variables de entorno de forma no temporal para la configuración: 

    RABBITMQ_PRO_URL="url"
    RABBITMQ_PRO_VHOST="vhost"
    RABBITMQ_PRO_PWD="contraseña"
    RABBITMQ_PRO_USER="usuario"


## Como usarlo
Después de clonar o descargar el ZIP, tenemos que compilar y crear el ejecutable (opcionalmente tras esto se puede eliminar todo el código fuente y dejar solo el binario compilado).

*Windows*

Desde el PowerShell, estando en el directorio donde se encuentra main.go, ejecutamos lo siguiente.

`$env:GOOS="windows"; $env:GOARCH="amd64"; go build -o RabbitGetter.exe`

`./RabbitGetter.exe QueueName NumberMessages`

Si no se tiene permisos, se puede ejecutar desde la carpeta raíz del proyecto.

`go run main.go QueueName NumberMessages`


*Mac*

`GOOS=darwin GOARCH=amd64 go build -o RabbitGetter`

`chmod +x RabbitGetter`

Ejecutar el programa

`./RabbitGetter QueueName NumberMessages`

