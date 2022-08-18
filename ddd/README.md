## DDD (Domain Design Drive).
Es la evolución de las arquitecturas convencionales de 1 o 3 capas.

![Arquitectura](guides/architecture.png)

Estas capas se divide en las siguientes:

1. __Presentación:__ No sufre cambios con la arquitectura de 3 capas.
2. __Aplicación:__ Reside los casos de usos.
3. __Dominio:__ Reside toda la lógica del dominio del problema.
4. __Infraestructura:__ La capa de datos se ha renombrado a la capa
                        de infraestructura, ya que no se va a tratar
                        con datos sino con todo lo relacionado a la 
                        infraestructura de nuestro sistema (Datos, 
                        Framework, Login, etc).


__NOTA.__ La capa de lógica de negocio se dividió en 2 que seria la capa de
aplicación y la capa de dominio.