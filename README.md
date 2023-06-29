# TO-DO Redactar nueva documentación completa


<strong>NOTA: este archivo es provicional, es para dar claridad sobre el uso del aplicativo para una revisión inicial. En el transcurso del día iré actualizando esta documentación.</strong>
## Método de uso

Inicialmente realicé el cargue de la aplicación lista para funcionar. Se debe ejecutar el comando docker-compose build y posteriormente un docker-compose up. La primera vez que inicia se va a genera un error, ya que la dependencia de contenedores contempla el inicio del contenedor y no del servicio que ejecuta tal contenedor. Cuando comienza el contenedor del aplicativo aún no ha terminado de inicializar el servicio de postgres, por lo que el aplicativo genera un error al no encontrar la base de datos disponible.

Al terminar la inicialización de la base de datos, se deben bajar los dos contenedores y volver a ejecutarlos, el problema no se volverá a presentar.

## Flujo de la aplicación

Al inicio de la aplicación se genera un usuario con las siguientes credenciales:

usuario: admin
contraseña: secret_password

este usuario se debe utilizar para la ruta de logín y posteriormente utilizar el token que retorna en el header Authorization de cada una de las otras peticiones.

Después de realizar al autenticación en el aplicativo, se debe generar una aplicación de prueba con la ruta de creación de apps. Después de tener una aplicación creada, se pueden comenzar a gestionar las reglas asociadas a la aplicación.