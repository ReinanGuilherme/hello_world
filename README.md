# hello_world

Este é um projeto utilizado para prática, que demonstra como realizar o Build and Push de uma imagem Docker para um aplicativo Go usando GitHub Actions e a abordagem de armazenamento seguro de valores sensíveis utilizando secrets.

## Objetivo

O objetivo deste projeto é praticar o processo de automação de construção e publicação de uma imagem Docker para um aplicativo Go. Usando GitHub Actions, configuramos um fluxo de trabalho que é acionado automaticamente sempre que há um push para a branch "main" do repositório. Este fluxo de trabalho realiza as seguintes etapas:

1. **Build do Aplicativo**: O código-fonte do aplicativo Go é baixado e o aplicativo é compilado em um contêiner Docker.

2. **Push da Imagem Docker**: Após a construção bem-sucedida, a imagem Docker é empurrada para um registro de contêiner, neste caso, o Docker Hub.
