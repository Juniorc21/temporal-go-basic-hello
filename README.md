# Temporal Go Hello World Project

Este proyecto es una aplicación de ejemplo que utiliza Temporal para ejecutar un flujo de trabajo simple en Go. La aplicación está diseñada para demostrar los conceptos básicos de Temporal, incluyendo Workflows, Activities, Workers y Starters. 

## Descripción del Proyecto

La aplicación implementa un flujo de trabajo simple que toma una cadena de texto, la formatea y la devuelve. A continuación, se describen los componentes clave de la aplicación:

### Componentes de Temporal

1. **Workflow**
   - Los **Workflows** son funciones que definen el flujo general de la aplicación y representan el aspecto de orquestación de la lógica de negocio. En este proyecto, el flujo de trabajo llamado `GreetingWorkFlow` toma una cadena de texto y la pasa a una actividad para ser formateada.

2. **Activity**
   - Las **Activities** son funciones que se llaman durante la ejecución del flujo de trabajo y representan el aspecto de ejecución de la lógica empresarial. En este proyecto, la actividad formatea la cadena de texto recibida y devuelve la versión formateada.

3. **Worker**
   - Los **Workers** alojan el código de las actividades y los flujos de trabajo y ejecutan el código pieza por pieza. En esta aplicación, un Worker se encarga de ejecutar el `GreetingWorkFlow` y la actividad asociada.

4. **Starter**
   - El **Starter** es un programa que inicia el flujo de trabajo enviando un mensaje al servidor Temporal para que realice un seguimiento del estado del flujo de trabajo. En este proyecto, el Starter también actúa como un servidor HTTP que inicia el flujo de trabajo y devuelve el resultado al cliente.

## Instalación

1. **Clona el repositorio:**

   ```bash
   git clone https://github.com/tu_usuario/temporal-go-hello-world.git
   cd temporal-go-hello-world
# temporal-go-basic-hello
