# Autodeploy Automatizado: GitOps com GitHub Actions e AWS

Este documento descreve o fluxo de **integração contínua (CI)** e **entrega contínua (CD)** utilizado para garantir que o código seja testado, empacotado e implantado automaticamente na infraestrutura AWS.

---

## 1. Estratégia de Branching (Git Flow Simplificado)

O projeto adota um modelo de **três branches principais**, garantindo que o código passe por estágios obrigatórios de validação antes de chegar à produção.

| **Branch** | **Propósito**                                                     | **Regra de Uso**                              |
| ---------- | ----------------------------------------------------------------- | --------------------------------------------- |
| `dev`      | Desenvolvimento ativo e funcionalidades em andamento.             | Única branch com commits diretos.             |
| `test`     | Ambiente de integração e execução de testes unitários/funcionais. | Recebe apenas *merges* da `dev`.              |
| `master`   | Código pronto para produção (base para imagem Docker).            | Recebe apenas *merges automáticos* da `test`. |

---

## 2. Pipeline de Integração Contínua (CI) – GitHub Actions

O processo de CI é iniciado automaticamente quando ocorre um **merge da branch `dev` para `test`**.

### Etapas do Pipeline

| **Etapa**                         | **Ação**                                                                                            | **Resultado em caso de sucesso**       |
| --------------------------------- | --------------------------------------------------------------------------------------------------- | -------------------------------------- |
| 1. Testes de Unidade e Integração | Executa o pacote completo de testes da aplicação.                                                   | Pipeline prossegue.                    |
| 2. Merge Automático (CD)          | Se todos os testes passarem, o GitHub Actions realiza automaticamente o merge de `test` → `master`. | Código testado disponível na `master`. |
| 3. Build e Containerização        | A partir da `master`, o código é construído e empacotado em uma imagem Docker.                      | Imagem Docker gerada localmente.       |
| 4. Teste em Container             | O container é iniciado em ambiente isolado para validação final.                                    | Container validado.                    |
| 5. Push para AWS ECR              | A imagem final é tagueada e enviada ao **AWS Elastic Container Registry (ECR)**.                    | Imagem disponível na AWS para deploy.  |

---

## 3. Fluxo de Entrega Contínua (CD) na AWS

Após o envio da imagem para o ECR, inicia-se automaticamente o processo de **deploy** na infraestrutura AWS.

### Etapas do Fluxo

1. **Gatilho do EventBridge**

   * O envio de uma nova imagem para o ECR gera um evento.
   * O **AWS EventBridge** está configurado para capturar este evento de *push*.

2. **Invocação da Lambda**

   * O EventBridge invoca uma função **AWS Lambda** previamente configurada.

3. **Acionamento da Rota de Deploy (EC2)**

   * A função Lambda faz uma requisição HTTP (GET) para uma rota de API hospedada na instância **EC2**.

4. **Deploy na EC2**
   A API na EC2 executa o script de deploy:

   * Faz o **pull da imagem mais recente** do ECR.
   * **Interrompe** o container antigo (se existir).
   * **Inicia** o novo container com a imagem atualizada.

O resultado é uma **implantação automatizada e validada**, sem necessidade de intervenção manual.

---

## 4. Diagrama do Fluxo Completo (ASCII)

```
[Desenvolvedor]
       |
       v
  +----------+
  |  Commit  |   <- apenas na branch `dev`
  +----------+
       |
       v
  +-----------------------------+
  | Merge: dev -> test (gatilho)|
  +-----------------------------+
       |
       v
  +--------------------------------------------------+
  |      [ CI/CD: GitHub Actions - Pipeline ]        |
  +--------------------------------------------------+
       |
       +--> 1. Executar Testes (Unitários/Integração)
       |
       +--> 2. Se sucesso: Merge automático test -> master
       |
       +--> 3. Build & Teste em Container (base master)
       |
       +--> 4. Push da Imagem Docker p/ AWS ECR
       |
       v
  +--------------------------------------------------+
  |                 [ AWS ECR ]                      |
  |        (Imagem armazenada e versionada)          |
  +--------------------------------------------------+
       |
       v
  +--------------------------------------------------+
  |              [ AWS EventBridge ]                 |
  |       (Detecta evento de push no ECR)            |
  +--------------------------------------------------+
       |
       v
  +--------------------------------------------------+
  |                [ AWS Lambda ]                    |
  |     (Invoca API de Deploy na instância EC2)      |
  +--------------------------------------------------+
       |
       v
  +--------------------------------------------------+
  |           [ AWS EC2 - API de Deploy ]            |
  |   - Pull da nova imagem                          |
  |   - Parada/remoção do container anterior          |
  |   - Execução do container atualizado              |
  +--------------------------------------------------+
       |
       v
     [ Aplicação Atualizada e em Execução ]
```
