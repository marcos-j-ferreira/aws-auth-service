s🚀 Autodeploy Automatizado: GitOps com GitHub Actions e AWS

Este documento descreve o fluxo de integração contínua (CI) e entrega contínua (CD) que garante que o código seja testado, empacotado e implantado automaticamente na infraestrutura AWS.

1. Estratégia de Branching (Git Flow Simplificado)

Adotamos um modelo de três branches principais, garantindo que o código passe por estágios de validação obrigatórios.

Branch

Propósito

Regra

dev

Desenvolvimento ativo e funcionalidades em andamento.

Única branch para commits diretos.

test

Ambiente de integração e execução dos testes de unidade/funcionais.

Recebe apenas merges da dev.

master

Código pronto para produção (fonte da imagem do Docker).

Recebe apenas merges automáticos da test.

2. Pipeline de Integração Contínua (CI) via GitHub Actions

O gatilho de todo o processo de CI é o merge da branch dev para a branch test.

Etapas do GitHub Actions:

Etapa

Ação

Resultado em Sucesso

1. Testes de Unidade/Integração

Execução completa do pacote de testes da aplicação.

O pipeline continua.

2. Merge Automático (CD)

Se todos os testes passarem, o GitHub Actions realiza automaticamente o merge de test para master.

O código testado está agora na master.

3. Build & Containerização

A partir da branch master, o código é construído e empacotado em uma imagem Docker.

Imagem Docker local criada.

4. Teste em Container

O container é iniciado em ambiente isolado para validação final.

Container validado.

5. Push para AWS ECR

A imagem Docker final é tagueada e enviada para o AWS Elastic Container Registry (ECR).

Imagem disponível na AWS para deploy.

3. Fluxo de Entrega Contínua (CD) na AWS

Uma vez que a imagem é enviada com sucesso para o ECR, o fluxo de deployment é acionado na infraestrutura AWS.

Gatilho do EventBridge: O push da nova imagem para o ECR gera um evento. O AWS EventBridge está configurado para capturar esse evento específico (mudança de estado no repositório ECR).

Invocação da Lambda: O EventBridge, ao detectar o novo evento, invoca uma função AWS Lambda pré-configurada.

Acionamento da Rota de Deploy: A função Lambda executa um código que faz uma chamada HTTP (via POST ou similar) para uma rota de API específica que está sendo executada na instância EC2.

Deploy na EC2: A API na EC2, ao receber a chamada, executa o script de deploy:

Faz o pull da imagem mais recente do ECR.

Para o container antigo (se houver).

Inicia o novo container com a imagem atualizada.

🖼️ Fluxo Completo em ASCII

[Desenvolvedor]
       |
       V
  +----------+
  |  Commit  | <----- APENAS NA
  | (Branch) |           DEV
  +----------+
       |
       V (PR / Merge)
  +----------+
  | Merge p/ |
  |   `dev`  |
  +----------+
       |
       V
  +-------------------------------------------------+
  |      [ CI/CD: GitHub Actions (Pipeline) ]       |
  +-------------------------------------------------+
       |
       +--> 1. Rodar Testes (Unitários/Integração)
       |
       +--> 2. SE SUCESSO: Merge Automático `test` -> `master`
       |
       +--> 3. Build & Teste em Container (usando `master`)
       |
       +--> 4. Push da Imagem
  +-------------------------------------------------+
       |
       V (Notificação de Nova Imagem)
  +-------------------------------------------------+
  |                [ AWS ECR ]                      |
  |                (Imagem Salva)                   |
  +-------------------------------------------------+
       |
       V (Evento: Imagem Pushed)
  +-------------------------------------------------+
  |              [ AWS EventBridge ]                |
  |           (Gatilho: ECR Push)                   |
  +-------------------------------------------------+
       |
       V (Invocação)
  +-------------------------------------------------+
  |              [ AWS Lambda ]                     |
  |           (Chama API da EC2)                    |
  +-------------------------------------------------+
       |
       V (Chamada HTTP p/ Rota de Deploy)
  +-------------------------------------------------+
  |        [ AWS EC2 - API de Deploy ]              |
  |  - Pull da Imagem (ECR)                         |
  |  - Parar/Remover Container Antigo                |
  |  - Rodar Container Novo                          |
  +-------------------------------------------------+
       |
       V
    [ Aplicação Atualizada e Rodando ]
