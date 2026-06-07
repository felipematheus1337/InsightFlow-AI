package spring_ai.v1.service;

import java.time.Instant;
import java.util.List;
import java.util.stream.Collectors;

import org.springframework.ai.chat.client.ChatClient;
import org.springframework.ai.openai.OpenAiChatOptions;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.stereotype.Service;

import software.amazon.awssdk.services.sqs.SqsClient;
import software.amazon.awssdk.services.sqs.model.GetQueueUrlRequest;
import software.amazon.awssdk.services.sqs.model.SendMessageRequest;
import tools.jackson.databind.ObjectMapper;

import spring_ai.v1.client.EmployeeClient;
import spring_ai.v1.client.dto.EmployeeResponse;
import spring_ai.v1.document.RelatorioDocument;
import spring_ai.v1.messaging.RelatorioMensagem;
import spring_ai.v1.repository.RelatorioRepository;

@Service
public class RelatorioJobService {

    private static final String SYSTEM_PROMPT = """
            Você é um analista de RH que escreve relatórios gerenciais em português,
            avaliando a carga e o perfil de tarefas de cada funcionário.
            """;

    private static final String USER_PROMPT = """
            Com base na lista de funcionários e nas tarefas de cada um, escreva um relatório
            em texto corrido (sem markdown) destacando: volume de tarefas por funcionário,
            possíveis sobrecargas ou ociosidade e sugestões de melhoria na distribuição do trabalho.

            Funcionários e suas tarefas:
            {funcionarios}
            """;

    private final EmployeeClient employeeClient;
    private final ChatClient chatClient;
    private final RelatorioRepository relatorioRepository;
    private final SqsClient sqsClient;
    private final ObjectMapper objectMapper;

    @Value("${app.aws.sqs.queue-name}")
    private String queueName;

    public RelatorioJobService(EmployeeClient employeeClient,
                               ChatClient chatClient,
                               RelatorioRepository relatorioRepository,
                               SqsClient sqsClient,
                               ObjectMapper objectMapper) {
        this.employeeClient = employeeClient;
        this.chatClient = chatClient;
        this.relatorioRepository = relatorioRepository;
        this.sqsClient = sqsClient;
        this.objectMapper = objectMapper;
    }

    public void executar() {
        var funcionarios = employeeClient.listarFuncionarios().data();
        var textoRelatorio = gerarTextoRelatorio(funcionarios);

        var documento = new RelatorioDocument(textoRelatorio, Instant.now(), "GERADO");
        var salvo = relatorioRepository.save(documento);

        publicarNaFila(salvo);
    }

    private String gerarTextoRelatorio(List<EmployeeResponse> funcionarios) {
        var resumoFuncionarios = funcionarios.stream()
                .map(funcionario -> "- %s (%d anos): %s".formatted(
                        funcionario.name(),
                        funcionario.age(),
                        String.join(", ", funcionario.tasks())))
                .collect(Collectors.joining("\n"));

        return chatClient.prompt()
                .system(SYSTEM_PROMPT)
                .user(u -> u.text(USER_PROMPT).param("funcionarios", resumoFuncionarios))
                .options(OpenAiChatOptions.builder().model("gpt-5.2"))
                .call()
                .content();
    }

    private void publicarNaFila(RelatorioDocument documento) {
        var mensagem = objectMapper.writeValueAsString(new RelatorioMensagem(documento.getRelatorio()));
        var queueUrl = sqsClient.getQueueUrl(GetQueueUrlRequest.builder()
                .queueName(queueName)
                .build())
                .queueUrl();

        sqsClient.sendMessage(SendMessageRequest.builder()
                .queueUrl(queueUrl)
                .messageBody(mensagem)
                .build());
    }

}
