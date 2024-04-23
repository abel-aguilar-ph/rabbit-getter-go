# RabbitGetter
Este ejecutable recupera los PaymentHubIds especificados de una cola de RabbitMQ, y genera directamente una consulta lista para utilizarse en Kibana, formateada de la siguiente manera: "ID" OR "ID" OR ...

Aparte, se agruparán los IDs por tipos de x-exception-message, para así saber cuantos tipos de errores hay.

En ningún caso este programa provocará cambios en la cola, simplemente leerá los N primeros mensajes de la cola y hará la lógica de extracción y transformación.

Se puede usar para los tres entornos (pro1,pro2,k8s)

### NO FUNCIONA EN TODAS LAS COLAS PORQUE HAY DISTINTOS TIPOS DE JSON.

### ¡¡Cualquier AMPLICACIÓN, MEJORA o IDEA para el código será bienvenida!!


## Instalar Go
Versión recomendable >=1.20

Para instalar en MAC se debe añadir la variable de entorno del ejecutable Go:

`open -e ~/.zshrc`

`export PATH="$PATH:/usr/local/go/bin"`

Cerrar terminal, volver a abrirlo y confirmar que funciona:

`go version`

## Configurar variables de entorno

Adicionalmente debemos de añadir las siguientes variables de entorno de forma no temporal para la configuración: 

    RABBITMQ_PRO1_URL="url"
    RABBITMQ_PRO1_VHOST="vhost"
    RABBITMQ_PRO1_PWD="contraseña"
    RABBITMQ_PRO1_USER="usuario"
    
    RABBITMQ_PRO2_URL="url"
    RABBITMQ_PRO2_VHOST="vhost"
    RABBITMQ_PRO2_PWD="contraseña"
    RABBITMQ_PRO2_USER="usuario"

    RABBITMQ_K8S_URL="url"
    RABBITMQ_K8S_VHOST="vhost"
    RABBITMQ_K8S_PWD="contraseña"
    RABBITMQ_K8S_USER="usuario"
    


## Como usarlo
Después de clonar o descargar el ZIP, tenemos que compilar y crear el ejecutable (opcionalmente tras esto se puede eliminar todo el código fuente y dejar solo el binario compilado).

*Windows*

Si no se tiene permisos, se puede ejecutar desde la carpeta raíz del proyecto, sin necesidad de hacer el ejecutable.

`go run main.go (pro1|pro2|k8s) QueueName NumberMessages`


Desde el PowerShell, estando en el directorio donde se encuentra main.go, ejecutamos lo siguiente.

`$env:GOOS="windows"; $env:GOARCH="amd64"; go build -o RabbitGetter.exe`

`./RabbitGetter.exe (pro1|pro2|k8s) QueueName NumberMessages`



*Mac*

`GOOS=darwin GOARCH=amd64 go build -o RabbitGetter`

`chmod +x RabbitGetter`

Ejecutar el programa

`./RabbitGetter (pro1|pro2|k8s) proQueueName NumberMessages`

